import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardBodyComponent } from './dashboard-body/dashboard-body.component';
import { DashboardRoutingModule } from './dashboard-routing.module';



@NgModule({
  declarations: [
    DashboardBodyComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule
  ]
})
export class DashboardModule { }
