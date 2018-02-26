package base

import (
	"encoding/json"
	"io"

	"github.com/mitchellh/mapstructure"
)

var (
	KEY_ERROR_CODE = "error_code"
)

func DecodeInto(reader io.Reader, value interface{}, err *Error) error {
	jsonDecoder := json.NewDecoder(reader)
	source := map[string]interface{}{}

	config := mapstructure.DecoderConfig{
		TagName: "json",
	}

	jsonDecoder.Decode(&source)
	if _, ok := source[KEY_ERROR_CODE]; ok {
		config.Result = err
	} else {
		config.Result = value
	}

	decoder, e := mapstructure.NewDecoder(&config)
	if e != nil {
		return e
	}
	return decoder.Decode(source)
}
