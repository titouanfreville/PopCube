// import {
//   it,
//   describe,
//   async,
//   inject,
//   beforeEachProviders
// } from '@angular/core/testing';

import { TestBed } from '@angular/core/testing';

import { ServerComponent } from './server.component';

describe('Server Component', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({declarations: [ServerComponent]});
  });

  it('should ...', () => {
    const fixture = TestBed.createComponent(ServerComponent);
    fixture.detectChanges();
  });
});
