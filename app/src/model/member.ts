import { User } from './user';
import { Channel } from './channel';


export class Member {
    _User: User;
    _Chanel: Channel[];
    role: number;
    constructor(User, Channel, role) {
        this._User = User;
        this._Chanel = Channel;
        this.role = role;
    }
}
