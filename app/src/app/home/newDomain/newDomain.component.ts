import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'my-new-domain',
  template: require('./newDomain.component.spec'),
  styles: [require('./NewDomain.component.scss')],
})
export class NewDomainComponent implements OnInit {

  constructor(public http: Http) {
    // Do stuff
  }

  ngOnInit() {
    console.log('Hello NewDomain');
  }

  find(event, domainName) {

  }
}
