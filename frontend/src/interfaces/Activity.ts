interface Activity {
    id: string
    description: string
    paymentMethod: string
    type: string
    userID: string
    categoryName: string
    value: number
    activityDate: Date | any
    actualParcel: number
    totalParcel: number
}

export default Activity
