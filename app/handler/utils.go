package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"

	. "app/computation"
	. "app/config"
)

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func publishOnTopic(payload Payload) {
	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	if res, err := TopicExec.Publish(context.Background(), &pubsub.Message{
		Data: data,
	}).Get(context.Background()); err != nil {
		panic(err)
	} else {
		log.Println(fmt.Sprintf("Published pubsub message: %s", res))
	}
}

func mapToKeyValue(mapValues map[string]interface{}) []KeyValue {
	var listValue []KeyValue
	for key, element := range mapValues {
		n := KeyValue{Key: key, Value: int64(element.(float64))}
		listValue = append(listValue, n)
	}
	return listValue
}

func (h *Handler) createComputation(idInt64 int64, listValue []KeyValue) (result Computation, err error) {
	result = Computation{
		WebhookID: idInt64,
		Values:    listValue,
		Processed: false,
	}

	err = h.ComputeStore.Put(&result)
	if err != nil {
		return result, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return result, nil
}