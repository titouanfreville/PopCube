import { Component, OnInit } from '@angular/core';
import { HttpModule, JsonpModule } from '@angular/http';
import { Router } from '@angular/router';

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
    private userSvc: UserService,
    private router: Router
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
    let i: number;
        if (localStorage.getItem('organisationSet')) {
          i = parseInt(localStorage.getItem('organisationSet'), 10) + 1;
        }else {
          i = 1;
        }
    this.router.navigate(['/login']);
  }
}
