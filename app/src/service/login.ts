import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class LoginService {

    private loginUrl = 'https://api-alpha.popcube.xyz';  // URL to web api

    constructor(private http: Http) { }

    login(login) {
        let headers = new Headers({
            'Content-Type': 'application/json',
        });
        return this.http
            .post(`${this.loginUrl + '/login'}`, JSON.stringify(login), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
