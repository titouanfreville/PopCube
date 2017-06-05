import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';
import { User } from '../model/user';

import { Injectable } from '@angular/core';

@Injectable()
export class UserService {
    private userKey = 'currentUser';

    constructor(private http: Http) { }

    getUsers(user) {
        let headers = new Headers({
            'Authorization': 'bearer ' + user,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${'https://' + localStorage.getItem('Stack') + '/user'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    newUser(user: User) {
        let headers = new Headers({
            'Content-Type': 'application/json'
        });
        let formatUser = {
            avatar: user.avatar,
            deleted: false,
            email: user.email,
            first_name: user.firstName,
            id_role: 1,
            last_name: user.lastName,
            nickname: user.nickName,
            password: user.password,
            username: user.userName,
        };
        console.log(JSON.stringify(formatUser));
        return this.http
            .post(`${'https://' + localStorage.getItem('Stack') + '/publicuser/new'}`, JSON.stringify(formatUser), {headers: headers})
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

    public formatUser(user): User {
        return new User(user.id, user.username, user.email, user.password, user.last_activity_at, user.last_password_update, 'fr', user.id_role, user.first_name, user.last_name, user.nickname, user.avartar);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
