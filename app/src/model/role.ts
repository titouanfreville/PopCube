export class Role {
    _idRole: number;
    roleName: string;
    canUserPrivate: boolean;
    caModerate: boolean;
    canArchive: boolean;
    canInvite: boolean;
    canManage: boolean;
    canManageUser: boolean;
    constructor(roleName, canUserPrivate, canModerate, canArchive, canInvite, canManage, canManageUser) {
        this._idRole = 1;
        this.roleName = roleName;
        this.canUserPrivate = canUserPrivate;
        this.caModerate = canModerate;
        this.canArchive = canArchive;
        this.canInvite = canInvite;
        this.canManage = canManage;
        this.canManageUser = canManageUser;
    }
}
