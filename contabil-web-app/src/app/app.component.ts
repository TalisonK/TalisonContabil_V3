import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';

import { PrimeNGConfig } from 'primeng/api';
import { User } from 'src/domains/User';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

	@Output() userLogado = new EventEmitter();

	constructor(private primengConfig: PrimeNGConfig, private router: Router) { }

	ngOnInit() {
		this.primengConfig.ripple = true;
		if(localStorage.getItem('user') != null){
      this.user = JSON.parse(localStorage.getItem('user') || '{}');
    }
		this.user? this.router.navigate(['/dashboard']) : this.router.navigate(['/']);
	}

	title = 'contabil-web-app';

	user: User | null = null;
	value: any;

	login(user: User) {
		this.user = user;
	}
}
