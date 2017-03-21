import { NgModule, ApplicationRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { OrganisationComponent} from './organisation/organisation.component';
import { LoginComponent} from './organisation/login/login.component';
import { RegisterComponent} from './organisation/register/register.component';
import { NewDomainComponent } from './home/newDomain/newDomain.component';

import { ApiService } from './shared';
import { routing } from './app.routing';

@NgModule({
  imports: [
    CommonModule,
    BrowserModule,
    FormsModule,
    HttpModule,
    routing
  ],
  declarations: [
    AppComponent,
    HomeComponent,
    AboutComponent,
    OrganisationComponent,
    LoginComponent,
    RegisterComponent,
    NewDomainComponent
  ],
  providers: [
    ApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(public appRef: ApplicationRef) {}
}
