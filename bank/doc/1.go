package doc

import (
	"fmt"
	"net/http"
	"testing"
)

type reqRespTuple struct {
	requestExpected *http.Request
	response        *http.Response
	netError        error
}

type httpClientMock struct {
	t      *testing.T
	tuples []reqRespTuple
}

func (c *httpClientMock) Do(req *http.Request) (resp *http.Response, err error) {
	if len(c.tuples) == 0 {
		err = fmt.Errorf("no more requests are expected, but got '%#v'", *req)
		c.t.Fatal(err)
		return
	}

	tuple := c.tuples[0]
	c.tuples = c.tuples[1:]

	var reqExpected *http.Request
	var netErr error
	reqExpected, resp, netErr = tuple.requestExpected, tuple.response, tuple.netError

	if reqExpected.Method != req.Method || reqExpected.URL.String() != req.URL.String() {
		err = fmt.Errorf("wrong request: want '%v %v', got '%v %v'", reqExpected.Method, reqExpected.URL.String(), req.Method, req.URL.String())
		c.t.Fatal(err)
		return
	}

	if req.Body != nil {
		err = fmt.Errorf("request body must be nil")
		c.t.Fatal(err)
		return
	}

	err = netErr
	return
}
