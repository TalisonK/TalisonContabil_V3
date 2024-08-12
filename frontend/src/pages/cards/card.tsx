import { Button, ButtonGroup } from "@mui/material";
import { DisplayFlex, IconFA } from "../../styles";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { useState } from "react";
import CreditCard from "../../interfaces/CreditCard";



interface CardPageProps {
    theme: string
}


const CardPage = (props: CardPageProps) => {

    const [cards, setCards] = useState<CreditCard[]>([])


    return (
        <DisplayFlex 
            card={true} 
            direction="column"
            width="50%" 
            height="90vh" 
            dark={props.theme === 'dark'} 
            marginTop="20px"
            style={{alignSelf:"center"}} 
            
            
        >
            
            <DisplayFlex 
                card={true} 
                width="100%" 
                height="50px" 
                justifyContent="end" 
                dark={props.theme === 'dark'}
            >

                <ButtonGroup variant="contained"
                        style={{margin: "3px"}}>
                    
                    <Button
                        color="success"
                        
                    >
                        <IconFA icon={faPlus} />
                    </Button>
                </ButtonGroup>

            </DisplayFlex>

            <DisplayFlex 
                backgroundColor="blue" 
                width="98%" 
                height="100%"
                marginTop="10px"
                marginBottom="10px"
                style={{alignSelf:"center"}}
            >

            </DisplayFlex>


        </DisplayFlex>
    );
    }

export default CardPage;