import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'
import React, { useEffect, useState } from 'react'
import CreditCard from '../../interfaces/CreditCard'
import { getCard } from '../../api/CreditCard'
import { IconFA } from '../../styles'
import { faPlus } from '@fortawesome/free-solid-svg-icons'

interface Props {
    error: boolean
    style?: any
    card: CreditCard
    paymentMethod: string
    setter: (value: any) => void
}

const CreditCardInput = (props: Props) => {


    const [cards, setCards] = useState<CreditCard[]>([])

    useEffect(() => {
        const user = JSON.parse(localStorage.getItem('user') || '{}')
        if (!user.id) {
            window.location.href = '/'
        }

        getCard(user.id).then((response) => {
            setCards(response.data)
        })
    },[])

    const redirectToCards = () => {
        window.location.href = "/Cards"
    }

    return (props.paymentMethod === "CREDIT_CARD" || props.paymentMethod === "DEBIT_CARD") ?
        <FormControl
            style={props.style}
            error={props.error}
        >
            <InputLabel id="demo-simple-select-label">Credit Card</InputLabel>
            <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={props.card.bank}
                label="Category"
            >
                <MenuItem value="" onClick={redirectToCards}><IconFA icon={faPlus}/> Add a Card</MenuItem>
                {cards.map((card) => (
                    <MenuItem value={card.bank} onClick={() => props.setter(card)}>{card.bank}</MenuItem>
                ))}
            </Select>
        </FormControl>
        :<></>
    
}

export default CreditCardInput
