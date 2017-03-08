// import {
//   it,
//   describe,
//   async,
//   inject,
//   beforeEachProviders
// } from '@angular/core/testing';

import { TestBed } from '@angular/core/testing';

import { OrganisationComponent } from './organisation.component';

describe('Organisation Component', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({declarations: [OrganisationComponent]});
  });

  it('should ...', () => {
    const fixture = TestBed.createComponent(OrganisationComponent);
    fixture.detectChanges();
  });
});
