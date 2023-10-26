import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/domains/User';


@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent {

  @Output() userEvent = new EventEmitter<User>();

  constructor(private router: Router){}

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
    
    
    const user = new User();
    user.nome = this.user;
    user.password = this.password;

    this.userEvent.emit(user);

    this.router.navigate(['/dashboard']);
  }

}
