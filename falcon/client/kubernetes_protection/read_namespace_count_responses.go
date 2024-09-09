// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// ReadNamespaceCountReader is a Reader for the ReadNamespaceCount structure.
type ReadNamespaceCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadNamespaceCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadNamespaceCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewReadNamespaceCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewReadNamespaceCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadNamespaceCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/namespaces/count/v1] ReadNamespaceCount", response, response.Code())
	}
}

// NewReadNamespaceCountOK creates a ReadNamespaceCountOK with default headers values
func NewReadNamespaceCountOK() *ReadNamespaceCountOK {
	return &ReadNamespaceCountOK{}
}

/*
ReadNamespaceCountOK describes a response with status code 200, with default header values.

OK
*/
type ReadNamespaceCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CommonCountResponse
}

// IsSuccess returns true when this read namespace count o k response has a 2xx status code
func (o *ReadNamespaceCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read namespace count o k response has a 3xx status code
func (o *ReadNamespaceCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read namespace count o k response has a 4xx status code
func (o *ReadNamespaceCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read namespace count o k response has a 5xx status code
func (o *ReadNamespaceCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read namespace count o k response a status code equal to that given
func (o *ReadNamespaceCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read namespace count o k response
func (o *ReadNamespaceCountOK) Code() int {
	return 200
}

func (o *ReadNamespaceCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountOK  %+v", 200, o.Payload)
}

func (o *ReadNamespaceCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountOK  %+v", 200, o.Payload)
}

func (o *ReadNamespaceCountOK) GetPayload() *models.CommonCountResponse {
	return o.Payload
}

func (o *ReadNamespaceCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.CommonCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNamespaceCountForbidden creates a ReadNamespaceCountForbidden with default headers values
func NewReadNamespaceCountForbidden() *ReadNamespaceCountForbidden {
	return &ReadNamespaceCountForbidden{}
}

/*
ReadNamespaceCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReadNamespaceCountForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this read namespace count forbidden response has a 2xx status code
func (o *ReadNamespaceCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read namespace count forbidden response has a 3xx status code
func (o *ReadNamespaceCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read namespace count forbidden response has a 4xx status code
func (o *ReadNamespaceCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this read namespace count forbidden response has a 5xx status code
func (o *ReadNamespaceCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this read namespace count forbidden response a status code equal to that given
func (o *ReadNamespaceCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the read namespace count forbidden response
func (o *ReadNamespaceCountForbidden) Code() int {
	return 403
}

func (o *ReadNamespaceCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadNamespaceCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountForbidden  %+v", 403, o.Payload)
}

func (o *ReadNamespaceCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadNamespaceCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNamespaceCountTooManyRequests creates a ReadNamespaceCountTooManyRequests with default headers values
func NewReadNamespaceCountTooManyRequests() *ReadNamespaceCountTooManyRequests {
	return &ReadNamespaceCountTooManyRequests{}
}

/*
ReadNamespaceCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ReadNamespaceCountTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this read namespace count too many requests response has a 2xx status code
func (o *ReadNamespaceCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read namespace count too many requests response has a 3xx status code
func (o *ReadNamespaceCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read namespace count too many requests response has a 4xx status code
func (o *ReadNamespaceCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this read namespace count too many requests response has a 5xx status code
func (o *ReadNamespaceCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this read namespace count too many requests response a status code equal to that given
func (o *ReadNamespaceCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the read namespace count too many requests response
func (o *ReadNamespaceCountTooManyRequests) Code() int {
	return 429
}

func (o *ReadNamespaceCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadNamespaceCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReadNamespaceCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ReadNamespaceCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNamespaceCountInternalServerError creates a ReadNamespaceCountInternalServerError with default headers values
func NewReadNamespaceCountInternalServerError() *ReadNamespaceCountInternalServerError {
	return &ReadNamespaceCountInternalServerError{}
}

/*
ReadNamespaceCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReadNamespaceCountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this read namespace count internal server error response has a 2xx status code
func (o *ReadNamespaceCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read namespace count internal server error response has a 3xx status code
func (o *ReadNamespaceCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read namespace count internal server error response has a 4xx status code
func (o *ReadNamespaceCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read namespace count internal server error response has a 5xx status code
func (o *ReadNamespaceCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read namespace count internal server error response a status code equal to that given
func (o *ReadNamespaceCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read namespace count internal server error response
func (o *ReadNamespaceCountInternalServerError) Code() int {
	return 500
}

func (o *ReadNamespaceCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadNamespaceCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/namespaces/count/v1][%d] readNamespaceCountInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadNamespaceCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *ReadNamespaceCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}