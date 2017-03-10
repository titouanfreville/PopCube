import { Component, OnInit } from '@angular/core';
import { HttpModule, JsonpModule } from '@angular/http';

import { User } from '../../../model/user';

import { UserService } from '../../../service/user';

@Component({
  selector: 'my-register',
  template: require('./register.component.html'),
  styles: [require('./register.component.scss')]
})
export class RegisterComponent implements OnInit {
  errorMessage: string;
  user: User;

  constructor(private userService: UserService) {
    // Do stuff
  }

  ngOnInit() {
    console.log('Hello Register');
  }

  addUser() {
    this.userService.addUser(1, this.user)
    .then()
    .catch();
  }
}
