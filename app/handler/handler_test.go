package handler

import (
	. "app/computation"
	. "app/webhook"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
)

type (
	MockedStoreWebhook struct {
		StoreWebhook
		mock.Mock
	}

	MockedStoreComputation struct {
		mock.Mock
		StoreComputation
	}
)

var (
	getResult = `{"a": -12, "b": 42, "c": 12}`
)

func (s *MockedStoreWebhook) Get(id int64) (*Webhook, error) {
	args := s.Called(id)
	return args.Get(0).(*Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) GetAll() ([]*Webhook, error) {
	args := s.Called()
	return args.Get(0).([]*Webhook), args.Error(1)
}

func (s *MockedStoreWebhook) Put(webhook *Webhook) error {
	args := s.Called(webhook)
	webhook.ID = 345678945678
	return args.Error(0)
}

func (s *MockedStoreWebhook) PutKey(webhook *Webhook, id int64) error {
	args := s.Called(webhook, id)
	return args.Error(0)
}

func (s *MockedStoreWebhook) Delete(id int64) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *MockedStoreComputation) Get(id int64) (*Computation, error) {
	args := s.Called(id)
	return args.Get(0).(*Computation), args.Error(1)
}

func (s *MockedStoreComputation) Put(computation *Computation) error {
	args := s.Called(computation)
	return args.Error(0)
}

func initMockedHandler() (h Handler, msw *MockedStoreWebhook, msc *MockedStoreComputation){
	msw = &MockedStoreWebhook{}
	msc = &MockedStoreComputation{}
	var sw StoreWebhook = msw
	var sc StoreComputation = msc
	h = Handler{
		StoreWebhook: sw,
		ComputeStore: sc,
	}
	return
}

func initEchoTest(path string, reader io.Reader, method string) (c echo.Context, req *http.Request, rec *httptest.ResponseRecorder){
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	req = httptest.NewRequest(method, path, reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath(path)
	return
}