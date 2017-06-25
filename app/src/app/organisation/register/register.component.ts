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
  private user: User;
  private passwordConfirm;
  private emailConfirm;
  private errorMsg = null;

  constructor(
    private userSvc: UserService,
    private router: Router
  ) {

  }

  ngOnInit() {
    this.user = new User();
    this.user._idUser = '';
    this.user.webId = null;
    this.user.userName = null;
    this.user.email = null;
    this.user.updatedAt = null;
    this.user.lastPasswordUpdate = null;
    this.user.locale = null;
    this.user.idRole = null;
    this.user.firstName = null;
    this.user.lastName = null;
    this.user.nickName = null;
    this.user.avatar = null;

    console.log(this.user);
  }

  addUser() {
    this.errorMsg = null;
    if (this.user.password !== null) {
      if (this.user.password === this.passwordConfirm) {
        if (this.user.password.length > 7) {
          if (this.user.email !== null) {
            if (this.user.email === this.emailConfirm) {
              this.user.avatar = 'default.svg';
              this.user.idRole = 1;
              this.user.lastPasswordUpdate = new Date().getTime();
              this.user.locale = 'fr';
              this.user.updatedAt = this.user.lastPasswordUpdate;
              try {
                let token = this.userSvc.newUser(this.user);
                let i: number;
                  if (localStorage.getItem('organisationSet')) {
                    i = parseInt(localStorage.getItem('organisationSet'), 10) + 1;
                  }else {
                    i = 1;
                  }
                  this.router.navigate(['/login']);
              } catch (e) {
                this.errorMsg = e;
              }
            }else {
              this.errorMsg = 'email don\'t match';
            }
          }else {
            this.errorMsg = 'email empty';
          }
        } else {
          this.errorMsg = 'password is too short';
        }
      }else {
        this.errorMsg = 'password don\'t match';
      }
    }else {
      this.errorMsg = 'password empty';
    }
  }
}
