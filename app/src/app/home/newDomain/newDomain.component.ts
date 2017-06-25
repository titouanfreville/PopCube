import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'my-new-domain',
  template: require('./newDomain.component.html'),
  styles: [require('./newDomain.component.scss')],
})
export class NewDomainComponent implements OnInit {

  constructor(public http: Http) {
    // Do stuff
  }

  ngOnInit() {

  }

  find(event, domainName) {

  }
}
