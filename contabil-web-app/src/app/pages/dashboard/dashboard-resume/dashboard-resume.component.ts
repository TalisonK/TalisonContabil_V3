import {Component, Input, OnInit} from '@angular/core';
import {Totals} from "../../../../domains/Totals.interface";
import {User} from "../../../../domains/User";

@Component({
    selector: 'app-dashboard-resume',
    templateUrl: './dashboard-resume.component.html',
    styleUrls: ['./dashboard-resume.component.scss']
})
export class DashboardResumeComponent implements OnInit{

    @Input() title: string = 'Titulo';
    @Input() month: string = 'Mês';
    @Input() year: string = 'Ano';
    @Input() values: Totals[] = [];
    user: User | null = null;

    value: number = 0;
    fraction: number = 0;

    chartOptions: any = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {

            y: [{
                grid: {
                    display: false,
                },
                border:{
                    display: false,
                }
            }],
            x: [{
              grid: {
                display: false,
              },
                border:{
                    display: false,
                }
            }]
        },
    };

    chartData: any = {
        labels: [],
        datasets: [
            {
                label: 'Sales',
                data: [],
            }
        ]
    }

    ngOnInit() {
        setTimeout(() => {
            this.value = Number.parseInt(this.values[7].value.toFixed(2));
            this.fraction = Number.parseInt(((this.values[6].value * 100) / this.values[7].value).toFixed(2)) - 100;
            },
            1000
        );

    }

    constructor() {
    }

    getValues(){

        this.values.forEach((value) => {
            this.chartData.labels.push(value.month);
            this.chartData.datasets[0].data.push(value.value);
            this.chartData.labels.push(value.month);
        });
    }

    protected readonly Number = Number;
}
