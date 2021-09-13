package handler

import (
	. "app/computation"
	. "app/config"
	. "app/webhook"
	"cloud.google.com/go/datastore"
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StoreComputation interface {
	Get(id int64) (computation *Computation, err error)
	Put(computation *Computation) (err error)
}

type (
	DatastoreComputationStore struct {
		StoreComputation
		client *datastore.Client
	}
)

func NewDatastoreComputationStore(client *datastore.Client) (StoreComputation, error) {
	var ds StoreComputation = &DatastoreComputationStore{
		client: client,
	}
	return ds, nil
}

func (s *DatastoreComputationStore) Get(id int64) (computation *Computation, err error) {
	err = s.client.Get(context.Background(), datastore.IDKey(ComputationKind, id, nil), &computation)
	computation.ID = id
	return
}

func (s *DatastoreComputationStore) Put(computation *Computation) (err error) {
	key, err := s.client.Put(context.Background(), datastore.IncompleteKey(ComputationKind, nil), computation)
	if err != nil {
		return
	}
	computation.ID = key.ID
	return
}

type StoreWebhook interface {
	Get(id int64) (webhook *Webhook, err error)
	GetAll() (list []*Webhook, err error)
	Put(webhook *Webhook) (err error)
	PutKey(webhook *Webhook, id int64) (err error)
	Delete(id int64) (err error)
}

type (
	DatastoreWebhookStore struct {
		StoreWebhook
		client *datastore.Client
	}
)

func NewDatastoreWebhookStore(client *datastore.Client) (StoreWebhook, error) {
	var ds StoreWebhook = &DatastoreWebhookStore{
		client: client,
	}
	return ds, nil
}

func (s *DatastoreWebhookStore) Get(id int64) (webhook *Webhook, err error) {
	err = s.client.Get(context.Background(), datastore.IDKey(WebhookKind, id, nil), webhook)
	webhook.ID = id
	return
}

func (s *DatastoreWebhookStore) GetAll() (webhook []*Webhook, err error) {
	var list []*Webhook

	if keys, err := DatastoreClient.GetAll(
		context.Background(),
		datastore.NewQuery(WebhookKind),
		&list,
	); err != nil {
		return nil, err
	} else {
		for i, key := range keys {
			list[i].ID = key.ID
		}
	}
	return
}

func (s *DatastoreWebhookStore) Put(webhook *Webhook) (err error) {
	key, err := s.client.Put(context.Background(), datastore.IncompleteKey(WebhookKind, nil), webhook)
	if err != nil {
		return
	}
	webhook.ID = key.ID
	return
}

func (s *DatastoreWebhookStore) PutKey(webhook *Webhook, id int64) (err error) {
	key, err := s.client.Put(context.Background(), datastore.IDKey(ComputationKind, id,nil), webhook)
	if err != nil {
		return
	}
	webhook.ID = key.ID
	return
}

func (s *DatastoreWebhookStore) Delete(id int64) (err error) {
	err = s.client.Delete(context.Background(), datastore.IDKey(WebhookKind, id, nil))
	if err != nil {
		return
	}
	return
}

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}