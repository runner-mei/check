package check

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/three-plus-three/modules/as"
)

func init() {

	strEquals := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("=", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := as.String(value)
			if err != nil {
				return false, ErrActualType("=", "string", value)
			}
			return actualValue == exceptedValue, nil
		}), nil
	})
	AddCheckFunc("=", "string", strEquals)
	AddCheckFunc("equals", "string", strEquals)

	strNotEquals := CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("!=", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("!=", "string", value)
			}
			return actualValue != exceptedValue, nil
		}), nil
	})
	AddCheckFunc("!=", "string", strNotEquals)
	AddCheckFunc("not_equals", "string", strNotEquals)

	AddCheckFunc("contains", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("contains", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("contains", "string", value)
			}
			return strings.Contains(actualValue, exceptedValue), nil
		}), nil
	}))
	AddCheckFunc("not_contains", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_contains", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_contains", "string", value)
			}
			return !strings.Contains(actualValue, exceptedValue), nil
		}), nil
	}))

	AddCheckFunc("contains", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("contains", "string", argValue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("contains", "string", value)
			}
			return strings.Contains(actualValue, exceptedValue), nil
		}), nil
	}))

	AddCheckFunc("contains_with_ignore_case", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("contains_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("contains_with_ignore_case", "string", value)
			}
			if strings.Contains(actualValue, exceptedValue) {
				return true, nil
			}
			return strings.Contains(strings.ToLower(actualValue), exceptedValue), nil
		}), nil
	}))
	AddCheckFunc("not_contains_with_ignore_case", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_contains_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_contains_with_ignore_case", "string", value)
			}
			if strings.Contains(actualValue, exceptedValue) {
				return false, nil
			}
			return !strings.Contains(strings.ToLower(actualValue), exceptedValue), nil
		}), nil
	}))

	AddCheckFunc("equals_with_ignore_case", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("equals_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("equals_with_ignore_case", "string", value)
			}
			if actualValue == exceptedValue {
				return true, nil
			}
			return strings.ToLower(actualValue) == exceptedValue, nil
		}), nil
	}))
	AddCheckFunc("not_equals_with_ignore_case", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedValue, err := toString(argValue)
		if err != nil {
			return nil, ErrArgumentType("not_equals_with_ignore_case", "string", argValue)
		}
		exceptedValue = strings.ToLower(exceptedValue)
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := toString(value)
			if err != nil {
				return false, ErrActualType("not_equals_with_ignore_case", "string", value)
			}
			if actualValue == exceptedValue {
				return false, nil
			}
			return strings.ToLower(actualValue) != exceptedValue, nil
		}), nil
	}))

	AddCheckFunc("in", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := as.Strings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "stringArray", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
			exceptedArray = append(exceptedArray, svalue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := as.String(value)
			if err != nil {
				return false, ErrActualType("in", "string", value)
			}
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					return true, nil
				}
			}
			return false, nil
		}), nil
	}))

	AddCheckFunc("nin", "string", CheckFactoryFunc(func(argValue interface{}) (Checker, error) {
		exceptedArray, err := as.Strings(argValue)
		if err != nil {
			svalue, ok := argValue.(string)
			if !ok {
				return nil, ErrArgumentType("in", "stringArray", argValue)
			}
			exceptedArray = strings.Split(svalue, ",")
			exceptedArray = append(exceptedArray, svalue)
		}
		return CheckFunc(func(value interface{}) (bool, error) {
			actualValue, err := as.String(value)
			if err != nil {
				return false, ErrActualType("nin", "string", value)
			}
			found := false
			for idx := range exceptedArray {
				if actualValue == exceptedArray[idx] {
					found = true
					break
				}
			}
			return !found, nil
		}), nil
	}))

}

func toString(value interface{}) (string, error) {
	if nil == value {
		return "", ErrValueNull
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case json.Number:
		return v.String(), nil
	case *json.Number:
		return v.String(), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'e', -1, 64), nil
	case float64:
		return strconv.FormatFloat(v, 'e', -1, 64), nil
	case bool:
		if v {
			return "true", nil
		} else {
			return "false", nil
		}
	}

	return "", errType(value, "string")
}

func toStrings(argValue interface{}) ([]string, error) {
	if argValue == nil {
		return nil, ErrValueNull
	}
	if ss, ok := argValue.([]string); ok {
		return ss, nil
	}

	rv := reflect.ValueOf(argValue)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Slice {
		return nil, errType(argValue, "string array")
	}
	aLen := rv.Len()
	results := make([]string, 0, aLen)
	for i := 0; i < aLen; i++ {
		v := rv.Index(i)

		if v.IsNil() {
			continue
		}
		s, e := toString(v)
		if e != nil {
			return nil, e
		}
		results = append(results, s)
	}
	return results, nil
}