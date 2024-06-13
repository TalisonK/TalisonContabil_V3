import React from 'react';

import {useAuth} from '../contexts/auth';

import AuthRoutes from './auth.routes';
import AppRoutes from './app.routes';

const Router:React.FC = () => {

    const {signed} = useAuth();

    return signed ? <AppRoutes/> : <AuthRoutes/>
}

export default Router;