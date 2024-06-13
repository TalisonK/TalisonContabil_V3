import React, { useState, useEffect } from 'react';
import { StatusBar } from 'expo-status-bar';
import { View, Image, StyleSheet, ActivityIndicator, Text } from 'react-native';
import { ping } from '../../services/login';

const Loading:React.FC = ({navigation}:any) => {

    const [loading, setLoading] = useState(false)
    const [crash, setCrash] = useState(false)

    useEffect(() => {
        ping().then((ret) => {
            if(ret){
                setLoading(true)
            }
            else{
                setCrash(true)
            }
        })
    }, [])

    if(crash){
        return (
            <View style={styles.body}>
                <StatusBar style="auto" />
                <Image style={styles.logo} source={require("../../../assets/logo.png")} />
                <Text style={styles.text}>Não foi possivel acessar o backEnd</Text>
            </View>
        )
    }
    else if (!loading) {
        return (
            <View style={styles.body}>
                <StatusBar style="auto" />
                <Image style={styles.logo} source={require("../../../assets/logo.png")} />
                <Text style={styles.text}>Contabilidades</Text>
                <ActivityIndicator size='large' style={styles.spinner} />
            </View>
        )
    }

    else{
        navigation.navigate('SignIn')
        return<View/>
    }
}

const styles = StyleSheet.create({
    logo: {
        width: 100,
        height: 100
    },
    body: {
        width: '100%',
        height: '100%',
        justifyContent: 'center',
        alignItems: 'center'
    },
    text:{
        textAlign: 'center',
        fontSize: 20,
        fontFamily: 'sans-serif'
    },
    spinner:{
        marginTop: 20
    }
})

export default Loading;
