import { FormControl, InputAdornment, InputLabel } from '@mui/material'
import React from 'react'
import { OutlinedFieldStyled } from '../../styles'

interface Props {
    value: string
    error: boolean
    setter: (value: string) => void
}

const numericOnly = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ',']

const ValueInput = (props: Props) => {
    const valueHandler = (event: any) => {
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
        props.setter(valuest)
    }

    const valueStart = (event: any) => {
        event.preventDefault()
        const { target } = event
        const extensionStart = target.value.indexOf(',')
        target.focus()
        target.setSelectionRange(0, extensionStart)
    }

    return (
        <FormControl style={{ width: '100%' }}>
            <InputLabel style={{ marginLeft: '7px' }} error={props.error}>
                Value
            </InputLabel>
            <OutlinedFieldStyled
                onFocus={(event) => {
                    valueStart(event)
                }}
                error={props.error}
                value={props.value}
                onChange={(e) => valueHandler(e)}
                id="outlined-adornment-amount"
                startAdornment={
                    <InputAdornment position="start">R$</InputAdornment>
                }
                label="Amount"
            />
        </FormControl>
    )
}

export default ValueInput
