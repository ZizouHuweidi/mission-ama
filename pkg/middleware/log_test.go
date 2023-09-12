package middleware

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"

	"github.com/zizouhuweidi/mission-ama/pkg/tests"
)

func TestLogRequestID(t *testing.T) {
	ctx, _ := tests.NewContext(c.Web, "/")
	_ = tests.ExecuteMiddleware(ctx, echomw.RequestID())
	_ = tests.ExecuteMiddleware(ctx, LogRequestID())

	var buf bytes.Buffer
	ctx.Logger().SetOutput(&buf)
	ctx.Logger().Info("test")
	rID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	assert.Contains(t, buf.String(), fmt.Sprintf(`id":"%s"`, rID))
}
