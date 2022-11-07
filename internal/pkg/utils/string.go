package utils

import (
	"encoding/json"
	"github.com/bilgehannal/harbctl/internal/pkg/errors"
)

func StringSliceContains(slice []string, element string) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

func PrettyPrintJson(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		errors.FatalPanic("Object cannot be shown as json!", err)
	}
	return string(s)
}
