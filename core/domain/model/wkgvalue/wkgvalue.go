package wkgvalue

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model/wkgkey"
)

type WorkingValue struct {
	Key          wkgkey.WorkingKey
	InitialValue int64
	Value        int64
	Created      time.Time
	Modified     time.Time
}
