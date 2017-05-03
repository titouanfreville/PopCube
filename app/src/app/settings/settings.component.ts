import { Component} from '@angular/core';

import { ApiService } from '../shared';

const remote = require('electron').remote;

@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
})
export class SettingsComponent {
url = 'https://github.com/preboot/angular2-webpack';
  public img = 'img/tab.svg';

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
