export class User {
    _idUser: number;
    webId: string;
    userName: string;
    email: string;
    emailVerified: boolean;
    password: string;
    updatedAt: number;
    deleted: boolean;
    lastPasswordUpdate: number;
    failedAttemprs: number;
    locale: string;
    idRole: number;
    avatar: string;
    nickName: string;
    firstName: string;
    lastName: string;
    constructor(idUser, webId, userName, email, password, updateAt, lastPasswordUpdate, locale, idRole, firstName, lastName, nickName, avatar) {
        this._idUser = idUser;
        this.webId = webId;
        this.userName = userName;
        this.email = email;
        this.updatedAt = updateAt;
        this.lastPasswordUpdate = lastPasswordUpdate;
        this.locale = locale;
        this.idRole = idRole;
        this.firstName = firstName;
        this.lastName = lastName;
        this.nickName = nickName;
        this.avatar = avatar;
    } 
}
