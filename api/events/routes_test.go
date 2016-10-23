package events

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Shopify/sarama/mocks"
	"github.com/rafaeljesus/event-tracker/lib/elastic"
	"github.com/rafaeljesus/event-tracker/lib/kafka"
	"github.com/rafaeljesus/event-tracker/models"
)

const (
	requestJSON = `{"cid":1, "status":"ok", "name":"order_created", "payload": {}}`
)

func TestMain(m *testing.M) {
	beforeEach()
	code := m.Run()
	os.Exit(code)
}

func beforeEach() {
	elastic.Connect()
	kafka.Connect()
}

func TestIndex(t *testing.T) {
	event := newEvent()
	event.Create()

	e := echo.New()
	q := make(url.Values)
	q.Set("cid", strconv.Itoa(event.Cid))
	q.Set("name", event.Name)
	q.Set("status", event.Status)

	req, _ := http.NewRequest(echo.GET, "/v1/events/?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	event1 := newEvent()
	expected, _ := json.Marshal([]*models.Event{event1})

	if assert.NoError(t, Index(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestCreate(t *testing.T) {
	e := echo.New()
	producer := mocks.NewSyncProducer(t, nil)

	req, _ := http.NewRequest(echo.POST, "/v1/events", strings.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, Create(ctx)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		producer.ExpectSendMessageAndSucceed()
	}
}

func newEvent() *models.Event {
	payload := []byte(`{"amount": 5}`)
	Payload := (*json.RawMessage)(&payload)

	return &models.Event{
		Cid:     1,
		Status:  "ok",
		Name:    "order_created",
		Payload: Payload,
	}
}
