import {Component, OnInit} from '@angular/core';
import {User} from 'src/domains/User';
import {DashboardServices} from "../../../../services/Dashboard.services";
import {Totals} from "../../../../domains/Totals.interface";

@Component({
    selector: 'app-dashboard-body',
    templateUrl: './dashboard-body.component.html',
    styleUrls: ['./dashboard-body.component.scss']
})
export class DashboardBodyComponent implements OnInit {

    resumes: any[] = ["IncomeInterface", "Expense", "Balance"];
    date: Date = new Date();
    user: User = {} as User;
    incomeArray: Totals[] = [];
    expenseArray: Totals[] = [];
    balanceArray: Totals[] = [];

    constructor(private dashboardServices: DashboardServices) {
    }

    ngOnInit() {
        // @ts-ignore
        const parse: User = JSON.parse(localStorage.getItem('user'));
        if (parse) {
            this.user = parse;
        }
        this.update();

    }

    update(){
        this.updateIncome();
        this.updateExpense();
    }

    updateDate(date: Date) {
        this.date = date;
        this.update();
    }

    updateIncome() {
        let date = this.date.toLocaleString('default', {month: 'short'});
        let month: string = date.slice(0, 1).toUpperCase() + date.slice(1, 3);
        this.user = JSON.parse(localStorage.getItem('user') || '{}');
        if (this.user.id != null) {
            this.dashboardServices.getIncomes(this.user.id, this.date.getFullYear().toString(), month).subscribe({
                next: (data) => {
                    //this.income = Number.parseInt(data[7].value.toFixed(2));
                    this.incomeArray = data;
                }, error: (error) => {
                    console.log(error);
                }
            });
        }
    }

    updateExpense() {
        let date = this.date.toLocaleString('default', {month: 'short'});
        let month: string = date.slice(0, 1).toUpperCase() + date.slice(1, 3);
        this.user = JSON.parse(localStorage.getItem('user') || '{}');
        if (this.user.id != null) {
            this.dashboardServices.getExpense(this.user.id, this.date.getFullYear().toString(), month).subscribe({
                next: (data) => {
                    //this.expense = Number.parseInt(data[7].value.toFixed(2));
                    this.expenseArray = data;
                }, error: (error) => {
                    console.log(error);
                }
            });
        }
    }
}
