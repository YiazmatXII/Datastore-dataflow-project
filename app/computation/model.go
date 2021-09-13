package computation

type (
	Computation struct {
		ID        int64      `json:"id" datastore:"-"`
		WebhookID int64      `json:"webhook_id"`
		Values    []KeyValue `json:"values"`
		Result    int64      `json:"result"`
		Processed bool		 `json:"processed"`
	}

	KeyValue struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	}

)
