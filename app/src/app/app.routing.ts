import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { OrganisationComponent} from './organisation/organisation.component';
import { LoginComponent} from './organisation/login/login.component';
import { RegisterComponent} from './organisation/register/register.component';
import { NewDomainComponent } from './home/newDomain/newDomain.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'about', component: AboutComponent},
  { path: 'organisation', component: OrganisationComponent},
  { path: 'login', component: LoginComponent},
  { path: 'register', component: RegisterComponent},
  { path: 'newDomain', component: NewDomainComponent}
];

export const routing = RouterModule.forRoot(routes, { useHash: true });

