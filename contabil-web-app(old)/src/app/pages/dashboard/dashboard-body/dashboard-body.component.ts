import {Component, OnInit} from '@angular/core';
import {User} from 'src/domains/User';
import {DashboardServices} from "../../../../services/Dashboard.services";
import {DashboardInterface} from "../../../../domains/dashboard.interface";

@Component({
    selector: 'app-dashboard-body',
    templateUrl: './dashboard-body.component.html',
    styleUrls: ['./dashboard-body.component.scss']
})
export class DashboardBodyComponent implements OnInit {

    resumes: any[] = ["Income", "Expense", "Balance"];
    date: Date = new Date();
    user: User = {} as User;

    content: DashboardInterface = {} as DashboardInterface;

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
        let date = this.date.toLocaleString('default', {month: 'short'});
        let month: string = date.slice(0, 1).toUpperCase() + date.slice(1, 3);
        this.user = JSON.parse(localStorage.getItem('user') || '{}');
        if (this.user.id != null) {
            this.dashboardServices.getDashboard(this.user.id, this.date.getFullYear().toString(), month).subscribe({
                next: (data) => {
                    this.content = data;
                    console.log(data)
                }, error: (error) => {
                    console.log(error);
                }
            });
        }
    }

    updateDate(date: Date) {
        this.date = date;
        this.update();
    }
}
