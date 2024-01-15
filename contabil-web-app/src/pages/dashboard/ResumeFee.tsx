import React from 'react'
import { DisplayFlex } from '../../styles'
import { Text } from '../../styles'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ICONS } from '../../constants/ICONS'

interface Props {
    conta: any
    size: string
}

const ResumeFee = (props: Props) => {
    return (
        <DisplayFlex
            direction="column"
            width={props.size === 'large' ? '200px' : '18%'}
            height={props.size === 'large' ? '150px' : '70px'}
            marginTop={props.size === 'large' ? '20px' : '0px'}
            marginBottom="20px"
            card={true}
            style={{
                borderRadius: '10px',
                textOverflow: 'clip',
            }}
        >
            <Text
                margin="10px"
                style={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignSelf: props.size === 'small' ? 'center' : '',
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
