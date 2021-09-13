package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	. "app/computation"
)

func TestEqual(t *testing.T) {
	tables := []struct {
		x []string
		y []string
		res bool
	}{
		{[]string{"test1", "test2", "test3"}, []string{"test1", "test2", "test3"}, true},
		{[]string{"test1", "test2"}, []string{"test3", "test4"}, false},
		{[]string{"test1", "test2", "test3"}, []string{"test1", "test2"}, false},
	}

	for _, table := range tables {
		assert.Equal(t, table.res, Equal(table.x, table.y))
	}
}

func TestMapToKeyValue(t *testing.T) {
	data := `{"a": 12, "b": -5, "c": 42}`
	var mapTest map[string]interface{}
	err := json.Unmarshal([]byte(data), &mapTest)
	if err != nil {
		panic(err)
	}
	var result = []KeyValue {
		{
			Key: "a",
			Value: 12,
		},
		{
			Key: "b",
			Value: -5,
		},
		{
			Key: "c",
			Value: 42,
		},
	}
	assert.Equal(t, result, mapToKeyValue(mapTest))
}

func TestPublishOnTopic(t *testing.T) {

}

/*func TestCreateComputation(t *testing.T) {
	var valuesTest = []KeyValue {
		{
			Key: "a",
			Value: 12,
		},
		{
			Key: "b",
			Value: -5,
		},
		{
			Key: "c",
			Value: 42,
		},
	}
	var testID int64 = 5631671361601536
	h := HandlerTest{StoreWebhook: MockedStoreWebhook{}, ComputeStore: MockedStoreComputation{}}
	key, err := h.createComputation(testID, valuesTest)
	want := datastore.Key{
		Kind: "result",
		ID: 5631671361601536,
		Name: "",
		Parent: nil,
		Namespace: "",
	}

	if assert.NotNil(t, key) {
		assert.Nil(t, err)
		assert.Equal(t, want, key)
	}

}
*/