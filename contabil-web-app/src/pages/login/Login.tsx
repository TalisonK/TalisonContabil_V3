import React, { useState } from "react";
import { FormContainer, LoginContainer, LoginForm, LoginFormLogo, LoginLogo, LoginLogoContainer, LoginWrapper } from "./loginStyle";
import { Button, FormControl, IconButton, InputAdornment, InputLabel, OutlinedInput, TextField } from "@mui/material";
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { DisplayFlex, IconFA, Text } from "../../styles";
import { faArrowRightToBracket } from "@fortawesome/free-solid-svg-icons";
import {login} from "../../api/LoginService";
import { VariantType, useSnackbar } from "notistack";

const Login = () => {

    const {enqueueSnackbar} = useSnackbar();

    const [showPassword, setShowPassword] = useState(false);
    
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");


    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };



    const handleNotification = (message: string) => {
        enqueueSnackbar(message);
    }

    const handleNotificationVariant = (messagee: string, variant: VariantType) => {
        enqueueSnackbar(messagee, { variant });
    };

    const submitEvent = ():void => {
        handleNotification("Teste");
        // login(name, password).then((response: any) => {
            
        //     localStorage.setItem("user", JSON.stringify(response.data))
            
        //     console.log(response);
        // }
        // ).catch((error: any) => {
        //     console.log(error);
        // });
        
    }

    return (
        <LoginWrapper imageURL="fundo.png">
            <LoginContainer>
                <LoginForm>
                    <LoginFormLogo src="logo-mini.png" alt="logo"/>

                    <DisplayFlex 
                        direction="row"
                        justifyContent="space-between" 
                        height="60px" 
                        width="85%"
                        marginBottom="80px"
                        marginTop="30px"
                    >
                        <div style={{marginTop:"-10px"}}>
                            <Text fontSize='2.2em' color="#5564a7" margin="0" style={{padding:'5px', paddingBottom:'0'}} textAlign="center">Welcome back!</Text>
                            <Text fontSize="0.8em" color="#636363" margin="0" style={{marginLeft:"10px"}}>Please fill the fields to sign-in</Text>
                        </div>
                        <Button href="#text-buttons">
                            Sign-up
                            
                        </Button>
                    </DisplayFlex>
                    <FormContainer>
                        <TextField id="standard-basic" label="Name" variant="outlined" value={name} onChange={e => {setName(e.target.value)}} />
                        <FormControl sx={{marginTop:'40px', width: '100%' }} variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">Password</InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-password"
                                type={showPassword ? 'text' : 'password'}
                                value={password}
                                onChange={e => {setPassword(e.target.value)}}
                                endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    edge="end"
                                    >
                                    {showPassword ? <VisibilityOff /> : <Visibility />}
                                    </IconButton>
                                </InputAdornment>
                                }
                                label="Password"
                            />
                        </FormControl>
                    </FormContainer>
                    <Button 
                        variant="outlined" 
                        endIcon={<IconFA fontSize="1.4em" color="#5564a7" icon={faArrowRightToBracket} style={{paddingBottom:'2px'}}/>}
                        onClick={submitEvent}
                        >Submit
                    </Button>
                </LoginForm>
            </LoginContainer>
            <LoginLogoContainer>
                <LoginLogo src="logo-invertido.png" alt="logo"/>
            </LoginLogoContainer>
        </LoginWrapper>
    )

}


export default Login;