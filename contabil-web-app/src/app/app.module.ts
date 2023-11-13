import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {LoginModule} from './pages/login/login.module';
import {DashboardModule} from './pages/dashboard/dashboard.module';
import {PrimengModule} from './shared/primeng.module';
import {LayoutModule} from './pages/layout/layout.module';
import {ChartModule} from "primeng/chart";
import { ActivityComponent } from './pages/forms/activity/activity.component';
import { FormsModule } from './pages/forms/forms.module';

@NgModule({
    declarations: [AppComponent, ActivityComponent],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        ChartModule,
        LoginModule,
        DashboardModule,
        PrimengModule,
        LayoutModule,
        FormsModule
    ],
    providers: [],
    bootstrap: [AppComponent],

})
export class AppModule {
}
