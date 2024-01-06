import React, { useEffect, useState } from 'react'
import {
    PieChart,
    Pie,
    Cell,
    ResponsiveContainer,
    Tooltip,
    Legend,
} from 'recharts'
import { CATEGORY_COLORS } from '../../constants/COLORS'
import { DisplayFlex, Text } from '../../styles'

const colors: any = {
    CREDIT_CARD: '#a724db',
    DEBIT_CARD: '#24dbb5',
    MONEY: '#db2424',
    PIX: '#dbb524',
    TRANSFER: '#24db24',
}

interface PaymentPieProps {
    data: any
}

const PaymentPie = (props: PaymentPieProps) => {
    const [data, setData] = useState([] as any[])

    useEffect(() => {
        const names = Object.keys(props.data)

        const aux = names.map((name) => ({ name, value: props.data[name] }))

        setData(aux as any[])
    }, [props.data])

    return (
        <DisplayFlex direction="column" width="100%" height="100%">
            <Text textAlign="center">Metodos do mês</Text>
            <ResponsiveContainer width="100%" height="100%">
                <PieChart width={800} height={800}>
                    <Pie data={data} fill="#8884d8" dataKey="value">
                        {data.map((entry, index) => (
                            <Cell
                                key={`cell-${index}`}
                                fill={colors[entry.name]}
                            />
                        ))}
                    </Pie>
                    <Tooltip />
                    <Legend />
                </PieChart>
            </ResponsiveContainer>
        </DisplayFlex>
    )
}

export default PaymentPie
