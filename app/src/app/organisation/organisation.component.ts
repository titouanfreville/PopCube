import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, AfterViewChecked } from '@angular/core';

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
  selector: 'my-organisation',
  template: require('./organisation.component.html'),
  styles: [require('./organisation.component.scss')],
  providers: [OrganisationService, ChannelService, MessageService, UserService, LocalOrganisationService, Stack]
})
export class OrganisationComponent implements OnInit, AfterViewInit, AfterViewChecked {

  @ViewChild('message') private myScrollContainer: ElementRef;
  @ViewChild('myVideo') private myVideo: any;

  organisations: Organisation[] = [];
  channels: Channel[] = [];
  messages: Message[] = [];
  users: User[] = [];

  storedInformations;

  token: String;
  messageSvc: MessageService;
  currentUser: User;

  currentOrganisation: Organisation;
  currentChannel: Channel;
  content: string;

  channelsText: Channel[] = [];
  channelsVoice: Channel[] = [];
  channelsVideo: Channel[] = [];

  isOrganisationLoad;
  isChannelLoad;
  isMessageLoad;

  // Peerjs
  peer;
  anotherid;
  mypeerid;

  storedInformationsTest: any[];

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService,
    private _stack: Stack
    ) {

    this.messageSvc = this._message;

    // retrieveAllOrganisation
    this.storedInformationsTest = this._localOrganisation.retrieveAllOrganisation();

    // Organisation

    // Organisations
    this.isOrganisationLoad = false;
    let i = 0;
    for (let org of this.storedInformationsTest) {
      i++;
      let requestOrganisation = this._organisation.getOrganisationWithStack(org.tokenKey, org.stack);
      requestOrganisation.then((data) => {
        this.organisations.push(new Organisation(data.id, data.name, data.description, data.avatar));
        this.organisations.find(o => o._idOrganisation === data.id).channels = this.channels;
        if (i === this.storedInformationsTest.length) {
        this.isOrganisationLoad = true;
       }
       }).catch((ex) => {
          console.error('Error fetching users', ex);
       });
    }

        // Peerjs
        this.peer = new Peer({
            config: {'iceServers': [
              { url: 'stun:stun.l.google.com:19302' },
              { url: 'turn:homeo@turn.bistri.com:80', credential: 'homeo' }
            ]}, key: 'tcgi4gqxdbcsor'});
          setTimeout(() => {
            this.mypeerid = this.peer.id;
            console.log(this.peer);
        });

        this.peer.on('connection', function(conn) {
          conn.on('data', function(data) {
            console.log(data);
          });
        });
    console.log(this.organisations);
    this.initStatus();
  }

  ngOnInit() {
    this.isOrganisationLoad = false;
  }

  connect() {
    let conn = this.peer.connect(this.anotherid);
    conn.on('open', function() {
      conn.send('hi');
    });
  }

  videoConnect() {
    let video = this.myVideo.nativeElement;
    let localvar = this.peer;
    let fname = this.anotherid;

    let n = <any>navigator;

    n.getUserMedia = n.getUserMedia || n.webkitGetUserMedia || n.mozGetUserMedia;

    n.getUserMedia({video: true, audio: true}, function(stream) {
      let call = localvar.call(fname, stream);
      call.on('stream', function(remotestream) {
        video.src = URL.createObjectURL(remotestream);
        video.play();
      });
    }, function(err) {
      console.log(err);
    });
  }

  ngAfterViewInit() {

  }

  ngAfterViewChecked() {
        this.scrollToBottom();

        if (this.myVideo) {
          let video = this.myVideo.nativeElement;
          let n = <any>navigator;

          n.getUserMedia = n.getUserMedia || n.webkitGetUserMedia || n.mozGetUserMedia;

          this.peer.on('call', function(call) {
            n.getUserMedia({video: true, audio: true}, function(stream){
              call.answer(stream);
              call.on('stream', function(remotestream) {
                video.src = URL.createObjectURL(remotestream);
                video.play();
              });
            }, function(err) {
              console.log(err);
            });
          });
        }
    }

    scrollToBottom(): void {
        try {
            this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
        } catch (err) { }
    }

  organisationClick(organisationName) {
    this.channels = [];
    this.channelsText = [];
    this.channelsVoice = [];
    this.channelsVideo = [];
    this.users = [];
    this.setToken(organisationName);
    this.setStack(organisationName);
    this.setUserList();
    for (let o of this.organisations) {
      if (o.organisationName === organisationName) {
        o.status = 'organisationFocus';
        // Channels
        this.isChannelLoad = false;
        let requestChannel = this._channel.getChannel(this.token);
        requestChannel.then((data) => {
          for (let d of data) {
            this.channels.push(new Channel(d.id, d.name, d.type, d.description));
          }
          this.sortChannelType();
          this.isChannelLoad = true;
        }).catch((ex) => {
        console.error('Error fetching channels', ex);
      });
        this.currentOrganisation = o;
      }else {
        o.status = '';
      }
    }
  }

  channelClick(channelId) {
    this.messages = [];
    this.isMessageLoad = false;
    let user: User = null;
    for (let c of this.channels) {
      if (c._idChannel === channelId) {
        c.status = 'channelFocus';
        this.currentChannel = c;
        // Messages
        let requestMessage = this._message.getMessage(this.token);
        requestMessage.then((data) => {
          for (let d of data){
            if (d.id_channel === channelId) {
              // Find correct user
              for (let u of this.users) {
                if (d.id_user === u._idUser) {
                  user = u;
                }
              }
              this.messages.push(new Message(d.id, d.date, d.content, user, channelId));
            }
          }
          this.channels.find(c => c._idChannel === this.currentChannel._idChannel).messages = this.messages;
          this.isMessageLoad = true;
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
      for (let u of this.users) {
         if ( this.currentUser._idUser === u._idUser) {
              user = u;
            }
      }
      let idMessage = this.channels.find(c => c._idChannel === this.currentChannel._idChannel)
      .messages.length + 1;
      let message = new Message(idMessage, (new Date()).getTime(), this.content, user, this.currentChannel._idChannel);
      this.channels.find(c => c._idChannel === this.currentChannel._idChannel).messages.push(message);
      this.messages = this.channels.find(c => c._idChannel === this.currentChannel._idChannel).messages;
      this.content = '';
      this.messageSvc.addMessage(this.token, message);
      try {
        this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight + 61;
      } catch (err) {
        console.log(err);
      }
     }
  }

  initStatus() {
    for (let o of this.organisations) {
      o.status = '';
      for (let c of o.channels) {
        c.status = '';
      }
    }
  }

  setToken(organisationName) {
    for (let o of this.storedInformationsTest) {
      if (organisationName === o.organisationName) {
        this.token = o.tokenKey;
      }
    }
  }

  setStack(organisationName) {
    for (let o of this.storedInformationsTest) {
      if (organisationName === o.organisationName) {
        this._stack.setStack(o.stack);
      }
    }
  }

  setUserList() {
    // Users list
      let requestUser = this._user.getUsers(this.token);
      requestUser.then((data) => {
          for (let d of data) {
             this.users.push(new User(d.id, d.username, d.email, null, d.updateAt, d.lastPasswordUpdate,
             d.locale, d.idRole, d.firstName, d.lastName, d.nickName, d.avatar));
          }
          // CurrentUser
          for (let u of this.users) {
            for (let st of this.storedInformationsTest){
                  if (parseInt(st.userKey, 10) === u._idUser) {
                    this.currentUser = u;
                  }
            }
          }
        }).catch((ex) => {
        console.error('Error fetching users', ex);
        });
  }

  setAllMembers() {

  }

}
