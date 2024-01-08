import React, { PureComponent } from 'react'
import {
    BarChart,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    Tooltip,
    Legend,
    ResponsiveContainer,
} from 'recharts'
import { DisplayFlex } from '../../styles'

const data = [
    {
        name: 'Page A',
        uv: 4000,
        pv: 2400,
        amt: 2400,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page B',
        uv: 3000,
        pv: 1398,
        amt: 2210,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page C',
        uv: 2000,
        pv: 9800,
        amt: 2290,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page D',
        uv: 2780,
        pv: 3908,
        amt: 2000,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page E',
        uv: 1890,
        pv: 4800,
        amt: 2181,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page F',
        uv: 2390,
        pv: 3800,
        amt: 2500,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
    {
        name: 'Page G',
        uv: 3490,
        pv: 4300,
        amt: 2100,
        coisa1: 420,
        coisa2: 920,
        coisa3: 150,
    },
]

const BarCategory = () => {
    return (
        <DisplayFlex width="100%" height="100%">
            <ResponsiveContainer width="100%" height="100%">
                <BarChart width={500} height={300} data={data}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" />
                    <YAxis />
                    <Tooltip />
                    <Legend verticalAlign="top" />
                    <Bar dataKey="pv" stackId="a" fill="#8884d8" />
                    <Bar dataKey="uv" stackId="a" fill="#82ca9d" />
                    <Bar dataKey="amt" stackId="a" fill="#005721" />
                    <Bar dataKey="coisa1" stackId="a" fill="#f8e217" />
                    <Bar dataKey="coisa2" stackId="a" fill="#1410d6" />
                    <Bar dataKey="coisa3" stackId="a" fill="#f03131" />
                </BarChart>
            </ResponsiveContainer>
        </DisplayFlex>
    )
}

export default BarCategory
