import React, { useEffect, useState } from "react";
import { StatusBar } from 'expo-status-bar';
import { View, StyleSheet, TouchableOpacity } from "react-native";
import { Button, Text } from "@react-native-material/core";
import Icon from "@expo/vector-icons/MaterialCommunityIcons";
import { useAuth } from "../../../contexts/auth";
import { monthHandler } from "../../../services/util";
import { getActivitiesByMonth, getExpenseByMonth, getIncomeByMonth } from "../../../services/activity";


const HomeHeader = ({ month, nextMonth, LastMonth, setIncome, setExpense, setLista }: any) => {

    const [monthDisplay, setMonth] = useState("")

    useEffect(() => {
        setTimeout(() => {
            setMonth(monthHandler(month));
            getIncomeByMonth(month).then(ret => { setIncome(ret.valor) })
            getExpenseByMonth(month).then(ret => { setExpense(ret.valor) })
            getActivitiesByMonth(month).then(ret => {setLista(ret)})
        }, 50)

    }, [monthDisplay])


    const nextHandler = () => {
        nextMonth();
        setMonth(monthHandler(month));
    }

    const lastHandler = () => {
        LastMonth();
        setMonth(monthHandler(month));
    }

    const { signOut } = useAuth();

    return (
        <View style={styles.container}>
            <StatusBar style="auto" />
            <View style={styles.body}>

                <TouchableOpacity activeOpacity={0.9} style={styles.button} onPress={signOut}>
                    <Icon name="close" size={30} color={'#FFF'} />
                </TouchableOpacity>
                <View style={styles.month}>
                    <TouchableOpacity activeOpacity={0.9} onPress={lastHandler} style={{ marginStart: 35 }}><Icon name="arrow-left" size={30} color={'#FFF'} /></TouchableOpacity>
                    <Text style={styles.titulo}>{monthDisplay ? monthDisplay : "Titulo"}</Text>
                    <TouchableOpacity activeOpacity={0.9} onPress={nextHandler} style={{ marginEnd: 35 }}><Icon name="arrow-right" size={30} color={'#FFF'} /></TouchableOpacity>
                </View>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: '#653780',
        height: 170,
        alignItems: 'center',
        marginTop: -10
    },
    button: {
        marginEnd: 10,
        marginStart: 'auto'
    },
    titulo: {
        color: '#FFF',
        marginStart: 20,
        textAlign: 'center',
        justifyContent: 'center',
        alignSelf: 'center'
    },
    body: {
        width: "100%",
        height: 40,
        justifyContent: 'space-between',
        alignContent: 'center',
        textAlign: 'center',
        marginTop: 44
    },
    month: {
        marginTop: 10,
        flexDirection: 'row',
        justifyContent: "space-between"
    }
})


export default HomeHeader;