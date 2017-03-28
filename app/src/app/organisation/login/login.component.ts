import { Component, OnInit } from '@angular/core'
import { Http } from '@angular/http'
import { Router } from '@angular/router'

import { LoginService } from '../../../service/login'
import { UserService } from '../../../service/user'
import { TokenManager } from '../../../service/tokenManager'

import { User } from '../../../model/user'

@Component({
  selector: 'my-login',
  template: require('./login.component.html'),
  styles: [require('./login.component.scss')],
  providers: [LoginService, TokenManager, UserService],
})
export class LoginComponent implements OnInit {

  loginVar = {login: 'devowner', password: 'popcube'};

  constructor(
    public http: Http,
    private _loginService: LoginService,
    private _token: TokenManager,
    private router: Router,
    private _user: UserService
  ) {
    // Do things
  }

  ngOnInit() {

  }

  login() {
    let request = this._loginService.login(this.loginVar);
    request.then((data) => {
        this._token.generateNewToken(data.token);
        this._user.generateNewUser(data.user.id);
        this.router.navigate(['/organisation']);
      }).catch((ex) => {
       console.error('Error fetching users', ex);
      });
  }
}
