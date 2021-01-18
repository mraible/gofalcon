// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// BatchActiveResponderCmdReader is a Reader for the BatchActiveResponderCmd structure.
type BatchActiveResponderCmdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchActiveResponderCmdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewBatchActiveResponderCmdCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchActiveResponderCmdBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchActiveResponderCmdForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewBatchActiveResponderCmdTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchActiveResponderCmdInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBatchActiveResponderCmdCreated creates a BatchActiveResponderCmdCreated with default headers values
func NewBatchActiveResponderCmdCreated() *BatchActiveResponderCmdCreated {
	return &BatchActiveResponderCmdCreated{}
}

/*BatchActiveResponderCmdCreated handles this case with default header values.

Created
*/
type BatchActiveResponderCmdCreated struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMultiCommandExecuteResponseWrapper
}

func (o *BatchActiveResponderCmdCreated) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-active-responder-command/v1][%d] batchActiveResponderCmdCreated  %+v", 201, o.Payload)
}

func (o *BatchActiveResponderCmdCreated) GetPayload() *models.DomainMultiCommandExecuteResponseWrapper {
	return o.Payload
}

func (o *BatchActiveResponderCmdCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMultiCommandExecuteResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchActiveResponderCmdBadRequest creates a BatchActiveResponderCmdBadRequest with default headers values
func NewBatchActiveResponderCmdBadRequest() *BatchActiveResponderCmdBadRequest {
	return &BatchActiveResponderCmdBadRequest{}
}

/*BatchActiveResponderCmdBadRequest handles this case with default header values.

Bad Request
*/
type BatchActiveResponderCmdBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *BatchActiveResponderCmdBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-active-responder-command/v1][%d] batchActiveResponderCmdBadRequest  %+v", 400, o.Payload)
}

func (o *BatchActiveResponderCmdBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchActiveResponderCmdBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchActiveResponderCmdForbidden creates a BatchActiveResponderCmdForbidden with default headers values
func NewBatchActiveResponderCmdForbidden() *BatchActiveResponderCmdForbidden {
	return &BatchActiveResponderCmdForbidden{}
}

/*BatchActiveResponderCmdForbidden handles this case with default header values.

Forbidden
*/
type BatchActiveResponderCmdForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *BatchActiveResponderCmdForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-active-responder-command/v1][%d] batchActiveResponderCmdForbidden  %+v", 403, o.Payload)
}

func (o *BatchActiveResponderCmdForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *BatchActiveResponderCmdForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchActiveResponderCmdTooManyRequests creates a BatchActiveResponderCmdTooManyRequests with default headers values
func NewBatchActiveResponderCmdTooManyRequests() *BatchActiveResponderCmdTooManyRequests {
	return &BatchActiveResponderCmdTooManyRequests{}
}

/*BatchActiveResponderCmdTooManyRequests handles this case with default header values.

Too Many Requests
*/
type BatchActiveResponderCmdTooManyRequests struct {
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

func (o *BatchActiveResponderCmdTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-active-responder-command/v1][%d] batchActiveResponderCmdTooManyRequests  %+v", 429, o.Payload)
}

func (o *BatchActiveResponderCmdTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *BatchActiveResponderCmdTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchActiveResponderCmdInternalServerError creates a BatchActiveResponderCmdInternalServerError with default headers values
func NewBatchActiveResponderCmdInternalServerError() *BatchActiveResponderCmdInternalServerError {
	return &BatchActiveResponderCmdInternalServerError{}
}

/*BatchActiveResponderCmdInternalServerError handles this case with default header values.

Internal Server Error
*/
type BatchActiveResponderCmdInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *BatchActiveResponderCmdInternalServerError) Error() string {
	return fmt.Sprintf("[POST /real-time-response/combined/batch-active-responder-command/v1][%d] batchActiveResponderCmdInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchActiveResponderCmdInternalServerError) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *BatchActiveResponderCmdInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}