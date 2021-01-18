// Code generated by go-swagger; DO NOT EDIT.

package malquery

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

// GetMalQueryQuotasV1Reader is a Reader for the GetMalQueryQuotasV1 structure.
type GetMalQueryQuotasV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMalQueryQuotasV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMalQueryQuotasV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMalQueryQuotasV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMalQueryQuotasV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMalQueryQuotasV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMalQueryQuotasV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMalQueryQuotasV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMalQueryQuotasV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMalQueryQuotasV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMalQueryQuotasV1OK creates a GetMalQueryQuotasV1OK with default headers values
func NewGetMalQueryQuotasV1OK() *GetMalQueryQuotasV1OK {
	return &GetMalQueryQuotasV1OK{}
}

/*GetMalQueryQuotasV1OK handles this case with default header values.

OK
*/
type GetMalQueryQuotasV1OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRateLimitsResponse
}

func (o *GetMalQueryQuotasV1OK) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1OK  %+v", 200, o.Payload)
}

func (o *GetMalQueryQuotasV1OK) GetPayload() *models.MalqueryRateLimitsResponse {
	return o.Payload
}

func (o *GetMalQueryQuotasV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRateLimitsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1BadRequest creates a GetMalQueryQuotasV1BadRequest with default headers values
func NewGetMalQueryQuotasV1BadRequest() *GetMalQueryQuotasV1BadRequest {
	return &GetMalQueryQuotasV1BadRequest{}
}

/*GetMalQueryQuotasV1BadRequest handles this case with default header values.

Bad Request
*/
type GetMalQueryQuotasV1BadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMalQueryQuotasV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMalQueryQuotasV1BadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryQuotasV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1Unauthorized creates a GetMalQueryQuotasV1Unauthorized with default headers values
func NewGetMalQueryQuotasV1Unauthorized() *GetMalQueryQuotasV1Unauthorized {
	return &GetMalQueryQuotasV1Unauthorized{}
}

/*GetMalQueryQuotasV1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetMalQueryQuotasV1Unauthorized struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMalQueryQuotasV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetMalQueryQuotasV1Unauthorized) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryQuotasV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1Forbidden creates a GetMalQueryQuotasV1Forbidden with default headers values
func NewGetMalQueryQuotasV1Forbidden() *GetMalQueryQuotasV1Forbidden {
	return &GetMalQueryQuotasV1Forbidden{}
}

/*GetMalQueryQuotasV1Forbidden handles this case with default header values.

Forbidden
*/
type GetMalQueryQuotasV1Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMalQueryQuotasV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMalQueryQuotasV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMalQueryQuotasV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1NotFound creates a GetMalQueryQuotasV1NotFound with default headers values
func NewGetMalQueryQuotasV1NotFound() *GetMalQueryQuotasV1NotFound {
	return &GetMalQueryQuotasV1NotFound{}
}

/*GetMalQueryQuotasV1NotFound handles this case with default header values.

Not Found
*/
type GetMalQueryQuotasV1NotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRateLimitsResponse
}

func (o *GetMalQueryQuotasV1NotFound) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1NotFound  %+v", 404, o.Payload)
}

func (o *GetMalQueryQuotasV1NotFound) GetPayload() *models.MalqueryRateLimitsResponse {
	return o.Payload
}

func (o *GetMalQueryQuotasV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRateLimitsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1TooManyRequests creates a GetMalQueryQuotasV1TooManyRequests with default headers values
func NewGetMalQueryQuotasV1TooManyRequests() *GetMalQueryQuotasV1TooManyRequests {
	return &GetMalQueryQuotasV1TooManyRequests{}
}

/*GetMalQueryQuotasV1TooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetMalQueryQuotasV1TooManyRequests struct {
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

func (o *GetMalQueryQuotasV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMalQueryQuotasV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMalQueryQuotasV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMalQueryQuotasV1InternalServerError creates a GetMalQueryQuotasV1InternalServerError with default headers values
func NewGetMalQueryQuotasV1InternalServerError() *GetMalQueryQuotasV1InternalServerError {
	return &GetMalQueryQuotasV1InternalServerError{}
}

/*GetMalQueryQuotasV1InternalServerError handles this case with default header values.

Internal Server Error
*/
type GetMalQueryQuotasV1InternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MalqueryRateLimitsResponse
}

func (o *GetMalQueryQuotasV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] getMalQueryQuotasV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMalQueryQuotasV1InternalServerError) GetPayload() *models.MalqueryRateLimitsResponse {
	return o.Payload
}

func (o *GetMalQueryQuotasV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MalqueryRateLimitsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMalQueryQuotasV1Default creates a GetMalQueryQuotasV1Default with default headers values
func NewGetMalQueryQuotasV1Default(code int) *GetMalQueryQuotasV1Default {
	return &GetMalQueryQuotasV1Default{
		_statusCode: code,
	}
}

/*GetMalQueryQuotasV1Default handles this case with default header values.

OK
*/
type GetMalQueryQuotasV1Default struct {
	_statusCode int

	Payload *models.MalqueryRateLimitsResponse
}

// Code gets the status code for the get mal query quotas v1 default response
func (o *GetMalQueryQuotasV1Default) Code() int {
	return o._statusCode
}

func (o *GetMalQueryQuotasV1Default) Error() string {
	return fmt.Sprintf("[GET /malquery/aggregates/quotas/v1][%d] GetMalQueryQuotasV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetMalQueryQuotasV1Default) GetPayload() *models.MalqueryRateLimitsResponse {
	return o.Payload
}

func (o *GetMalQueryQuotasV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MalqueryRateLimitsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}