package jwt

import (
	"encoding/pem"
	"os"

	"github.com/hotelbear/garm/internal/logger"
)

var key []byte

//openssl ecparam -name prime256v1 -genkey -noout -out garm.pem
// openssl ec -in ec_private_key.pem -text -noout

func GetKey() []byte {
	return key
}

func Initialize() {
	log := logger.Get()
	file := "garm.pem"

	pemData, err := os.ReadFile(file)
	if err != nil {
		log.Errorw("startup", "package", "jwt", "error", err.Error())
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Errorw("startup", "package", "jwt", "error", "failed to decode pem block")
		return
	}

	key = block.Bytes

	log.Infow("startup", "package", "jwt", "status", "ok")
}
