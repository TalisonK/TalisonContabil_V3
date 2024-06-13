import { Text, Surface, Divider } from "@react-native-material/core";
import React from "react";
import { View, StyleSheet, TouchableOpacity } from "react-native";
import Buttons from "./buttons";
import { valueHandler } from "../../../services/util";
import { faChartLine, faPlus, faTableList, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";

const Balance = ({navigation, income, expense, update}: any) => {
    return (
        <Surface style={styles.container} elevation={5} category="medium">
            <View style={styles.resume}>
                <View>
                    <Text>Incomes</Text>
                    <View style={styles.values}>
                        <Text style={styles.coin}>R$ </Text>
                        <Text style={styles.income}>{valueHandler(income)}</Text>
                    </View>
                </View>
                <View>
                    <Text>Expenses</Text>
                    <View style={styles.values}>
                        <Text style={styles.coin}>R$ </Text>
                        <Text style={styles.expense}>{valueHandler(expense)}</Text>
                    </View>
                </View>
            </View>
            <Divider/>
            <View style={styles.buttonsContainer}>
            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => {navigation.navigate("Explorer")}}>
                <FontAwesomeIcon icon={faTableList} size={30} style={styles.icon} />
                <Text style={styles.texto}>Filter</Text>
            </TouchableOpacity>

            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => {navigation.navigate("InsertExpense", {update})}}>
                <FontAwesomeIcon icon={faPlus} size={30} style={styles.icon} />
                <Text style={styles.texto}>Add</Text>
            </TouchableOpacity>
            
            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => {navigation.navigate("Chart", {update})}}>
                <FontAwesomeIcon icon={faChartLine} size={30} style={styles.icon} />
                <Text style={styles.texto}>Chart</Text>
            </TouchableOpacity>

            <TouchableOpacity activeOpacity={0.5} style={styles.botao}>
                <FontAwesomeIcon icon={faUser} size={30} style={styles.icon} />
                <Text style={styles.texto}>Users</Text>
            </TouchableOpacity>
            
        </View>
        </Surface>
    )
}


const styles = StyleSheet.create({
    container: {
        backgroundColor: 'rgba(247, 247, 247, 1)',
        width: '80%',
        marginTop: -40,
        zIndex: 100,
        alignSelf: "center",
        padding: 10,
        paddingBottom: 0,
        borderRadius: 10
    },
    resume: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        margin: 5
    },
    values: {
        flexDirection: 'row',
        alignContent: 'center',
        justifyContent: "center"
    },
    income: {
        color: "#4de039"
    },
    expense: {
        color: "#f75a5a"
    },
    coin: {
        fontSize: 11,
        color: '#bdbcbc',
        alignSelf: 'center'
    },
    buttonsContainer: {
        flexDirection: 'row',
        justifyContent: 'space-evenly',
        marginTop:4
    },
    botao: {
        justifyContent: 'center',
        alignContent: "center",
        textAlign: "center",
        width: 60,
        height: 60,
    },
    icon: {
        color: '#000',
        alignSelf: 'center',
    },
    texto: {
        justifyContent: "center",
        fontSize: 14,
        fontFamily: 'Roboto',
        textAlign: 'center'
    },
})

export default Balance;