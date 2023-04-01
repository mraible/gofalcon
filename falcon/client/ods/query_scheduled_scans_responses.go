// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// QueryScheduledScansReader is a Reader for the QueryScheduledScans structure.
type QueryScheduledScansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryScheduledScansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryScheduledScansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryScheduledScansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryScheduledScansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryScheduledScansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryScheduledScansOK creates a QueryScheduledScansOK with default headers values
func NewQueryScheduledScansOK() *QueryScheduledScansOK {
	return &QueryScheduledScansOK{}
}

/*
QueryScheduledScansOK describes a response with status code 200, with default header values.

OK
*/
type QueryScheduledScansOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query scheduled scans o k response has a 2xx status code
func (o *QueryScheduledScansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query scheduled scans o k response has a 3xx status code
func (o *QueryScheduledScansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scheduled scans o k response has a 4xx status code
func (o *QueryScheduledScansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query scheduled scans o k response has a 5xx status code
func (o *QueryScheduledScansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query scheduled scans o k response a status code equal to that given
func (o *QueryScheduledScansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query scheduled scans o k response
func (o *QueryScheduledScansOK) Code() int {
	return 200
}

func (o *QueryScheduledScansOK) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansOK  %+v", 200, o.Payload)
}

func (o *QueryScheduledScansOK) String() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansOK  %+v", 200, o.Payload)
}

func (o *QueryScheduledScansOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryScheduledScansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryScheduledScansForbidden creates a QueryScheduledScansForbidden with default headers values
func NewQueryScheduledScansForbidden() *QueryScheduledScansForbidden {
	return &QueryScheduledScansForbidden{}
}

/*
QueryScheduledScansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryScheduledScansForbidden struct {

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

// IsSuccess returns true when this query scheduled scans forbidden response has a 2xx status code
func (o *QueryScheduledScansForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scheduled scans forbidden response has a 3xx status code
func (o *QueryScheduledScansForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scheduled scans forbidden response has a 4xx status code
func (o *QueryScheduledScansForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scheduled scans forbidden response has a 5xx status code
func (o *QueryScheduledScansForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query scheduled scans forbidden response a status code equal to that given
func (o *QueryScheduledScansForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query scheduled scans forbidden response
func (o *QueryScheduledScansForbidden) Code() int {
	return 403
}

func (o *QueryScheduledScansForbidden) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansForbidden  %+v", 403, o.Payload)
}

func (o *QueryScheduledScansForbidden) String() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansForbidden  %+v", 403, o.Payload)
}

func (o *QueryScheduledScansForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryScheduledScansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryScheduledScansNotFound creates a QueryScheduledScansNotFound with default headers values
func NewQueryScheduledScansNotFound() *QueryScheduledScansNotFound {
	return &QueryScheduledScansNotFound{}
}

/*
QueryScheduledScansNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryScheduledScansNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query scheduled scans not found response has a 2xx status code
func (o *QueryScheduledScansNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scheduled scans not found response has a 3xx status code
func (o *QueryScheduledScansNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scheduled scans not found response has a 4xx status code
func (o *QueryScheduledScansNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scheduled scans not found response has a 5xx status code
func (o *QueryScheduledScansNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query scheduled scans not found response a status code equal to that given
func (o *QueryScheduledScansNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query scheduled scans not found response
func (o *QueryScheduledScansNotFound) Code() int {
	return 404
}

func (o *QueryScheduledScansNotFound) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansNotFound  %+v", 404, o.Payload)
}

func (o *QueryScheduledScansNotFound) String() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansNotFound  %+v", 404, o.Payload)
}

func (o *QueryScheduledScansNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryScheduledScansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryScheduledScansTooManyRequests creates a QueryScheduledScansTooManyRequests with default headers values
func NewQueryScheduledScansTooManyRequests() *QueryScheduledScansTooManyRequests {
	return &QueryScheduledScansTooManyRequests{}
}

/*
QueryScheduledScansTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryScheduledScansTooManyRequests struct {

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

// IsSuccess returns true when this query scheduled scans too many requests response has a 2xx status code
func (o *QueryScheduledScansTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scheduled scans too many requests response has a 3xx status code
func (o *QueryScheduledScansTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scheduled scans too many requests response has a 4xx status code
func (o *QueryScheduledScansTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scheduled scans too many requests response has a 5xx status code
func (o *QueryScheduledScansTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query scheduled scans too many requests response a status code equal to that given
func (o *QueryScheduledScansTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query scheduled scans too many requests response
func (o *QueryScheduledScansTooManyRequests) Code() int {
	return 429
}

func (o *QueryScheduledScansTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryScheduledScansTooManyRequests) String() string {
	return fmt.Sprintf("[GET /ods/queries/scheduled-scans/v1][%d] queryScheduledScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryScheduledScansTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryScheduledScansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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