import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardBodyComponent } from './dashboard-body/dashboard-body.component';

const routes: Routes = [
  {path: '', component: DashboardBodyComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
