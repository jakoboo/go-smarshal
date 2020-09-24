package smarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Marshal(src ...interface{}) ([]byte, error) {
	out := map[string]interface{}{}

	for _, i := range src {
		t, err := json.Marshal(i)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(t, &out)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(out)
}

func Unmarshal(data []byte, dst ...interface{}) error {
	for _, v := range dst {
		err := json.Unmarshal(data, v)
		if err != nil {
			return fmt.Errorf("Unmarshaling data to interface: %s, %w", data, err)
		}

		vRealValue := reflect.Indirect(reflect.Indirect(reflect.ValueOf(v)))
		vRealValueInf := vRealValue.Interface()
		vRealValueZeroInf := reflect.Zero(vRealValue.Type()).Interface()

		if reflect.DeepEqual(vRealValueInf, vRealValueZeroInf) {
			vVal := reflect.ValueOf(v).Elem()
			vValType := vVal.Type()
			vVal.Set(reflect.Zero(vValType))
		}
	}

	return nil
}
