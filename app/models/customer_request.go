package models

type CustomerRequest struct {
	CustomerID uint   `json:"customerID" binding:"required"`
	TagID      int    `json:"tagID" binding:"required"`
	UserID     string `json:"userID" binding:"required"`
	RemoteIP   string `json:"remoteIP" binding:"required"`
	Timestamp  int64  `json:"timestamp" binding:"required"`
}
