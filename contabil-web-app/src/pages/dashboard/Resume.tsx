import React from 'react'
import { DisplayFlex, Text } from '../../styles'
import Totals from '../../interfaces/Totals'
import ResumeBundle from '../../interfaces/Resume'

interface Props {
    entradas: ResumeBundle
    type: string
}

const Resume = (props: Props) => {
    return (
        <DisplayFlex direction="row" justifyContent="space-between" width="95%">
            <DisplayFlex direction="column">
                <Text fontSize="1.5em" margin="3px">
                    {props.type}
                </Text>
                <Text marginLeft="10px">
                    R$ {props.entradas.actual.toFixed(2)}
                </Text>
            </DisplayFlex>
            <DisplayFlex
                direction="column"
                justifyContent="center"
                marginTop="20px"
            >
                <Text fontSize="1em" style={{ alignSelf: 'center' }}>
                    {`${props.entradas.balance} %`}
                </Text>
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default Resume
