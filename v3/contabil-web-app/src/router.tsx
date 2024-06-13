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

export const RouterApp = (props: any) => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Dashboard theme={props.theme} />} />
                <Route
                    path="/insert"
                    element={<Insert theme={props.theme} />}
                />
                <Route
                    path="/list"
                    element={<ListPage theme={props.theme} />}
                />
            </Routes>
        </BrowserRouter>
    )
}
