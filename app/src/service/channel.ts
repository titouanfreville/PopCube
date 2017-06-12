import 'rxjs/add/operator/toPromise';

import { Headers, Http } from '@angular/http';

import { Injectable } from '@angular/core';
import { Channel } from '../model/channel';

@Injectable()
export class ChannelService {

    constructor(private http: Http) { }

    getChannel(token){
        console.log(localStorage.getItem('Stack'));
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

    newChannel(token, channel:Channel) {

        let formatChannel = this.formatChannel(channel);
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .post(`${'https://' + localStorage.getItem('Stack') + '/channel/new'}`, JSON.stringify(formatChannel), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    updateChannel(token, channel:Channel) {
        let formatChannel = this.formatChannel(channel);
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .put(`${'https://' + localStorage.getItem('Stack') + '/channel/' + channel._idChannel.toString()}`, JSON.stringify(formatChannel), { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    deleteChannel(token, channel:Channel) {
        let headers = new Headers({
            'Authorization': 'bearer ' + token,
            'Content-Type': 'application/json'
        });
        return this.http
            .delete(`${'https://' + localStorage.getItem('Stack') + '/channel/' + channel._idChannel.toString()}`, { headers: headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.handleError);
    }

    formatChannel(channel:Channel) {
       return {
            description: channel.description,
            type: channel.type,
            name: channel.channelName,
            id: channel._idChannel
        }
    }

    private handleError(error: any) {
        console.error('An error occurred', error);
        return Promise.reject(error.message || error);
    }
}
