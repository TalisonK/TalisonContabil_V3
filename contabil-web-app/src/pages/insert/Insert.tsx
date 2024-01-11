import React, { useEffect, useState } from 'react'
import { DisplayFlex, OutlinedFieldStyled, TextFieldStyled } from '../../styles'
import {
    Button,
    Divider,
    FormControl,
    FormHelperText,
    InputAdornment,
    InputLabel,
    MenuItem,
    OutlinedInput,
    Select,
    TextField,
} from '@mui/material'
import {
    DateCalendar,
    DatePicker,
    LocalizationProvider,
    StaticDatePicker,
} from '@mui/x-date-pickers'
import dayjs from 'dayjs'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { categoryList, submitActivity } from '../../api/insert'
import Category from '../../interfaces/Category'
import User from '../../interfaces/User'
import { VariantType, useSnackbar } from 'notistack'

const numericOnly = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',']

const Insert = () => {
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

    const [categories, setCategories] = useState<Category[]>([])
    const [paymentMethods, _] = useState<string[]>([
        'CREDIT_CARD',
        'DEBIT_CARD',
        'MONEY',
        'PIX',
    ])

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

        return validated
    }

    const resetValues = () => {
        setDescription('')
        setValue('0,00')
        setCategory('')
        setPaymentMethod('')
        setActualParcel(1)
        setTotalParcel(1)
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
            user: user.name,
            value: parseFloat(value.replace(',', '.')),
            paidAt,
            category,
            paymentMethod,
            actualParcel,
            totalParcel,
            type,
        }

        if (type === 'Income') {
            data = { ...data, receivedAt: paidAt }
            delete data.paidAt
            delete data.actualParcel
            delete data.totalParcel
            delete data.paymentMethod
            delete data.category
        }

        submitActivity(data).then((res) => {
            handleNotificationVariant(
                'Atividade inserida com sucesso!',
                'success'
            )
            resetValues()
        })
    }

    const ValueHandler = (event: any) => {
        event.preventDefault()
        const { target } = event
        let valuest = target.value

        if (!valuest.split('').includes(',')) return

        const [real, cents] = valuest.split(',')

        for (let i = 0; i < valuest.length; i++) {
            if (!numericOnly.includes(valuest[i])) {
                return
            }
        }

        if (valuest.match(/,/g)?.length === 2) {
            const extensionStart = target.value.indexOf(',')
            setTimeout(() => {
                target.focus()
                target.setSelectionRange(extensionStart + 1, extensionStart + 1)
            }, 50)
            return
        }

        if (cents.length > 2) {
            const centsArray = cents.split('')
            centsArray.splice(2, 1)

            valuest = real.concat(',', centsArray.join(''))
            const extensionStart = target.value.indexOf(',')
            setTimeout(() => {
                target.focus()
                target.setSelectionRange(extensionStart + 2, extensionStart + 2)
            }, 20)
        }

        if (cents.length < 2) {
            for (let i = cents.length; i < 2; i++) {
                cents.concat('0')
            }

            valuest = real.concat(',', cents)
        }

        if (valuest[valuest.length - 1] === ',') return
        setValue(valuest)
    }

    const valueStart = (event: any) => {
        event.preventDefault()
        const { target } = event
        const extensionStart = target.value.indexOf(',')
        target.focus()
        target.setSelectionRange(0, extensionStart)
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
        >
            <DisplayFlex
                width="100%"
                height="80px"
                card={true}
                justifyContent="center"
            >
                <Select
                    style={{
                        height: '80%',
                        alignSelf: 'center',
                        width: '200px',
                        textAlign: 'center',
                    }}
                    labelId="demo-simple-select-label"
                    id="demo-simple-select"
                    value={type}
                    label="Type"
                    notched={false}
                    onChange={(e) => setType(e.target.value)}
                >
                    <MenuItem value={'Expense'}>Expense</MenuItem>
                    <MenuItem value={'Income'}>Income</MenuItem>
                </Select>
            </DisplayFlex>
            <DisplayFlex direction="row" width="100%" height="100%">
                <DisplayFlex
                    width={type === 'Expense' ? '50%' : '100%'}
                    height="100%"
                    direction="column"
                    justifyContent="space-evenly"
                    marginTop="40px"
                >
                    <TextFieldStyled
                        id="outlined-basic"
                        label="Description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        variant="outlined"
                        marginRight="70px"
                        marginLeft="70px"
                        error={descriptionError}
                    />
                    <FormControl
                        style={{
                            marginRight: '70px',
                            marginLeft: '70px',
                            marginTop: '50px',
                            marginBottom: '50px',
                        }}
                    >
                        <InputLabel
                            style={{ marginLeft: '7px' }}
                            error={valueError}
                        >
                            Value
                        </InputLabel>
                        <OutlinedFieldStyled
                            onFocus={(event) => {
                                valueStart(event)
                            }}
                            error={valueError}
                            value={value}
                            onChange={(e) => ValueHandler(e)}
                            id="outlined-adornment-amount"
                            startAdornment={
                                <InputAdornment position="start">
                                    R$
                                </InputAdornment>
                            }
                            label="Amount"
                        />
                    </FormControl>
                    <DisplayFlex direction="column" width="100%">
                        <DisplayFlex
                            width="100%"
                            justifyContent="center"
                            marginBottom="20px"
                        >
                            {paidAt.toLocaleDateString()}
                        </DisplayFlex>
                        <LocalizationProvider dateAdapter={AdapterDayjs}>
                            <DateCalendar
                                value={dayjs(paidAt)}
                                onChange={(newValue: any) =>
                                    setPaidAt(newValue.toDate())
                                }
                            />
                        </LocalizationProvider>
                    </DisplayFlex>
                </DisplayFlex>
                {type === 'Expense' ? (
                    <Divider orientation="vertical" flexItem />
                ) : (
                    <></>
                )}
                {type === 'Expense' ? (
                    <>
                        <DisplayFlex
                            direction="column"
                            justifyContent="space-evenly"
                            width="50%"
                            height="100%"
                            marginTop="50px"
                        >
                            <FormControl
                                style={{
                                    marginRight: '70px',
                                    marginLeft: '70px',
                                }}
                                error={categoryError}
                            >
                                <InputLabel id="demo-simple-select-label">
                                    Category
                                </InputLabel>
                                <Select
                                    labelId="demo-simple-select-label"
                                    id="demo-simple-select"
                                    value={category}
                                    label="Category"
                                    onChange={(e) =>
                                        setCategory(e.target.value)
                                    }
                                >
                                    {categories.map((category) => (
                                        <MenuItem value={category.name}>
                                            {category.name}
                                        </MenuItem>
                                    ))}
                                </Select>
                            </FormControl>

                            <FormControl
                                style={{
                                    marginRight: '70px',
                                    marginLeft: '70px',
                                    marginTop: '70px',
                                    marginBottom: `${
                                        paymentMethod === 'CREDIT_CARD'
                                            ? '0px'
                                            : '55px'
                                    }`,
                                }}
                                error={paymentMethodError}
                            >
                                <InputLabel id="demo-simple-select-label">
                                    Payment Method
                                </InputLabel>
                                <Select
                                    labelId="demo-simple-select-label"
                                    id="demo-simple-select"
                                    value={paymentMethod}
                                    label="paymentMethod"
                                    onChange={(e) =>
                                        setPaymentMethod(e.target.value)
                                    }
                                >
                                    {paymentMethods.map((paymentMethod) => (
                                        <MenuItem value={paymentMethod}>
                                            {paymentMethod}
                                        </MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                            <DisplayFlex
                                justifyContent="center"
                                marginBottom="250px"
                                marginTop="150px"
                            >
                                {paymentMethod === 'CREDIT_CARD' ? (
                                    <>
                                        <TextField
                                            id="outlined-basic"
                                            label="Actual Parcel"
                                            variant="outlined"
                                            value={actualParcel}
                                            onChange={(e) =>
                                                setActualParcel(
                                                    parseInt(e.target.value)
                                                )
                                            }
                                            style={{ width: '200px' }}
                                        />
                                        <TextField
                                            id="outlined-basic"
                                            label="Total Parcel"
                                            variant="outlined"
                                            value={totalParcel}
                                            onChange={(e) =>
                                                setTotalParcel(
                                                    parseInt(e.target.value)
                                                )
                                            }
                                            style={{ width: '200px' }}
                                        />
                                    </>
                                ) : (
                                    <></>
                                )}
                            </DisplayFlex>
                        </DisplayFlex>
                    </>
                ) : (
                    <></>
                )}
            </DisplayFlex>
            <Divider variant="middle" flexItem />
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

/*
    private LocalDateTime paidAt;
    private String category;
    private EnumPaymentMethod paymentMethod;
    private Integer actualParcel;
    private Integer totalParcel;



    private String user;

    private String description;
    private Double value;






*/
