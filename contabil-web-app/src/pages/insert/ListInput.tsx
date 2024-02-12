import React, { useState } from 'react'
import { DisplayFlex, Text, TextFieldStyled } from '../../styles'
import { DataGrid, GridColDef } from '@mui/x-data-grid'
import { Button } from '@mui/material'
import ValueInput from './ValueInput'

const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 100 },
    { field: 'name', headerName: 'name', width: 350 },
    { field: 'price', headerName: 'Price', width: 350 },
]

interface item {
    id: number
    name: string
    price: string
}

interface ListInputProps {
    rows: item[]
    setRows: (value: item[]) => void
}

const ListInput = (props: ListInputProps) => {
    const [id, setId] = useState<number>(1)
    const [name, setName] = useState<string>('')
    const [price, setPrice] = useState<string>('0,00')

    const [selectionModel, setSelectionModel] = useState<any[]>([])
    const [nameError, setNameError] = useState<boolean>(false)
    const [priceError, setPriceError] = useState<boolean>(false)

    const addItem = () => {
        if (name === '') {
            setNameError(true)
            return
        }
        setNameError(false)

        if (price === '0,00') {
            setPriceError(true)
            return
        }
        setPriceError(false)

        props.setRows([...props.rows, { id: id, name: name, price: price }])
        setId(id + 1)
    }

    const deleteSelected = () => {
        props.setRows(
            props.rows.filter((row) => !selectionModel.includes(row.id))
        )
    }

    return (
        <DisplayFlex width="100%" direction="column">
            <DisplayFlex
                width="100%"
                height="50px"
                justifyContent="center"
                card={true}
            >
                <Text style={{ fontWeight: 'bold' }}>List of Itens</Text>
            </DisplayFlex>
            <DisplayFlex
                width="100%"
                height="70px"
                justifyContent="space-between"
                style={{ alignItems: 'center' }}
                marginTop="5px"
            >
                <Button
                    style={{ marginLeft: '10px' }}
                    color="error"
                    variant="contained"
                    onClick={deleteSelected}
                >
                    Remove Selected
                </Button>
                <DisplayFlex>
                    <TextFieldStyled
                        id="outlined-basic"
                        label="Name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        variant="outlined"
                        error={nameError}
                        className="description-input-space"
                        width="50%"
                    />
                    <DisplayFlex width="50%">
                        <ValueInput
                            error={priceError}
                            setter={setPrice}
                            value={price}
                        />
                    </DisplayFlex>
                </DisplayFlex>
                <Button
                    style={{ marginRight: '10px' }}
                    variant="contained"
                    onClick={addItem}
                >
                    Submit
                </Button>
            </DisplayFlex>
            <DisplayFlex>
                <DataGrid
                    rows={props.rows}
                    columns={columns}
                    style={{ minHeight: '200px' }}
                    initialState={{
                        pagination: {
                            paginationModel: { page: 0, pageSize: 5 },
                        },
                    }}
                    pageSizeOptions={[5]}
                    checkboxSelection
                    onRowSelectionModelChange={(e) => setSelectionModel(e)}
                    rowSelectionModel={selectionModel}
                />
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default ListInput
