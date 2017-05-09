import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { User } from '../model/user';
import { Organisation } from '../model/organisation';
import { Injectable } from '@angular/core';

@Injectable()
export class InitOrganisationService {

    private loginUrl = 'alpha-api.popcube.xyz';  // URL to web api

    constructor(private http: Http) { }
    /* Object
    {
    "organisation": {
        "avatar": "string",
        "description": "string",
        "docker_stack": "integer",
        "domain": "string",
        "id": "integer",
        "name": "string",
        "public": "boolean"
    },
    "user": {
        "avatar": "string",
        "deleted": "boolean",
        "email": "string",
        "email_verified": "boolean",
        "failed_attempts": "integer",
        "first_name": "string",
        "id": "integer",
        "id_role": "integer",
        "last_activity_at": "integer",
        "last_name": "string",
        "last_password_update": "integer",
        "last_update": "integer",
        "nickname": "string",
        "password": "string",
        "username": "string",
        "web_id": "string"
        }
    }
    */
    newOrganiation(organisation: Organisation, user: User) {
        let headers = new Headers({
            'Content-Type': 'application/json',
        });
        console.log(this.http
            .post(`${this.loginUrl + '/login'}`, JSON.stringify(organisation) + JSON.stringify(user), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError)
            );
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
