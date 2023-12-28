import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardBodyComponent } from './dashboard-body/dashboard-body.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { PrimengModule } from 'src/app/shared/primeng.module';
import { DashboardHeaderComponent } from './dashboard-header/dashboard-header.component';
import { DashboardDatePickerComponent } from './dashboard-date-picker/dashboard-date-picker.component';
import { DashboardResumeComponent } from './dashboard-resume/dashboard-resume.component';
import {DashboardServices} from "../../../services/Dashboard.services";
import { DashboardTimelineComponent } from './dashboard-timeline/dashboard-timeline.component';


@NgModule({
  declarations: [
    DashboardBodyComponent,
    DashboardHeaderComponent,
    DashboardDatePickerComponent,
    DashboardResumeComponent,
    DashboardTimelineComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    PrimengModule
  ],
    providers: [
        DashboardServices
    ],
})
export class DashboardModule { }
