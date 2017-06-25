import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class LoginService {

    constructor(private http: Http) { }

    login(login) {
        let headers = new Headers({
            'Content-Type': 'application/json',
        });
        return this.http
            .post(`${'https://' + localStorage.getItem('Stack') + '/login'}`, JSON.stringify(login), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
