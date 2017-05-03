// import {
//   it,
//   describe,
//   async,
//   inject,
//   beforeEachProviders
// } from '@angular/core/testing';

import { TestBed } from '@angular/core/testing';

import { SettingsComponent } from './settings.component';

describe('Settings Component', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({declarations: [SettingsComponent]});
  });

  it('should ...', () => {
    const fixture = TestBed.createComponent(SettingsComponent);
    fixture.detectChanges();
  });
});
