import {Component, OnInit} from '@angular/core';
import { User } from 'src/domains/User';

@Component({
  selector: 'app-dashboard-body',
  templateUrl: './dashboard-body.component.html',
  styleUrls: ['./dashboard-body.component.scss']
})
export class DashboardBodyComponent implements OnInit{

  date: Date = new Date();

  user: User = new User();

  constructor() {
  }

  ngOnInit() {
    // @ts-ignore
    const parse: User = JSON.parse(localStorage.getItem('user'));
    // @ts-ignore
    console.log("oi" + JSON.parse(localStorage.getItem('user')));
    if(parse){
      this.user = parse;
    }
  }
}
