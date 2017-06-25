import { Message } from './message';

export class Channel {
    _idChannel: number;
    webId: string;
    channelName: string;
    type: string;
    private: boolean;
    updatedAt: number;
    description: string;
    avatar: string;
    subject: string;
    status: string;
    messages: Message[];
    constructor(idChannel, channelName, type, description) {
        this._idChannel = idChannel;
        this.channelName = channelName;
        this.type = type;
        this.description = description;
        this.messages = [];
    }
}
