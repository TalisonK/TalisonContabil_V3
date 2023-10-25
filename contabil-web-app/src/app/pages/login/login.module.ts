import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PrimengModule } from 'src/app/shared/primeng.module';
import { LoginComponent } from './login/login.component';
import { LoginFormComponent } from './login-form/login-form.component';


@NgModule({
  declarations: [
    LoginComponent,
    LoginFormComponent
  ],
  imports: [
    CommonModule,
    PrimengModule
  ],
  exports:[
    LoginComponent
  ]
})
export class LoginModule { }
