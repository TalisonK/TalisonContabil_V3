import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';

import Home from "../pages/home";
import InsertExpense from '../pages/activity';
import Explorer from '../pages/explorer';
import Chart from '../pages/chart';

const AuthStack = createStackNavigator();


const AppRoutes: React.FC = () => (
    <AuthStack.Navigator>
        <AuthStack.Screen name="Home" component={Home} options={{headerShown:false}}/>
        <AuthStack.Screen name="InsertExpense" component={InsertExpense}/>
        <AuthStack.Screen name="Explorer" component={Explorer}/>
        <AuthStack.Screen name="Chart" component={Chart}/>
    </AuthStack.Navigator>
)

export default AppRoutes;