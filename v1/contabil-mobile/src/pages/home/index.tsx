import React, { useEffect, useState } from 'react';
import { Button, View } from 'react-native';
import { useAuth } from '../../contexts/auth';
import HomeHeader from './components/header';
import Balance from './components/balance';
import Movimentacao from './components/movimentacao';
import Tabs from '../components/tabNav';
import { getActivitiesByMonth, getExpenseByMonth, getIncomeByMonth } from '../../services/activity';
import { Activity } from '../../interfaces/Activity.interface';

const Home: React.FC = ({navigation}:any) => {

    const [month, setMonth] = useState(new Date());
    const [income, setIncome] = useState(0.0)
    const [expense, setExpense] = useState(0.0)
    const [lista, setLista] = useState([] as Activity[])

    const {verifyToken} = useAuth();

    useEffect(() => {
        console.log("OI") 
        update();
    },[])

    const update = () => {
        getIncomeByMonth(month).then(ret => {setIncome(ret.valor) })
        getExpenseByMonth(month).then(ret => {setExpense(ret.valor) })
        getActivitiesByMonth(month).then(ret => setLista(ret))
    }
    

    const nextMonth = () => {
        const aux = month;
        aux.setMonth(aux.getMonth() + 1);
        setMonth(aux)
    }

    const LastMonth = () => {
        const aux = month;
        aux.setMonth(aux.getMonth() - 1);
        setMonth(aux)
    }

    return(
        <View>
            <HomeHeader month={month} nextMonth={nextMonth} LastMonth={LastMonth} setIncome={setIncome} setExpense={setExpense} setLista={setLista}/>
            <Balance navigation={navigation} income={income} expense={expense} update={update}/>
            <Movimentacao activities={lista} update={update}/>
        </View>
    )
}

export default Home;