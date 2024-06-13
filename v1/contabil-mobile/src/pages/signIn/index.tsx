import React, { useState } from 'react';
import { View, Image, StyleSheet } from 'react-native';
import {useAuth} from '../../contexts/auth';
import LoginHead from './components/header';
import FormHead from './components/form';
import Butons from './components/button';

const SignIn: React.FC = () => {

    const { signed, signIn } = useAuth();
    const [user, setUser] = useState("")
    const [password, setPassword] = useState("")

    async function handleSignIn() {
        await signIn(user,password);
    }


    return (
        <View style={styles.container}>
            <LoginHead />
            <FormHead user={user} setUser={setUser} password={password} setPassword={setPassword} />
            <Butons handleSignIn={handleSignIn} />
        </View>
    );
}

const styles = StyleSheet.create({
    container :{
        height:'100%'
    },
    form: {
        zIndex: 100
    },
    icon: {
        width: 200,
        height: 200,
        marginStart: "auto",
        marginTop:10
    }
})

export default SignIn;