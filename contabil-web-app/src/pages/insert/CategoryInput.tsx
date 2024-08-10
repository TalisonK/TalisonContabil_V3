import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'
import React from 'react'

interface Props {
    error: boolean
    category: string
    setter: (value: string) => void
    categories: { name: string; id: number }[]
    style: any
}

const CategoryInput = (props: Props) => {
    return (
        <FormControl
            style={props.style}
            error={props.error}
        >
            <InputLabel id="demo-simple-select-label">Category</InputLabel>
            <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={props.category}
                label="Category"
                onChange={(e) => props.setter(e.target.value)}
            >
                {props.categories.map((category) => (
                    <MenuItem value={category.name}>{category.name}</MenuItem>
                ))}
            </Select>
        </FormControl>
    )
}

export default CategoryInput
