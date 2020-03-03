package httpserver

const (
	URIPrefix = "/api/v1"

	URIPathClientGetUser = URIPrefix + "/some/%s"
	URIPathClientGetOrders = URIPrefix + "/some/%s"
	URIPathClientGetUserCount = URIPrefix + "/some/%s"
	URIPathClientGetOrdersWithoutParams = URIPrefix + "/some/%s"

	URIPathGetUser = URIPrefix + "/user"
	URIPathGetOrders = URIPrefix + "/orders"
	URIPathGetUserCount = URIPrefix + "/user/:id/count"
	URIPathGetOrdersWithoutParams = URIPrefix + "/orders"

	HTTPMethodGetUser = "fill me!"
	HTTPMethodGetOrders = "fill me!"
	HTTPMethodGetUserCount = "fill me!"
	HTTPMethodGetOrdersWithoutParams = "fill me!"
)
