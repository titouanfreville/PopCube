import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Message } from '../model/message';
import { Injectable } from '@angular/core';

@Injectable()
export class MessageService {

    private messageUrl = 'https://api-alpha.popcube.xyz';  // URL to web api

    constructor(private http: Http) { }

    getMessage(token){
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${this.messageUrl + '/message'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    addMessage(token, message: Message) {
        console.log("message : ");
        console.log(message);
        let formatMessage =  {
            content: message.content,
            date: message.date,
            id: message._idMessage,
            id_channel: message.idChannel,
            id_user: message.creator._idUser
        };
        console.log(formatMessage);
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .post(`${this.messageUrl + '/message'}`, JSON.stringify(formatMessage), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }
    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
