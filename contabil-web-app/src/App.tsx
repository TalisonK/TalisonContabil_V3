import React, { useEffect, useState } from 'react'
import { RouterApp } from './router'
import { AppContainer, DisplayFlex } from './styles'
import Login from './pages/login/Login'
import { SnackbarProvider } from 'notistack'
import Sidebar from './components/sidebar'
import './style.css'

function App() {
    const [user, setUser] = React.useState(null)

    useEffect(() => {
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        }
    }, [])

    const updateUser = () => {
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        }
    }

    return (
        <SnackbarProvider maxSnack={3} autoHideDuration={3000}>
            <AppContainer>
                {user ? (
                    <DisplayFlex direction="column" overflow="hidden">
                        <Sidebar />
                        <RouterApp />
                    </DisplayFlex>
                ) : (
                    <Login update={updateUser} />
                )}
            </AppContainer>
        </SnackbarProvider>
    )
}

export default App
