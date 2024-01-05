interface Activity {
    id: string
    description: string
    method: string
    type: string
    userId: string
    activityCategory: string
    value: number
    activityDate: Date | any
    actualParcel: number
    totalParcel: number
}

export default Activity
