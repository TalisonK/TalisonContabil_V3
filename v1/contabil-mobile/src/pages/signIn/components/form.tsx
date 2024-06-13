import React from 'react';
import { View, StyleSheet, SafeAreaView, Text } from 'react-native';

import { Surface, TextInput } from '@react-native-material/core';
import Icon from "@expo/vector-icons/MaterialCommunityIcons";


interface FormObjects{
    user:string,
    setUser(user:string):any,
    password:string,
    setPassword(password:string):any
}

const FormHead = (props:FormObjects) => {
    return (
        <Surface style={styles.container} elevation={5} category="medium">
            <SafeAreaView style={styles.textin}>
                <View>
                    <TextInput
                        label="Nome"
                        leading={props => <Icon name="account" {...props} />}
                        onChangeText={props.setUser}
                        value={props.user}
                        style={styles.entries}
                    />
                </View>
                <View>
                    <TextInput
                        label="Senha"
                        leading={props => <Icon name="lock" {...props} />}
                        secureTextEntry={true}
                        onChangeText={props.setPassword}
                        value={props.password}
                        style={styles.entries}
                    />
                </View>

            </SafeAreaView>
        </Surface>
    );
}


const styles = StyleSheet.create({
    container: {
        backgroundColor: 'rgba(247, 247, 247, 1)',
        height: 150,
        borderRadius: 10,
        justifyContent: 'center',
        marginTop: -90,
        marginStart: 30,
        marginEnd: 30,
        marginBottom:30

    },
    entries: {
        marginStart: 6,
        marginEnd: 6,
        backgroundColor: 'rgba(247, 247, 247, 1)',
    },
    textin:{
        margin:16
    }
})


export default FormHead;