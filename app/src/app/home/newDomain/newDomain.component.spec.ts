// This shows a different way of testing a component, check about for a simpler one
import { Component } from '@angular/core';

import { TestBed } from '@angular/core/testing';

import { NewDomainComponent } from './newDomain.component';

describe('NewDomain Component', () => {
  const html = '<my-new-domain></my-new-domain>';

  beforeEach(() => {
    TestBed.configureTestingModule({declarations: [NewDomainComponent, TestComponent]});
    TestBed.overrideComponent(TestComponent, { set: { template: html }});
  });

  it('should ...', () => {
    const fixture = TestBed.createComponent(TestComponent);
    fixture.detectChanges();
    expect(fixture.nativeElement.children[0].textContent).toContain('NewDomain Works!');
  });

});

@Component({selector: 'my-test', template: ''})
class TestComponent { }
