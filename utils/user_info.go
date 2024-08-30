package utils

import (
	"context"
	"errors"
)

type userKeyType int

var userKey userKeyType

type UserInfo struct {
	Id uint
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	if !ok {
		return nil, errors.New("get user info wrong")
	} else {
		return u, nil
	}
}
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}
