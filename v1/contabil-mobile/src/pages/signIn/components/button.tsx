import React from 'react';
import { Button } from "@react-native-material/core";
import { StyleSheet, View } from "react-native";

interface Logar{
    handleSignIn:any
}

const Butons = (props:Logar) => {
    return(
        <View >
            <Button title="Entrar" style={styles.botao} color="#7322a3" onPress={props.handleSignIn}/>
        </View>
    )
}

const styles = StyleSheet.create({
    botao:{
        width: 150,
        marginStart: 'auto',
        marginEnd: 40,
        marginTop: 'auto'
    }
})

export default Butons;