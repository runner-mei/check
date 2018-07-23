package check

import (
	"encoding/json"
	"testing"
)

type check_data struct {
	t        string
	operator string
	operant  interface{}
	value    interface{}

	excepted_error  string
	excepted_status bool
	excepted_value  interface{}
}

func TestChecker(t *testing.T) {
	var all_check_data []check_data

	n11 := json.Number("11")
	n12 := json.Number("12")
	n13 := json.Number("13")
	for _, class := range []string{"integer", "biginteger", "decimal"} {
		for _, v := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
			for _, operant := range []interface{}{n11, &n11, "11", []byte("11"), uint(11), uint8(11), uint16(11), uint32(11), uint64(11), int(11), int8(11), int16(11), int32(11), int64(11), float32(11), float64(11)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: ">", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: ">", operant: operant, value: v, excepted_status: false})
			}
			for _, operant := range []interface{}{n11, &n11, "11", []byte("11"), uint(11), uint8(11), uint16(11), uint32(11), uint64(11), int(11), int8(11), int16(11), int32(11), int64(11), float32(11), float64(11)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: ">=", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: ">=", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n13, &n13, "13", []byte("13"), uint(13), uint8(13), uint16(13), uint32(13), uint64(13), int(13), int8(13), int16(13), int32(13), int64(13), float32(13), float64(13)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: ">=", operant: operant, value: v, excepted_status: false})
			}
			for _, operant := range []interface{}{n13, &n13, "13", []byte("13"), uint(13), uint8(13), uint16(13), uint32(13), uint64(13), int(13), int8(13), int16(13), int32(13), int64(13), float32(13), float64(13)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "<", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "<", operant: operant, value: v, excepted_status: false})
			}
			for _, operant := range []interface{}{n11, &n11, "11", []byte("11"), uint(11), uint8(11), uint16(11), uint32(11), uint64(11), int(11), int8(11), int16(11), int32(11), int64(11), float32(11), float64(11)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "<=", operant: operant, value: v, excepted_status: false})
			}
			for _, operant := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "<=", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n13, &n13, "13", []byte("13"), uint(13), uint8(13), uint16(13), uint32(13), uint64(13), int(13), int8(13), int16(13), int32(13), int64(13), float32(13), float64(13)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "<=", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n11, &n11, "11", []byte("11"), uint(11), uint8(11), uint16(11), uint32(11), uint64(11), int(11), int8(11), int16(11), int32(11), int64(11), float32(11), float64(11)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "=", operant: operant, value: v, excepted_status: false})
				all_check_data = append(all_check_data, check_data{t: class, operator: "!=", operant: operant, value: v, excepted_status: true})
			}
			for _, operant := range []interface{}{n12, &n12, "12", []byte("12"), uint(12), uint8(12), uint16(12), uint32(12), uint64(12), int(12), int8(12), int16(12), int32(12), int64(12), float32(12), float64(12)} {
				all_check_data = append(all_check_data, check_data{t: class, operator: "=", operant: operant, value: v, excepted_status: true})
				all_check_data = append(all_check_data, check_data{t: class, operator: "!=", operant: operant, value: v, excepted_status: false})
			}
		}
	}

	for _, class := range []string{"ipAddress", "string"} {
		for _, test := range []check_data{
			{t: class, operator: "=", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: true},
			{t: class, operator: "==", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: true},
			{t: class, operator: "!=", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: true},
			{t: class, operator: "<>", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: true},

			{t: class, operator: "=", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: false},
			{t: class, operator: "==", operant: "192.168.1.1", value: "192.168.1.3", excepted_status: false},
			{t: class, operator: "!=", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: false},
			{t: class, operator: "<>", operant: "192.168.1.1", value: "192.168.1.1", excepted_status: false},
		} {
			all_check_data = append(all_check_data, test)
		}
	}

	for _, class := range []string{"physicalAddress"} {
		for _, test := range []check_data{
			{t: class, operator: "=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: true},
			{t: class, operator: "==", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: true},
			{t: class, operator: "!=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: true},
			{t: class, operator: "<>", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: true},

			{t: class, operator: "=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: false},
			{t: class, operator: "==", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:03", excepted_status: false},
			{t: class, operator: "!=", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: false},
			{t: class, operator: "<>", operant: "19:16:11:ab:01:01", value: "19:16:11:ab:01:01", excepted_status: false},
		} {
			all_check_data = append(all_check_data, test)
		}
	}
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "true", value: "true", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "true", value: "true", excepted_status: true})

	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "false", value: "false", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "false", value: "false", excepted_status: true})

	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "true", value: "false", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "true", value: "false", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "=", operant: "true", value: "false", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "==", operant: "true", value: "false", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "true", value: "true", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "true", value: "true", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "!=", operant: "false", value: "false", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "boolean", operator: "<>", operant: "false", value: "false", excepted_status: false})

	all_check_data = append(all_check_data, check_data{t: "string", operator: "equals", operant: "a", value: "a", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_equals", operant: "a", value: "a", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "contains", operant: "a", value: "abc", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_contains", operant: "a", value: "abc", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "equals_with_ignore_case", operant: "a", value: "A", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_equals_with_ignore_case", operant: "a", value: "A", excepted_status: false})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "contains_with_ignore_case", operant: "a", value: "Abc", excepted_status: true})
	all_check_data = append(all_check_data, check_data{t: "string", operator: "not_contains_with_ignore_case", operant: "a", value: "Abc", excepted_status: false})

	for _, test := range all_check_data {
		check, e := MakeChecker(test.t, test.operator, test.operant)
		if nil != e {
			if 0 == len(test.excepted_error) {
				t.Errorf("[%#v] failed, %v", test, e)
			} else if test.excepted_error != e.Error() {
				t.Errorf("[%#v] failed, excepted is '%v', actual is '%v'", test, test.excepted_error, e)
			}
			continue
		}

		status, e := check.Check(test.value)
		if nil != e {
			if 0 == len(test.excepted_error) {
				t.Errorf("[%#v] failed, %v", test, e)
			} else if test.excepted_error != e.Error() {
				t.Errorf("[%#v] failed, excepted error is '%v', actual error is '%v'", test, test.excepted_error, e)
			}
			continue
		}
		if status != test.excepted_status {
			t.Errorf("[%#v] failed, excepted status is %v, actual status is %v", test, test.excepted_status, status)
		}
	}
}
