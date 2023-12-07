package circuitbreaker

type RequestsResults struct {
	ConsecutiveSuccessCount uint32
	ConsecutiveFailCount    uint32
	MaxTries                uint32
}

type CircuitBreaker struct {
	ServiceName string
	Mode        uint8 // 0 - Closed, 1 - Half-Open, 2 - Open
	ReqRes      RequestsResults
}
