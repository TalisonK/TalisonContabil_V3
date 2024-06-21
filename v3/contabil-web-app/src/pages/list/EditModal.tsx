import React, { useEffect } from "react";
import Activity from '../../interfaces/Activity'
import { Box, Button, Modal, TextField, Typography } from "@mui/material";
import { DisplayFlex, Text } from "../../styles";
import CalendarInput from "../insert/CalendarInput";
import CategoryInput from "../insert/CategoryInput";
import PaymentMethodInput from "../insert/PaymentMethodInput";
import { categoryList } from "../../api/insert";
import Category from "../../interfaces/Category";
import { updateActivity } from "../../api/List";
import { VariantType, useSnackbar } from "notistack";

interface Props {
    atual : Activity | null
    openEdit: boolean
    setOpenEdit: (value: boolean) => void
}


const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 600,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
  };

const EditModal = (props: Props) => {
    const { enqueueSnackbar } = useSnackbar()


    const [paidAt, setPaidAt] = React.useState<Date>(new Date());

    const [description, setDescription] = React.useState<string>('');
    const [paymentMethod, setPaymentMethod] = React.useState<string>('');
    const [categoryName, setCategoryName] = React.useState<string>('');
    const [value, setValue] = React.useState<number>(0);

    const [categories, setCategories] = React.useState<Category[]>([]);



    useEffect(() => {
        console.log(props.atual);
        if (props.atual) {
            setPaidAt(new Date(props.atual.activityDate));
            setDescription(props.atual.description);
            setPaymentMethod(props.atual.paymentMethod);
            setCategoryName(props.atual.categoryName);
            setValue(props.atual.value);

            categoryList()
            .then((res) => {
                setCategories(res)
            })
        }
    }, [props.atual])

    const submit = () => {

        if (!props.atual || description === "" || value === 0) {
            return;
        }
        
        const body: Activity = {
            id: props.atual.id,
            description: description,
            paymentMethod: paymentMethod,
            type: props.atual.type,
            userID: props.atual.userID,
            categoryName: categoryName,
            value: value,
            activityDate: paidAt,
            actualParcel: props.atual.actualParcel,
            totalParcel: props.atual.totalParcel
        }

        updateActivity(body).then((response) => {
            if (response){
                props.setOpenEdit(false)
            } else {
                handleNotificationVariant("Fail to update activity", "error")
            }
        })
        
    }

    const handleNotificationVariant = (
        messagee: string,
        variant: VariantType
    ) => {
        enqueueSnackbar(messagee, { variant })
    }

    return(
    <Modal
    open={props.openEdit}
    onClose={() => {props.setOpenEdit(false)}}
    aria-labelledby="modal-modal-title"
    aria-describedby="modal-modal-description"
    >
        <Box sx={style}>
            <Text fontSize="1.5rem" textAlign="center" width="100%" marginBottom="40px">Edit Activity</Text>

            <DisplayFlex direction="row" justifyContent="space-evenly" marginTop="15px">
                <TextField required id="outlined-required" label="Description" defaultValue={props.atual?.description} onChange={(e) => setDescription(e.target.value)}/>
                <TextField required id="outlined-required" label="Value" defaultValue={props.atual?.value} onChange={(e) => setValue(Number.parseFloat(e.target.value))} />
            </DisplayFlex>

            
            {props.atual?.type !== "Income" ?
            <DisplayFlex direction="row" justifyContent="space-evenly" marginTop="35px">
                <CategoryInput
                    error={false}
                    category={categoryName}
                    setter={setCategoryName}
                    categories={categories}
                    style={{width:"230px"}}
                    
                />

                <PaymentMethodInput
                    error={false}
                    paymentMethod={paymentMethod}
                    setter={setPaymentMethod}
                    style={{width:"230px"}}
                />
            </DisplayFlex>
            :<></>}
            
            <DisplayFlex direction="row" justifyContent="space-evenly" marginTop="35px">
                <CalendarInput paidAt={paidAt} setter={setPaidAt} />
            </DisplayFlex>
            <DisplayFlex width="100%" justifyContent="center">
                <Button
                    variant="contained"
                    style={{
                        width: '200px',
                        height: '50px',
                        marginTop: '10px',
                        marginBottom: '10px',
                    }}
                    onClick={() => submit()}
                >
                    Update
                </Button>
            </DisplayFlex>
            
        </Box>
    </Modal>)
}

export default EditModal