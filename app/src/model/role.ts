export class Role {
    id: number;
    name: string;
    can_use_private: boolean;
    can_moderate: boolean;
    can_archive: boolean;
    can_invite: boolean;
    can_manage: boolean;
    can_manage_user: boolean;
    constructor(id, roleName, canUserPrivate, canModerate, canArchive, canInvite, canManage, canManageUser) {
        this.id = id;
        this.name = roleName;
        this.can_use_private = canUserPrivate;
        this.can_moderate = canModerate;
        this.can_archive = canArchive;
        this.can_invite = canInvite;
        this.can_manage = canManage;
        this.can_manage_user = canManageUser;
    }
}
