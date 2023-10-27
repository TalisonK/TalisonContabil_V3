import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardBodyComponent } from './dashboard-body/dashboard-body.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { PrimengModule } from 'src/app/shared/primeng.module';
import { DashboardHeaderComponent } from './dashboard-header/dashboard-header.component';



@NgModule({
  declarations: [
    DashboardBodyComponent,
    DashboardHeaderComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    PrimengModule
  ]
})
export class DashboardModule { }
