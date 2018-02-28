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

// DcimDeviceBaysUpdateReader is a Reader for the DcimDeviceBaysUpdate structure.
type DcimDeviceBaysUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimDeviceBaysUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceBaysUpdateOK creates a DcimDeviceBaysUpdateOK with default headers values
func NewDcimDeviceBaysUpdateOK() *DcimDeviceBaysUpdateOK {
	return &DcimDeviceBaysUpdateOK{}
}

/*DcimDeviceBaysUpdateOK handles this case with default header values.

DcimDeviceBaysUpdateOK dcim device bays update o k
*/
type DcimDeviceBaysUpdateOK struct {
	Payload *models.WritableDeviceBay
}

func (o *DcimDeviceBaysUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bays/{id}/][%d] dcimDeviceBaysUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBaysUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableDeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}