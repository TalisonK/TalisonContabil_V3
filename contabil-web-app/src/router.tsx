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
import ListPage from './pages/list/ListPage'

export const RouterApp = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" Component={Dashboard} />
                <Route path="/insert" Component={Insert} />
                <Route path="/list" Component={ListPage} />
            </Routes>
        </BrowserRouter>
    )
}
