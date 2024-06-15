import React from 'react'
import { DisplayFlex } from '../../styles'
import IncomeVSExpense from '../../interfaces/IncomeVSExpense'
import {
    CartesianGrid,
    Legend,
    Line,
    LineChart,
    ResponsiveContainer,
    Tooltip,
    XAxis,
    YAxis,
} from 'recharts'

interface Props {
    lista: IncomeVSExpense[]
}

const IncomeExpense = (props: Props) => {
    const data = () => {
        return props.lista
    }

    return (
        <DisplayFlex
            width="100%"
            height="100%"
            direction="column"
            justifyContent="center"
        >
            <ResponsiveContainer width="100%" height="97%">
                <LineChart
                    data={data()}
                    margin={{
                        top: 5,
                        right: 30,
                        bottom: 5,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="month" />
                    <YAxis />
                    <Tooltip />
                    <Legend verticalAlign="top" iconType="circle" />
                    <Line
                        type="monotone"
                        dataKey="income"
                        stroke="#048063"
                        strokeWidth={3}
                    />
                    <Line
                        type="monotone"
                        dataKey="expense"
                        stroke="#db5151"
                        strokeWidth={3}
                    />
                </LineChart>
            </ResponsiveContainer>
        </DisplayFlex>
    )
}

export default IncomeExpense
