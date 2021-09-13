package webhook

import "cloud.google.com/go/datastore"

type Webhook struct {
	ID     int64    `json:"id" datastore:"-"`
	Fields []string `json:"fields" datastore:"fields"`
	Op     string   `json:"op" datastore:"op"`
}

func (w *Webhook) Load(ps []datastore.Property) error {
	return datastore.LoadStruct(w, ps)
}

func (w *Webhook) Save() ([]datastore.Property, error) {
	return datastore.SaveStruct(w)
}

func (w *Webhook) LoadKey(k *datastore.Key) error {
	w.ID = k.ID
	return nil
}
