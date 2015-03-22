// +build OMIT

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

var key *rsa.PrivateKey

func init() {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generate key: %v", err)
	}
	key = k
}

type Person struct {
	Name string `json:"name"`
	SSN  secret `json:"ssn"`
}

type secret string

func (s secret) MarshalJSON() ([]byte, error) {
	m, err := rsa.EncryptOAEP(crypto.SHA512.New(), rand.Reader, key.Public().(*rsa.PublicKey), []byte(s), nil)
	if err != nil {
		return nil, err
	}
	return json.Marshal(base64.StdEncoding.EncodeToString(m))
}

func (s *secret) UnmarshalJSON(data []byte) error {
	var text string
	if err := json.Unmarshal(data, &text); err != nil { // HL
		return fmt.Errorf("deocde secret string: %v", err)
	}
	cypher, err := base64.StdEncoding.DecodeString(text) // HL
	if err != nil {
		return err
	}
	raw, err := rsa.DecryptOAEP(crypto.SHA512.New(), rand.Reader, key, cypher, nil) // HL
	if err == nil {
		*s = secret(raw)
	}
	return err
}

func main() {
	p := Person{
		Name: "Francesc",
		SSN:  "123456789",
	}

	b, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		log.Fatalf("Encode person: %v", err)
	}
	fmt.Printf("%s\n", b)

	var d Person
	if err := json.Unmarshal(b, &d); err != nil {
		log.Fatalf("Decode person: %v", err)
	}
	fmt.Println(d)
}
