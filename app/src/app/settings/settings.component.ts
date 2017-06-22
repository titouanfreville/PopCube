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

  private newChannelT: Channel = new Channel(null, null, null, null);
  private newChannelV: Channel = new Channel(null, null, null, null);
  private newChannelVi: Channel = new Channel(null, null, null, null);

  private hideT: Boolean = true;
  private hideV: Boolean = true;
  private hideVi: Boolean = true;
  private hAvatar: Boolean = true;

  private loadUser = null;
  private loadRole = null;

  private port = "80";

  private currentRole: Role;

  private avatarList: String[] = [
    "boy.svg",
    "boy-1.svg",
    "girl.svg",
    "girl-1.svg",
    "default.svg",
    "man-1.svg",
    "man-2.svg",
    "man-3.svg",
    "man-4.svg"
  ];

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
          this.setChannelsList();
          this.loadUser = 1;
      }).catch((ex) => {
      console.error('Error fetching users', ex);
      });
  }

  setChannelsList() {
    this.channels = [];
    this.channelsText = [];
    this.channelsVideo = [];
    this.channelsVoice = [];

    // Channels
    let requestChannel = this._channel.getChannel(this.token);
      requestChannel.then((data) => {
        for (let d of data) {
          this.channels.push(new Channel(d.id, d.name, d.type, d.description));
        }
        this.sortChannelType();
        this.setRolesList();
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
        for (let role of this.roles) {
          if (this.currentUser.idRole === role.id) {
            this.currentRole = role;
          }
        }
        this.loadRole = 1;
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
    this.setChannelsList();
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

  hideAvatar() {
    if (this.hAvatar === false) { this.hAvatar = true; } else { this.hAvatar = false; }
  }

  modifyChannel(channel: Channel) {
    this._channel.updateChannel(this.token, channel);
  }

  deleteChannel(channel: Channel) {
    this._channel.deleteChannel(this.token, channel);
  }

  setAvatar(avatar) {
    this.currentUser.avatar = avatar;
  }

  updateUser(user: User) {
    this._user.updateUser(this.token, user);
  }

  updateRole(user: User, role) {
    user.idRole = parseInt(role[3], 10);
    this._user.updateUser(this.token, user);
  }
}
