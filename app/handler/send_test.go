package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"

	. "app/computation"
	. "app/webhook"
)

func TestSendCalc(t *testing.T) {

}

func TestGeneratePayload(t *testing.T) {
	var testString = strings.NewReader(`{"a":"8", "b":"16"}`)
	c, _, rec := initEchoTest("/", testString, http.MethodGet)
	h, msw, msc := initMockedHandler()
	c.SetPath("/:id/computation")
	c.SetParamNames("id")
	c.SetParamValues("5632499082330112")

	var testFields = []string{"a", "b", "c"}
	elem := Webhook{
		ID: 5632499082330112,
		Fields: testFields,
		Op:     "sub",
	}

	keys := []KeyValue{
		{
			Key: "a",
			Value: 5,
		},
		{
			Key: "b",
			Value: 7,
		},
	}
	result := Computation{
		WebhookID: 7575757575757575,
		Values:    keys,
		Processed: false,
	}
	msw.On("Get", int64(5632499082330112)).Return(&elem, nil)
	msc.On("Put", &result).Return(nil)
	payload, err := h.generatePayload(c)
	fmt.Printf("%#v", payload)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"id":5632499082330112,"fields":["a","b","c"],"op":"sub"}` + "\n", rec.Body.String())
	}
}
