import { Component, OnInit, ViewChild, ElementRef, AfterViewInit, AfterViewChecked } from '@angular/core';

import { Organisation } from '../../model/organisation';
import { Channel } from '../../model/channel';
import { Message } from '../../model/message';
import { User } from '../../model/user';
import { Role } from '../../model/role';

import { Router } from '@angular/router';

import { OrganisationService } from '../../service/organisation';
import { ChannelService } from '../../service/channel';
import { MessageService } from '../../service/message';
import { UserService } from '../../service/user';
import { LocalOrganisationService } from '../../service/localOrganisationService';
import { RoleService } from '../../service/role';

import { Stack } from '../../service/external/stack';

@Component({
  selector: 'my-organisation',
  template: require('./organisation.component.html'),
  styles: [require('./organisation.component.scss')],
  providers: [OrganisationService, ChannelService, MessageService, UserService, LocalOrganisationService, Stack, RoleService]
})
export class OrganisationComponent implements OnInit, AfterViewInit, AfterViewChecked {

  @ViewChild('message') private myScrollContainer: ElementRef;
  @ViewChild('myVideo1') private myVideo: any;
  @ViewChild('myVideo5') private myVideo5: any;

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
  private currentRole: Role;
  content: string;

  channelsText: Channel[] = [];
  channelsVoice: Channel[] = [];
  channelsVideo: Channel[] = [];

  isOrganisationLoad;
  isChannelLoad;
  isMessageLoad;

  storedInformationsTest: any[];

  // PeerJS
  peer;

  constructor(
    private _organisation: OrganisationService,
    private _channel: ChannelService,
    private _message: MessageService,
    private _user: UserService,
    private _localOrganisation: LocalOrganisationService,
    private _stack: Stack,
    private _router: Router,
    private _role: RoleService
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

    console.log(this.organisations);
    this.initStatus();
  }

