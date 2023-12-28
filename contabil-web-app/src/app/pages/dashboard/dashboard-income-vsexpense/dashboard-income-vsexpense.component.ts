import {Component, Input, OnInit} from '@angular/core';
import {Totals} from "../../../../domains/Totals.interface";

@Component({
  selector: 'app-dashboard-income-vsexpense',
  templateUrl: './dashboard-income-vsexpense.component.html',
  styleUrls: ['./dashboard-income-vsexpense.component.scss']
})
export class DashboardIncomeVSexpenseComponent implements OnInit{

    @Input() incomes: Totals[] = [];
    @Input() expenses: Totals[] = [];

    data: any;

    options: any;

    constructor() {
    }

    ngOnInit() {
        const documentStyle = getComputedStyle(document.documentElement);
        const textColor = documentStyle.getPropertyValue('--text-color');
        const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
        const surfaceBorder = documentStyle.getPropertyValue('--surface-border');

        this.data = {
            labels: this.expenses.map(expense => expense.month),
            datasets: [
                {
                    label: 'Incomes',
                    fill: false,
                    borderColor: documentStyle.getPropertyValue('--green-500'),
                    yAxisID: 'y',
                    tension: 0.4,
                    data: this.incomes.map(income => income.value)
                },
                {
                    label: 'Expenses',
                    fill: false,
                    borderColor: documentStyle.getPropertyValue('--red-500'),
                    yAxisID: 'y1',
                    tension: 0.4,
                    data: this.expenses.map(expense => expense.value)
                }
                ]
        };
        console.log(this.data)
        this.options = {
            stacked: false,
            maintainAspectRatio: false,
            aspectRatio: 0.6,
            plugins: {
                legend: {
                    labels: {
                        color: textColor
                    }
                }
            },
            scales: {
                x: {
                    ticks: {
                        color: textColorSecondary
                    },
                    grid: {
                        color: surfaceBorder
                    }
                },
                y: {
                    type: 'linear',
                    display: true,
                    position: 'left',
                    ticks: {
                        color: textColorSecondary
                    },
                    grid: {
                        color: surfaceBorder
                    }
                },
                y1: {
                    type: 'linear',
                    display: true,
                    position: 'right',
                    ticks: {
                        color: textColorSecondary
                    },
                    grid: {
                        drawOnChartArea: false,
                        color: surfaceBorder
                    }
                }
            }
        };
    }
}
