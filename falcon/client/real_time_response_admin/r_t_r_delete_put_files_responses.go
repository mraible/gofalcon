// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRDeletePutFilesReader is a Reader for the RTRDeletePutFiles structure.
type RTRDeletePutFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeletePutFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRDeletePutFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeletePutFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeletePutFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRDeletePutFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeletePutFilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRDeletePutFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRDeletePutFilesOK creates a RTRDeletePutFilesOK with default headers values
func NewRTRDeletePutFilesOK() *RTRDeletePutFilesOK {
	return &RTRDeletePutFilesOK{}
}

/*RTRDeletePutFilesOK handles this case with default header values.

OK
*/
type RTRDeletePutFilesOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRDeletePutFilesOK) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRDeletePutFilesOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesBadRequest creates a RTRDeletePutFilesBadRequest with default headers values
func NewRTRDeletePutFilesBadRequest() *RTRDeletePutFilesBadRequest {
	return &RTRDeletePutFilesBadRequest{}
}

/*RTRDeletePutFilesBadRequest handles this case with default header values.

Bad Request
*/
type RTRDeletePutFilesBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRDeletePutFilesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeletePutFilesBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesForbidden creates a RTRDeletePutFilesForbidden with default headers values
func NewRTRDeletePutFilesForbidden() *RTRDeletePutFilesForbidden {
	return &RTRDeletePutFilesForbidden{}
}

/*RTRDeletePutFilesForbidden handles this case with default header values.

Forbidden
*/
type RTRDeletePutFilesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRDeletePutFilesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeletePutFilesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesNotFound creates a RTRDeletePutFilesNotFound with default headers values
func NewRTRDeletePutFilesNotFound() *RTRDeletePutFilesNotFound {
	return &RTRDeletePutFilesNotFound{}
}

/*RTRDeletePutFilesNotFound handles this case with default header values.

Not Found
*/
type RTRDeletePutFilesNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRDeletePutFilesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeletePutFilesNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesTooManyRequests creates a RTRDeletePutFilesTooManyRequests with default headers values
func NewRTRDeletePutFilesTooManyRequests() *RTRDeletePutFilesTooManyRequests {
	return &RTRDeletePutFilesTooManyRequests{}
}

/*RTRDeletePutFilesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type RTRDeletePutFilesTooManyRequests struct {
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

func (o *RTRDeletePutFilesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeletePutFilesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesDefault creates a RTRDeletePutFilesDefault with default headers values
func NewRTRDeletePutFilesDefault(code int) *RTRDeletePutFilesDefault {
	return &RTRDeletePutFilesDefault{
		_statusCode: code,
	}
}

/*RTRDeletePutFilesDefault handles this case with default header values.

OK
*/
type RTRDeletePutFilesDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// Code gets the status code for the r t r delete put files default response
func (o *RTRDeletePutFilesDefault) Code() int {
	return o._statusCode
}

func (o *RTRDeletePutFilesDefault) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] RTR-DeletePut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRDeletePutFilesDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}