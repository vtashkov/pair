package pair_test

import (
	"testing"

	"github.com/vtashkov/pair"
)

func TestCreateKeyValuePair(t *testing.T) {
	pair := pair.KeyValue[int, string]{Key: 1, Value: "test"}

	if pair.Key != 1 {
		t.Errorf("When create KeyValuePair expect the Key to be 1, got %d", pair.Key)
	}

	if pair.Value != "test" {
		t.Errorf("When create KeyValuePair expect the Value to be \"test\", got %s", pair.Value)
	}
}

func TestChangeKeyValuePair(t *testing.T) {
	pair := pair.KeyValue[int, string]{Key: 1, Value: "test"}

	if pair.Key != 1 {
		t.Errorf("When create KeyValuePair expect the Key to be 1, got %d", pair.Key)
	}

	if pair.Value != "test" {
		t.Errorf("When create KeyValuePair expect the Value to be \"test\", got %s", pair.Value)
	}

	pair.Key = 2
	pair.Value = "test2"

	if pair.Key != 2 {
		t.Errorf("When change KeyValuePair expect the Key to be 2, got %d", pair.Key)
	}

	if pair.Value != "test2" {
		t.Errorf("When change KeyValuePair expect the Value to be \"test2\", got %s", pair.Value)
	}
}

func TestDeconstructKeyValuePair(t *testing.T) {
	pair := pair.KeyValue[string, string]{Key: "key", Value: "value"}

	key, value := pair.Deconstruct()

	if key != "key" {
		t.Errorf("When deconstruct KeyValuePair expect the Key to be \"key\", got %s", key)
	}

	if value != "value" {
		t.Errorf("When deconstruct KeyValuePair expect the Value to be \"value\", got %s", pair.Value)
	}
}

func TestStringKeyValuePairOfStrings(t *testing.T) {
	pair := pair.KeyValue[string, string]{Key: "key", Value: "value"}

	want := "{key value}"
	got := pair.String()

	if want != got {
		t.Errorf("KeyValuePair.String() expect to return %s, got %s", want, got)
	}
}
