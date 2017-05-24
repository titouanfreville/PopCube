import { Component, OnInit } from '@angular/core';
import { HttpModule, JsonpModule } from '@angular/http';

import { User } from '../../../model/user';

import { UserService } from '../../../service/user';

@Component({
  selector: 'my-register',
  template: require('./register.component.html'),
  styles: [require('./register.component.scss')],
  providers: [UserService]
})
export class RegisterComponent implements OnInit {
  errorMessage: string;
  user: User = new User('', null, null, null, null, null, null, null, null, null, null, null);
  constructor(
    private userSvc: UserService
  ) {

  }

  ngOnInit() {

  }

  addUser() {
    this.user.avatar = 'default.svg';
    this.user.idRole = 0;
    this.user.lastPasswordUpdate = new Date().getTime();
    this.user.locale = 'fr';
    this.user.updatedAt = this.user.lastPasswordUpdate;
    let token = this.userSvc.newUser(this.user);
    console.log(token);
  }
}
