export class Organisation {
    _idOrganisation: number;
    dockerStack: number;
    organisationName: string;
    description: string;
    avatar: string;
    domain: string;
    status: string;
    constructor(id, organisationName, description, avatar) {
        this._idOrganisation = id;
        this.organisationName = organisationName;
        this.description = description;
        this.avatar = avatar;
    }
}
