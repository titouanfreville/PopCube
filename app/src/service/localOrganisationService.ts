import {Injectable} from '@angular/core';

@Injectable()
export class localOrganisationService {

    private localOrganisationKey: string = 'local_organisation';

    private store(content: Object) {
        localStorage.setItem(this.localOrganisationKey, JSON.stringify(content));
    }

    private retrieve(organisationId) {
        let storedOrganisation: string = localStorage.getItem(this.localOrganisationKey += organisationId);
        if(!storedOrganisation) throw 'no organisation found';
        return storedOrganisation;
    }

    public generateNewOrganisation(organisation: number, user: number, token: string) {
        this.localOrganisationKey += organisation.toString();
        let organisationKey = JSON.stringify(organisation);
        let userKey =  JSON.stringify(user);
        let tokenKey = token;
        let currentTime: number = (new Date()).getTime() + 60 * 60 * 12;
        this.store({ttl: currentTime, organisationKey, userKey, tokenKey});
    }

    public retrieveOrganisation(organisationId) {
        const currentTime2: number = (new Date()).getTime();
        let organisation = null;
        try {
            let storedOrganisation = JSON.parse(this.retrieve(organisationId));
            if (storedOrganisation.ttl < currentTime2) throw 'invalid organisation found';
            organisation = storedOrganisation;
        }
        catch (err) {
            console.error(err);
        }
        return organisation;
    }
}