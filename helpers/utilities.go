package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/fatih/structs"
	"github.com/golang/gddo/httputil/header"
	"github.com/twinj/uuid"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return errors.New(msg)
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

	return err
}

func ModelToInterface(model interface{}) map[string]interface{} {
	s := structs.New(model)
	_map := s.Map()
	return _map
}

func GenUUID() string {
	return uuid.NewV4().String()
}

