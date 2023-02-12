package core_test

// import (
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	assert "github.com/stretchr/testify/assert"
// 	core "github.com/da-moon/northern-labs-interview/api/core"
// 	handler "github.com/da-moon/northern-labs-interview/sdk/api/handler"
// 	logger "github.com/da-moon/northern-labs-interview/internal/logger"
// )

// func TestRouterFromConfig(t *testing.T) {
// 	assert := assert.New(t)
// 	log := logger.DefaultWrappedLogger(string(logger.ErrorLevel))
// 	conf, err := core.DefaultConfig(log)
// 	assert.NoError(err)
// 	assert.NotEmpty(conf)
// 	expected := "hello http world"
// 	prefix := "/v1"
// 	conf.AppendRoute(handler.Route{
// 		PathPrefix: "/",
// 		Name:       "foo",
// 		Method:     http.MethodGet,
// 		Path:       "/foo",
// 		Queries:    []string{},
// 		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte(expected))
// 		},
// 	})
// 	router := conf.Router(prefix)
// 	s := httptest.NewServer(router)
// 	defer s.Close()
// 	resp, err := s.Client().Get(s.URL + prefix + "/foo")
// 	assert.NoError(err)
// 	assert.NotEmpty(resp)
// 	assert.Equal(http.StatusOK, resp.StatusCode)
// 	body, err := ioutil.ReadAll(resp.Body)
// 	assert.NoError(err)
// 	assert.NotEmpty(body)
// 	assert.Equal(expected, string(body))
// }
