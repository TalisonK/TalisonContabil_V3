import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-dashboard-date-picker',
  templateUrl: './dashboard-date-picker.component.html',
  styleUrls: ['./dashboard-date-picker.component.scss']
})
export class DashboardDatePickerComponent implements OnInit{

  @Output() dateOut: EventEmitter<Date> = new EventEmitter<Date>();

  date: Date = new Date();

  dates: Date[] = [];

  constructor() {
    this.dates = this.getDates();
  }

  ngOnInit(): void {
    this.coiso();
  }

  coiso() {
    setTimeout(() => {
      let objDiv: HTMLElement | null = document.getElementById("picker-body");
      let objDiv2: HTMLElement | null = document.getElementById("selected");

      const difYear = this.date.getFullYear() - 2020;
      const difMonth = (difYear * 12) + this.date.getMonth() - 2;

        // @ts-ignore
        objDiv.scrollLeft = difMonth * objDiv2?.clientWidth;
    }, 2700);
  }


  getDates(): Date[] {

    const date: Date = new Date(this.date);
    date.setFullYear(date.getFullYear() + 2);

    const dates: Date[] = [];

    for (let i = 0;; i++) {
      const novo = new Date('2020-01-15');
      novo.setMonth(novo.getMonth() + i);
      if(novo.valueOf() > date.valueOf()){
        break;
      }
      dates.push(novo);
    }

    return dates;

  }


  dateComparator(date1: Date): boolean {
    return date1.getMonth() === this.date.getMonth() && date1.getFullYear() === this.date.getFullYear();
  }

  setDate(item: Date) {
    this.date = item;
    this.dateOut.emit(this.date);
  }
}
