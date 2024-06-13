import React, {createContext, useState, useEffect, useContext} from 'react';
import AsyncStorage from '@react-native-community/async-storage';

import { logar } from '../services/login';

interface AuthContextData{
    signed: boolean,
    user: string | null,
    signIn(nome:string, senha:string): Promise<void>,
    signOut():void,
    verifyToken():Promise<void>
}

const AuthContext = createContext<AuthContextData>({} as AuthContextData);

export const AuthProvider:React.FC = ({children}:any) => {
    
    const [user, setUser] = useState<string | null>(null);

    useEffect(() => {
        async function loadStorage() {
            const storageUser = await AsyncStorage.getItem('@Auth:user');
            const storageToken = await AsyncStorage.getItem('@Auth:token');
            const storageExpiresIn = await AsyncStorage.getItem('@Auth:expiresIn');

            if(storageUser && storageToken && storageExpiresIn){                
                if( Number.parseInt(storageExpiresIn) > (Date.now() / 1000)){
                    setUser(JSON.parse(storageUser));
                }
                else{
                    setUser(null);
                }
            }
        }
        loadStorage();
    }, [])

    async function signOut(){
        setUser(null);
        await AsyncStorage.setItem('@Auth:user', JSON.stringify(""));
        await AsyncStorage.setItem('@Auth:token',JSON.stringify(""));
        await AsyncStorage.setItem('@Auth:expiresIn',JSON.stringify(""));
    }

    async function verifyToken() {
        const storageExpiresIn = await AsyncStorage.getItem('@Auth:expiresIn');

        if(Number(storageExpiresIn) <= Date.now() / 1000){
            signOut();
        }
    }

    async function signIn(nome:string, senha:string) {
        const response = await logar(nome, senha);
        const {token, userId, expiresIn} = response;

        setUser(userId);

        await AsyncStorage.setItem('@Auth:user', JSON.stringify(userId));
        await AsyncStorage.setItem('@Auth:token',JSON.stringify(token));
        await AsyncStorage.setItem('@Auth:expiresIn',JSON.stringify(expiresIn));
    }

    return(
    <AuthContext.Provider value={{signed: !!user,user:user, signIn, signOut, verifyToken}}>
        {children}
    </AuthContext.Provider>
    )
}

export function useAuth(){
    const context = useContext(AuthContext);

    return context;
}
