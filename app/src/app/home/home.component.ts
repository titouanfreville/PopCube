import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'my-home',
  template: require('./home.component.html'),
  styles: [require('./home.component.scss')],
})
export class HomeComponent implements OnInit {

  constructor(public http: Http) {
    // Do stuff
  }

  ngOnInit() {
    console.log('Hello Home');
  }

  find(event, domainName) {
    
  }
}
