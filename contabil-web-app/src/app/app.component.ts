import { Component, EventEmitter, OnInit, Output } from '@angular/core';

import { PrimeNGConfig } from 'primeng/api';
import { User } from 'src/domains/User';

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

	@Output() userLogado = new EventEmitter();

	constructor(private primengConfig: PrimeNGConfig) { }

	ngOnInit() {
		this.primengConfig.ripple = true;
	}

	title = 'contabil-web-app';

	user: User | null = null;
	value: any;

	login(user: User) {
		this.user = user;
		console.log(this.user);
	}
}
