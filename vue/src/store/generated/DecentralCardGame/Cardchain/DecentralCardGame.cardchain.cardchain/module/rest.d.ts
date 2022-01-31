export declare type CardchainMsgBuyCardSchemeResponse = object;
export declare type CardchainMsgCreateuserResponse = object;
export declare type CardchainMsgDonateToCardResponse = object;
export declare type CardchainMsgSaveCardContentResponse = object;
export declare type CardchainMsgTransferCardResponse = object;
export declare type CardchainMsgVoteCardResponse = object;
/**
 * Params defines the parameters for the module.
 */
export declare type CardchainParams = object;
/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface CardchainQueryParamsResponse {
    /** params holds all the parameters of this module. */
    params?: CardchainParams;
}
export interface CardchainQueryQCardContentResponse {
    content?: string;
}
export interface CardchainQueryQCardResponse {
    owner?: string;
    content?: string;
    artist?: string;
    image?: string;
    fullArt?: boolean;
    notes?: string;
    status?: string;
    votePool?: string;
    /** @format uint64 */
    fairEnoughVotes?: string;
    /** @format uint64 */
    overpoweredVotes?: string;
    /** @format uint64 */
    underpoweredVotes?: string;
    /** @format uint64 */
    inappropriateVotes?: string;
    /** @format int64 */
    nerflevel?: string;
}
export interface CardchainQueryQCardchainInfoResponse {
    cardAuctionPrice?: string;
}
export interface CardchainQueryQCardsResponse {
    cardsList?: string[];
}
export interface CardchainQueryQUserResponse {
    alias?: string;
    ownedCardSchemes?: string[];
    ownedCards?: string[];
    voteRights?: CardchainVoteRight[];
}
export interface CardchainQueryQVotableCardsResponse {
    unregistered?: boolean;
    noVoteRights?: boolean;
    voteRights?: CardchainVoteRight[];
}
export interface CardchainQueryQVotingResultsResponse {
    lastVotingResults?: CardchainVotingResults;
}
export interface CardchainVoteRight {
    /** @format uint64 */
    cardId?: string;
    /** @format int64 */
    expireBlock?: string;
}
export interface CardchainVotingResult {
    /** @format uint64 */
    cardId?: string;
    /** @format uint64 */
    fairEnoughVotes?: string;
    /** @format uint64 */
    overpoweredVotes?: string;
    /** @format uint64 */
    underpoweredVotes?: string;
    /** @format uint64 */
    inappropriateVotes?: string;
    result?: string;
}
export interface CardchainVotingResults {
    /** @format uint64 */
    totalVotes?: string;
    /** @format uint64 */
    totalFairEnoughVotes?: string;
    /** @format uint64 */
    totalOverpoweredVotes?: string;
    /** @format uint64 */
    totalUnderpoweredVotes?: string;
    /** @format uint64 */
    totalInappropriateVotes?: string;
    cardResults?: CardchainVotingResult[];
    notes?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title cardchain/card.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryParams
     * @summary Parameters queries the parameters of the module.
     * @request GET:/DecentralCardGame/cardchain/cardchain/params
     */
    queryParams: (params?: RequestParams) => Promise<HttpResponse<CardchainQueryParamsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQCard
     * @summary Queries a list of QCard items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_card/{cardId}
     */
    queryQCard: (cardId: string, params?: RequestParams) => Promise<HttpResponse<CardchainQueryQCardResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQCardContent
     * @summary Queries a list of QCardContent items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_card_content/{cardId}
     */
    queryQCardContent: (cardId: string, params?: RequestParams) => Promise<HttpResponse<CardchainQueryQCardContentResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQCardchainInfo
     * @summary Queries a list of QCardchainInfo items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_cardchain_info
     */
    queryQCardchainInfo: (params?: RequestParams) => Promise<HttpResponse<CardchainQueryQCardchainInfoResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQCards
     * @summary Queries a list of QCards items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_cards/{owner}/{status}/{cardType}/{classes}/{sortBy}/{nameContains}/{keywordsContains}/{notesContains}
     */
    queryQCards: (owner: string, status: string, cardType: string, classes: string, sortBy: string, nameContains: string, keywordsContains: string, notesContains: string, params?: RequestParams) => Promise<HttpResponse<CardchainQueryQCardsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQUser
     * @summary Queries a list of QUser items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_user/{address}
     */
    queryQUser: (address: string, params?: RequestParams) => Promise<HttpResponse<CardchainQueryQUserResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQVotableCards
     * @summary Queries a list of QVotableCards items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_votable_cards/{address}
     */
    queryQVotableCards: (address: string, params?: RequestParams) => Promise<HttpResponse<CardchainQueryQVotableCardsResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryQVotingResults
     * @summary Queries a list of QVotingResults items.
     * @request GET:/DecentralCardGame/cardchain/cardchain/q_voting_results
     */
    queryQVotingResults: (params?: RequestParams) => Promise<HttpResponse<CardchainQueryQVotingResultsResponse, RpcStatus>>;
}
export {};
