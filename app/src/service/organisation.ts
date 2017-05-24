import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class OrganisationService {

    private organisationsUrl = 'https://' + localStorage.getItem('Stack');  // URL to web api

    constructor(private http: Http) { }

    getOrganisation(token){
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${this.organisationsUrl + '/organisation'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
