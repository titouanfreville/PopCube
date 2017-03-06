import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { ServerComponent} from './server/server.component';
import { LoginComponent} from './server/login/login.component';
import { RegisterComponent} from './server/register/register.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'about', component: AboutComponent},
  { path: 'server', component: ServerComponent},
  { path: 'login', component: LoginComponent},
  { path: 'register', component: RegisterComponent}
];

export const routing = RouterModule.forRoot(routes, { useHash: true });

