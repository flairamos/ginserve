package uuid_util

import "github.com/google/uuid"

func Uuid36() string {
	uuidWithHyphen := uuid.New()
	return uuidWithHyphen.String()
}

func Uuid32() string {
	uuidWithHyphen := uuid.New()
	return uuidWithHyphen.String()[:32]
}
