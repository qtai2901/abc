package v2board

type UserTraffic struct {
	UID   int `json:"user_id"`
	Upload   int64  `json:"u"`
	Download int64  `json:"d"`
	Ip       string `json:"ip"`
}
type OnlineUser struct {
	UID int    `json:"user_id"`
	IP  string `json:"ip"`
}
// UserResponse is the response of user
type UserResponse struct {
	ID          int     `json:"id"`
	Passwd      string  `json:"passwd"`
	SpeedLimit  float64 `json:"nodeSpeedlimit"`
	DeviceLimit int     `json:"nodeConnector"`
}
