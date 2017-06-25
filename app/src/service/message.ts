import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Message } from '../model/message';
import { Injectable } from '@angular/core';

@Injectable()
export class MessageService {

    constructor(private http: Http) { }

    getMessage(token) {
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .get(`${'https://' + localStorage.getItem('Stack') + '/message'}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    addMessage(token, message: Message) {
        let formatMessage =  {
            content: message.content,
            // date: message.date,
            // id: message._idMessage,
            id_channel: message.idChannel,
            id_user: message.creator._idUser
        };
        console.log(formatMessage);
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .post(`${'https://' + localStorage.getItem('Stack') + '/message'}`, JSON.stringify(formatMessage), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }
    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
