package models

type DailyStats struct {
	CustomerValidRequests   int64
	CustomerInvalidRequests int64
	TotalDailyRequests      int64
}
