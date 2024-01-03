import React from 'react'
import { DisplayFlex, Text } from '../../styles'
import Totals from '../../interfaces/Totals'

interface Props {
    entradas: Totals[]
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
                    R$ {Number.parseInt(props.entradas[1].value.toFixed(4))}
                </Text>
            </DisplayFlex>
            <DisplayFlex>
                <Text fontSize="1em" style={{ alignSelf: 'center' }}>
                    {100 -
                        Number.parseInt(
                            (
                                (props.entradas[0].value * 100) /
                                props.entradas[1].value
                            ).toFixed(2)
                        )}
                    %
                </Text>
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default Resume
