// Code generated by go-swagger; DO NOT EDIT.

package sensor_download

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

// DownloadSensorInstallerByIDReader is a Reader for the DownloadSensorInstallerByID structure.
type DownloadSensorInstallerByIDReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DownloadSensorInstallerByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadSensorInstallerByIDOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDownloadSensorInstallerByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadSensorInstallerByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadSensorInstallerByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDownloadSensorInstallerByIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDownloadSensorInstallerByIDDefault(response.Code(), o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownloadSensorInstallerByIDOK creates a DownloadSensorInstallerByIDOK with default headers values
func NewDownloadSensorInstallerByIDOK(writer io.Writer) *DownloadSensorInstallerByIDOK {
	return &DownloadSensorInstallerByIDOK{
		Payload: writer,
	}
}

/*DownloadSensorInstallerByIDOK handles this case with default header values.

OK
*/
type DownloadSensorInstallerByIDOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload io.Writer
}

func (o *DownloadSensorInstallerByIDOK) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdOK  %+v", 200, o.Payload)
}

func (o *DownloadSensorInstallerByIDOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadSensorInstallerByIDBadRequest creates a DownloadSensorInstallerByIDBadRequest with default headers values
func NewDownloadSensorInstallerByIDBadRequest() *DownloadSensorInstallerByIDBadRequest {
	return &DownloadSensorInstallerByIDBadRequest{}
}

/*DownloadSensorInstallerByIDBadRequest handles this case with default header values.

Bad Request
*/
type DownloadSensorInstallerByIDBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DownloadSensorInstallerByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdBadRequest  %+v", 400, o.Payload)
}

func (o *DownloadSensorInstallerByIDBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadSensorInstallerByIDForbidden creates a DownloadSensorInstallerByIDForbidden with default headers values
func NewDownloadSensorInstallerByIDForbidden() *DownloadSensorInstallerByIDForbidden {
	return &DownloadSensorInstallerByIDForbidden{}
}

/*DownloadSensorInstallerByIDForbidden handles this case with default header values.

Forbidden
*/
type DownloadSensorInstallerByIDForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DownloadSensorInstallerByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdForbidden  %+v", 403, o.Payload)
}

func (o *DownloadSensorInstallerByIDForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadSensorInstallerByIDNotFound creates a DownloadSensorInstallerByIDNotFound with default headers values
func NewDownloadSensorInstallerByIDNotFound() *DownloadSensorInstallerByIDNotFound {
	return &DownloadSensorInstallerByIDNotFound{}
}

/*DownloadSensorInstallerByIDNotFound handles this case with default header values.

Not Found
*/
type DownloadSensorInstallerByIDNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DownloadSensorInstallerByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdNotFound  %+v", 404, o.Payload)
}

func (o *DownloadSensorInstallerByIDNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadSensorInstallerByIDTooManyRequests creates a DownloadSensorInstallerByIDTooManyRequests with default headers values
func NewDownloadSensorInstallerByIDTooManyRequests() *DownloadSensorInstallerByIDTooManyRequests {
	return &DownloadSensorInstallerByIDTooManyRequests{}
}

/*DownloadSensorInstallerByIDTooManyRequests handles this case with default header values.

Too Many Requests
*/
type DownloadSensorInstallerByIDTooManyRequests struct {
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

func (o *DownloadSensorInstallerByIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] downloadSensorInstallerByIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DownloadSensorInstallerByIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDownloadSensorInstallerByIDDefault creates a DownloadSensorInstallerByIDDefault with default headers values
func NewDownloadSensorInstallerByIDDefault(code int, writer io.Writer) *DownloadSensorInstallerByIDDefault {
	return &DownloadSensorInstallerByIDDefault{
		_statusCode: code,
		Payload:     writer,
	}
}

/*DownloadSensorInstallerByIDDefault handles this case with default header values.

OK
*/
type DownloadSensorInstallerByIDDefault struct {
	_statusCode int

	Payload io.Writer
}

// Code gets the status code for the download sensor installer by Id default response
func (o *DownloadSensorInstallerByIDDefault) Code() int {
	return o._statusCode
}

func (o *DownloadSensorInstallerByIDDefault) Error() string {
	return fmt.Sprintf("[GET /sensors/entities/download-installer/v1][%d] DownloadSensorInstallerById default  %+v", o._statusCode, o.Payload)
}

func (o *DownloadSensorInstallerByIDDefault) GetPayload() io.Writer {
	return o.Payload
}

func (o *DownloadSensorInstallerByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}