import styled from "styled-components";


interface IBackgroundImageProps {
    imageURL: string;
  }
  
  export const LoginWrapper = styled.div<IBackgroundImageProps>`
    display: flex;
    flex-direction: row;
    background-image: url(${props => props.imageURL});
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    height: 100vh;
    align-items: center;
  `;

export const LoginContainer = styled.div`
    width: 40%;
    height: 95vh;
    margin-left: 30px;
    border-radius: 30px;
    background-color: #ffffffcc;
    display: flex;
    justify-content: center;
`;

export const LoginForm = styled.div`

    display: flex;
    flex-direction: column;
    align-self: center;
    width: 600px;
    height: 1000px;
    align-self: center;
    align-items: center;

`;

export const FormContainer = styled.div`

  display: flex;
  flex-direction: column;
  width: 80%;
  margin-bottom: 100px;

`;

export const LoginFormLogo = styled.img`
  height: 100px;
  width: 100px;
  margin-bottom: 70px;
  margin-top: 90px;

`;

export const LoginLogoContainer = styled.div `
    width: 60%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
`;

export const LoginLogo = styled.img`
    height: 10vw;
    align-self: center;
`;
