package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DiogoJorge1401/golang-api/handler"
	"github.com/gin-gonic/gin"
)

func MalformedBodyMiddleware(c *gin.Context) {
	bodyBytes, err := c.GetRawData()

	if len(bodyBytes) == 0 {
		c.Next()
		return
	}

	if err != nil {
		handler.SendErrorJSONResponse(c, http.StatusBadRequest, "Unable to read request body")
		c.Abort()
		return
	}

	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		handler.SendErrorJSONResponse(c, http.StatusBadRequest, fmt.Sprintf("Malformed request body: %v", err.Error()))
		c.Abort()
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	c.Next()
}
