package jwt

import (
	_ "embed"
	"encoding/pem"

	"github.com/hotelbear/garm/internal/logger"
)

var key []byte

//openssl ecparam -name prime256v1 -genkey -noout -out garm.pem
// openssl ec -in ec_private_key.pem -text -noout

func GetKey() []byte {
	return key
}

func Initialize(pemData []byte) {
	log := logger.Get()

	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Errorw("startup", "package", "jwt", "error", "failed to decode pem block")
		return
	}

	key = block.Bytes

	log.Infow("startup", "package", "jwt", "status", "ok")
}
