package httpserver

const (
	URIPrefix = "/api/v1"

	URIPathGetUser                = URIPrefix + "/user"
	URIPathGetOrders              = URIPrefix + "/orders"
	URIPathGetUserCount           = URIPrefix + "/user/:id/count"
	URIPathGetOrdersWithoutParams = URIPrefix + "/orders"

	HTTPMethodGetUser                = "GET"
	HTTPMethodGetOrders              = "POST"
	HTTPMethodGetUserCount           = "GET"
	HTTPMethodGetOrdersWithoutParams = "GET"
)
