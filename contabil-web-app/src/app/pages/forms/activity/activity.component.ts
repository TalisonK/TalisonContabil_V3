import {Component, OnInit} from '@angular/core';
import {CategoryServices} from "../../../../services/Category.services";
import {MethodsUtil} from "../../../shared/utils/methods.util";
import {IncomeServices} from "../../../../services/Income.services";
import {ExpenseServices} from "../../../../services/Expense.services";

@Component({
    selector: 'app-income', templateUrl: './activity.component.html', styleUrls: ['./activity.component.scss']
})
export class ActivityComponent implements OnInit {

    type: boolean = false;

    description: string = "";
    value: number | null = null;
    method: string = "";
    category: string = "";
    date: Date = new Date();
    actualParcel: number = 1;
    totalParcel: number = 1;

    categories: string[] = []
    catDescription: any = {};

    constructor(private categoryServices: CategoryServices, private methodUtil: MethodsUtil, private incomeServices: IncomeServices, private expenseServices: ExpenseServices) {
    }

    ngOnInit(): void {

        this.categoryServices.getCategories().subscribe({
            next: (categories: any) => {
                categories.forEach((category: any) => {
                    this.categories.push(category.name);
                    this.catDescription[category.name] = category.description;
                })
            }, error: (err: any) => {
                console.log(err);
            }
        })
    }

    getMethods(): string[] {
        return this.methodUtil.getMethods();
    }

    submit(){

        const user = JSON.parse(localStorage.getItem('user') || '{}');

        if(user.name){
            if(this.type){
                console.log("oi")
                this.incomeServices.createIncome({
                    description: this.description,
                    value: this.value,
                    receivedAt: this.date,
                    user:user.name
                }).subscribe({
                    next: (income: any) => {
                        console.log(income);
                    }, error: (err: any) => {
                        console.log(err);
                    }
                })
            }else{
                var aux: any = {
                    description: this.description,
                    value: this.value,
                    paymentMethod: this.methodUtil.translateMethod(this.method),
                    category: this.category,
                    user:user.name,
                    paidAt: this.date
                }
                if(aux.paymentMethod === "CREDIT_CARD"){
                    aux = {...aux, actualParcel: this.actualParcel, totalParcel: this.totalParcel};
                }
                console.log(aux)
                this.expenseServices.createExpense(aux).subscribe({
                    next: (expense: any) => {
                        console.log(expense);
                    }, error: (err: any) => {
                        console.log(err);
                    }
                })
            }
        }


    }


}
