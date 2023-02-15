package shared

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	response "github.com/da-moon/northern-labs-interview/sdk/api/response"
	stacktrace "github.com/palantir/stacktrace"
	redis "github.com/redis/go-redis/v9"
)

// PreflightErrorResponse represents the error response
// of the server when Redis is not available
//
//go:generate gomodifytags -override -file $GOFILE -struct PreflightErrorResponse -add-tags json,yaml,mapstructure -w -transform snakecase
type PreflightErrorResponse struct {
	Code    int    `json:"code" yaml:"code" mapstructure:"code"`
	Message string `json:"message" yaml:"message" mapstructure:"message"`
}

// RedisClient returns a redis client for handlers under this API group
// Note that a mutex lock is not needed as GetRedisOptions locks the state
// In case redis is not available, it returns the common error response
func RedisClient(ctx context.Context, w http.ResponseWriter, r *http.Request) *redis.Client {
	var (
		opts = Group.GetRedisOptions()
		err  error
	)
	defer func() {
		if err != nil {
			msg := fmt.Sprintf("%s API group (%s) request preflight check failed", Name, GroupPrefix)
			response.LogErrorResponse(r, err, msg)
			return
		}
		response.LogSuccessfulResponse(r, "preflight check successful")
		return
	}()
	if opts == nil {
		err = stacktrace.NewError("Redis client option is nil")
		HandlerFn(w, r)
		return nil
	}
	client := redis.NewClient(opts)
	if client == nil {
		err = stacktrace.NewError("Redis client is nil")
		HandlerFn(w, r)
		return nil
	}
	status := client.Ping(ctx)
	err = status.Err()
	if err != nil {
		fmt.Println("*** ping err", err)
		HandlerFn(w, r)
		return nil
	}
	return client
}

// HandlerFn function handles incoming HTTP request
// it satisfies golang's stdlib
// request handler interface (http.HandlerFunc)
var HandlerFn = func(w http.ResponseWriter, r *http.Request) { //nolint:gochecknoglobals //this function is scoped only to this package
	body, err := io.ReadAll(r.Body)
	defer func() {
		if err != nil {
			err = stacktrace.NewErrorWithCode(stacktrace.ErrorCode(http.StatusInternalServerError), err.Error())
			response.LogErrorResponse(r, err, "")
			return
		}
		if body != nil && len(body) > 0 {
			compact := &bytes.Buffer{}
			bodyEnc, _ := json.Marshal(body) //nolint:gosec // failure here does not matter
			// NOTE: this is just to stay safe from nil pointer dereference
			if bodyEnc != nil {
				json.Compact(compact, bodyEnc) //nolint:gosec // failure
				err := stacktrace.NewErrorWithCode(stacktrace.ErrorCode(http.StatusInternalServerError), compact.String())
				response.LogErrorResponse(r, err, "")
				return
			}
		}
		response.LogSuccessfulResponse(r, string(body))
		return
	}()
	code := http.StatusOK
	payload := &PreflightErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Redis is offline",
	}
	response.WriteJSON(
		w,
		r,
		code,
		make(map[string]string, 0),
		payload,
	)
}
