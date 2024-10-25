package ksuid

import (
	"github.com/sathirak/garm/pkg/logger"
	"github.com/sathirak/garm/repository"
	"github.com/segmentio/ksuid"
)

func Gen() ksuid.KSUID {
	for {
		id, err := safeNewKSUID()
		if err != nil {
			logger.Get().Errorw("onprocess", "package", "ksuid", "error", err.Error())
			continue
		}

		isAvailable, err := repository.IsIDAvailable(id.String())

		if err != nil {
			logger.Get().Errorw("onprocess", "package", "ksuid", "error", err.Error())
			continue
		}

		if isAvailable {
			return id
		}

	}
}

func safeNewKSUID() (ksuid.KSUID, error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Get().Infow("onprocess", "package", "ksuid", "msg", "recovered from panic")
		}
	}()

	id := ksuid.New()

	return id, nil
}
