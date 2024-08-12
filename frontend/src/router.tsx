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
import ErrorPage from './pages/error/ErrorPage'
import CardPage from './pages/cards/card'

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
                    path="/activities"
                    element={<ListPage theme={props.theme} />}
                />
                <Route
                    path="/cards"
                    element={<CardPage theme={props.theme} />}
                />
                <Route path="*" element={<ErrorPage theme={props.theme} />}/>
            </Routes>
        </BrowserRouter>
    )
}
