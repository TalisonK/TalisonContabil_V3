import React from 'react'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { DisplayFlex } from '../../styles'
import { DateCalendar, LocalizationProvider } from '@mui/x-date-pickers'
import dayjs from 'dayjs'

interface Props {
    paidAt: Date
    setter: (date: Date) => void
}

const CalendarInput = (props: Props) => {
    return (
        <DisplayFlex direction="column" width="100%">
            <DisplayFlex
                width="100%"
                justifyContent="center"
                marginBottom="20px"
            >
                {props.paidAt.toLocaleDateString()}
            </DisplayFlex>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DateCalendar
                    value={dayjs(props.paidAt)}
                    onChange={(newValue: any) =>
                        props.setter(newValue.toDate())
                    }
                />
            </LocalizationProvider>
        </DisplayFlex>
    )
}

export default CalendarInput
