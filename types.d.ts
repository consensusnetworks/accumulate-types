// generated using github.com/32leaves/bel on 2022-01-17 00:59:10.97634 -0500 EST m=+0.006253834
// DO NOT MODIFY
// Generated TypeScript types for Accumulate API 

export interface APIDataResponse {
    type: string
    merkleState?: MerkleState
    data: number[]
    sponsor: string
    keyPage: APIRequestKeyPage
    txid: number[]
    signer?: Signer
    sig?: number[]
    status?: number[]
}

export interface APIDataResponsePagination {
    data: APIDataResponse[]
    type: string
    start: number
    limit: number
    total: number
}

export interface APIRequestChainId {
    chainId: number[]
    wait: boolean
}

export interface APIRequestKeyPage {
    height: number
    index: number
}

export interface APIRequestRaw {
    wait: boolean
    tx: APIRequestRawTx
}

export interface APIRequestURL {
    url: string
    wait: boolean
}

export interface APIRequestRawTx {
    sponsor: string
    data: number[]
    signer: Signer
    sig: number[]
    keyPage: APIRequestKeyPage
}

export interface APIRequestTxId {
    txid: number[]
    wait: boolean
}

export interface APIRequestURLPagination {
    APIRequestURL: APIRequestURL
    start: number
    limit: number
}

export interface MerkleState {
    count?: number
    roots?: number[][]
}

export interface Signer {
    publicKey: number[]
    nonce: number
}

export interface Query {
    Type: number
    RouteId: number
    Content: number[]
}

export interface RequestDataEntry {
    RequestDataEntry: RequestDataEntry
}

export interface RequestTxHistory {
    ChainId: number[]
    Start: number
    Limit: number
}

export interface ResponseTxHistory {
    txs: ResponseByTxId[]
    total: number
}

export interface ResponseByTxId {
    TxId: number[]
    TxState: number[]
    TxPendingState: number[]
    TxSynthTxIds: number[]
}

export interface ResponseByChainId {
    Object: Object
}

export interface ResponseKeyPageIndex {
    keyBook?: string
    keyPage?: string
    index: number
}
