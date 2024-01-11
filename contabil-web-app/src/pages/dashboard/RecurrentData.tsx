import React, { useState } from 'react'
import { DisplayFlex } from '../../styles'
import { Text } from '../../styles'
import ResumeFee from './ResumeFee'
import Activity from '../../interfaces/Activity'

const fixedFee = ['energia', 'agua', 'internet', 'vivo']

interface Props {
    contas: Activity[]
    streaming: Activity[]
}

const RecurrentData = (props: Props) => {
    const feeShuffler = (fee: string) => {
        for (let i = 0; i < props.contas.length; i++) {
            if (props.contas[i].description === fee) {
                return props.contas[i]
            }
        }
        return {
            description: fee,
            method: 'Waiting',
            value: 0,
        }
    }

    return (
        <DisplayFlex
            direction="column"
            width="100%"
            height="100%"
            justifyContent="center"
        >
            <DisplayFlex width="100%" height="50px" card={true}>
                <Text fontSize="1.3em" margin="10px">
                    Reccurent fees
                </Text>
            </DisplayFlex>
            <DisplayFlex
                width="100%"
                height="100%"
                direction="column"
                style={{ alignItems: 'center' }}
            >
                {/* Contas */}
                <DisplayFlex
                    width="100%"
                    height="220px"
                    justifyContent="space-evenly"
                    style={{ alignItems: 'center' }}
                >
                    {fixedFee.map((fee) => (
                        <ResumeFee conta={feeShuffler(fee)} />
                    ))}
                </DisplayFlex>

                {/* Streaming */}
                <DisplayFlex></DisplayFlex>
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default RecurrentData
