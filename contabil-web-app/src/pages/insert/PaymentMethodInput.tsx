import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'
import React from 'react'

interface Props {
    error: boolean
    paymentMethod: string
    setter: (value: string) => void
    style?: any
}

const paymentMethods: string[] = ['CREDIT_CARD', 'DEBIT_CARD', 'MONEY', 'PIX']

const PaymentMethodInput = (props: Props) => {
    return (
        <FormControl
            style={props.style}
            error={props.error}
        >
            <InputLabel id="demo-simple-select-label">
                Payment Method
            </InputLabel>
            <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={props.paymentMethod}
                label="paymentMethod"
                onChange={(e) => props.setter(e.target.value)}
            >
                {paymentMethods.map((paymentMethod) => (
                    <MenuItem value={paymentMethod}>{paymentMethod}</MenuItem>
                ))}
            </Select>
        </FormControl>
    )
}

export default PaymentMethodInput
