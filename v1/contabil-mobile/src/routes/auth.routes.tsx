import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';

import SignIn from "../pages/signIn";
import Loading from "../pages/loading";

const AuthStack = createStackNavigator();


const AuthRoutes: React.FC = () => (
    <AuthStack.Navigator>
        <AuthStack.Screen name="Loading" component={Loading} options={{headerShown:false}}/>
        <AuthStack.Screen name="SignIn" component={SignIn} options={{headerShown:false, headerBackVisible: false}}/>
    </AuthStack.Navigator>
)

export default AuthRoutes;