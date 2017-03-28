import { Component, OnInit, ViewChild, ElementRef, AfterViewChecked } from '@angular/core'

import { Organisation } from '../../model/organisation'
import { Channel } from '../../model/channel'
import { Message } from '../../model/message'
import { User } from '../../model/user'

import { OrganisationService } from '../../service/organisation'
import { ChannelService } from '../../service/channel'
import { MessageService } from '../../service/message'
import { UserService } from '../../service/user'
import { TokenManager } from '../../service/tokenManager'

@Component({
  selector: 'my-organisation',
  template: require('./organisation.component.html'),
  styles: [require('./organisation.component.scss')],
  providers: [OrganisationService, TokenManager, ChannelService, MessageService, UserService]
})
export class OrganisationComponent implements OnInit, AfterViewChecked {

  @ViewChild('message') private myScrollContainer: ElementRef;

  organisations: Organisation[] = [];
  channels: Channel[] = [];
  messages: Message[] = [];
  users: User[] = [];

  token: String;
  messageSvc: MessageService;
  currentUser;

  currentOrganisation: number;
  currentChannel: number;
  content: string;

  channelsText: Channel[] = [];
  channelsVoice: Channel[] = [];
  channelsVideo: Channel[] = [];

  channelTitle: string;

  constructor(
    private _organisation: OrganisationService,
    private _token: TokenManager, 
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService
    ) {

    this.token = this._token.retrieveToken();
    this.messageSvc = this._message;
    this.currentUser = this._user.retrieveUser();
    // Organisations
    let requestOrganisation = this._organisation.getOrganisation(this.token);
    requestOrganisation.then((data) => {
        this.organisations.push(new Organisation(data.id, data.name, data.description, data.avatar));

      this.organisations.find(o => o._idOrganisation === data.id).channels = this.channels;
      }).catch((ex) => {
       console.error('Error fetching users', ex);
      });

      // Users list
      let requestUser = this._user.getUsers(this.token);
      requestUser.then((data) => {
          for(let d of data) {
             this.users.push(new User(d.id, d.username, d.password, d.email, d.firstName, d.lastName, d.avatar));
          }
        }).catch((ex) => {
        console.error('Error fetching users', ex);
        });

    // Local test
    //this.initAllLocal();

    this.initStatus();
  }

  ngOnInit() {
    this.currentOrganisation = null;
    this.currentChannel = null;
    this.scrollToBottom();
  }

  ngAfterViewChecked() {
        this.scrollToBottom();
    } 

    scrollToBottom(): void {
        try {
            this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
        } catch(err) { }
    }

  organisationClick(organisationId) {
    this.channels = [];
    this.channelsText = [];
    this.channelsVoice = [];
    this.channelsVideo = [];
    for (let o of this.organisations) {
      if (o._idOrganisation === organisationId) {
        o.status = 'organisationFocus';
        // Channels
        let requestChannel = this._channel.getChannel(this.token);
        requestChannel.then((data) => {
          for(let d of data){
            this.channels.push(new Channel(d.id, d.name, d.type, d.description));
          }
          this.sortChannelType();
        }).catch((ex) => {
        console.error('Error fetching channels', ex);
      });
        this.currentOrganisation = o._idOrganisation;
      }else {
        o.status = '';
      }
    }
  }

  channelClick(channelId) {
    this.messages = [];
    let user: User = null;
    for (let c of this.channels) {
      if (c._idChannel === channelId) {
        c.status = 'channelFocus';
        this.channelTitle = c.channelName;
        this.currentChannel = c._idChannel;
        // Messages
        let requestMessage = this._message.getMessage(this.token);
        requestMessage.then((data) => {
          for(let d of data){
            if(d.id_channel === channelId) {
              // Find correct user
              for(let u of this.users) {
                if(d.id_user === u._idUser){
                  user = u;
                }
              }
              this.messages.push(new Message(d.id, d.date, d.content, user, channelId));
            }
          }
          this.channels.find(c => c._idChannel === this.currentChannel).messages = this.messages;
          }).catch((ex) => {
          console.error('Error fetching messages', ex);
        });
      }else {
        c.status = '';
      }
    }
    this.channelsText = [];
    this.channelsVoice = [];
    this.channelsVideo = [];
    this.sortChannelType();
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

  addMessage() {
    let user = null;
    if (this.content != null) {
      for(let u of this.users){
         if(this.currentUser.idUser === u._idUser){
              user = u;
              console.log(user)
            }
      }
      let idMessage = this.channels.find(c => c._idChannel === this.currentChannel)
      .messages.length + 1;
      let message = new Message(idMessage, (new Date()).getTime(), this.content, user, this.currentChannel)
      this.channels.find(c => c._idChannel === this.currentChannel).messages.push(message);
      this.messages = this.channels.find(c => c._idChannel === this.currentChannel).messages;
      this.content = '';
      this.messageSvc.addMessage(this.token, message);
      try {
        this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight + 61;
      } catch (err) {
        console.log(err);
      }
     }
  }

  initStatus(){
    for (let o of this.organisations) {
      o.status = '';
      for (let c of o.channels) {
        c.status = '';
      }
    }
  }

  initAllLocal() {
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
    .messages.push(new Message(1, 'Content', 1));

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
  }
}
