import React from 'react'
import { DisplayFlex, TextFieldStyled } from '../../styles'
import { Switch, Tooltip } from '@mui/material'

interface Props {
    description: string
    setDescription: (value: string) => void
    error: boolean
    type: string
    list: boolean
    setList: (value: boolean) => void
}

const DescricaoInput = (props: Props) => {
    return (
        <DisplayFlex
            marginRight="70px"
            marginLeft="70px"
            style={{ alignItems: 'center' }}
        >
            <TextFieldStyled
                id="outlined-basic"
                label="Description"
                value={props.description}
                onChange={(e) => props.setDescription(e.target.value)}
                variant="outlined"
                error={props.error}
                width="100%"
                className="description-input-space"
            />
            {props.type === 'Expense' ? (
                <DisplayFlex marginLeft="-60px">
                    <Tooltip title="List">
                        <Switch
                            value={props.list}
                            size="medium"
                            onChange={() => {
                                props.setList(!props.list)
                            }}
                        />
                    </Tooltip>
                </DisplayFlex>
            ) : (
                <></>
            )}
        </DisplayFlex>
    )
}

export default DescricaoInput
