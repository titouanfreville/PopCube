import { Component} from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';
import { Message } from '../../model/message';
import { User } from '../../model/user';

import { OrganisationService } from '../../service/organisation';
import { ChannelService } from '../../service/channel';
import { MessageService } from '../../service/message';
import { UserService } from '../../service/user';
import { LocalOrganisationService } from '../../service/localOrganisationService';
import { Stack } from '../../service/external/stack';


@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
  providers: [OrganisationService, ChannelService, MessageService, UserService, LocalOrganisationService, Stack]
})
export class SettingsComponent {

  private nav;
  currentUser: User;
  private token;
  private users: User[] = [];
  private currentOrganisation;

  private channels: Channel[] = [];
  private channelsText: Channel[] = [];
  private channelsVoice: Channel[] = [];
  private channelsVideo: Channel[] = [];

  private currentChannel: Channel;

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService,
    private _stack: Stack
  ) {

    this.nav = localStorage.getItem('settingsNav');

    let organisations = _localOrganisation.retrieveAllOrganisation();
    for (let o of organisations) {
      if (localStorage.getItem('Stack') === o.stack) {
        this.currentOrganisation = o;
        this.token = o.tokenKey;
      }
    }
    this.setUserList();
    this.setChannelsList();
  }

  profilClick() {
    this.nav = 'profil'
    localStorage.setItem('settingsNav', 'profil');
  }

  organisationClick() {
    localStorage.setItem('settingsNav', 'organisation');
    this.nav = 'organisation';
  }

  channelClick() {
    localStorage.setItem('settingsNav', 'channel');
    this.nav = 'channel';
  }

  clientClick() {
    localStorage.setItem('settingsNav', 'client');
    this.nav = 'client';
  }

  rightClick() {
    localStorage.setItem('settingsNav', 'right');
    this.nav = 'right';
  }

  setUserList() {
    // Users list
      let requestUser = this._user.getUsers(this.token);
      requestUser.then((data) => {
          for (let d of data) {
             this.users.push(this._user.formatUser(d));
          }
          // CurrentUser
          for (let u of this.users) {
                  if (parseInt(this.currentOrganisation.userKey, 10) === u._idUser) {
                    this.currentUser = u;
                  }
          }
        }).catch((ex) => {
        console.error('Error fetching users', ex);
        });
  }

  setChannelsList() {
        // Channels
        let requestChannel = this._channel.getChannel(this.token);
        requestChannel.then((data) => {
          for (let d of data) {
            this.channels.push(new Channel(d.id, d.name, d.type, d.description));
          }
          this.sortChannelType();
        }).catch((ex) => {
        console.error('Error fetching channels', ex);
      });
    }

  sortChannelType() {
    for (let c of this.channels){
      if (c.type === 'text') {
        this.channelsText.push(c);
      }
      if (c.type === 'audio') {
        this.channelsVoice.push(c);
      }
      if (c.type === 'video') {
        this.channelsVideo.push(c);
      }
    }
  }

  oneChannelClick(id ) {
    for (let c of this.channels) {
      if( c._idChannel === id) {
        this.currentChannel = c;
      }
    }
  }
}
