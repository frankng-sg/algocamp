package main

import (
	"reflect"
	"testing"
)

func TestTimeMap_1(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("foo", "bar1", 1)
	timeMap.Set("foo", "bar2", 2)
	timeMap.Set("foo", "bar3", 3)

	tests := []struct {
		key       string
		timestamp int
		want      string
	}{
		{"foo", 1, "bar1"},
		{"foo", 2, "bar2"},
		{"foo", 3, "bar3"},
		{"foo", 4, "bar3"},
	}

	for _, test := range tests {
		got := timeMap.Get(test.key, test.timestamp)
		if got != test.want {
			t.Errorf("Get(%s, %d) = %s; want %s", test.key, test.timestamp, got, test.want)
		}
	}
}

func TestTimeMap_2(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	tests := []struct {
		ops  string
		args []interface{}
		want string
	}{
		{"get", []interface{}{"love", 5}, ""},
		{"get", []interface{}{"love", 10}, "high"},
		{"get", []interface{}{"love", 15}, "high"},
		{"get", []interface{}{"love", 20}, "low"},
		{"get", []interface{}{"love", 25}, "low"},
	}

	for _, test := range tests {
		var got string
		switch test.ops {
		case "get":
			got = timeMap.Get(test.args[0].(string), test.args[1].(int))
		}
		if got != test.want {
			t.Errorf("%s(%v) = %s; want %s", test.ops, test.args, got, test.want)
		}
	}
}

func TestTimeMap_3(t *testing.T) {
	timeMap := Constructor()

	tests := []struct {
		ops  string
		args []interface{}
		want interface{}
	}{
		{"TimeMap", []interface{}{}, nil},
		{"set", []interface{}{"foo", "bar", 1}, nil},
		{"get", []interface{}{"foo", 1}, "bar"},
		{"get", []interface{}{"foo", 3}, "bar"},
		{"set", []interface{}{"foo", "bar2", 4}, nil},
		{"get", []interface{}{"foo", 4}, "bar2"},
		{"get", []interface{}{"foo", 5}, "bar2"},
	}

	for _, test := range tests {
		switch test.ops {
		case "TimeMap":
			timeMap = Constructor()
			if timeMap.cache == nil {
				timeMap.cache = make(map[string][]TimeBasedValue)
			}
		case "set":
			timeMap.Set(test.args[0].(string), test.args[1].(string), test.args[2].(int))
		case "get":
			got := timeMap.Get(test.args[0].(string), test.args[1].(int))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Get(%s, %d) = %v; want %v", test.args[0], test.args[1], got, test.want)
			}
		}
	}
}

func TestTimeMap_4(t *testing.T) {
	timeMap := Constructor()

	tests := []struct {
		ops  string
		args []interface{}
		want interface{}
	}{
		{"TimeMap", []interface{}{}, nil},
		{"set", []interface{}{"a", "bar", 1}, nil},
		{"set", []interface{}{"x", "b", 3}, nil},
		{"get", []interface{}{"b", 3}, ""},
		{"set", []interface{}{"foo", "bar2", 4}, nil},
		{"get", []interface{}{"foo", 4}, "bar2"},
		{"get", []interface{}{"foo", 5}, "bar2"},
	}

	for _, test := range tests {
		switch test.ops {
		case "TimeMap":
			timeMap = Constructor()
			if timeMap.cache == nil {
				timeMap.cache = make(map[string][]TimeBasedValue)
			}
		case "set":
			timeMap.Set(test.args[0].(string), test.args[1].(string), test.args[2].(int))
		case "get":
			got := timeMap.Get(test.args[0].(string), test.args[1].(int))
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Get(%s, %d) = %v; want %v", test.args[0], test.args[1], got, test.want)
			}
		}
	}
}
