import {Injectable} from '@angular/core';

@Injectable()
export class TokenManager {

    private tokenKey: string = 'app_token';

    private store(content: Object) {
        localStorage.setItem(this.tokenKey, JSON.stringify(content));
    }

    private retrieve() {
        let storedToken: string = localStorage.getItem(this.tokenKey);
        if (!storedToken) throw 'no token found';
        return storedToken;
    }

    public generateNewToken(token) {
        let currentTime: number = (new Date()).getTime() + 60 * 60;
        this.store({ttl: currentTime, token});
    }


    public retrieveToken() {
        let currentTime: number = (new Date()).getTime(), token = null;
        try {
            let storedToken = JSON.parse(this.retrieve());
            console.log('Token :');
            console.log(storedToken.ttl - currentTime);
            console.log(storedToken.ttl);
            if (storedToken.ttl < currentTime) throw 'invalid token found';
            token = storedToken.token;
        }
        catch (err) {
            console.error(err);
        }
        return token;
    }
}