import {Injectable} from '@angular/core';

@Injectable()
export class localOrganisationService {

    private localOrganisationKey: string = 'local_organisation';
    private isConnected: string = 'isConnected';
    private organisationSet: string = 'organisationSet';

    private store(content: Object) {
        let i: number;
         if (localStorage.getItem(this.organisationSet)) {
            console.log("OrganisationSet is set at :")
            i = parseInt(localStorage.getItem(this.organisationSet), 10) + 1;
            console.log(i);
         }else {
            i = 1;
         }
        localStorage.setItem(this.organisationSet, JSON.stringify(i));
        localStorage.setItem(this.localOrganisationKey + i, JSON.stringify(content));
        if (localStorage.getItem(this.isConnected)) { }else { localStorage.setItem(this.isConnected, JSON.stringify(1)); }
    }

    private retrieve(i) {
        let storedOrganisation: string = localStorage.getItem(this.localOrganisationKey + i);
        if(!storedOrganisation) throw 'no organisation found';
        return storedOrganisation;
    }

    public generateNewOrganisation(i: number, user: number, token: string) {
        let organisationKey = JSON.stringify(i);
        let y: string;
        let organisationName: string = '';
        let x: number = 0;
        do {
            organisationName += localStorage.getItem('Stack').charAt(x);
            y = localStorage.getItem('Stack').charAt(x+1);
            x++;
            console.log('x : ' + x + 'y : ' + y );

        } while (y != '.');
        console.log(organisationName);
        let stack = localStorage.getItem('Stack');
        let userKey =  JSON.stringify(user);
        let tokenKey = token;
        let currentTime: number = (new Date()).getTime() + 1000 * 60 * 60 * 24 * 30 * 12;
        this.store({ttl: currentTime, organisationKey, organisationName, userKey, tokenKey, stack,});
    }

    public retrieveOrganisation(i: number) {
        const currentTime2: number = (new Date()).getTime();
        let organisation = null;
        try {
            let storedOrganisation = JSON.parse(this.retrieve(i));
            if (storedOrganisation.ttl < currentTime2) throw 'invalid organisation found';
            organisation = storedOrganisation;
        }
        catch (err) {
            console.error(err);
        }
        return organisation;
    }

    public retrieveAllOrganisation() {
        const currentTime2: number = (new Date()).getTime();
        let organisations = [];
        let max: number;
        if (localStorage.getItem('organisationSet')) {
            max = parseInt(localStorage.getItem('organisationSet'), 10);
        }else {
            max = 1;
        }
        for (let i = 1; i<= max; i++) {
            try {
                let storedOrganisation = JSON.parse(this.retrieve(i));
                console.log('all Organisation : ' + storedOrganisation);
                if (storedOrganisation.ttl < currentTime2) throw 'invalid organisation found';
                organisations.push(storedOrganisation);
            }
            catch (err) {
                console.error(err);
            }
        }
        return organisations;
    }
}