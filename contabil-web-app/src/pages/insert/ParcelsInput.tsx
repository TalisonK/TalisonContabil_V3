import React from 'react'
import { DisplayFlex } from '../../styles'
import { TextField } from '@mui/material'

interface Props {
    paymentMethod: string
    actualParcel: number
    setActualParcel: (value: number) => void
    totalParcel: number
    setTotalParcel: (value: number) => void
}

const ParcelsInput = (props: Props) => {
    return (
        <DisplayFlex
            justifyContent="center"
            marginBottom="250px"
            marginTop="150px"
        >
            {props.paymentMethod === 'CREDIT_CARD' ? (
                <>
                    <TextField
                        id="outlined-basic"
                        label="Actual Parcel"
                        variant="outlined"
                        value={props.actualParcel}
                        onChange={(e) =>
                            props.setActualParcel(parseInt(e.target.value))
                        }
                        style={{ width: '200px' }}
                    />
                    <TextField
                        id="outlined-basic"
                        label="Total Parcel"
                        variant="outlined"
                        value={props.totalParcel}
                        onChange={(e) =>
                            props.setTotalParcel(parseInt(e.target.value))
                        }
                        style={{ width: '200px' }}
                    />
                </>
            ) : (
                <></>
            )}
        </DisplayFlex>
    )
}

export default ParcelsInput
