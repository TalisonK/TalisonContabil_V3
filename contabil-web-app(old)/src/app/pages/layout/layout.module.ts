import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HeaderComponent } from './header/header.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { FooterComponent } from './footer/footer.component';
import { PrimengModule } from 'src/app/shared/primeng.module';
import { BreadcrumbComponent } from './breadcrumb/breadcrumb.component';



@NgModule({
  declarations: [
    HeaderComponent,
    SidebarComponent,
    FooterComponent,
    BreadcrumbComponent
  ],
  imports: [
    CommonModule,
    PrimengModule
  ],
  exports: [
    HeaderComponent,
    SidebarComponent,
    FooterComponent,
    BreadcrumbComponent
  ]
})
export class LayoutModule { }
