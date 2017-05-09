import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { Router } from '@angular/router';

import { LoginService } from '../../../service/login';
import { UserService } from '../../../service/user';
import { localOrganisationService } from '../../../service/localOrganisationService';
import { TokenManager } from '../../../service/tokenManager';

@Component({
  selector: 'my-login',
  template: require('./login.component.html'),
  styles: [require('./login.component.scss')],
  providers: [LoginService, TokenManager, UserService, localOrganisationService],
})
export class LoginComponent implements OnInit {

  loginVar = {login: 'devowner', password: 'popcube'};

  constructor(
    public http: Http,
    private _loginService: LoginService,
    private _token: TokenManager,
    private router: Router,
    private _user: UserService,
    private _localOrganisation: localOrganisationService
  ) {
    // Do things
  }

  ngOnInit() {
    console.log(localStorage);
  }

  login() {
    let request = this._loginService.login(this.loginVar);
    request.then((data) => {
        // this._token.generateNewToken(data.token);
        // this._user.generateNewUser(data.user.id);
        this._localOrganisation.generateNewOrganisation(1, data.user.id, data.token);
        this.router.navigate(['/organisation']);
      }).catch((ex) => {
       console.error('Error login', ex);
      });
  }
}
