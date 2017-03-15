import { Component , Input } from '@angular/core';

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
  @Input() channelTitle;

  constructor(private api: ApiService) {
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
