import React, { useEffect, useState } from 'react'
import { RouterApp } from './router'
import { AppContainer, DisplayFlex } from './styles'
import Login from './pages/login/Login'
import { SnackbarProvider } from 'notistack'
import TopBar from './components/TopBar'
import './style.css'

function App() {
    const [user, setUser] = useState(null)
    const [theme, setTheme] = useState('light')

    useEffect(() => {
        const userStorage = localStorage.getItem('user')
        const themeStorage = localStorage.getItem('theme')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        }
        if (themeStorage) {
            setTheme(themeStorage)
        }
    }, [])

    const updateUser = () => {
        const userStorage = localStorage.getItem('user')
        if (userStorage) {
            setUser(JSON.parse(userStorage))
        }
    }

    const updateTheme = (theme: boolean) => {
        if (theme) {
            setTheme('dark')
            localStorage.setItem('theme', 'dark')
        }
        if (!theme) {
            setTheme('light')
            localStorage.setItem('theme', 'light')
        }
    }

    return (
        <SnackbarProvider maxSnack={3} autoHideDuration={3000}>
            <AppContainer>
                {!localStorage.getItem('user') ? (
                    <Login update={updateUser} />
                ) : (
                    <DisplayFlex
                        direction="column"
                        overflow="hidden"
                        height="100vh"
                        backgroundColor={
                            theme === 'light' ? '#f5f5f5' : '#1b1b1b'
                        }
                    >
                        <TopBar theme={theme} setTheme={updateTheme} />
                        <RouterApp theme={theme} />
                    </DisplayFlex>
                )}
            </AppContainer>
        </SnackbarProvider>
    )
}

export default App
