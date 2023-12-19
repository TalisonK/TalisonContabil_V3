import {Component, Input, OnInit} from '@angular/core';
import {TotalsServices} from "../../../../services/Totals.services";

@Component({
  selector: 'app-dashboard-timeline',
  templateUrl: './dashboard-timeline.component.html',
  styleUrls: ['./dashboard-timeline.component.scss']
})
export class DashboardTimelineComponent{

    @Input() activities: any[] = [];

    constructor() {}

}
