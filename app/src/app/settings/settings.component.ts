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


@Component({
  selector: 'my-settings',
  template: require('./settings.component.html'),
  styles: [require('./settings.component.scss')],
  providers: [OrganisationService, ChannelService, MessageService, UserService, LocalOrganisationService]
})
export class SettingsComponent {

  private nav;
  currentUser;

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService
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
