package util

import (
	"github.com/satori/go.uuid"
	"fmt"
	"log"
	"strings"
)

func GenUUID() (string, error){
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(fmt.Sprintf("Uuid went wrong: %s", err))
		return "", err
	}
	str, err := id.Value()
	if err != nil {
		log.Fatal(fmt.Sprintf("Uuid went wrong: %s", str))
		return "", err
	}
	if strValue, ok := str.(string); ok {
		return strValue, err
	}
	return "", err
}

func GenUUID32() (string, error){
	strUuid, err := GenUUID()
	if err != nil {
		return "", err
	}
	strUuid = strings.Replace(strUuid, "-", "", -1)
	return strUuid, err
}
