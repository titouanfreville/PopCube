import { Component } from '@angular/core';

import { ApiService } from './shared';

import { User } from '../model/user';

import { UserService } from '../service/user';

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
  providers: [UserService]
})
export class AppComponent {
  url = 'https://github.com/preboot/angular2-webpack';
  public img = 'img/tab.svg';

  currentUser;

  constructor(
    private api: ApiService,
    private _user: UserService
    ) {
    this.currentUser = this._user.retrieveUser();
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
