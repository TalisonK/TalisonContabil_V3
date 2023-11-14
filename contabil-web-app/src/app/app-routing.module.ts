import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ActivityComponent} from "./pages/forms/activity/activity.component";
import {ChartComponent} from "./pages/forms/chart/chart.component";

const routes: Routes = [
    {
        path: 'dashboard',
        loadChildren: () => import('./pages/dashboard/dashboard.module').then(m => m.DashboardModule)
    },
    {
        path: 'forms',
        component: ActivityComponent
    },
    {
        path: 'chart',
        component: ChartComponent
    }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)], exports: [RouterModule]
})
export class AppRoutingModule {
}
