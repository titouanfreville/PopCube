import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { Stack } from '../../service/external/stack';
import { localOrganisationService } from '../../service/localOrganisationService'

@Component({
  selector: 'my-home',
  template: require('./home.component.html'),
  styles: [require('./home.component.scss')],
  providers: [Stack, FormsModule, localOrganisationService]
})
export class HomeComponent implements OnInit {

  organisation;
  errorMsg = null;

  constructor(
    private _stack: Stack,
    private _router: Router,
    private _localOrg: localOrganisationService
  ) {

  }

  ngOnInit() {
  }

  findOrganisation() {
    if (this.organisation != null) {
      let findOrganisation = this.organisation + '.popcube.xyz';
      let localOrganisations = this._localOrg.retrieveAllOrganisation();
      let externalOrganisations = this._stack.getOrg();
      let isAlreadySet = false;
      for(let l of localOrganisations) {
        if(l.stack === findOrganisation) {
          isAlreadySet = true;
        }
      }
      if(!isAlreadySet){
        for (let o of externalOrganisations) {
          if (o ===  findOrganisation) {
            this._stack.setStack(findOrganisation);
            this._router.navigate(['/login']);
          }else {
            this.errorMsg = 'Couldn\'t find the domain';
          }
        }
      }else{
        this.errorMsg = 'Already log in';
      }
    }
  }

  resetLocal(){
    localStorage.clear();
    console.log('Clear');
  }
}
