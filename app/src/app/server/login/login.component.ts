import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'my-login',
  template: require('./login.component.html'),
  styles: [require('./login.component.scss')],
})
export class LoginComponent implements OnInit {

  constructor(public http: Http) {
    // Do stuff
  }

  ngOnInit() {
    console.log('Hello Home');
  }

  login(event, username, password) {
    event.preventDefault();
    let body = JSON.stringify({username, password});
  }
}
