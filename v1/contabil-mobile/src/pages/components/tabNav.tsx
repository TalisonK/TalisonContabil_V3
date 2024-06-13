import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import React from "react";
import { View, StyleSheet, Text, TouchableOpacity} from "react-native";
import { faTableList, faUser, faPlus } from "@fortawesome/free-solid-svg-icons";

const Tabs = ({navigation, update}:any) => {
    return (
        <View style={styles.container}>
            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => {navigation.navigate("Explorer")}}>
                <FontAwesomeIcon icon={faTableList} size={30} style={styles.icon} />
                <Text style={styles.texto}>Filter</Text>
            </TouchableOpacity>

            <TouchableOpacity activeOpacity={0.5} style={styles.botao} onPress={() => {navigation.navigate("InsertExpense", {update})}}>
                <FontAwesomeIcon icon={faPlus} size={30} style={styles.icon} />
                <Text style={styles.texto}>Add</Text>
            </TouchableOpacity>
            
            <TouchableOpacity activeOpacity={0.5} style={styles.botao}>
                <FontAwesomeIcon icon={faUser} size={30} style={styles.icon} />
                <Text style={styles.texto}>Users</Text>
            </TouchableOpacity>
        </View>
    )
}


const styles = StyleSheet.create({
    container: {
        position: "absolute",
        zIndex: 10,
        marginTop: 650,
        flexDirection: 'row',
        backgroundColor: '#ffffffab',
        justifyContent: 'center',
        alignContent: 'center',
        alignSelf: 'center',
        borderRadius: 30,
        padding: 4
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

export default Tabs;