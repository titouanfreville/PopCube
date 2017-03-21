import { Component, OnInit } from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';
import { Message } from '../../model/message';
import { User } from '../../model/user';

@Component({
  selector: 'my-organisation',
  template: require('./organisation.component.html'),
  styles: [require('./organisation.component.scss')]
})
export class OrganisationComponent implements OnInit {

  organisations: Organisation[] = [];
  channels: Channel[] = [];
  messages: Message[] = [];

  user: User = new User(1, 'Davaï', '1', '1', '1', '1', '1');

  currentOrganisation: number;
  currentChannel: number;
  content: string;

  channelsText: Channel[] = [];
  channelsVoice: Channel[] = [];

  channelTitle: string;

  constructor() {
    // Init organisation 1
    this.organisations.push(new Organisation(1, 'Pop', 'Serveur de développement', 'Pop'));
    
    // Init channels of organisation 1
    this.channels.push(new Channel(1, 'Developpement', 'Text', 'chanel'));
    this.channels.push(new Channel(2, 'Infrastructure', 'Text', 'chanel'));
    this.channels.push(new Channel(3, 'Marketing', 'Text', 'chanel'));

    this.channels.push(new Channel(4, 'Developpement', 'Voice', 'chanel'));
    this.channels.push(new Channel(5, 'Infrastructure', 'Voice', 'chanel'));
    this.channels.push(new Channel(6, 'Everyones', 'Voice', 'chanel'));

    this.organisations.find(o => o._idOrganisation === 1)
    .channels = this.channels;

    this.organisations.find(o => o._idOrganisation === 1)
    .channels.find(c => c._idChannel === 1)
    .messages.push(new Message(1, 'Content', this.user));

    this.channels = [];

    // Init organisation 2
    this.organisations.push(new Organisation(2, 'Cube', 'Serveur de test', 'Cub'));

    // Init channels of organisation 2
    this.channels.push(new Channel(1, 'Les sodomites', 'Text', 'chanel'));
    this.channels.push(new Channel(2, 'FAQ', 'Text', 'chanel'));
    this.channels.push(new Channel(3, 'What did you expect', 'Text', 'chanel'));

    this.channels.push(new Channel(4, 'Fuck', 'Voice', 'chanel'));
    this.channels.push(new Channel(5, 'The', 'Voice', 'chanel'));
    this.channels.push(new Channel(6, 'Police mothafoka', 'Voice', 'chanel'));

    this.organisations.find(or => or._idOrganisation === 2)
    .channels = this.channels;

    this.channels = [];
  }

  ngOnInit() {
    this.currentOrganisation = null;
    this.currentChannel = null;
  }

  organisationClick(organisationId) {
    for (let o of this.organisations) {
      if (o._idOrganisation === organisationId) {
        o.status = 'organisationFocus';
        this.channels = o.channels;
        this.currentOrganisation = o._idOrganisation;
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
        this.channelTitle = c.channelName;
        this.currentChannel = c._idChannel;
        this.messages = c.messages;
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

  addMessage() {
    console.log('corp du message:' + this.content);
    let idMessage = this.channels.find(c => c._idChannel === this.currentChannel)
    .messages.length + 2;
    this.channels.find(c => c._idChannel === this.currentChannel)
    .messages.push(new Message(idMessage, this.content, this.user));
    this.messages = this.channels.find(c => c._idChannel === this.currentChannel)
    .messages;
    this.content = '';

  }
}
