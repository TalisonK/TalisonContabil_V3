import {Component, EventEmitter, Input, Output} from '@angular/core';
import {User} from 'src/domains/User';

@Component({
    selector: 'app-dashboard-header',
    templateUrl: './dashboard-header.component.html',
    styleUrls: ['./dashboard-header.component.scss']
})
export class DashboardHeaderComponent {

    @Output() dateOut: EventEmitter<Date> = new EventEmitter<Date>();
    @Input() user: User = {} as User;
    @Input() date: Date = new Date();

    constructor() {
    }

    updateDate() {
        this.dateOut.emit(this.date);
    }

}
