import { Component } from '@angular/core';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent {

  invalidUser: boolean = false;
  invalidPass: boolean = false;

  userError: string = 'Usuário inválido';
  passwordError: string = 'Senha inválida';

  user: string = '';
  password: string = '';


  login() {
    this.invalidUser = this.user === '';
    this.invalidPass = this.password === '';

    if (this.invalidUser || this.invalidPass) return;
  }

}
