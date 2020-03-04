package httpclient

const (
	URIPrefix = "/api/v1"

	URIPathClientGetUser                = URIPrefix + "/user"
	URIPathClientGetOrders              = URIPrefix + "/orders"
	URIPathClientGetUserCount           = URIPrefix + "/user/%s/count"
	URIPathClientGetOrdersWithoutParams = URIPrefix + "/orders"

	HTTPMethodGetUser                = "GET"
	HTTPMethodGetOrders              = "POST"
	HTTPMethodGetUserCount           = "GET"
	HTTPMethodGetOrdersWithoutParams = "GET"
)
