import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import React from "react";
import { View, StyleSheet, TouchableOpacity, Text, Image } from "react-native";
import { faMoneyCheckDollar, faTableList, faUser } from "@fortawesome/free-solid-svg-icons";

const Buttons = (props:any) => {
    return (
        <View style={styles.container}>
            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => props.navigation.navigate("InsertExpense")}>
                <FontAwesomeIcon icon={faMoneyCheckDollar} size={33} style={styles.icon}/>
                <Text style={styles.texto}>Add</Text>
            </TouchableOpacity>
            <TouchableOpacity activeOpacity={0.5} style={styles.botao}>
                <FontAwesomeIcon icon={faTableList} size={30} style={styles.icon}/>
                <Text style={styles.texto}>Filter</Text>
            </TouchableOpacity>
            <TouchableOpacity activeOpacity={0.5} style={styles.botao}>
                <FontAwesomeIcon icon={faUser} size={30} style={styles.icon}/>
                <Text style={styles.texto}>Users</Text>
            </TouchableOpacity>
        </View>
    )
}


const styles = StyleSheet.create({
    container: {
        justifyContent: 'space-evenly',
        flexDirection: 'row',
        alignContent: "center",
        margin: 'auto',
    },
    botao: {
        justifyContent: 'center',
        alignContent: "center",
        textAlign: "center",
        width: 60,
        height: 60,
        marginTop: 10,
    },
    texto: {
        justifyContent: "center",
        fontSize: 14,
        fontFamily: 'Roboto',
        textAlign: 'center'
    },
    icon: {
        color: '#000',
        alignSelf: 'center',
    }
})

export default Buttons;