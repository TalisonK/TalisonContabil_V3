import styled from 'styled-components'

interface IBackgroundImageProps {
    imageURL: string
}

export const LoginWrapper = styled.div<IBackgroundImageProps>`
    display: flex;
    flex-direction: row;
    background-image: url(${(props) => props.imageURL});
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;
    align-items: center;
`

export const LoginContainer = styled.div`
    width: 40%;
    height: 95vh;
    margin-left: 30px;
    border-radius: 30px;
    background-color: #ffffffcc;
    display: flex;
    flex-direction: column;
`

export const LoginForm = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    align-self: center;
    width: 80%;
    max-width: 600px;
    min-width: 350px;
    height: 100%;
    align-self: center;
    align-items: center;
`

export const FormContainer = styled.div`
    display: flex;
    flex-direction: column;
    min-height: 150px;
    max-height: 1000px;
    justify-content: space-between;
    width: 100%;
`

export const LoginFormLogo = styled.img`
    height: 100px;
    width: 100px;
`

export const LoginLogoContainer = styled.div`
    width: 60%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
`

export const LoginLogo = styled.img`
    height: 10vw;
    align-self: center;
`
