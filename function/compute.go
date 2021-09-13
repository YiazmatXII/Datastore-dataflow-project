package p

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type (
	PubSubEvent struct {
		Subscription string `json:"subscription"`
		Message      struct {
			Attributes  map[string]string `json:"attributes"`
			Data        string            `json:"data"`
			PublishTime time.Time         `json:"publish_time"`
		} `json:"message"`
	}

	Payload struct {
		ComputationID 	int64      `json:"computation_id"`
		WebhookID 		int64      `json:"webhook_id"`
		Values        	[]KeyValue `json:"values"`
		Fields        	[]string   `json:"fields"`
		Op            	string   	`json:"op"`
	}

	Computation struct {
		ID        int64      `json:"id"`
		WebhookID int64      `json:"webhook_id"`
		Values    []KeyValue `json:"values"`
		Result    int64      `json:"result"`
	}

	KeyValue struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
	}

	ComputationFunction func(a int64, b int64) (res int64)
)

const (
	ProjectId               = "pierre-test-321108"
	PubsubTopicResult 		= "Result"
)

var (
	PubsubClient *pubsub.Client
	PubsubTopic  *pubsub.Topic

	ComputationFunctions map[string]ComputationFunction
)

func init() {
	ComputationFunctions = map[string]ComputationFunction{
		"add": func(a int64, b int64) (res int64) {
			return a + b
		},
		"sub": func(a int64, b int64) (res int64) {
			return a - b
		},
	}

	var err error
	PubsubClient, err = pubsub.NewClient(context.Background(), ProjectId)
	if err != nil {
		panic(err)
	}
	PubsubTopic = PubsubClient.Topic(PubsubTopicResult)
}

func DoCompute(values map[string]int64, fields []string, f ComputationFunction) (res int64) {
	res = values[fields[0]]
	for _, key := range fields[1:] {
		if value, ok := values[key]; ok {
			res = f(res, value)
		} else {
			panic("Error: Empty value in Computation.")
		}
	}
	return res
}

func convertToMap(listValue []KeyValue) map[string]int64 {
	m := make(map[string]int64)
	for _, structure := range listValue {
		m[structure.Key] = int64(structure.Value)
	}
	return m
}

func createComputation(result int64, p Payload) (computation Computation){
	return Computation{
		ID:        p.ComputationID,
		WebhookID: p.WebhookID,
		Values:    p.Values,
		Result:    result,
	}
}

func publishResult(computation *Computation) {
	data, err := json.Marshal(computation)
	if err != nil {
		panic(err)
	}
	if res, err := PubsubTopic.Publish(context.Background(), &pubsub.Message{
		Data: data,
	}).Get(context.Background()); err != nil {
		panic(err)
	} else {
		log.Println(fmt.Sprintf("Published pubsub message: %s", res))
	}
}

func Compute(w http.ResponseWriter, r *http.Request) {
	var e PubSubEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		panic(err)
	}

	log.Printf("event: %#v\n", e)

	data, err := base64.StdEncoding.DecodeString(e.Message.Data)
	if err != nil {
		panic(err)
	}

	var p Payload
	err = json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}

	log.Printf("payload: %#v\n", p)
	if f, ok := ComputationFunctions[p.Op]; ok {
		result := DoCompute(convertToMap(p.Values), p.Fields, f)
		log.Printf("result: %#v\n", result)
		_, _ = fmt.Fprintln(w, "OK")
		response := createComputation(result, p)
		publishResult(&response)
	} else {
		panic("Error: Unknown Operator.")
	}
}
