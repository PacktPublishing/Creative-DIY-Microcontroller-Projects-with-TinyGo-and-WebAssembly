package login

import "time"

type UserInfo struct {
	LoggedIn   bool
	UserName   string
	LoggedInAt time.Time
}
