import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { ApiService } from './shared';

import '../style/app.scss';

const remote = require('electron').remote;

/*
 * App Component
 * Top Level Component
 */
@Component({
  selector: 'my-app', // <my-app></my-app>
  template: require('./app.component.html'),
  styles: [require('./app.component.scss')],
})
export class AppComponent {
  url = 'https://github.com/preboot/angular2-webpack';
  public img = 'img/tab.svg';

  currentUser;

  constructor(
    private api: ApiService,
    private router: Router
    ) {
      if (localStorage.getItem('isConnected') === '1') {
        this.router.navigate(['/organisation']);
      }
  }

  popcubeNavigate() {
    if (localStorage.getItem('isConnected') === '1') {
        this.router.navigate(['/organisation']);
    }else {
      this.router.navigate(['']);
    }
  }

  minimize() {
    let window = remote.getCurrentWindow();
    window.minimize();
  }

  maximize() {
    let window = remote.getCurrentWindow();
    if (!window.isMaximized()) {
           window.maximize();
           this.img = 'img/multi-tab.svg';
       } else {
           window.unmaximize();
           this.img = 'img/tab.svg';
    }
  }

  close() {
    let window = remote.getCurrentWindow();
    window.close();
  }
}
