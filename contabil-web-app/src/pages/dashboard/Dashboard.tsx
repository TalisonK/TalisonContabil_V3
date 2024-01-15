import React, { useState, useEffect } from 'react'
import { DisplayFlex, Text } from '../../styles'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers'
import dayjs from 'dayjs'
import User from '../../interfaces/User'
import DashboardBundle from '../../interfaces/Dashboard'
import { getDashboard } from '../../api/Dashboard'
import Resume from './Resume'
import { Button, CircularProgress, Divider, Skeleton } from '@mui/material'
import Timeline from './timeline'
import IncomeExpense from './IncomeExpense'
import ExpensePie from './ExpensePie'
import PaymentPie from './PaymentPie'
import RecurrentData from './RecurrentData'

const Dashboard = () => {
    const [user, setUser] = useState<User>({} as User)
    const [date, setDate] = useState<Date>(new Date())

    const [bundle, setBundle] = useState<DashboardBundle>({} as DashboardBundle)
    const [loading, setLoading] = useState<boolean>(true)
    const [skeleton, setSkeleton] = useState<boolean>(false)

    useEffect(() => {
        updater()
    }, [date])

    const updater = () => {
        setSkeleton(true)
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
        )
            .then((res) => {
                setBundle(res.data)
                setLoading(false)
                setSkeleton(false)
            })
            .catch((err) => {
                console.log(err)
                localStorage.removeItem('user')
                window.location.href = '/'
            })
    }

    const updateDate = (newDate: any) => {
        const nd = newDate.toDate()
        //setLoading(true)
        if (
            nd.getMonth() !== date.getMonth() ||
            nd.getFullYear() !== date.getFullYear()
        ) {
            setDate(nd)
        } else {
            //setLoading(false)
        }
    }

    const isNow = (compare: Date) => {
        const now = new Date()
        return (
            now.getFullYear() === date.getFullYear() &&
            now.getMonth() === date.getMonth()
        )
    }

    return (
        <>
            {!loading ? (
                <DisplayFlex
                    overflow="auto"
                    width="100%"
                    height="100%"
                    justifyContent="center"
                    style={{ padding: '20px' }}
                >
                    <DisplayFlex direction="column" width="80%" height="90vh">
                        {/* Header */}
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
                                width={isNow(date) ? '135px' : '199px'}
                            >
                                <LocalizationProvider
                                    dateAdapter={AdapterDayjs}
                                >
                                    {!isNow(date) ? (
                                        <Button
                                            style={{ padding: 0, margin: 0 }}
                                            onClick={() => {
                                                setDate(new Date())
                                            }}
                                        >
                                            Today
                                        </Button>
                                    ) : (
                                        <></>
                                    )}
                                    <DatePicker
                                        className="date-picker"
                                        views={['month', 'year']}
                                        format="MMM-YYYY"
                                        value={dayjs(date)}
                                        onYearChange={updateDate}
                                    />
                                </LocalizationProvider>
                            </DisplayFlex>
                        </DisplayFlex>

                        {/* Body */}
                        <DisplayFlex direction="row">
                            {/* Resume */}
                            <DisplayFlex direction="column" width="65%">
                                <DisplayFlex
                                    direction="row"
                                    height="100px"
                                    justifyContent="space-between"
                                    style={{ paddingRight: '10px' }}
                                >
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="32%"
                                            height="100px"
                                            card={true}
                                        >
                                            <Resume
                                                entradas={
                                                    bundle.resumes.incomes
                                                }
                                                type="Income"
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'32%'}
                                            height={'100px'}
                                        />
                                    )}
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="32%"
                                            height="100px"
                                            card={true}
                                        >
                                            <Resume
                                                entradas={
                                                    bundle.resumes.expenses
                                                }
                                                type="Expense"
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'32%'}
                                            height={'100px'}
                                        />
                                    )}
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="32%"
                                            height="100px"
                                            card={true}
                                        >
                                            <Resume
                                                entradas={
                                                    bundle.resumes.balances
                                                }
                                                type="Balance"
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'32%'}
                                            height={'100px'}
                                        />
                                    )}
                                </DisplayFlex>

                                {/* IncomeVSExpense and pie */}
                                <DisplayFlex
                                    direction="row"
                                    width="100%"
                                    height="500px"
                                    justifyContent="space-between"
                                    marginTop="10px"
                                >
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="70%"
                                            height="100%"
                                            card={true}
                                            marginRight="10px"
                                        >
                                            <IncomeExpense
                                                lista={bundle.incomeVSexpense}
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'70%'}
                                            height={'100%'}
                                            style={{ marginRight: '10px' }}
                                        />
                                    )}
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="29%"
                                            height="100%"
                                            card={true}
                                            marginRight="10px"
                                        >
                                            <ExpensePie
                                                data={bundle.expenseByCategory}
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'29%'}
                                            height={'100%'}
                                            style={{ marginRight: '10px' }}
                                        />
                                    )}
                                </DisplayFlex>

                                {/* Chart fixed fees and pie */}
                                <DisplayFlex
                                    direction="row"
                                    width="100%"
                                    height="480px"
                                    justifyContent="space-between"
                                    marginTop="10px"
                                >
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="29%"
                                            height="100%"
                                            card={true}
                                            marginRight="10px"
                                        >
                                            <PaymentPie
                                                data={bundle.expenseByMethod}
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'29%'}
                                            height={'100%'}
                                            style={{ marginRight: '10px' }}
                                        />
                                    )}
                                    {!skeleton ? (
                                        <DisplayFlex
                                            width="70%"
                                            height="100%"
                                            card={true}
                                            marginRight="10px"
                                        >
                                            <RecurrentData
                                                contas={
                                                    bundle.fixatedExpenses
                                                        .contas
                                                }
                                                streaming={
                                                    bundle.fixatedExpenses
                                                        .streaming
                                                }
                                            />
                                        </DisplayFlex>
                                    ) : (
                                        <Skeleton
                                            variant="rectangular"
                                            width={'70%'}
                                            height={'100%'}
                                            style={{ marginRight: '10px' }}
                                        />
                                    )}
                                </DisplayFlex>
                            </DisplayFlex>

                            {/* Timeline */}
                            {!skeleton ? (
                                <DisplayFlex
                                    direction="column"
                                    width="35%"
                                    height="1100px"
                                    card={true}
                                >
                                    <Text
                                        fontSize="1.5em"
                                        margin="0"
                                        style={{ paddingLeft: '10px' }}
                                    >
                                        Timeline
                                    </Text>
                                    <Divider />
                                    <Timeline activities={bundle.timeline} />
                                </DisplayFlex>
                            ) : (
                                <Skeleton
                                    variant="rectangular"
                                    width={'35%'}
                                    height={'1100px'}
                                />
                            )}
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
