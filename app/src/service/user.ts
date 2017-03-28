import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class UserService {

    private usersUrl = 'https://api-alpha.popcube.xyz';  // URL to web api
    private userKey = 'currentUser';

    constructor(private http: Http) { }

    getUsers(user) {
        let headers = new Headers({
            'Authorization': 'bearer ' + user,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${this.usersUrl + '/user'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    // Local Storage User
    private store(content:Object) {
        localStorage.setItem(this.userKey, JSON.stringify(content));
    }

    private retrieve() {
        let storedUser:string = localStorage.getItem(this.userKey);
        if(!storedUser) throw 'no user found';
        return storedUser;
    }

    public generateNewUser(idUser: number) {
        let currentTime:number = (new Date()).getTime() + 60*60;
        this.store({ttl: currentTime, idUser});
    }


    public retrieveUser() {
        let currentTime:number = (new Date()).getTime(), user = null;
        try {
            let storedUser = JSON.parse(this.retrieve());
            if(storedUser.ttl < currentTime) throw 'invalid user found';
            user = storedUser;
        }
        catch(err) {
            console.error(err);
        }
        return user;
    }


    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
