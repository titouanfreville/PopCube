import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { Stack } from '../../service/external/stack';

@Component({
  selector: 'my-home',
  template: require('./home.component.html'),
  styles: [require('./home.component.scss')],
  providers: [Stack, FormsModule]
})
export class HomeComponent implements OnInit {

  organisation;
  errorMsg = null;

  constructor(
    private stack: Stack,
    private router: Router
  ) {

  }

  ngOnInit() {
  }

  findOrganisation() {
    if (this.organisation != null) {
        let organisations = this.stack.getOrg();
        for (let o of organisations) {
          if (o ===  this.organisation) {
            localStorage.setItem('Stack', this.organisation + '.popcube.xyz');
            this.router.navigate(['/login']);
          }else {
            this.errorMsg = 'Couldn\'t find the domain';
          }
        }
    }
  }

  resetLocal(){
    localStorage.clear();
    console.log('ok');
  }
}
