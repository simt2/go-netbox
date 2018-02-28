// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// DcimSitesReadReader is a Reader for the DcimSitesRead structure.
type DcimSitesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimSitesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimSitesReadOK creates a DcimSitesReadOK with default headers values
func NewDcimSitesReadOK() *DcimSitesReadOK {
	return &DcimSitesReadOK{}
}

/*DcimSitesReadOK handles this case with default header values.

DcimSitesReadOK dcim sites read o k
*/
type DcimSitesReadOK struct {
	Payload *models.Site
}

func (o *DcimSitesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/sites/{id}/][%d] dcimSitesReadOK  %+v", 200, o.Payload)
}

func (o *DcimSitesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}