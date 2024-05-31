package helpers

import (
	"encoding/json"
	"errors"
	_ "github.com/pquerna/ffjson/ffjson"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type Envelope map[string]any

func WriteJson(w http.ResponseWriter, status int, data Envelope, lg zerolog.Logger) error{
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil{
		return err
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func ReadJson(w http.ResponseWriter, r *http.Request, dst any) error{
	maxBytes := 1_048_576 // set max bytes to protect web server
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() //cannot pass in json fields that aren't defined in the interface

	if err := dec.Decode(dst); err != nil{
		//TODO alex edwards lets go further chapter 4
		return err
	}

	err := dec.Decode(&struct {}{})
	if err != io.EOF{
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}
//todo test - check each endpoint for incorrect requests type
