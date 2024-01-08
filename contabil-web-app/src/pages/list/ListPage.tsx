import React, { useEffect, useState } from 'react'
import { DisplayFlex } from '../../styles'
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid'
import { getList } from '../../api/List'
import Activity from '../../interfaces/Activity'

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

const ListPage = () => {
    const [rows, setRows] = useState<Activity[]>([])

    useEffect(() => {
        const user = JSON.parse(localStorage.getItem('user') || '{}')
        if (!user.id) {
            window.location.href = '/'
        }

        getList(user.id).then((response) => {
            setRows(response)
        })
    }, [])

    return (
        <DisplayFlex
            direction="column"
            width="100%"
            height="90vh"
            overflow="auto"
            style={{ alignItems: 'center' }}
        >
            <DisplayFlex direction="column" width="80%" marginTop="30px">
                <DataGrid
                    rows={rows}
                    columns={columns}
                    initialState={{
                        pagination: {
                            paginationModel: { page: 0, pageSize: 10 },
                        },
                    }}
                    pageSizeOptions={[10, 20, 30, 50, 100]}
                    checkboxSelection
                />
            </DisplayFlex>
        </DisplayFlex>
    )
}

export default ListPage
