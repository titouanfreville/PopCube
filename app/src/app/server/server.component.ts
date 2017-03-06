import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'my-server',
  template: require('./server.component.html'),
  styles: [require('./server.component.scss')]
})
export class ServerComponent implements OnInit {

  constructor() {
    // Do stuff
  }

  ngOnInit() {
    console.log('Hello Server');
  }

}
