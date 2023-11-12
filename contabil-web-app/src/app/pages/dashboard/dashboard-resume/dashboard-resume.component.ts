import {Component, Input, OnInit} from '@angular/core';
import {Totals} from "../../../../domains/Totals";
import {User} from "../../../../domains/User";

@Component({
    selector: 'app-dashboard-resume',
    templateUrl: './dashboard-resume.component.html',
    styleUrls: ['./dashboard-resume.component.scss']
})
export class DashboardResumeComponent{

    @Input() title: string = 'Titulo';
    @Input() month: string = 'Mês';
    @Input() year: string = 'Ano';
    @Input() values: Totals[] = [];
    user: User | null = null;

    constructor() {
    }

    getValues(): any[]{
        var val: any = {
            labels: [],
            datasets: [{
                label: this.title,
                data: []
            }]
        };

        this.values.forEach((value) => {
            val.labels.push(value.month);
            val.datasets[0].data.push(value.value);
        });

        return val;
    }

    protected readonly Number = Number;
}
