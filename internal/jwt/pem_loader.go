package jwt

import (
	"encoding/pem"
	"os"

	"github.com/sathirak/garm/pkg/logger"
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
		log.Errorf("error reading .pem file: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Errorf("failed to parse .pem block")
		return
	}

	key = block.Bytes

	log.Infof("garm.pem loaded")
}
