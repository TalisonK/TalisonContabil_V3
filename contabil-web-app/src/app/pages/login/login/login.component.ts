import { Component, EventEmitter, Input, Output } from '@angular/core';
import { User } from 'src/domains/User';

@Component({
  selector: 'login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {

  @Output() userEvent = new EventEmitter<User>();

  login: boolean = true;
  cadastro: boolean = false;
  forgot: boolean = false;

  signUpMenu(){
    this.login = false;
    this.cadastro = true;
    this.forgot = false;
  }

  forgotMenu() {

    this.login = false;
    this.cadastro = false;
    this.forgot = true;
  }

  loginMenu() {
    this.login = true;
    this.cadastro = false;
    this.forgot = false;
  }

  loginHandler(user: User) {
    this.userEvent.emit(user);
  }

}
