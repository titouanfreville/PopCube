import { Component} from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';
import { Message } from '../../model/message';
import { User } from '../../model/user';
import { Role } from '../../model/role';

import { OrganisationService } from '../../service/organisation';
import { ChannelService } from '../../service/channel';
import { MessageService } from '../../service/message';
import { UserService } from '../../service/user';
import { LocalOrganisationService } from '../../service/localOrganisationService';
import { Stack } from '../../service/external/stack';
import { RoleService } from '../../service/role';


@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
  providers: [OrganisationService, ChannelService, MessageService, UserService, LocalOrganisationService, Stack, RoleService]
})
export class SettingsComponent {

  private nav;
  currentUser: User;
  private token;
  private users: User[] = [];
  private currentOrganisation: Organisation;
private roles: Role[] = [];

  private channels: Channel[] = [];
  private channelsText: Channel[] = [];
  private channelsVoice: Channel[] = [];
  private channelsVideo: Channel[] = [];

  private currentChannel: Channel;

  private newChannelT: Channel = new Channel(null, null, null, null);
  private newChannelV: Channel = new Channel(null, null, null, null);
  private newChannelVi: Channel = new Channel(null, null, null, null);

  private hideT: Boolean = true;
  private hideV: Boolean = true;
  private hideVi: Boolean = true;

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService,
    private _stack: Stack,
    private _role: RoleService
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
    this.setRolesList();
    console.log(this.roles);
  }

  profilClick() {
    this.nav = 'profil';
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

    setRolesList() {
      let requestRole = this._role.getAllRole(this.token);
      requestRole.then((data) => {
        for (let d of data) {
          this.roles.push(new Role(d.id, d.name, d.can_use_private, d.can_moderate, d.can_archive, d.can_invite, d.can_manage, d.can_manage_user));
        }
        console.log(this.roles);
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
      if ( c._idChannel === id) {
        this.currentChannel = c;
      }
    }
  }

  // Channel
  newChannel() {
    if (this.newChannelT.channelName !== null &&  this.newChannelT.description !== null) {
      this.newChannelT.type = 'text';
      this._channel.newChannel(this.token, this.newChannelT);
    }else {
      if (this.newChannelV.channelName !== null &&  this.newChannelV.description !== null) {
        this.newChannelV.type = 'voice';
        this._channel.newChannel(this.token, this.newChannelV);
      }else {
        if (this.newChannelVi.channelName !== null &&  this.newChannelVi.description !== null) {
          this.newChannelT.type = 'video';
          this._channel.newChannel(this.token, this.newChannelVi);
        }else {

        }
      }
    }
  }

  hideText() {
    if (this.hideT === false) { this.hideT = true; } else { this.hideT = false; }
  }

  hideVoice() {
    if (this.hideV === false) { this.hideV = true; } else { this.hideV = false; }
  }

  hideVideo() {
    if (this.hideVi === false) { this.hideVi = true; } else { this.hideVi = false; }
  }

  modifyChannel(channel: Channel) {
    this._channel.updateChannel(this.token, channel);
  }

  deleteChannel(channel: Channel) {
    this._channel.deleteChannel(this.token, channel);
  }
}
