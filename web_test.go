package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "launchpad.net/gocheck"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//
// This file is the "driver" for the test suite. We're using gocheck.
// This file will contain helpers and general things the rest of the suite needs
//

//
// gocheck: hook into "go test"
//
func Test(t *testing.T) { TestingT(t) }

// Make a testing request
func newTestRequest(method, path string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()

	return recorder, request
}

func assertResponse(c *C, rr *httptest.ResponseRecorder, body string, code int) {
	c.Assert(strings.TrimSpace(string(rr.Body.Bytes())), Equals, body)
	c.Assert(rr.Code, Equals, code)
}

// TODO: refactor this to do proper
// TODO: change tests in error test to use this
func assertResponseT(t *testing.T, rr *httptest.ResponseRecorder, body string, code int) {
	assert.Equal(t, body, strings.TrimSpace(string(rr.Body.Bytes())))
	assert.Equal(t, code, rr.Code)
}

//
// Some default contexts and possible error handlers / actions
//
type Context struct{}

type AdminContext struct {
	*Context
}

type APIContext struct {
	*Context
}

type SiteContext struct {
	*Context
}

type TicketsContext struct {
	*AdminContext
}

func (c *Context) ErrorMiddleware(w ResponseWriter, r *Request, next NextMiddlewareFunc) {
	var x, y int
	fmt.Fprintln(w, x/y)
}

func (c *Context) ErrorHandler(w ResponseWriter, r *Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "My Error")
}

func (c *Context) ErrorHandlerSecondary(w ResponseWriter, r *Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "My Secondary Error")
}

func (c *Context) ErrorAction(w ResponseWriter, r *Request) {
	var x, y int
	fmt.Fprintln(w, x/y)
}

func (c *AdminContext) ErrorMiddleware(w ResponseWriter, r *Request, next NextMiddlewareFunc) {
	var x, y int
	fmt.Fprintln(w, x/y)
}

func (c *AdminContext) ErrorHandler(w ResponseWriter, r *Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Admin Error")
}

func (c *AdminContext) ErrorAction(w ResponseWriter, r *Request) {
	var x, y int
	fmt.Fprintln(w, x/y)
}

func (c *APIContext) ErrorHandler(w ResponseWriter, r *Request, err interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Api Error")
}

func (c *APIContext) ErrorAction(w ResponseWriter, r *Request) {
	var x, y int
	fmt.Fprintln(w, x/y)
}

func (c *TicketsContext) ErrorAction(w ResponseWriter, r *Request) {
	var x, y int
	fmt.Fprintln(w, x/y)
}
