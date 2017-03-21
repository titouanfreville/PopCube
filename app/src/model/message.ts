import { User } from './user';

export class Message {
    _idMessage: number;
    date: Date;
    content: string;
    creator: User;
    status: string;
    constructor(idMessage, content, user) {
        this._idMessage = idMessage;
        this.content = content;
        this.creator = user;
    }
}
