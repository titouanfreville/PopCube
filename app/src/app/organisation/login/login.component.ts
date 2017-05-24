import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';
import { Router } from '@angular/router';

import { LoginService } from '../../../service/login';
import { UserService } from '../../../service/user';
import { localOrganisationService } from '../../../service/localOrganisationService';

@Component({
  selector: 'my-login',
  template: require('./login.component.html'),
  styles: [require('./login.component.scss')],
  providers: [LoginService, UserService, localOrganisationService],
})
export class LoginComponent implements OnInit {

  loginVar = {login: 'devowner', password: 'popcube'};

  constructor(
    public http: Http,
    private _loginService: LoginService,
    private router: Router,
    private _user: UserService,
    private _localOrganisation: localOrganisationService
  ) {

  }

  ngOnInit() {

  }

  login() {
    let request = this._loginService.login(this.loginVar);
    request.then((data) => {
        let i: number;
        if (localStorage.getItem('organisationSet')) {
          i = parseInt(localStorage.getItem('organisationSet'), 10) + 1;
        }else {
          i = 1;
        }
        this._localOrganisation.generateNewOrganisation(i, data.user.id, data.token);
        this.router.navigate(['/organisation']);
      }).catch((ex) => {
       console.error('Error login', ex);
      });
  }
}
