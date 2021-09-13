package handler

import (
	. "app/computation"
)

type (
	CreateInfoDto struct {
		Fields []string `json:"fields" validate:"required,min=2"`
		Op     string   `json:"op" validate:"required,oneof=add sub"`
	}

	GetInfoDto struct {
		ID int64 `param:"id" validate:"required"`
	}

	Payload struct {
		ComputationID 	int64      `json:"computation_id"`
		WebhookID 		int64      `json:"webhook_id"`
		Values        	[]KeyValue `json:"values"`
		Fields        	[]string   `json:"fields"`
		Op            	string     `json:"op"`
	}

	GetResultDto struct {
		ComputationID int64 `param:"computation_id" validate:"required"`
	}

	Handler struct {
		StoreWebhook	StoreWebhook
		ComputeStore 	StoreComputation
	}
)

