package ksuid

import "github.com/segmentio/ksuid"

func Gen() ksuid.KSUID {
	return ksuid.New()
}