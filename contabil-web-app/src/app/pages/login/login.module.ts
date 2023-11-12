import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PrimengModule } from 'src/app/shared/primeng.module';
import { LoginComponent } from './login/login.component';
import { LoginFormComponent } from './login-form/login-form.component';
import { CadastroFormComponent } from './cadastro-form/cadastro-form.component';
import { ForgotFormComponent } from './forgot-form/forgot-form.component';
import { UserServices } from 'src/services/User.services';


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
  providers: [
    UserServices
  ],
  exports:[
    LoginComponent
  ]
})
export class LoginModule { }
