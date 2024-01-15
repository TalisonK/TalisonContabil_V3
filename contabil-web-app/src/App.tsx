import React, { useEffect, useState } from 'react'
import { RouterApp } from './router'
import { AppContainer, DisplayFlex } from './styles'
import Login from './pages/login/Login'
import { SnackbarProvider } from 'notistack'
import TopBar from './components/TopBar'
import './style.css'

function App() {
    const [user, setUser] = useState(null)

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
                {!localStorage.getItem('user') ? (
                    <Login update={updateUser} />
                ) : (
                    <DisplayFlex direction="column" overflow="hidden">
                        <TopBar />
                        <RouterApp />
                    </DisplayFlex>
                )}
            </AppContainer>
        </SnackbarProvider>
    )
}

export default App
