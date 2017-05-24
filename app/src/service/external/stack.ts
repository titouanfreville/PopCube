import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class Stack {

    private externalUrl = '';  // Waiting for external API
    private auth = ''; // Waiting for auth

    constructor(private http: Http) { }

    getOrganisation(organisation) {
        let headers = new Headers({
            'Authorization': 'bearer ' + this.auth,
            'Content-Type': 'application/json'
        });
        this.http
            .get(`${this.externalUrl + '/'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }

    getOrg() {
        return ['maxime', 'society'];
    }
}
