export interface IApiResponse {
    took: number
    timed_out: boolean
    _shards: IShards
    hits: IHits

}

export interface IShards {
    total: number
    successful: number
    skipped: number
    failed: number
}

export interface IHits {
    total: ITotal
    max_score: number
    hits: IEmail[]
}

export interface ITotal {
    value: number
}

export interface IEmail {
    _index: string
    _type: string
    _id: string
    _score: number
    timestamp: string
    _source: ISource
}

export interface ISource {
    timestamp: string
    content: string
    content_transfer_encoding: string
    content_type: string
    date: string
    from: string
    message_id: string
    mime_version: string
    subject: string
    to: string
    x_bcc: string
    x_cc: string
    x_filename: string
    x_folder: string
    x_from: string
    x_origin: string
    x_to: string
}