import { Component } from '@angular/core';

@Component({
  selector: 'app-cadastro-form',
  templateUrl: './cadastro-form.component.html',
  styleUrls: ['./cadastro-form.component.scss']
})
export class CadastroFormComponent {

  invalidUser: boolean = false;
  invalidPass: boolean = false;

  userError: string = 'Usuário inválido';
  passwordError: string = 'Senha inválida';

  user: string = '';
  password: string = '';


  create() {
    this.invalidUser = this.user === '';
    this.invalidPass = this.password === '';

    if (this.invalidUser || this.invalidPass) return;

    console.log("cria");
  }

}
