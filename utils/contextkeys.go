package utils

type contextKey string

const UserKey = contextKey("user")
const TokenCookieKey = contextKey("tokencookie")
const ResponseWriterKey = contextKey("responsewriter")
