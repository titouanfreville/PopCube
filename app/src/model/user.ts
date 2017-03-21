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
    constructor(idUser, userName, password, email, firstName, lastName, avatar) {
        this._idUser = idUser;
        this.userName = userName;
        this.password = password;
        this.email = email;
        this.firstName = firstName;
        this.lastName = lastName;
        this.avatar = avatar;
    }
}
