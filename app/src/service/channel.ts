import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';

@Injectable()
export class ChannelService {

    constructor(private http: Http) { }

    getChannel(token) {
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${'https://' + localStorage.getItem('Stack') + '/channel'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
