import React from 'react'
import { DisplayFlex } from '../../styles'
import { Text } from '../../styles'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ICONS } from '../../constants/ICONS'

const ResumeFee = (props: any) => {
    return (
        <DisplayFlex
            direction="column"
            width="200px"
            height="150px"
            marginTop="10px"
            marginBottom="30px"
            card={true}
            style={{ borderRadius: '10px' }}
        >
            <Text
                margin="10px"
                style={{
                    display: 'flex',
                    justifyContent: 'space-between',
                }}
            >
                {props.conta.description}
                {Object.keys(ICONS).includes(props.conta.description) ? (
                    <FontAwesomeIcon
                        icon={
                            ICONS[props.conta.description as keyof typeof ICONS]
                        }
                    />
                ) : (
                    <></>
                )}
            </Text>
            <DisplayFlex
                direction="column"
                justifyContent="center"
                width="100%"
                height="100%"
                style={{
                    borderRadius: '10px',
                    alignItems: 'center',
                }}
                card={true}
            >
                <Text margin="10px">{props.conta.method}</Text>
                <Text margin="10px">R${props.conta.value.toFixed(2)}</Text>
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default ResumeFee
