import { Button, ButtonGroup } from "@mui/material";
import { DisplayFlex, IconFA } from "../../styles";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";
import { useEffect, useState } from "react";
import CreditCard from "../../interfaces/CreditCard";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { getCard } from "../../api/CreditCard";



interface CardPageProps {
    theme: string
}

const columns: GridColDef[] = [
    { field: 'flag', headerName: 'Flag', width: 150 },
    { field: 'bank', headerName: 'Bank', width: 250 },
    { field: 'closesAt', headerName: 'Closes At', width: 180 },
    { field: 'expiresAt', headerName: 'Expires At', width: 200 },
    { field: 'createdAt', headerName: 'Created At', width: 250 },
    { field: 'updatedAt', headerName: 'Updated At', width: 250 },
]


const CardPage = (props: CardPageProps) => {

    const [cards, setCards] = useState<CreditCard[]>([])

    useEffect(() => {

        const user = JSON.parse(localStorage.getItem('user') || '{}')
        if (!user.id) {
            window.location.href = '/'
        }

        getCard(user.id).then((response) => {
            if (response !== null) {
                setCards(response.data)
            } else {
                setCards([])
            }
        })
    }, [])

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
                        style={{margin: "3px", marginRight: "10px"}}>
                    
                    <Button
                        color="success"
                        
                    >
                        <IconFA icon={faPlus} />
                    </Button>
                </ButtonGroup>

            </DisplayFlex>

            <DisplayFlex 
                className="theme-dinamic"
                width="98%" 
                height="100%"
                marginTop="10px"
                marginBottom="10px"
                style={{alignSelf:"center"}}
            >
                <DataGrid
                    showCellVerticalBorder={true}
                    rows={cards}
                    columns={columns}
                    style={{
                        color: `${props.theme === 'dark' ? 'white' : 'black'}`,
                        minHeight: '200px',
                    }}
                    initialState={{
                        pagination: {
                            paginationModel: { page: 0, pageSize: 10 },
                        },
                    }}
                    pageSizeOptions={[10, 20, 30, 50, 100]}
                />
            </DisplayFlex>

            
        </DisplayFlex>
    );
    }

export default CardPage;