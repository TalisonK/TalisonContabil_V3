import {Component, OnInit} from '@angular/core';
import {User} from 'src/domains/User';
import {DashboardServices} from "../../../../services/Dashboard.services";
import {Totals} from "../../../../domains/Totals.interface";
import {TotalsServices} from "../../../../services/Totals.services";

@Component({
    selector: 'app-dashboard-body',
    templateUrl: './dashboard-body.component.html',
    styleUrls: ['./dashboard-body.component.scss']
})
export class DashboardBodyComponent implements OnInit {

    resumes: any[] = ["Income", "Expense", "Balance"];
    date: Date = new Date();
    user: User = {} as User;
    incomeArray: Totals[] = [];
    expenseArray: Totals[] = [];
    balanceArray: Totals[] = [];
    timeline: any = [];

    constructor(private dashboardServices: DashboardServices, private totalServices: TotalsServices) {
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
        let date = this.date.toLocaleString('default', {month: 'short'});
        let month: string = date.slice(0, 1).toUpperCase() + date.slice(1, 3);
        this.user = JSON.parse(localStorage.getItem('user') || '{}');
        this.updateIncome(month, this.date.getFullYear());
        this.updateExpense(month, this.date.getFullYear());
        this.updateTimeline(month, this.date.getFullYear());
    }

    updateDate(date: Date) {
        this.date = date;
        this.update();
    }

    updateIncome(month: string, year: number) {

        if (this.user.id != null) {
            this.dashboardServices.getIncomes(this.user.id, year.toString(), month).subscribe({
                next: (data) => {
                    //this.income = Number.parseInt(data[7].value.toFixed(2));
                    this.incomeArray = data;
                }, error: (error) => {
                    console.log(error);
                }
            });
        }
    }

    updateExpense(month: string, year: number) {
        if (this.user.id != null) {
            this.dashboardServices.getExpense(this.user.id, year.toString(), month).subscribe({
                next: (data) => {
                    //this.expense = Number.parseInt(data[7].value.toFixed(2));
                    this.expenseArray = data;
                }, error: (error) => {
                    console.log(error);
                }
            });
        }
    }

    updateTimeline(month: string, year: number) {
        if(this.user.id != null) {
            this.totalServices.getTimeline(this.user.id, year.toString(), month).subscribe({
                next: (data: any[]) => {
                    this.timeline = data
                    console.log(data)
                    console.log(year)
                    console.log(month)
                },
                error: (error) => {
                    console.log(error);
                }
            })
        }
    }
}
