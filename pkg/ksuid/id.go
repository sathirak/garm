package ksuid

import (
	"github.com/sathirak/garm/repository"
	"github.com/segmentio/ksuid"
)

func Gen() ksuid.KSUID {
	for {
		id := ksuid.New()

		if repository.IsIDAvailable(id.String()) {
			return id
		}
	}
}
