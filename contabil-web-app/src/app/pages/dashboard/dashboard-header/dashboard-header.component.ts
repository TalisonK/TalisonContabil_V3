import { Component, Input } from '@angular/core';
import { User } from 'src/domains/User';

@Component({
  selector: 'app-dashboard-header',
  templateUrl: './dashboard-header.component.html',
  styleUrls: ['./dashboard-header.component.scss']
})
export class DashboardHeaderComponent {

  constructor() {}

  @Input() user: User = {} as User;

  @Input() date: Date = new Date();

}
