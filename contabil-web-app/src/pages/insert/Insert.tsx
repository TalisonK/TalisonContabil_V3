import React, { useEffect, useState } from 'react'
import { DisplayFlex } from '../../styles'
import {
    Button,
    Divider,
} from '@mui/material'
import { categoryList, submitActivity } from '../../api/insert'
import Category from '../../interfaces/Category'
import User from '../../interfaces/User'
import { VariantType, useSnackbar } from 'notistack'
import CategoryInput from './CategoryInput'
import PaymentMethodInput from './PaymentMethodInput'
import ParcelsInput from './ParcelsInput'
import DescricaoInput from './DescricaoInput'
import ValueInput from './ValueInput'
import CalendarInput from './CalendarInput'
import TypeInput from './TypeInput'
import ListInput from './ListInput'

interface item {
    id: number
    name: string
    price: string
}

const Insert = (props: any) => {
    const { enqueueSnackbar } = useSnackbar()

    const [user, setUser] = useState<User>({} as User)
    const [description, setDescription] = useState<string>('')
    const [value, setValue] = useState<string>('0,00')
    const [paidAt, setPaidAt] = useState<Date>(new Date())
    const [type, setType] = useState<string>('Expense')
    const [category, setCategory] = useState<string>('')
    const [paymentMethod, setPaymentMethod] = useState<string>('')
    const [actualParcel, setActualParcel] = useState<number>(1)
    const [totalParcel, setTotalParcel] = useState<number>(1)
    const [list, setList] = useState<boolean>(false)
    const [listItens, setListItens] = useState<item[]>([])

    const [categories, setCategories] = useState<Category[]>([])

    //Errors

    const [descriptionError, setDescriptionError] = useState<boolean>(false)
    const [valueError, setValueError] = useState<boolean>(false)
    const [categoryError, setCategoryError] = useState<boolean>(false)
    const [paymentMethodError, setPaymentMethodError] = useState<boolean>(false)

    useEffect(() => {
        categoryList()
            .then((res) => {
                setCategories(res)
            })

            .catch((err) => {
                console.log(err)
                localStorage.removeItem('user')
                window.location.href = '/'
            })
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        } else {
            window.location.href = '/'
        }
    }, [])

    useEffect(() => {
        console.log(list)
    }, [list])

    const handleNotificationVariant = (
        messagee: string,
        variant: VariantType
    ) => {
        enqueueSnackbar(messagee, { variant })
    }

    const resetValidation = () => {
        setDescriptionError(false)
        setValueError(false)
        setCategoryError(false)
        setPaymentMethodError(false)
    }

    const validate = () => {
        resetValidation()
        let validated: boolean = true
        if (!description) {
            setDescriptionError(true)
            validated = false
        }

        if (value === '0,00') {
            setValueError(true)
            validated = false
        }

        if (type === 'Expense' && !category) {
            setCategoryError(true)
            validated = false
        }

        if (type === 'Expense' && !paymentMethod) {
            setPaymentMethodError(true)
            validated = false
        }

        if (list && listItens.length === 0) {
            handleNotificationVariant(
                'Please fill in at least one list item!',
                'error'
            )
            validated = false
        }

        return validated
    }

    const resetValues = () => {
        setDescription('')
        setValue('0,00')
        setCategory('')
        setPaymentMethod('')
        setActualParcel(1)
        setTotalParcel(1)
        setList(false)
    }

    const submit = () => {
        if (!validate()) {
            handleNotificationVariant(
                'Please fill in all required fields!',
                'error'
            )
            return
        }

        let data: any = {
            description,
            userId: user.id,
            value: parseFloat(value.replace(',', '.')),
            type,
        }

        if (type === 'Expense') {
            data = {
                ...data,
                categoryName: category,
                paymentMethod,
                actualParcel,
                totalParcel,
                paidAt,
            }
        } else {
            data = { ...data, receivedAt: paidAt }
        }

        if (list) {
            data = { ...data, list: listItens }
        }

        submitActivity(data).then((res) => {
            handleNotificationVariant(
                'Atividade inserida com sucesso!',
                'success'
            )
            resetValues()
        })
    }

    const activityChange = (event: any) => {
        setType(event.target.value)
        resetValues()
    }

    return (
        <DisplayFlex
            card={true}
            direction="column"
            width={type === 'Expense' ? '50%' : '25%'}
            height="90vh"
            marginBottom="10px"
            marginTop="10px"
            style={{ alignSelf: 'center' }}
            overflow="auto"
            dark={props.theme === 'dark'}
            className={props.theme === 'dark' ? 'theme-dinamic' : ''}
        >
            <DisplayFlex
                width="100%"
                height="80px"
                card={true}
                justifyContent="center"
                dark={props.theme === 'dark'}
            >
                <TypeInput
                    theme={props.theme}
                    type={type}
                    activityChange={activityChange}
                />
            </DisplayFlex>
            <DisplayFlex
                direction="row"
                width="100%"
                height="100%"
                style={{ minHeight: '1000px' }}
            >
                {/*ESQUERDA*/}
                <DisplayFlex
                    width={type === 'Expense' ? '50%' : '100%'}
                    height="100%"
                    direction="column"
                    justifyContent="space-evenly"
                    marginTop="40px"
                >
                    <DescricaoInput
                        description={description}
                        error={descriptionError}
                        list={list}
                        setDescription={setDescription}
                        setList={setList}
                        type={type}
                        theme={props.theme}
                    />

                    <DisplayFlex
                        marginRight="70px"
                        marginLeft="70px"
                        marginTop="50px"
                        marginBottom="50px"
                    >
                        <ValueInput
                            theme={props.theme}
                            error={valueError}
                            value={value}
                            setter={setValue}
                        />
                    </DisplayFlex>

                    <CalendarInput paidAt={paidAt} setter={setPaidAt} />
                </DisplayFlex>
                {type === 'Expense' ? (
                    <Divider orientation="vertical" flexItem />
                ) : (
                    <></>
                )}
                {/*DIREITA*/}
                {type === 'Expense' ? (
                    <>
                        <DisplayFlex
                            direction="column"
                            justifyContent="space-evenly"
                            width="50%"
                            height="100%"
                            marginTop="50px"
                        >
                            <CategoryInput
                                error={categoryError}
                                category={category}
                                setter={setCategory}
                                categories={categories}
                                style={{marginRight: "70px", marginLeft:'70px'}}
                                
                            />

                            <PaymentMethodInput
                                error={paymentMethodError}
                                paymentMethod={paymentMethod}
                                setter={setPaymentMethod}
                                style={{marginRight: '70px',
                                    marginLeft: '70px',
                                    marginTop: '70px',
                                    marginBottom: `${
                                        paymentMethod === 'CREDIT_CARD' ? '0px' : '55px'
                                    }`,}}
                            />

                            <ParcelsInput
                                paymentMethod={paymentMethod}
                                actualParcel={actualParcel}
                                setActualParcel={setActualParcel}
                                totalParcel={totalParcel}
                                setTotalParcel={setTotalParcel}
                            />
                        </DisplayFlex>
                    </>
                ) : (
                    <></>
                )}
            </DisplayFlex>
            <Divider variant="middle" flexItem />

            {list ? (
                <ListInput
                    theme={props.theme}
                    rows={listItens}
                    setRows={setListItens}
                />
            ) : (
                <></>
            )}
            <Button
                variant="contained"
                style={{
                    width: '200px',
                    height: '50px',
                    alignSelf: 'center',
                    marginTop: '10px',
                    marginBottom: '10px',
                }}
                onClick={() => submit()}
            >
                Insert
            </Button>
        </DisplayFlex>
    )
}

export default Insert
