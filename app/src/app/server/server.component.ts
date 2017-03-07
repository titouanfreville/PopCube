import { Component, OnInit } from '@angular/core';

import { Server } from '../../model/server';
import { ChanelText } from '../../model/chanelText';
import { ChanelVoice } from '../../model/chanelVoice';

@Component({
  selector: 'my-server',
  template: require('./server.component.html'),
  styles: [require('./server.component.scss')]
})
export class ServerComponent implements OnInit {

  public server1 = {status};
  public server2 = {status};

  public servers = [];
  public chanelsText = [];
  public chanelsVoice = [];

  constructor() {

    this.servers.push(new Server(1, 'Pop', 'P0P', 'server'));
    this.servers.push(new Server(2, 'Cube', 'CuB', 'server'));

  }

  ngOnInit() {
    console.log('Hello Server');
  }

  serverClick(serverId) {
    for (let s of this.servers) {
      if (s.id === serverId) {
        s.status = 'server serverFocus';

        if (serverId === 1) {
          this.chanelsText.push(new ChanelText(1, 1, 'Developpement', 'chanel'));
          this.chanelsText.push(new ChanelText(2, 1, 'Infrastructure', 'chanel'));
          this.chanelsText.push(new ChanelText(3, 1, 'Marketing', 'chanel'));

          this.chanelsVoice.push(new ChanelVoice(1, 1, 'Developpement', 'chanel'));
          this.chanelsVoice.push(new ChanelVoice(2, 1, 'Infrastructure', 'chanel'));
          this.chanelsVoice.push(new ChanelVoice(3, 1, 'Everyones', 'chanel'));
        }else {
        this.chanelsText = [];
          this.chanelsVoice = [];
        }
      }else {
        s.status = 'server';
      }
    }
  }

  chanelTextClick(chanelTextId) {
    for (let ct of this.chanelsText) {
      if (ct.id === parseInt(chanelTextId, 10)) {
        ct.status = 'chanel chanelFocus';
      }else {
        ct.status = 'chanel';
      }
    }
    for (let cv of this.chanelsVoice) {
      cv.status = 'chanel';
    }
  }

  chanelVoiceClick(chanelVoiceId) {
    for (let cv of this.chanelsVoice) {
      if (cv.id === parseInt(chanelVoiceId, 10)) {
        cv.status = 'chanel chanelFocus';
      }else {
        cv.status = 'chanel';
      }
    }
    for (let ct of this.chanelsText) {
      ct.status = 'chanel';
    }
  }
}
