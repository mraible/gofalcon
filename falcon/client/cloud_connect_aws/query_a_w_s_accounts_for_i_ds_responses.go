// Code generated by go-swagger; DO NOT EDIT.

package cloud_connect_aws

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

// QueryAWSAccountsForIDsReader is a Reader for the QueryAWSAccountsForIDs structure.
type QueryAWSAccountsForIDsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryAWSAccountsForIDsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryAWSAccountsForIDsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryAWSAccountsForIDsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryAWSAccountsForIDsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryAWSAccountsForIDsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryAWSAccountsForIDsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryAWSAccountsForIDsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryAWSAccountsForIDsOK creates a QueryAWSAccountsForIDsOK with default headers values
func NewQueryAWSAccountsForIDsOK() *QueryAWSAccountsForIDsOK {
	return &QueryAWSAccountsForIDsOK{}
}

/*QueryAWSAccountsForIDsOK handles this case with default header values.

OK
*/
type QueryAWSAccountsForIDsOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryAWSAccountsForIDsOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] queryAWSAccountsForIDsOK  %+v", 200, o.Payload)
}

func (o *QueryAWSAccountsForIDsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryAWSAccountsForIDsBadRequest creates a QueryAWSAccountsForIDsBadRequest with default headers values
func NewQueryAWSAccountsForIDsBadRequest() *QueryAWSAccountsForIDsBadRequest {
	return &QueryAWSAccountsForIDsBadRequest{}
}

/*QueryAWSAccountsForIDsBadRequest handles this case with default header values.

Bad Request
*/
type QueryAWSAccountsForIDsBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryAWSAccountsForIDsBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] queryAWSAccountsForIDsBadRequest  %+v", 400, o.Payload)
}

func (o *QueryAWSAccountsForIDsBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryAWSAccountsForIDsForbidden creates a QueryAWSAccountsForIDsForbidden with default headers values
func NewQueryAWSAccountsForIDsForbidden() *QueryAWSAccountsForIDsForbidden {
	return &QueryAWSAccountsForIDsForbidden{}
}

/*QueryAWSAccountsForIDsForbidden handles this case with default header values.

Forbidden
*/
type QueryAWSAccountsForIDsForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryAWSAccountsForIDsForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] queryAWSAccountsForIDsForbidden  %+v", 403, o.Payload)
}

func (o *QueryAWSAccountsForIDsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryAWSAccountsForIDsTooManyRequests creates a QueryAWSAccountsForIDsTooManyRequests with default headers values
func NewQueryAWSAccountsForIDsTooManyRequests() *QueryAWSAccountsForIDsTooManyRequests {
	return &QueryAWSAccountsForIDsTooManyRequests{}
}

/*QueryAWSAccountsForIDsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type QueryAWSAccountsForIDsTooManyRequests struct {
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

func (o *QueryAWSAccountsForIDsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] queryAWSAccountsForIDsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryAWSAccountsForIDsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryAWSAccountsForIDsInternalServerError creates a QueryAWSAccountsForIDsInternalServerError with default headers values
func NewQueryAWSAccountsForIDsInternalServerError() *QueryAWSAccountsForIDsInternalServerError {
	return &QueryAWSAccountsForIDsInternalServerError{}
}

/*QueryAWSAccountsForIDsInternalServerError handles this case with default header values.

Internal Server Error
*/
type QueryAWSAccountsForIDsInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryAWSAccountsForIDsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] queryAWSAccountsForIDsInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryAWSAccountsForIDsInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryAWSAccountsForIDsDefault creates a QueryAWSAccountsForIDsDefault with default headers values
func NewQueryAWSAccountsForIDsDefault(code int) *QueryAWSAccountsForIDsDefault {
	return &QueryAWSAccountsForIDsDefault{
		_statusCode: code,
	}
}

/*QueryAWSAccountsForIDsDefault handles this case with default header values.

OK
*/
type QueryAWSAccountsForIDsDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query a w s accounts for i ds default response
func (o *QueryAWSAccountsForIDsDefault) Code() int {
	return o._statusCode
}

func (o *QueryAWSAccountsForIDsDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/queries/accounts/v1][%d] QueryAWSAccountsForIDs default  %+v", o._statusCode, o.Payload)
}

func (o *QueryAWSAccountsForIDsDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryAWSAccountsForIDsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}