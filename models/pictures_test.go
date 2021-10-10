package models

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name string
		from string
		to   string
		err  error
	}{
		{
			name: "Correct dates",
			from: "2020-01-02",
			to:   "2020-01-05",
			err:  nil,
		},
		{
			name: "End before start",
			from: "2020-02-06",
			to:   "2020-01-03",
			err:  ErrEndDateBefore,
		},
		{
			name: "In the future",
			from: "2024-01-05",
			to:   "2010-02-01",
			err:  ErrDateInFuture,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request, _ = http.NewRequest("GET", "/", nil)
			q := ctx.Request.URL.Query()
			q.Add("from", test.from)
			q.Add("to", test.to)
			ctx.Request.URL.RawQuery = q.Encode()
			var req Date
			err := ctx.ShouldBind(&req)
			assert.NoError(t, err, "Should bind should not return error")
			err = req.ValidateDate()
			assert.Equal(t, test.err, err, "Expected error: %s, actual error: %s", test.err, err)
		})
	}
}
