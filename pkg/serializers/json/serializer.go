package json

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	validator "github.com/ressley/test_task_go_middle/pkg/validators"
)

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Decode[T any](object *T, bytes []byte) error {
	if err := json.Unmarshal(bytes, object); err != nil {
		return errors.Wrap(err, "serializer.Decode")
	}
	return nil
}

func Encode[T any](data *T) ([]byte, error) {
	rawMsg, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Encode")
	}
	return rawMsg, nil
}

func FromBody[T any](object *T, body io.ReadCloser) any {
	requestBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := Decode(object, requestBody); err != nil {
		return err
	}

	return validator.ValidateStruct(object)
}
