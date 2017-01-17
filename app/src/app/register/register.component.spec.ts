// This shows a different way of testing a component, check about for a simpler one
import { Component } from '@angular/core';

import { TestBed } from '@angular/core/testing';

import { RegisterComponent } from './register.component';

describe('Register Component', () => {
  const html = '<my-register></my-register>';

  beforeEach(() => {
    TestBed.configureTestingModule({declarations: [RegisterComponent, TestComponent]});
    TestBed.overrideComponent(TestComponent, { set: { template: html }});
  });

  it('should ...', () => {
    const fixture = TestBed.createComponent(TestComponent);
    fixture.detectChanges();
    expect(fixture.nativeElement.children[0].textContent).toContain('<form>');
  });

});

@Component({selector: 'my-test', template: ''})
class TestComponent { }
