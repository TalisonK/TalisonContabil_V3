import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PrimengModule } from 'src/app/shared/primeng.module';
import { LoginComponent } from './login/login.component';
import { LoginFormComponent } from './login-form/login-form.component';
import { CadastroFormComponent } from './cadastro-form/cadastro-form.component';
import { ForgotFormComponent } from './forgot-form/forgot-form.component';


@NgModule({
  declarations: [
    LoginComponent,
    LoginFormComponent,
    CadastroFormComponent,
    ForgotFormComponent
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
