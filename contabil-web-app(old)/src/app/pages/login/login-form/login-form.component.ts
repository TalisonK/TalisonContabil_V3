import { Component, EventEmitter, Output } from '@angular/core';
import { Router } from '@angular/router';
import { User } from 'src/domains/User';
import { UserServices } from 'src/services/User.services';


@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent {

  @Output() userEvent = new EventEmitter<User>();

  constructor(private router: Router, private userService: UserServices){}

  invalidUser: boolean = false;
  invalidPass: boolean = false;

  userError: string = 'Usuário inválido';
  passwordError: string = 'Senha inválida';

  user: string = '';
  password: string = '';


  async login() {
    this.invalidUser = this.user === '';
    this.invalidPass = this.password === '';

    if (this.invalidUser || this.invalidPass) return;

    this.userService.login(this.user, this.password).subscribe({
      next: (user: User) => {
        localStorage.setItem('user', JSON.stringify(user));
        this.userEvent.emit(user);
        this.router.navigate(['/dashboard']);
      },
      error: () => {
        this.invalidUser = true;
        this.invalidPass = true;
      }
    })
  }

}
