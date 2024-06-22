import React from 'react';
import { Typography } from "@mui/material";
import ErrorOutlineIcon from '@mui/icons-material/ErrorOutline';
import { DisplayFlex, Text } from "../../styles";

const ErrorPage = (props: any) => {
    return (
        <DisplayFlex width="100vw" height="90vh" justifyContent="center" style={{alignItems:"center"}} direction="column">
            <ErrorOutlineIcon style={{ fontSize: 80, color: "#d32f2f" }} />
            <Typography variant="h4" style={{ margin: "20px 0", fontWeight: "bold" }}>
                <Text color={props.theme === 'dark' ? 'white' : 'black'}>404 - Página Não Encontrada</Text>
            </Typography>
            <Text color={props.theme === 'dark' ? 'white' : 'black'}>
                Oops! Parece que a página que você está procurando não existe.
            </Text>
        </DisplayFlex>
    );
}

export default ErrorPage;