import React, { useEffect, useState } from 'react'
import {
    FormContainer,
    LoginContainer,
    LoginForm,
    LoginFormLogo,
    LoginLogo,
    LoginLogoContainer,
    LoginWrapper,
} from './loginStyle'
import {
    Button,
    FormControl,
    IconButton,
    InputAdornment,
    InputLabel,
    OutlinedInput,
    TextField,
} from '@mui/material'
import Visibility from '@mui/icons-material/Visibility'
import VisibilityOff from '@mui/icons-material/VisibilityOff'
import { DisplayFlex, IconFA, Text } from '../../styles'
import { faArrowRightToBracket } from '@fortawesome/free-solid-svg-icons'
import { login, signUp } from '../../api/LoginService'
import { VariantType, useSnackbar } from 'notistack'
import LinearProgress from '@mui/joy/LinearProgress'
import Stack from '@mui/joy/Stack'
import Typography from '@mui/joy/Typography'
import { Key } from '@mui/icons-material'
import Input from '@mui/joy/Input'

interface LoginProps {
    update: () => void
}

const Login = (props: LoginProps) => {
    const { enqueueSnackbar } = useSnackbar()

    const [showPassword, setShowPassword] = useState(false)
    const [singUp, setSingUp] = useState(false)
    const [passwordStrength, setPasswordStrength] = useState(0)
    const [passwordStrengthLabel, setPasswordStrengthLabel] =
        useState('Very weak')

    const [name, setName] = useState('')
    const [password, setPassword] = useState('')

    useEffect(() => {
        switch (passwordStrength) {
            case 0:
                setPasswordStrengthLabel('Very weak')
                break
            case 1:
                setPasswordStrengthLabel('Very weak')
                break
            case 2:
                setPasswordStrengthLabel('Weak')
                break
            case 3:
                setPasswordStrengthLabel('Weak')
                break
            case 4:
                setPasswordStrengthLabel('Strong')
                break
            case 5:
                setPasswordStrengthLabel('Strong')
                break
            default:
                setPasswordStrengthLabel('Very strong')
                break
        }
    }, [passwordStrength])

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword)
    }

    const handleMouseDownPassword = (
        event: React.MouseEvent<HTMLButtonElement>
    ) => {
        event.preventDefault()
    }

    const handleNotificationVariant = (
        messagee: string,
        variant: VariantType
    ) => {
        enqueueSnackbar(messagee, { variant })
    }

    const submitEvent = (): void => {
        if (name === '' || password === '') {
            handleNotificationVariant('Preencha todos os campos!', 'warning')
            return
        }
        if (singUp) {
            signUp(name, password).then((response: any) => {
                handleNotificationVariant(
                    'successful user registration',
                    'success'
                )
                changeSingUp()
            })
        } else {
            login(name, password)
                .then((response: any) => {
                    localStorage.setItem('user', JSON.stringify(response.data))
                    handleNotificationVariant('Login successfully!', 'success')
                    props.update()
                })
                .catch((error: any) => {
                    handleNotificationVariant('Login Failed!', 'error')
                    console.log(error)
                })
        }
    }

    const passwordStrengthMeter = (value: string) => {
        let strength = 0
        if (value.length > 3) {
            strength += 1
        }
        if (value.length > 10) {
            strength += 1
        }
        if (value.match(/[a-z]/)) {
            strength += 1
        }
        if (value.match(/[A-Z]/)) {
            strength += 1
        }
        if (value.match(/[0-9]/)) {
            strength += 1
        }
        if (value.match(/[$@#&!;]/)) {
            strength += 1
        }
        if (value.length < 3) {
            strength = 0
        }

        setPasswordStrength(strength)
    }

    const changeSingUp = () => {
        setName('')
        setPassword('')
        setPasswordStrength(0)
        setSingUp(!singUp)
    }

    return (
        <LoginWrapper imageURL="fundo.png">
            <LoginContainer>
                <LoginForm>
                    <LoginFormLogo src="logo-mini.png" alt="logo" />
                    <DisplayFlex direction="column" width="100%">
                        <DisplayFlex
                            direction="row"
                            justifyContent="space-between"
                            height="60px"
                            width="100%"
                            marginBottom="30px"
                        >
                            <div style={{ marginTop: '-10px' }}>
                                <Text
                                    fontSize="2.2em"
                                    color="#5564a7"
                                    margin="0"
                                    style={{
                                        padding: '5px',
                                        paddingBottom: '0',
                                    }}
                                    textAlign="center"
                                >
                                    Welcome{singUp ? '' : ' back'}!
                                </Text>
                                <Text
                                    fontSize="0.8em"
                                    color="#636363"
                                    margin="0"
                                    style={{ marginLeft: '10px' }}
                                >
                                    Please fill the fields to sign-in
                                </Text>
                            </div>
                            <Button
                                href="#text-buttons"
                                onClick={() => {
                                    changeSingUp()
                                }}
                            >
                                Sign-up
                            </Button>
                        </DisplayFlex>
                        <FormContainer>
                            <TextField
                                id="standard-basic"
                                label="Name"
                                variant="outlined"
                                value={name}
                                onChange={(e) => {
                                    setName(e.target.value)
                                }}
                            />
                            <FormControl
                                sx={{ width: '100%' }}
                                variant="outlined"
                            >
                                <InputLabel htmlFor="outlined-adornment-password">
                                    Password
                                </InputLabel>
                                <OutlinedInput
                                    id="outlined-adornment-password"
                                    type={showPassword ? 'text' : 'password'}
                                    value={password}
                                    onChange={(e) => {
                                        setPassword(e.target.value)
                                        singUp &&
                                            passwordStrengthMeter(
                                                e.target.value
                                            )
                                    }}
                                    endAdornment={
                                        <InputAdornment position="end">
                                            <IconButton
                                                aria-label="toggle password visibility"
                                                onClick={
                                                    handleClickShowPassword
                                                }
                                                onMouseDown={
                                                    handleMouseDownPassword
                                                }
                                                edge="end"
                                            >
                                                {showPassword ? (
                                                    <VisibilityOff />
                                                ) : (
                                                    <Visibility />
                                                )}
                                            </IconButton>
                                        </InputAdornment>
                                    }
                                    label="Password"
                                />
                            </FormControl>
                            {singUp ? (
                                <Stack
                                    sx={{
                                        '--hue': Math.min(
                                            passwordStrength * 16.666,
                                            120
                                        ),
                                    }}
                                >
                                    <LinearProgress
                                        determinate
                                        size="md"
                                        value={passwordStrength * 16.666}
                                        sx={{
                                            bgcolor: 'background.level3',
                                            color: 'hsl(var(--hue) 80% 40%)',
                                        }}
                                    />
                                    <Typography
                                        level="body-xs"
                                        sx={{
                                            alignSelf: 'flex-end',
                                            color: 'hsl(var(--hue) 80% 30%)',
                                        }}
                                    >
                                        {passwordStrengthLabel}
                                    </Typography>
                                </Stack>
                            ) : (
                                <></>
                            )}
                        </FormContainer>
                    </DisplayFlex>
                    <Button
                        variant="outlined"
                        endIcon={
                            <IconFA
                                fontSize="1.4em"
                                color="#5564a7"
                                icon={faArrowRightToBracket}
                                style={{ paddingBottom: '2px' }}
                            />
                        }
                        onClick={submitEvent}
                    >
                        Submit
                    </Button>
                </LoginForm>
            </LoginContainer>
            <LoginLogoContainer>
                <LoginLogo src="logo-invertido.png" alt="logo" />
            </LoginLogoContainer>
        </LoginWrapper>
    )
}

export default Login
