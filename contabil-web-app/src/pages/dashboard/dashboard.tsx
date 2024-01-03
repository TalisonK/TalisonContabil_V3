import React, { useState, useEffect } from 'react'
import { DisplayFlex, Text } from '../../styles'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers'
import dayjs from 'dayjs'
import User from '../../interfaces/User'
import DashboardBundle from '../../interfaces/Dashboard'
import { getDashboard } from '../../api/Dashboard'
import Resume from './resume'
import { CircularProgress } from '@mui/material'

const Dashboard = () => {
    const [user, setUser] = useState<User>({} as User)
    const [date, setDate] = useState<Date>(new Date())

    const [bundle, setBundle] = useState<DashboardBundle>({} as DashboardBundle)

    useEffect(() => {
        updater()
    }, [date])

    const updater = () => {
        updateBundle(updateUser())
    }

    const updateUser = () => {
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
            return '' + JSON.parse(userStorage).id
        }
        return ''
    }

    const updateBundle = (id: string) => {
        getDashboard(
            id,
            '' + date.getFullYear(),
            date
                .toLocaleString('default', { month: 'short' })
                .slice(0, 1)
                .toLocaleUpperCase() +
                date.toLocaleString('default', { month: 'short' }).slice(1, 3)
        ).then((res) => {
            setBundle(res.data)
            console.log(res.data)
        })
    }

    const updateDate = (newDate: any) => {
        const nd = newDate.toDate()
        if (
            nd.getMonth() !== date.getMonth() ||
            nd.getFullYear() !== date.getFullYear()
        ) {
            setDate(nd)
            updateBundle(user.id)
        }
    }

    return (
        <>
            {bundle.resumes ? (
                <DisplayFlex
                    overflow="auto"
                    width="100%"
                    height="100%"
                    justifyContent="center"
                    style={{ padding: '20px' }}
                >
                    <DisplayFlex direction="column" width="80%" height="100vh">
                        <DisplayFlex
                            direction="row"
                            width="100%"
                            height="60px"
                            marginBottom="10px"
                            card={true}
                            style={{
                                alignItems: 'center',
                            }}
                        >
                            <Text
                                fontSize="2em"
                                margin="0"
                                style={{ paddingLeft: '10px' }}
                            >
                                Hi, {user.name}!
                            </Text>
                            <Text
                                fontSize="0.7em"
                                marginTop="auto"
                                marginBottom="13px"
                            >
                                {user.role === 'ROLE_ADMIN' ? 'Admin' : ''}
                            </Text>
                            <DisplayFlex
                                marginLeft="auto"
                                marginRight="5px"
                                width="135px"
                            >
                                <LocalizationProvider
                                    dateAdapter={AdapterDayjs}
                                >
                                    <DatePicker
                                        views={['month', 'year']}
                                        format="MMM-YYYY"
                                        value={dayjs(date)}
                                        onYearChange={(newDate) =>
                                            updateDate(newDate)
                                        }
                                    />
                                </LocalizationProvider>
                            </DisplayFlex>
                        </DisplayFlex>
                        <DisplayFlex direction="row">
                            <DisplayFlex direction="column" width="65%">
                                <DisplayFlex
                                    direction="row"
                                    height="100px"
                                    justifyContent="space-between"
                                    style={{ paddingRight: '10px' }}
                                >
                                    <DisplayFlex
                                        width="32%"
                                        height="100px"
                                        card={true}
                                    >
                                        <Resume
                                            entradas={bundle.resumes.incomes}
                                            type="Income"
                                        />
                                    </DisplayFlex>
                                    <DisplayFlex
                                        width="32%"
                                        height="100px"
                                        card={true}
                                    >
                                        <Resume
                                            entradas={bundle.resumes.expenses}
                                            type="Expense"
                                        />
                                    </DisplayFlex>
                                    <DisplayFlex
                                        width="32%"
                                        height="100px"
                                        card={true}
                                    >
                                        <Resume
                                            entradas={bundle.resumes.balances}
                                            type="Balance"
                                        />
                                    </DisplayFlex>
                                </DisplayFlex>
                                <DisplayFlex
                                    direction="row"
                                    width="100%"
                                    height="500px"
                                    justifyContent="space-between"
                                    marginTop="10px"
                                >
                                    <DisplayFlex
                                        width="60%"
                                        height="100%"
                                        card={true}
                                        marginRight="10px"
                                    ></DisplayFlex>
                                    <DisplayFlex
                                        width="39%"
                                        height="100%"
                                        card={true}
                                        marginRight="10px"
                                    ></DisplayFlex>
                                </DisplayFlex>
                                <DisplayFlex
                                    direction="row"
                                    width="100%"
                                    height="480px"
                                    justifyContent="space-between"
                                    marginTop="10px"
                                >
                                    <DisplayFlex
                                        width="39%"
                                        height="100%"
                                        card={true}
                                        marginRight="10px"
                                    ></DisplayFlex>
                                    <DisplayFlex
                                        width="60%"
                                        height="100%"
                                        card={true}
                                        marginRight="10px"
                                    ></DisplayFlex>
                                </DisplayFlex>
                            </DisplayFlex>
                            <DisplayFlex
                                width="35%"
                                height="1100px"
                                card={true}
                            ></DisplayFlex>
                        </DisplayFlex>
                    </DisplayFlex>
                </DisplayFlex>
            ) : (
                <DisplayFlex width="100%" height="100vh">
                    <CircularProgress style={{ margin: 'auto' }} />
                </DisplayFlex>
            )}
        </>
    )
}

export default Dashboard
