import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { Stack } from '../../service/external/stack';
import { LocalOrganisationService } from '../../service/localOrganisationService';

@Component({
  selector: 'my-home',
  template: require('./home.component.html'),
  styles: [require('./home.component.scss')],
  providers: [Stack, FormsModule, LocalOrganisationService]
})
export class HomeComponent implements OnInit {

  organisation;
  errorMsg = null;

  constructor(
    private _stack: Stack,
    private _router: Router,
    private _localOrg: LocalOrganisationService
  ) {

  }

  ngOnInit() {
  }

  findOrganisation() {
    if (this.organisation != null) {
      let findOrganisation = this.organisation + '.popcube.xyz';
      let isAlreadySet = false;
      try {
         let localOrganisations = this._localOrg.retrieveAllOrganisation();
         for (let l of localOrganisations) {
         if (l.stack === findOrganisation) {
            isAlreadySet = true;
          }
        }
      }catch (e) {
        console.log(e);
      }
      let externalOrganisations = this._stack.getOrg();
      if (!isAlreadySet) {
        for (let o of externalOrganisations) {
          if (o ===  findOrganisation) {
            this._stack.setStack(findOrganisation);
            this._router.navigate(['/login']);
          }else {
            this.errorMsg = 'Couldn\'t find the domain';
          }
        }
      }else {
        this.errorMsg = 'Already log in';
      }
    }
  }

  resetLocal() {
    localStorage.clear();
    console.log('Clear');
  }
}
