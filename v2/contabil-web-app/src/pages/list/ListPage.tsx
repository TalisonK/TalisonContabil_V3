import React, { useEffect, useState } from 'react'
import { DisplayFlex } from '../../styles'
import { DataGrid, GridColDef } from '@mui/x-data-grid'
import { deleteActivity, getList } from '../../api/List'
import Activity from '../../interfaces/Activity'
import { Button, ButtonGroup } from '@mui/material'

const columns: GridColDef[] = [
    { field: 'id', headerName: 'ID', width: 220 },
    { field: 'description', headerName: 'Description', width: 350 },
    { field: 'type', headerName: 'Type', width: 180 },
    { field: 'method', headerName: 'Payment Method', width: 180 },
    { field: 'activityCategory', headerName: 'Category', width: 200 },
    { field: 'value', headerName: 'Value', width: 150, type: 'number' },
    { field: 'actualParcel', headerName: 'Actual Parcel', width: 100 },
    { field: 'totalParcel', headerName: 'Total Parcel', width: 100 },
    { field: 'activityDate', headerName: 'Activity Date', width: 250 },
    { field: 'createdAt', headerName: 'Created At', width: 250 },
]

const ListPage = (props: any) => {
    const [rows, setRows] = useState<Activity[]>([])

    const [selectionModel, setSelectionModel] = useState<any[]>([])

    useEffect(() => {
        const user = JSON.parse(localStorage.getItem('user') || '{}')
        if (!user.id) {
            window.location.href = '/'
        }

        getList(user.id).then((response) => {
            setRows(response)
        })
    }, [])

    const deleteSelected = () => {
        const bucket = rows.filter((row) => selectionModel.includes(row.id))

        deleteActivity(bucket).then((res) => {
            if (res) {
                const user = JSON.parse(localStorage.getItem('user') || '{}')
                getList(user.id).then((response) => {
                    setRows(response)
                })
            }
        })
    }

    return (
        <DisplayFlex
            direction="column"
            width="100%"
            height="90vh"
            overflow="auto"
            style={{ alignItems: 'center' }}
            className={props.theme === 'dark' ? 'theme-dinamic' : ''}
        >
            <DisplayFlex direction="column" width="80%" marginTop="30px">
                <DisplayFlex
                    direction="row"
                    width="100%"
                    justifyContent="end"
                    marginBottom="10px"
                >
                    <ButtonGroup variant="contained">
                        <Button
                            color="warning"
                            disabled={!(selectionModel.length === 1)}
                            style={{
                                color: `${
                                    props.theme === 'dark' &&
                                    selectionModel.length !== 1
                                        ? '#ffffff30'
                                        : 'black'
                                }`,
                            }}
                        >
                            Edit
                        </Button>
                        <Button
                            color="error"
                            onClick={deleteSelected}
                            disabled={!(selectionModel.length >= 1)}
                            style={{
                                color: `${
                                    props.theme === 'dark' &&
                                    selectionModel.length === 0
                                        ? '#ffffff30'
                                        : 'black'
                                }`,
                            }}
                        >
                            Delete
                        </Button>
                    </ButtonGroup>
                </DisplayFlex>
                <DataGrid
                    showCellVerticalBorder={true}
                    rows={rows}
                    columns={columns}
                    style={{
                        color: `${props.theme === 'dark' ? 'white' : 'black'}`,
                    }}
                    initialState={{
                        pagination: {
                            paginationModel: { page: 0, pageSize: 10 },
                        },
                    }}
                    pageSizeOptions={[10, 20, 30, 50, 100]}
                    checkboxSelection
                    onRowSelectionModelChange={(e) => setSelectionModel(e)}
                    rowSelectionModel={selectionModel}
                />
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default ListPage
