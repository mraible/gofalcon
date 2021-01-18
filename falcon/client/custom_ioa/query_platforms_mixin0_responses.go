// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryPlatformsMixin0Reader is a Reader for the QueryPlatformsMixin0 structure.
type QueryPlatformsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPlatformsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPlatformsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryPlatformsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryPlatformsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryPlatformsMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryPlatformsMixin0OK creates a QueryPlatformsMixin0OK with default headers values
func NewQueryPlatformsMixin0OK() *QueryPlatformsMixin0OK {
	return &QueryPlatformsMixin0OK{}
}

/*QueryPlatformsMixin0OK handles this case with default header values.

OK
*/
type QueryPlatformsMixin0OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryPlatformsMixin0OK) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/platforms/v1][%d] queryPlatformsMixin0OK  %+v", 200, o.Payload)
}

func (o *QueryPlatformsMixin0OK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPlatformsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPlatformsMixin0Forbidden creates a QueryPlatformsMixin0Forbidden with default headers values
func NewQueryPlatformsMixin0Forbidden() *QueryPlatformsMixin0Forbidden {
	return &QueryPlatformsMixin0Forbidden{}
}

/*QueryPlatformsMixin0Forbidden handles this case with default header values.

Forbidden
*/
type QueryPlatformsMixin0Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryPlatformsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/platforms/v1][%d] queryPlatformsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *QueryPlatformsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPlatformsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPlatformsMixin0TooManyRequests creates a QueryPlatformsMixin0TooManyRequests with default headers values
func NewQueryPlatformsMixin0TooManyRequests() *QueryPlatformsMixin0TooManyRequests {
	return &QueryPlatformsMixin0TooManyRequests{}
}

/*QueryPlatformsMixin0TooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryPlatformsMixin0TooManyRequests struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
	/*Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryPlatformsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/platforms/v1][%d] queryPlatformsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPlatformsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPlatformsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-RetryAfter
	xRateLimitRetryAfter, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-RetryAfter"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", response.GetHeader("X-RateLimit-RetryAfter"))
	}
	o.XRateLimitRetryAfter = xRateLimitRetryAfter

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPlatformsMixin0Default creates a QueryPlatformsMixin0Default with default headers values
func NewQueryPlatformsMixin0Default(code int) *QueryPlatformsMixin0Default {
	return &QueryPlatformsMixin0Default{
		_statusCode: code,
	}
}

/*QueryPlatformsMixin0Default handles this case with default header values.

OK
*/
type QueryPlatformsMixin0Default struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query platforms mixin0 default response
func (o *QueryPlatformsMixin0Default) Code() int {
	return o._statusCode
}

func (o *QueryPlatformsMixin0Default) Error() string {
	return fmt.Sprintf("[GET /ioarules/queries/platforms/v1][%d] query-platformsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *QueryPlatformsMixin0Default) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPlatformsMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}