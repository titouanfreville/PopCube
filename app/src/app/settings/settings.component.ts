import { Component} from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';
import { Message } from '../../model/message';
import { User } from '../../model/user';

import { OrganisationService } from '../../service/organisation';
import { ChannelService } from '../../service/channel';
import { MessageService } from '../../service/message';
import { UserService } from '../../service/user';
import { TokenManager } from '../../service/tokenManager';
import { localOrganisationService } from '../../service/localOrganisationService';


@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
  providers: [OrganisationService, TokenManager, ChannelService, MessageService, UserService, localOrganisationService]
})
export class SettingsComponent {

  private nav;
  currentUser;

  constructor(
    private _organisation: OrganisationService,
    private _token: TokenManager,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: localOrganisationService
  ) {
    this.currentUser = this._localOrganisation.retrieveOrganisation(1).userKey;
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
