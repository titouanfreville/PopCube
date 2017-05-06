import { Component} from '@angular/core';

import { ApiService } from '../shared';

const remote = require('electron').remote;

@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
})
export class SettingsComponent {

  private nav;

  constructor(private api: ApiService) {
  }

  profilClick() {
    this.nav = 'profil';
  }

  organisationClick() {
    this.nav = 'organisation';
  }

  channelClick() {
    this.nav = 'channel';
  }

  clientClick() {
    this.nav = 'client';
  }
}
