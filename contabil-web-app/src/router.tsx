import React from 'react'

import {
    BrowserRouter,
    Routes,
    Route,
    createBrowserRouter,
} from 'react-router-dom'
import Login from './pages/login/Login'
import Dashboard from './pages/dashboard/Dashboard'
import Insert from './pages/insert/Insert'

export const RouterApp = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" Component={Dashboard} />
                <Route path="/insert" Component={Insert} />
            </Routes>
        </BrowserRouter>
    )
}
