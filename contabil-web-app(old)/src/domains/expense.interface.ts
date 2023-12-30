

export interface Expense{

    id:string;
    description:string;
    paymentMethod:string;
    category:string;
    value:number;
    user:string;
    createdAt:Date;
    paidAt:Date;
    actualParcel:number;
    totalParcel:number;

}