  ngOnInit() {
    this.isOrganisationLoad = false;

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

  ngAfterViewInit() {

  }

  ngAfterViewChecked() {
        this.scrollToBottom();
    }

    scrollToBottom(): void {
        try {
            this.myScrollContainer.nativeElement.scrollTop = this.myScrollContainer.nativeElement.scrollHeight;
        } catch (err) { }
    }

    setRolesList() {
      let requestRole = this._role.getAllRole(this.token);
      requestRole.then((data) => {
        for (let d of data) {
          if (this.currentUser.idRole === d.id) {
            this.currentRole = d;
          }
        }
      });
    }

  organisationClick(organisationName) {
    // Set to empty all the global for an organisation
    this.channels = [];
    this.channelsText = [];
    this.channelsVoice = [];
    this.channelsVideo = [];
    this.users = [];
    this.currentRole = [];

    // Set the settings to call the API of this organisation
    this.setToken(organisationName);
    this.setStack(organisationName);

    // Get all the users of this organisations and set it to global var
    this.setUserList();

    // Get all the channels for this organisation and set it to global
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
    // Close previous Peer if exist
    if (this.peer) {
      //this.closePeer();
    }

    // Reset messages to set it with the new messages of this channel
    this.messages = [];
    this.isMessageLoad = false;

    let user: User = null;

    // Get all the messages of this channels
    for (let c of this.channels) {
      if (c._idChannel === channelId) {
        c.status = 'channelFocus';
        this.currentChannel = c;
        // Messages
        let requestMessage = this._message.getMessage(this.token);
        requestMessage.then((data) => {
          for (let d of data){
            if (d.id_channel === channelId) {
              // Find correct user of the message
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

    // Automatique call all the members connected to the channels

    // Call audio

    // Call Video
    // if (this.currentChannel.type === 'video') {
    //     this.newPeer();
    //     this.videoConnect();
    // }

    // this.connect();

    // Reload the channels with the messages
    this.channelsText = [];
    this.channelsVoice = [];
    this.channelsVideo = [];
    this.sortChannelType();
  }

  userClick(id) {
      for (let u of this.users) {
        if (id === u._idUser) {
          user = u;
        }
      }
    }

  // Set all the channels by type.
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

  // Send messages to the API
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
             this.users.push(this._user.formatUser(d));
          }
          // CurrentUser
          for (let u of this.users) {
            for (let st of this.storedInformationsTest){
                  if (parseInt(st.userKey, 10) === u._idUser) {
                    this.currentUser = u;
                  }
            }
          }
          this.setRolesList();
          console.log(this.currentUser);
          this.peer = new Peer([this.currentUser.webId], {
            config: {'iceServers': [
              { url: 'stun:stun.l.google.com:19302' },
              { url: 'turn:homeo@turn.bistri.com:80', credential: 'homeo' }
            ]}, key: 'tcgi4gqxdbcsor'});
          setTimeout(() => {
            console.log(this.peer);
        });

        this.peer.on('connection', function(conn) {
          conn.on('data', function(data) {
            console.log(data);
          });
        });
        }).catch((ex) => {
        console.error('Error fetching users', ex);
      });
      console.log(this.users);
  }

  navigateToChannel() {
    localStorage.setItem('settingsNav', 'channel');
    this._router.navigate(['/settings']);
  }

  // newPeer() {
  //       // Peerjs
  //       // this.peer = new Peer([this.currentUser.webId + this.currentChannel._idChannel], {
  //       //     config: {'iceServers': [
  //       //       { url: 'stun:stun.l.google.com:19302' },
  //       //       { url: 'turn:homeo@turn.bistri.com:80', credential: 'homeo' }
  //       //     ]}, key: 'tcgi4gqxdbcsor'});
  //       //   setTimeout(() => {
  //       //     console.log(this.peer);
  //       // });

  //       // this.peer.on('connection', function(conn) {
  //       //   conn.on('data', function(data) {
  //       //     console.log(data);
  //       //   });
  //       // });
  // }

  // closePeer() {
  //   this.peer.destroy();
  // }

  connect() {
    for (let u of this.users) {
        if (u._idUser !== this.currentUser._idUser) {
        let conn = this.peer.connect(u.webId + this.currentChannel._idChannel);
        conn.on('open', function() {
          conn.send('hi');
        });
      }
    }
  }

  videoConnect() {

    console.log(this.users);
    let video = this.myVideo.nativeElement;
    let localvar = this.peer;
    let fname = "3sn77n9uhpbspf4kyj8f9773cc";

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

    // If myVideo1 div exist
    // if (this.myVideo1) {
    //     let video1 = this.myVideo1.nativeElement;
    //     let video2 = this.myVideo2.nativeElement;
    //     let video3 = this.myVideo3.nativeElement;
    //     let video4 = this.myVideo4.nativeElement;
    //     let video5 = this.myVideo5.nativeElement;

    //     let n = <any>navigator;
    //     let localPeer = this.peer;
    //     let localChanId = this.currentChannel._idChannel;
    //     let localCurU = this.currentUser._idUser;
    //     let i: number = 0;

    //     n.getUserMedia = n.getUserMedia || n.webkitGetUserMedia || n.mozGetUserMedia;

    //     for (let u of this.users) {
    //       n.getUserMedia({video: true, audio: true}, function(stream) {
    //         if (u._idUser !== localCurU) {
    //           let call = localPeer.call(u.webId + localChanId, stream);
    //           console.log('Dest id is : ' + u.webId + localChanId);
    //           call.on('stream', function(remotestream) {
    //             i++;
    //             console.log(i);
    //             switch (i) {
    //               case 1: video1.src = URL.createObjectURL(remotestream);
    //                       video1.play();
    //                       console.log('stream');
    //                       console.log('1');
    //                     break;
    //               case 2: video2.src = URL.createObjectURL(remotestream);
    //                       video2.play();
    //                       console.log('stream');
    //                       console.log('2');
    //                     break;
    //               case 3: video3.src = URL.createObjectURL(remotestream);
    //                       video3.play();
    //                       console.log('stream');
    //                       console.log('3');
    //                     break;
    //               case 4: video4.src = URL.createObjectURL(remotestream);
    //                       video4.play();
    //                       console.log('stream');
    //                     break;
    //               default: console.log('bitch');
    //                     break;
    //             }
    //           });
    //         }else{
    //           // ToDo myVideo5 is set to local cam
    //         }
    //     }, function(err) {
    //       console.log(err);
    //     });
    //     }

    // this.peer.on('call', function(call) {
    //       n.getUserMedia({video: true, audio: true}, function(stream){
    //         call.answer(stream);
    //         call.on('stream', function(remotestream) {
    //           console.log(remotestream);
    //           video5.src = URL.createObjectURL(remotestream);
    //           video5.play();
    //         });
    //       }, function(err) {
    //         console.log(err);
    //       });
    //     });
    //   }
  }
}
