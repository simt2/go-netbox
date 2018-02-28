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

// DcimDevicesListReader is a Reader for the DcimDevicesList structure.
type DcimDevicesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimDevicesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDevicesListOK creates a DcimDevicesListOK with default headers values
func NewDcimDevicesListOK() *DcimDevicesListOK {
	return &DcimDevicesListOK{}
}

/*DcimDevicesListOK handles this case with default header values.

DcimDevicesListOK dcim devices list o k
*/
type DcimDevicesListOK struct {
	Payload *models.DcimDevicesListOKBody
}

func (o *DcimDevicesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/][%d] dcimDevicesListOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DcimDevicesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}