// This package is responsible for loading our public key as
// mentioned in our properties file in DER format.

package certsec

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"ardyngolang/src/env"
)

//-------------------------------------------------------------

var publicKey []byte

//-------------------------------------------------------------

func InitSecurity() bool {

	raw, err := ioutil.ReadFile(env.Config.PublicKeyFile)

	if err != nil {

		log.Panicln("Failed to read certificate " + err.Error())

		return false

	}

	pubKey, err := x509.ParsePKIXPublicKey(raw)

	if pubKey == nil {

		log.Fatalln("Error parsing the public key certificate: ", err)

	}

	publicKeyDer, err := x509.MarshalPKIXPublicKey(pubKey)

	if err != nil {

		log.Fatalln("Error marshaling the public key certificate: ", err)

	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}

	publicKey = pem.EncodeToMemory(&publicKeyBlock)

	return true

}

//-------------------------------------------------------------

func GetPublicKey() []byte {

	return publicKey

}

//-------------------------------------------------------------
