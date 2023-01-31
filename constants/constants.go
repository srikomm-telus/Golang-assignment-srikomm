package constants

var EXPIRY_SECONDS float64 = 10

var US_CRYPTO_PRICE_CACHE_KEY = "crypto_price_USD"
var EU_CRYPTO_PRICE_CACHE_KEY = "crypto_price_EUR"

var PRICE_IS_FROM_CACHE = true
var PRICE_IS_NOT_FROM_CACHE = false

var USD_CURRENCY_IDENTIFIER = "USD"
var EUR_CURRENCY_IDENTIFIER = "EUR"

var BITCOIN_IDENTIFIER = "BTC"
var ETHEREUM_IDENTIFIER = "ETH"

var COINDESK_ENDPOINT = "https://api.coindesk.com/v1/bpi/currentprice.json"
var CRYPTONATOR_ENDPOINT = "https://api.cryptonator.com/api/ticker/eth-usd"

var DEFAULT_CRYPTO = "BTC"

var ERROR_MESSAGE = "Error Message"
var INVALID_CACHE_ERROR_MESSAGE = "invalid cache"
var PRICE_FETCH_SUCCESSFUL_MESSAGE = "Crypto Price fetched successfully"
var COINDESK_FETCH_ERROR_MESSAGE = "Error while fetching crypto price from CoinDesk"
var CRYPTONATOR_FETCH_ERROR_MESSAGE = "Error while fetching crypto price from Cryptonator"
var ERROR_WHILE_DECODING_RESPONSE = "Error while decoding response"

var SET_CACHE_VALUE = "Set cache value"
var EMPTY_STRING = ""
var CRYPTO_NAME_PARAM_KEY = "cryptoName"

var ERROR = "error"
var PRICE = "price"
var DATA = "data"
