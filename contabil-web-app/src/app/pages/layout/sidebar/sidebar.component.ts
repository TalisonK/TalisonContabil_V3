import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { MenuItem } from 'primeng/api';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.scss']
})
export class SidebarComponent {

  constructor(private router: Router) {}

  move(item: MenuItem) {
    this.router.navigate([item.routerLink]);
  }

  items: MenuItem[] = [
    { "label": "Dashboards", icon: "pi-home", routerLink:'/dashboard' }, 
    { "label": "Segundo" , icon: "pi-user", routerLink:'/dashboard' }, 
    { "label": "Terceiro", icon: "pi-cross", routerLink:'/dashboard' }
  ];
}
