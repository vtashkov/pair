package pair

import "fmt"

type KeyValue[TKey any, TValue any] struct {
	Key   TKey
	Value TValue
}

func (pair KeyValue[TKey, TValue]) Deconstruct() (TKey, TValue) {
	return pair.Key, pair.Value
}

func (pair KeyValue[TKey, TValue]) String() string {
	return fmt.Sprintf("{%v %v}", pair.Key, pair.Value)
}
