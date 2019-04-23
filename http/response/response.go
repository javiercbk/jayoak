package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const successMessage = "success"

// filled by compiler flag -X http.response.jayOakVersion=value
var jayOakVersion string

// Status is the status of the response
type Status struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Version string `json:"version"`
}

// ServiceResponse is a generic service response
type ServiceResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// NewSuccessResponseWithCode sends a successful response with code
func NewSuccessResponseWithCode(c *gin.Context, code int, data interface{}) {
	resp := ServiceResponse{
		Status: Status{
			Error:   false,
			Code:    code,
			Message: successMessage,
			Version: jayOakVersion,
		},
	}
	if data != nil {
		resp.Data = data
	}
	c.JSON(code, resp)
}

// NewSuccessResponse sends a successful response
func NewSuccessResponse(c *gin.Context, data interface{}) {
	NewSuccessResponseWithCode(c, http.StatusOK, data)
}

// NewErrorResponse sends an error response
func NewErrorResponse(c *gin.Context, code int, message string) {
	resp := ServiceResponse{
		Status: Status{
			Error:   true,
			Code:    code,
			Message: message,
			Version: jayOakVersion,
		},
	}
	if resp.Status.Message == "" {
		resp.Status.Message = http.StatusText(code)
	}
	c.AbortWithStatusJSON(code, resp)
}

// NewNotFoundResponse sends a not found response
func NewNotFoundResponse(c *gin.Context) {
	code := http.StatusNotFound
	resp := ServiceResponse{
		Status: Status{
			Error:   false,
			Code:    code,
			Message: http.StatusText(http.StatusNotFound),
			Version: jayOakVersion,
		},
	}
	c.AbortWithStatusJSON(code, resp)
}
