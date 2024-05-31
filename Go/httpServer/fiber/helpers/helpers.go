package helpers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type Envelope map[string]any

func WriteJson(c *fiber.Ctx, status int, data any, lg zerolog.Logger) error{
	//defer time_completion.LogTimer(&lg, fmt.Sprintf("Write Json"))()

	//js, err := json.MarshalIndent(data, "", "  ")
	js, err := json.Marshal(data)
	if err != nil{
		return err
	}
	js = append(js, '\n')

	c.Set("Content-Type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")
	//c.Set("Access-Control-Allow-Origin", fmt.Sprintf("%s:%s", c.IP(), c.Port()))
	c.Status(status)
	err = c.Send(js)
	if err != nil {
		return err
	}
	return nil
}

//func ReadJson(c *fiber.Ctx, req *http.Request, dst any) error{
//
//	maxBytes := 1_048_576 // set max bytes to protect web server
//
//	err := c.Request().ReadLimitBody( ,int64(maxBytes))
//	if err != nil {
//		return err
//	}
//	req.Body = http.MaxBytesReader(w , req.Body, int64(maxBytes))
//
//	dec := json.NewDecoder(req.Body)
//	dec.DisallowUnknownFields() //cannot pass in json fields that aren't defined in the interface
//
//	if err := dec.Decode(dst); err != nil{
//		//TODO alex edwards lets go further chapter 4
//		return err
//	}
//
//	err := dec.Decode(&struct {}{})
//	if err != io.EOF{
//		return errors.New("body must only contain a single JSON object")
//	}
//
//	return nil
//}

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

func ReadFile(){}
func WriteFile(){}


