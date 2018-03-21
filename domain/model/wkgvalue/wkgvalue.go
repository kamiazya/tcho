package wkgvalue

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model/wkgkey"
)

type WorkingValue struct {
	Key          wkgkey.WorkingKey
	InitialValue int64
	Value        int64
	Created      time.Time
	Modified     time.Time
}
