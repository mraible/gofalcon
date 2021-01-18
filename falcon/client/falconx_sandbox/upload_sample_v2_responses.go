// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

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

// UploadSampleV2Reader is a Reader for the UploadSampleV2 structure.
type UploadSampleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadSampleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadSampleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadSampleV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadSampleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUploadSampleV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadSampleV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUploadSampleV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadSampleV2OK creates a UploadSampleV2OK with default headers values
func NewUploadSampleV2OK() *UploadSampleV2OK {
	return &UploadSampleV2OK{}
}

/*UploadSampleV2OK handles this case with default header values.

OK
*/
type UploadSampleV2OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV2OK) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] uploadSampleV2OK  %+v", 200, o.Payload)
}

func (o *UploadSampleV2OK) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV2BadRequest creates a UploadSampleV2BadRequest with default headers values
func NewUploadSampleV2BadRequest() *UploadSampleV2BadRequest {
	return &UploadSampleV2BadRequest{}
}

/*UploadSampleV2BadRequest handles this case with default header values.

Bad Request
*/
type UploadSampleV2BadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] uploadSampleV2BadRequest  %+v", 400, o.Payload)
}

func (o *UploadSampleV2BadRequest) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV2Forbidden creates a UploadSampleV2Forbidden with default headers values
func NewUploadSampleV2Forbidden() *UploadSampleV2Forbidden {
	return &UploadSampleV2Forbidden{}
}

/*UploadSampleV2Forbidden handles this case with default header values.

Forbidden
*/
type UploadSampleV2Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UploadSampleV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] uploadSampleV2Forbidden  %+v", 403, o.Payload)
}

func (o *UploadSampleV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV2TooManyRequests creates a UploadSampleV2TooManyRequests with default headers values
func NewUploadSampleV2TooManyRequests() *UploadSampleV2TooManyRequests {
	return &UploadSampleV2TooManyRequests{}
}

/*UploadSampleV2TooManyRequests handles this case with default header values.

Too Many Requests
*/
type UploadSampleV2TooManyRequests struct {
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

func (o *UploadSampleV2TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] uploadSampleV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadSampleV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV2InternalServerError creates a UploadSampleV2InternalServerError with default headers values
func NewUploadSampleV2InternalServerError() *UploadSampleV2InternalServerError {
	return &UploadSampleV2InternalServerError{}
}

/*UploadSampleV2InternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadSampleV2InternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SamplestoreSampleMetadataResponseV2
}

func (o *UploadSampleV2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] uploadSampleV2InternalServerError  %+v", 500, o.Payload)
}

func (o *UploadSampleV2InternalServerError) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV2Default creates a UploadSampleV2Default with default headers values
func NewUploadSampleV2Default(code int) *UploadSampleV2Default {
	return &UploadSampleV2Default{
		_statusCode: code,
	}
}

/*UploadSampleV2Default handles this case with default header values.

OK
*/
type UploadSampleV2Default struct {
	_statusCode int

	Payload *models.SamplestoreSampleMetadataResponseV2
}

// Code gets the status code for the upload sample v2 default response
func (o *UploadSampleV2Default) Code() int {
	return o._statusCode
}

func (o *UploadSampleV2Default) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v2][%d] UploadSampleV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UploadSampleV2Default) GetPayload() *models.SamplestoreSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SamplestoreSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}