/**
 * Created by Lzientek on 28-10-2016
 */

import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { User } from '../model/user';
import { Injectable } from '@angular/core';

@Injectable()
export class UserService {

    private usersUrl = 'alpha-api.popcube.xyz/user';  // URL to web api

    constructor(private http: Http) { }

    addUser(id: number, user: User) {
        let headers = new Headers({
            'Content-Type': 'application/json',
            Authorization: ''
        });
        return this.http
            .post(`${this.usersUrl}/${id}/user`, JSON.stringify(user), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    updateUser(userId: string, user: User) {
        let headers = new Headers({
            'Content-Type': 'application/json',
            Authorization: ''
        });
        return this.http
            .put(`${this.usersUrl}/${userId}/user/${user._idUser}`, JSON.stringify(user), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
