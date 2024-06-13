import React from 'react';
import { StatusBar } from 'expo-status-bar';
import { View, Image, StyleSheet, Text } from 'react-native';


const LoginHead: React.FC = () => {
    return (
        <View style={styles.container}>
            <StatusBar style="auto" />
            <Image style={styles.logo} source={require("../../../../assets/logo.png")} />
            <Text style={styles.text}>SING IN</Text>
        </View>
    );
}


const styles = StyleSheet.create({
    container: {
        backgroundColor: '#653780',
        height: 350,
        alignItems: 'center',
        justifyContent: 'center',
        marginTop: -10
    },
    logo: {
        width: 80,
        height: 80
    },

    top: {
        marginStart: 14
    },
    text: {
        fontWeight: 'bold',
        textAlign: 'center',
        fontSize: 20,
        fontFamily: 'sans-serif',
        marginBottom: 70,
        textDecorationColor: '#fff'
    }
})


export default LoginHead;