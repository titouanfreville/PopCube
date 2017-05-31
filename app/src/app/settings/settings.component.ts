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

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService,
    private _stack: Stack
  ) {
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
    console.log(this.currentUser);
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

  setUserList() {
    // Users list
      let requestUser = this._user.getUsers(this.token);
      requestUser.then((data) => {
          for (let d of data) {
             this.users.push(new User(d.id, d.username, d.email, null, d.updateAt, d.lastPasswordUpdate,
             d.locale, d.idRole, d.firstName, d.lastName, d.nickName, d.avatar));
          }
          console.log(this.users);
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
}
