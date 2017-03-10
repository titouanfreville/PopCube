import { Component, OnInit } from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';

@Component({
  selector: 'my-organisation',
  template: require('./organisation.component.html'),
  styles: [require('./organisation.component.scss')]
})
export class OrganisationComponent implements OnInit {

  public organisation1 = {status};
  public organisation2 = {status};

  public organisations = [];
  public channels = [];
  public channelsText = [];
  public channelsVoice = [];

  constructor() {

    this.organisations.push(new Organisation(1, 'Pop', 'Serveur de développement', 'Pop'));
    this.organisations.push(new Organisation(2, 'Cube', 'Serveur de test', 'Cub'));

  }

  ngOnInit() {
    console.log('Hello Organisation');
  }

  organisationClick(organisationId) {
    for (let o of this.organisations) {
      if (o._idOrganisation === organisationId) {
        o.status = 'organisationFocus';

        if (organisationId === 1) {
          this.channels.push(new Channel(1, 'Developpement', 'Text', 'chanel'));
          this.channels.push(new Channel(2, 'Infrastructure', 'Text', 'chanel'));
          this.channels.push(new Channel(3, 'Marketing', 'Text', 'chanel'));

          this.channels.push(new Channel(4, 'Developpement', 'Voice', 'chanel'));
          this.channels.push(new Channel(5, 'Infrastructure', 'Voice', 'chanel'));
          this.channels.push(new Channel(6, 'Everyones', 'Voice', 'chanel'));
        }else {
          this.channels = [];
        }
      }else {
        o.status = '';
      }
      this.channelsText = [];
      this.channelsVoice = [];
      this.sortChannelType();
    }
  }

  channelClick(channelId) {
    for (let c of this.channels) {
      if (c._idChannel === parseInt(channelId, 10)) {
        c.status = 'channelFocus';
      }else {
        c.status = '';
      }
    }
    this.channelsText = [];
    this.channelsVoice = [];
    this.sortChannelType();
  }

  sortChannelType() {
    for (let c of this.channels){
      if (c.type === 'Text') {
        this.channelsText.push(c);
      }
      if (c.type === 'Voice') {
        this.channelsVoice.push(c);
      }
    }
  }
}
