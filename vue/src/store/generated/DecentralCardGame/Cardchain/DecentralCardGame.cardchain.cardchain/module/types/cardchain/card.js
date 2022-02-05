/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
export var Status;
(function (Status) {
    Status[Status["scheme"] = 0] = "scheme";
    Status[Status["prototype"] = 1] = "prototype";
    Status[Status["trial"] = 2] = "trial";
    Status[Status["permanent"] = 3] = "permanent";
    Status[Status["suspended"] = 4] = "suspended";
    Status[Status["banned"] = 5] = "banned";
    Status[Status["bannedSoon"] = 6] = "bannedSoon";
    Status[Status["bannedVerySoon"] = 7] = "bannedVerySoon";
    Status[Status["none"] = 8] = "none";
    Status[Status["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(Status || (Status = {}));
export function statusFromJSON(object) {
    switch (object) {
        case 0:
        case "scheme":
            return Status.scheme;
        case 1:
        case "prototype":
            return Status.prototype;
        case 2:
        case "trial":
            return Status.trial;
        case 3:
        case "permanent":
            return Status.permanent;
        case 4:
        case "suspended":
            return Status.suspended;
        case 5:
        case "banned":
            return Status.banned;
        case 6:
        case "bannedSoon":
            return Status.bannedSoon;
        case 7:
        case "bannedVerySoon":
            return Status.bannedVerySoon;
        case 8:
        case "none":
            return Status.none;
        case -1:
        case "UNRECOGNIZED":
        default:
            return Status.UNRECOGNIZED;
    }
}
export function statusToJSON(object) {
    switch (object) {
        case Status.scheme:
            return "scheme";
        case Status.prototype:
            return "prototype";
        case Status.trial:
            return "trial";
        case Status.permanent:
            return "permanent";
        case Status.suspended:
            return "suspended";
        case Status.banned:
            return "banned";
        case Status.bannedSoon:
            return "bannedSoon";
        case Status.bannedVerySoon:
            return "bannedVerySoon";
        case Status.none:
            return "none";
        default:
            return "UNKNOWN";
    }
}
const baseCard = {
    owner: "",
    artist: "",
    fullArt: false,
    notes: "",
    status: 0,
    votePool: "",
    fairEnoughVotes: 0,
    overpoweredVotes: 0,
    underpoweredVotes: 0,
    inappropriateVotes: 0,
    nerflevel: 0,
};
export const Card = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.artist !== "") {
            writer.uint32(18).string(message.artist);
        }
        if (message.content.length !== 0) {
            writer.uint32(26).bytes(message.content);
        }
        if (message.image.length !== 0) {
            writer.uint32(34).bytes(message.image);
        }
        if (message.fullArt === true) {
            writer.uint32(40).bool(message.fullArt);
        }
        if (message.notes !== "") {
            writer.uint32(50).string(message.notes);
        }
        if (message.status !== 0) {
            writer.uint32(56).int32(message.status);
        }
        if (message.votePool !== "") {
            writer.uint32(66).string(message.votePool);
        }
        if (message.fairEnoughVotes !== 0) {
            writer.uint32(72).uint64(message.fairEnoughVotes);
        }
        if (message.overpoweredVotes !== 0) {
            writer.uint32(80).uint64(message.overpoweredVotes);
        }
        if (message.underpoweredVotes !== 0) {
            writer.uint32(88).uint64(message.underpoweredVotes);
        }
        if (message.inappropriateVotes !== 0) {
            writer.uint32(96).uint64(message.inappropriateVotes);
        }
        if (message.nerflevel !== 0) {
            writer.uint32(104).int64(message.nerflevel);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCard };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = reader.string();
                    break;
                case 2:
                    message.artist = reader.string();
                    break;
                case 3:
                    message.content = reader.bytes();
                    break;
                case 4:
                    message.image = reader.bytes();
                    break;
                case 5:
                    message.fullArt = reader.bool();
                    break;
                case 6:
                    message.notes = reader.string();
                    break;
                case 7:
                    message.status = reader.int32();
                    break;
                case 8:
                    message.votePool = reader.string();
                    break;
                case 9:
                    message.fairEnoughVotes = longToNumber(reader.uint64());
                    break;
                case 10:
                    message.overpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 11:
                    message.underpoweredVotes = longToNumber(reader.uint64());
                    break;
                case 12:
                    message.inappropriateVotes = longToNumber(reader.uint64());
                    break;
                case 13:
                    message.nerflevel = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCard };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = String(object.artist);
        }
        else {
            message.artist = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = bytesFromBase64(object.content);
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = bytesFromBase64(object.image);
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = Boolean(object.fullArt);
        }
        else {
            message.fullArt = false;
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = String(object.notes);
        }
        else {
            message.notes = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = statusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        if (object.votePool !== undefined && object.votePool !== null) {
            message.votePool = String(object.votePool);
        }
        else {
            message.votePool = "";
        }
        if (object.fairEnoughVotes !== undefined &&
            object.fairEnoughVotes !== null) {
            message.fairEnoughVotes = Number(object.fairEnoughVotes);
        }
        else {
            message.fairEnoughVotes = 0;
        }
        if (object.overpoweredVotes !== undefined &&
            object.overpoweredVotes !== null) {
            message.overpoweredVotes = Number(object.overpoweredVotes);
        }
        else {
            message.overpoweredVotes = 0;
        }
        if (object.underpoweredVotes !== undefined &&
            object.underpoweredVotes !== null) {
            message.underpoweredVotes = Number(object.underpoweredVotes);
        }
        else {
            message.underpoweredVotes = 0;
        }
        if (object.inappropriateVotes !== undefined &&
            object.inappropriateVotes !== null) {
            message.inappropriateVotes = Number(object.inappropriateVotes);
        }
        else {
            message.inappropriateVotes = 0;
        }
        if (object.nerflevel !== undefined && object.nerflevel !== null) {
            message.nerflevel = Number(object.nerflevel);
        }
        else {
            message.nerflevel = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.owner !== undefined && (obj.owner = message.owner);
        message.artist !== undefined && (obj.artist = message.artist);
        message.content !== undefined &&
            (obj.content = base64FromBytes(message.content !== undefined ? message.content : new Uint8Array()));
        message.image !== undefined &&
            (obj.image = base64FromBytes(message.image !== undefined ? message.image : new Uint8Array()));
        message.fullArt !== undefined && (obj.fullArt = message.fullArt);
        message.notes !== undefined && (obj.notes = message.notes);
        message.status !== undefined && (obj.status = statusToJSON(message.status));
        message.votePool !== undefined && (obj.votePool = message.votePool);
        message.fairEnoughVotes !== undefined &&
            (obj.fairEnoughVotes = message.fairEnoughVotes);
        message.overpoweredVotes !== undefined &&
            (obj.overpoweredVotes = message.overpoweredVotes);
        message.underpoweredVotes !== undefined &&
            (obj.underpoweredVotes = message.underpoweredVotes);
        message.inappropriateVotes !== undefined &&
            (obj.inappropriateVotes = message.inappropriateVotes);
        message.nerflevel !== undefined && (obj.nerflevel = message.nerflevel);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCard };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = object.artist;
        }
        else {
            message.artist = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = new Uint8Array();
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = object.image;
        }
        else {
            message.image = new Uint8Array();
        }
        if (object.fullArt !== undefined && object.fullArt !== null) {
            message.fullArt = object.fullArt;
        }
        else {
            message.fullArt = false;
        }
        if (object.notes !== undefined && object.notes !== null) {
            message.notes = object.notes;
        }
        else {
            message.notes = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        if (object.votePool !== undefined && object.votePool !== null) {
            message.votePool = object.votePool;
        }
        else {
            message.votePool = "";
        }
        if (object.fairEnoughVotes !== undefined &&
            object.fairEnoughVotes !== null) {
            message.fairEnoughVotes = object.fairEnoughVotes;
        }
        else {
            message.fairEnoughVotes = 0;
        }
        if (object.overpoweredVotes !== undefined &&
            object.overpoweredVotes !== null) {
            message.overpoweredVotes = object.overpoweredVotes;
        }
        else {
            message.overpoweredVotes = 0;
        }
        if (object.underpoweredVotes !== undefined &&
            object.underpoweredVotes !== null) {
            message.underpoweredVotes = object.underpoweredVotes;
        }
        else {
            message.underpoweredVotes = 0;
        }
        if (object.inappropriateVotes !== undefined &&
            object.inappropriateVotes !== null) {
            message.inappropriateVotes = object.inappropriateVotes;
        }
        else {
            message.inappropriateVotes = 0;
        }
        if (object.nerflevel !== undefined && object.nerflevel !== null) {
            message.nerflevel = object.nerflevel;
        }
        else {
            message.nerflevel = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
const atob = globalThis.atob ||
    ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa ||
    ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(""));
}
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
