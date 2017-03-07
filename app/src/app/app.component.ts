import { Component } from '@angular/core';

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

  constructor(private api: ApiService) {
  }

  minimize(){
    var window = remote.getCurrentWindow();
    window.minimize();
  }

  maximize(){
    var window = remote.getCurrentWindow();
    if (!window.isMaximized()) {
           window.maximize();          
       } else {
           window.unmaximize();
    }
  }

  close(){
    var window = remote.getCurrentWindow();
    window.close();
  }
}
