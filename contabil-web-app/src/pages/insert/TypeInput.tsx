import { MenuItem, Select } from '@mui/material'
import React from 'react'

interface Props {
    type: string
    activityChange: (e: any) => void
}

const TypeInput = (props: Props) => {
    return (
        <Select
            style={{
                height: '80%',
                alignSelf: 'center',
                width: '200px',
                textAlign: 'center',
            }}
            labelId="demo-simple-select-label"
            id="demo-simple-select"
            value={props.type}
            label="Type"
            notched={false}
            onChange={(e) => props.activityChange(e)}
        >
            <MenuItem value={'Expense'}>Expense</MenuItem>
            <MenuItem value={'Income'}>Income</MenuItem>
        </Select>
    )
}

export default TypeInput
