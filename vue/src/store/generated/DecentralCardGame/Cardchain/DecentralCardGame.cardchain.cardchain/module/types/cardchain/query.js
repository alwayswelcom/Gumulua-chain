/* eslint-disable */
import { statusFromJSON, statusToJSON } from "../cardchain/card";
import { councilStatusFromJSON, councilStatusToJSON, } from "../cardchain/user";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../cardchain/params";
import { VoteRight } from "../cardchain/vote_right";
import { VotingResults } from "../cardchain/voting_results";
export const protobufPackage = "DecentralCardGame.cardchain.cardchain";
const baseQueryParamsRequest = {};
export const QueryParamsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryParamsRequest };
        return message;
    },
};
const baseQueryParamsResponse = {};
export const QueryParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryParamsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
};
const baseQueryQCardRequest = { cardId: "" };
export const QueryQCardRequest = {
    encode(message, writer = Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQCardRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQCardRequest };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = String(object.cardId);
        }
        else {
            message.cardId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQCardRequest };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = object.cardId;
        }
        else {
            message.cardId = "";
        }
        return message;
    },
};
const baseQueryQCardResponse = {
    owner: "",
    content: "",
    artist: "",
    image: "",
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
export const QueryQCardResponse = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.content !== "") {
            writer.uint32(18).string(message.content);
        }
        if (message.artist !== "") {
            writer.uint32(26).string(message.artist);
        }
        if (message.image !== "") {
            writer.uint32(34).string(message.image);
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
        const message = { ...baseQueryQCardResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = reader.string();
                    break;
                case 2:
                    message.content = reader.string();
                    break;
                case 3:
                    message.artist = reader.string();
                    break;
                case 4:
                    message.image = reader.string();
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
        const message = { ...baseQueryQCardResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = String(object.content);
        }
        else {
            message.content = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = String(object.artist);
        }
        else {
            message.artist = "";
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = String(object.image);
        }
        else {
            message.image = "";
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
        message.content !== undefined && (obj.content = message.content);
        message.artist !== undefined && (obj.artist = message.artist);
        message.image !== undefined && (obj.image = message.image);
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
        const message = { ...baseQueryQCardResponse };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = "";
        }
        if (object.artist !== undefined && object.artist !== null) {
            message.artist = object.artist;
        }
        else {
            message.artist = "";
        }
        if (object.image !== undefined && object.image !== null) {
            message.image = object.image;
        }
        else {
            message.image = "";
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
const baseQueryQCardContentRequest = { cardId: "" };
export const QueryQCardContentRequest = {
    encode(message, writer = Writer.create()) {
        if (message.cardId !== "") {
            writer.uint32(10).string(message.cardId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardContentRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQCardContentRequest,
        };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = String(object.cardId);
        }
        else {
            message.cardId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cardId !== undefined && (obj.cardId = message.cardId);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQCardContentRequest,
        };
        if (object.cardId !== undefined && object.cardId !== null) {
            message.cardId = object.cardId;
        }
        else {
            message.cardId = "";
        }
        return message;
    },
};
const baseQueryQCardContentResponse = { content: "" };
export const QueryQCardContentResponse = {
    encode(message, writer = Writer.create()) {
        if (message.content !== "") {
            writer.uint32(10).string(message.content);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardContentResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.content = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQCardContentResponse,
        };
        if (object.content !== undefined && object.content !== null) {
            message.content = String(object.content);
        }
        else {
            message.content = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.content !== undefined && (obj.content = message.content);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQCardContentResponse,
        };
        if (object.content !== undefined && object.content !== null) {
            message.content = object.content;
        }
        else {
            message.content = "";
        }
        return message;
    },
};
const baseQueryQUserRequest = { address: "" };
export const QueryQUserRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQUserRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQUserRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQUserRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryQUserResponse = {
    alias: "",
    ownedCardSchemes: 0,
    ownedCards: 0,
    councilStatus: 0,
};
export const QueryQUserResponse = {
    encode(message, writer = Writer.create()) {
        if (message.alias !== "") {
            writer.uint32(10).string(message.alias);
        }
        writer.uint32(18).fork();
        for (const v of message.ownedCardSchemes) {
            writer.uint64(v);
        }
        writer.ldelim();
        writer.uint32(26).fork();
        for (const v of message.ownedCards) {
            writer.uint64(v);
        }
        writer.ldelim();
        for (const v of message.voteRights) {
            VoteRight.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.councilStatus !== 0) {
            writer.uint32(40).int32(message.councilStatus);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQUserResponse };
        message.ownedCardSchemes = [];
        message.ownedCards = [];
        message.voteRights = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.alias = reader.string();
                    break;
                case 2:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedCardSchemes.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.ownedCards.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.ownedCards.push(longToNumber(reader.uint64()));
                    }
                    break;
                case 4:
                    message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.councilStatus = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQUserResponse };
        message.ownedCardSchemes = [];
        message.ownedCards = [];
        message.voteRights = [];
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = String(object.alias);
        }
        else {
            message.alias = "";
        }
        if (object.ownedCardSchemes !== undefined &&
            object.ownedCardSchemes !== null) {
            for (const e of object.ownedCardSchemes) {
                message.ownedCardSchemes.push(Number(e));
            }
        }
        if (object.ownedCards !== undefined && object.ownedCards !== null) {
            for (const e of object.ownedCards) {
                message.ownedCards.push(Number(e));
            }
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromJSON(e));
            }
        }
        if (object.councilStatus !== undefined && object.councilStatus !== null) {
            message.councilStatus = councilStatusFromJSON(object.councilStatus);
        }
        else {
            message.councilStatus = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.alias !== undefined && (obj.alias = message.alias);
        if (message.ownedCardSchemes) {
            obj.ownedCardSchemes = message.ownedCardSchemes.map((e) => e);
        }
        else {
            obj.ownedCardSchemes = [];
        }
        if (message.ownedCards) {
            obj.ownedCards = message.ownedCards.map((e) => e);
        }
        else {
            obj.ownedCards = [];
        }
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        message.councilStatus !== undefined &&
            (obj.councilStatus = councilStatusToJSON(message.councilStatus));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQUserResponse };
        message.ownedCardSchemes = [];
        message.ownedCards = [];
        message.voteRights = [];
        if (object.alias !== undefined && object.alias !== null) {
            message.alias = object.alias;
        }
        else {
            message.alias = "";
        }
        if (object.ownedCardSchemes !== undefined &&
            object.ownedCardSchemes !== null) {
            for (const e of object.ownedCardSchemes) {
                message.ownedCardSchemes.push(e);
            }
        }
        if (object.ownedCards !== undefined && object.ownedCards !== null) {
            for (const e of object.ownedCards) {
                message.ownedCards.push(e);
            }
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromPartial(e));
            }
        }
        if (object.councilStatus !== undefined && object.councilStatus !== null) {
            message.councilStatus = object.councilStatus;
        }
        else {
            message.councilStatus = 0;
        }
        return message;
    },
};
const baseQueryQCardchainInfoRequest = {};
export const QueryQCardchainInfoRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardchainInfoRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseQueryQCardchainInfoRequest,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseQueryQCardchainInfoRequest,
        };
        return message;
    },
};
const baseQueryQCardchainInfoResponse = { cardAuctionPrice: "" };
export const QueryQCardchainInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.cardAuctionPrice !== "") {
            writer.uint32(10).string(message.cardAuctionPrice);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQCardchainInfoResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cardAuctionPrice = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQCardchainInfoResponse,
        };
        if (object.cardAuctionPrice !== undefined &&
            object.cardAuctionPrice !== null) {
            message.cardAuctionPrice = String(object.cardAuctionPrice);
        }
        else {
            message.cardAuctionPrice = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cardAuctionPrice !== undefined &&
            (obj.cardAuctionPrice = message.cardAuctionPrice);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQCardchainInfoResponse,
        };
        if (object.cardAuctionPrice !== undefined &&
            object.cardAuctionPrice !== null) {
            message.cardAuctionPrice = object.cardAuctionPrice;
        }
        else {
            message.cardAuctionPrice = "";
        }
        return message;
    },
};
const baseQueryQVotingResultsRequest = {};
export const QueryQVotingResultsRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQVotingResultsRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = {
            ...baseQueryQVotingResultsRequest,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseQueryQVotingResultsRequest,
        };
        return message;
    },
};
const baseQueryQVotingResultsResponse = {};
export const QueryQVotingResultsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.lastVotingResults !== undefined) {
            VotingResults.encode(message.lastVotingResults, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQVotingResultsResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.lastVotingResults = VotingResults.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQVotingResultsResponse,
        };
        if (object.lastVotingResults !== undefined &&
            object.lastVotingResults !== null) {
            message.lastVotingResults = VotingResults.fromJSON(object.lastVotingResults);
        }
        else {
            message.lastVotingResults = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.lastVotingResults !== undefined &&
            (obj.lastVotingResults = message.lastVotingResults
                ? VotingResults.toJSON(message.lastVotingResults)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQVotingResultsResponse,
        };
        if (object.lastVotingResults !== undefined &&
            object.lastVotingResults !== null) {
            message.lastVotingResults = VotingResults.fromPartial(object.lastVotingResults);
        }
        else {
            message.lastVotingResults = undefined;
        }
        return message;
    },
};
const baseQueryQVotableCardsRequest = { address: "" };
export const QueryQVotableCardsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQVotableCardsRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQVotableCardsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQVotableCardsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryQVotableCardsResponse = {
    unregistered: false,
    noVoteRights: false,
};
export const QueryQVotableCardsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.unregistered === true) {
            writer.uint32(8).bool(message.unregistered);
        }
        if (message.noVoteRights === true) {
            writer.uint32(16).bool(message.noVoteRights);
        }
        for (const v of message.voteRights) {
            VoteRight.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryQVotableCardsResponse,
        };
        message.voteRights = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.unregistered = reader.bool();
                    break;
                case 2:
                    message.noVoteRights = reader.bool();
                    break;
                case 3:
                    message.voteRights.push(VoteRight.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryQVotableCardsResponse,
        };
        message.voteRights = [];
        if (object.unregistered !== undefined && object.unregistered !== null) {
            message.unregistered = Boolean(object.unregistered);
        }
        else {
            message.unregistered = false;
        }
        if (object.noVoteRights !== undefined && object.noVoteRights !== null) {
            message.noVoteRights = Boolean(object.noVoteRights);
        }
        else {
            message.noVoteRights = false;
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.unregistered !== undefined &&
            (obj.unregistered = message.unregistered);
        message.noVoteRights !== undefined &&
            (obj.noVoteRights = message.noVoteRights);
        if (message.voteRights) {
            obj.voteRights = message.voteRights.map((e) => e ? VoteRight.toJSON(e) : undefined);
        }
        else {
            obj.voteRights = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryQVotableCardsResponse,
        };
        message.voteRights = [];
        if (object.unregistered !== undefined && object.unregistered !== null) {
            message.unregistered = object.unregistered;
        }
        else {
            message.unregistered = false;
        }
        if (object.noVoteRights !== undefined && object.noVoteRights !== null) {
            message.noVoteRights = object.noVoteRights;
        }
        else {
            message.noVoteRights = false;
        }
        if (object.voteRights !== undefined && object.voteRights !== null) {
            for (const e of object.voteRights) {
                message.voteRights.push(VoteRight.fromPartial(e));
            }
        }
        return message;
    },
};
const baseQueryQCardsRequest = {
    owner: "",
    status: 0,
    cardType: "",
    classes: "",
    sortBy: "",
    nameContains: "",
    keywordsContains: "",
    notesContains: "",
};
export const QueryQCardsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.status !== 0) {
            writer.uint32(16).int32(message.status);
        }
        if (message.cardType !== "") {
            writer.uint32(26).string(message.cardType);
        }
        if (message.classes !== "") {
            writer.uint32(34).string(message.classes);
        }
        if (message.sortBy !== "") {
            writer.uint32(42).string(message.sortBy);
        }
        if (message.nameContains !== "") {
            writer.uint32(50).string(message.nameContains);
        }
        if (message.keywordsContains !== "") {
            writer.uint32(58).string(message.keywordsContains);
        }
        if (message.notesContains !== "") {
            writer.uint32(66).string(message.notesContains);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQCardsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.owner = reader.string();
                    break;
                case 2:
                    message.status = reader.int32();
                    break;
                case 3:
                    message.cardType = reader.string();
                    break;
                case 4:
                    message.classes = reader.string();
                    break;
                case 5:
                    message.sortBy = reader.string();
                    break;
                case 6:
                    message.nameContains = reader.string();
                    break;
                case 7:
                    message.keywordsContains = reader.string();
                    break;
                case 8:
                    message.notesContains = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQCardsRequest };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = statusFromJSON(object.status);
        }
        else {
            message.status = 0;
        }
        if (object.cardType !== undefined && object.cardType !== null) {
            message.cardType = String(object.cardType);
        }
        else {
            message.cardType = "";
        }
        if (object.classes !== undefined && object.classes !== null) {
            message.classes = String(object.classes);
        }
        else {
            message.classes = "";
        }
        if (object.sortBy !== undefined && object.sortBy !== null) {
            message.sortBy = String(object.sortBy);
        }
        else {
            message.sortBy = "";
        }
        if (object.nameContains !== undefined && object.nameContains !== null) {
            message.nameContains = String(object.nameContains);
        }
        else {
            message.nameContains = "";
        }
        if (object.keywordsContains !== undefined &&
            object.keywordsContains !== null) {
            message.keywordsContains = String(object.keywordsContains);
        }
        else {
            message.keywordsContains = "";
        }
        if (object.notesContains !== undefined && object.notesContains !== null) {
            message.notesContains = String(object.notesContains);
        }
        else {
            message.notesContains = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.owner !== undefined && (obj.owner = message.owner);
        message.status !== undefined && (obj.status = statusToJSON(message.status));
        message.cardType !== undefined && (obj.cardType = message.cardType);
        message.classes !== undefined && (obj.classes = message.classes);
        message.sortBy !== undefined && (obj.sortBy = message.sortBy);
        message.nameContains !== undefined &&
            (obj.nameContains = message.nameContains);
        message.keywordsContains !== undefined &&
            (obj.keywordsContains = message.keywordsContains);
        message.notesContains !== undefined &&
            (obj.notesContains = message.notesContains);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQCardsRequest };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.status !== undefined && object.status !== null) {
            message.status = object.status;
        }
        else {
            message.status = 0;
        }
        if (object.cardType !== undefined && object.cardType !== null) {
            message.cardType = object.cardType;
        }
        else {
            message.cardType = "";
        }
        if (object.classes !== undefined && object.classes !== null) {
            message.classes = object.classes;
        }
        else {
            message.classes = "";
        }
        if (object.sortBy !== undefined && object.sortBy !== null) {
            message.sortBy = object.sortBy;
        }
        else {
            message.sortBy = "";
        }
        if (object.nameContains !== undefined && object.nameContains !== null) {
            message.nameContains = object.nameContains;
        }
        else {
            message.nameContains = "";
        }
        if (object.keywordsContains !== undefined &&
            object.keywordsContains !== null) {
            message.keywordsContains = object.keywordsContains;
        }
        else {
            message.keywordsContains = "";
        }
        if (object.notesContains !== undefined && object.notesContains !== null) {
            message.notesContains = object.notesContains;
        }
        else {
            message.notesContains = "";
        }
        return message;
    },
};
const baseQueryQCardsResponse = { cardsList: 0 };
export const QueryQCardsResponse = {
    encode(message, writer = Writer.create()) {
        writer.uint32(10).fork();
        for (const v of message.cardsList) {
            writer.uint64(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryQCardsResponse };
        message.cardsList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.cardsList.push(longToNumber(reader.uint64()));
                        }
                    }
                    else {
                        message.cardsList.push(longToNumber(reader.uint64()));
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryQCardsResponse };
        message.cardsList = [];
        if (object.cardsList !== undefined && object.cardsList !== null) {
            for (const e of object.cardsList) {
                message.cardsList.push(Number(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.cardsList) {
            obj.cardsList = message.cardsList.map((e) => e);
        }
        else {
            obj.cardsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryQCardsResponse };
        message.cardsList = [];
        if (object.cardsList !== undefined && object.cardsList !== null) {
            for (const e of object.cardsList) {
                message.cardsList.push(e);
            }
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Params(request) {
        const data = QueryParamsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "Params", data);
        return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
    }
    QCard(request) {
        const data = QueryQCardRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCard", data);
        return promise.then((data) => QueryQCardResponse.decode(new Reader(data)));
    }
    QCardContent(request) {
        const data = QueryQCardContentRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardContent", data);
        return promise.then((data) => QueryQCardContentResponse.decode(new Reader(data)));
    }
    QUser(request) {
        const data = QueryQUserRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QUser", data);
        return promise.then((data) => QueryQUserResponse.decode(new Reader(data)));
    }
    QCardchainInfo(request) {
        const data = QueryQCardchainInfoRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCardchainInfo", data);
        return promise.then((data) => QueryQCardchainInfoResponse.decode(new Reader(data)));
    }
    QVotingResults(request) {
        const data = QueryQVotingResultsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotingResults", data);
        return promise.then((data) => QueryQVotingResultsResponse.decode(new Reader(data)));
    }
    QVotableCards(request) {
        const data = QueryQVotableCardsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QVotableCards", data);
        return promise.then((data) => QueryQVotableCardsResponse.decode(new Reader(data)));
    }
    QCards(request) {
        const data = QueryQCardsRequest.encode(request).finish();
        const promise = this.rpc.request("DecentralCardGame.cardchain.cardchain.Query", "QCards", data);
        return promise.then((data) => QueryQCardsResponse.decode(new Reader(data)));
    }
}
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
