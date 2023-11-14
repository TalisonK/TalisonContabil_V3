import {Component, Input, OnInit} from '@angular/core';
import {Activity} from "../../../../domains/timeline.interface";
import {TotalsServices} from "../../../../services/Totals.services";

@Component({
  selector: 'app-dashboard-timeline',
  templateUrl: './dashboard-timeline.component.html',
  styleUrls: ['./dashboard-timeline.component.scss']
})
export class DashboardTimelineComponent implements OnInit{

    @Input() year: string = "";
    @Input() month: string = "";

    activities: any[] = [];

    constructor(private totalsServices: TotalsServices) {}

    ngOnInit() {
        const user = JSON.parse(localStorage.getItem('user') || '{}');

        let month: string = this.month.slice(0, 1).toUpperCase() + this.month.slice(1, 3);

        this.totalsServices.getTimeline(user.id, this.year, month).subscribe({
            next: (activity: any) => {
                this.activities = activity;
            },
            error: (error) => {
                console.log(error);
            }
        })

    }


}
