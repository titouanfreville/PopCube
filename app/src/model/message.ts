import { User } from './user';

export class Message {
    _idMessage: number;
    date: number;
    content: String;
    creator: User;
    idChannel: number;
    status: string;
    constructor(idMessage: number, date: number, content: String, user: User, idChannel: number) {
        this._idMessage = idMessage;
        this.date = date;
        this.content = content;
        this.creator = user;
        this.idChannel = idChannel;
    }
}
