import {Injectable} from '@angular/core';

import { Organisation } from '../model/organisation'
import { User } from '../model/user'

@Injectable()
export class localOrganisationService {

    private localOrganisationKey: string = 'local_organisation';
    private organisations: number[] = [];
    private users: number[] = [];
    private tokens: string[] = [];

    private store(content:Object) {
        localStorage.setItem(this.localOrganisationKey, JSON.stringify(content));
    }

    private retrieve() {
        let storedOrganisation:string = localStorage.getItem(this.localOrganisationKey);
        if(!storedOrganisation) throw 'no organisation found';
        return storedOrganisation;
    }

    public generateNewOrganisation(organisation: number, user: number, token: string) {
        if(this.retrieveOrganisation()){
            this.initAll();
        }
        this.organisations.push(organisation);
        this.users.push(user);
        this.tokens.push(token);
        let organisationKey = JSON.stringify(this.organisations);
        let userKey =  JSON.stringify(this.users);
        let tokenKey = JSON.stringify(this.tokens);
        let currentTime:number = (new Date()).getTime() + 60*60*12;
        this.store({ttl: currentTime, organisationKey, userKey, tokenKey});
    }

    private initAll() {
        let storedOrganisation = this.retrieveOrganisation();
        this.organisations = storedOrganisation.organisation;
        this.users = storedOrganisation.user;
        this.tokens = storedOrganisation.token;
    }

    public retrieveOrganisation() {
        const currentTime2:number = (new Date()).getTime();
        let organisation = null;
        try {
            let storedOrganisation = JSON.parse(this.retrieve());
            console.log("organisation")
            console.log(storedOrganisation.ttl - currentTime2)
            console.log(storedOrganisation.ttl)
            if(storedOrganisation.ttl < currentTime2) throw 'invalid organisation found';
            organisation = storedOrganisation.organisation;
        }
        catch(err) {
            console.error(err);
        }
        return organisation;
    }
}