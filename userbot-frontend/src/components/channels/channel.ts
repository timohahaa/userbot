export class Channel {
    public readonly id: number
    public readonly userount: number
    public readonly name: string

    constructor(id: number, usercount: number, name: string) {
        this.id = id
        this.userount = usercount
        this.name = name
    }
}