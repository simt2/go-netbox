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

// DcimPowerPortsListReader is a Reader for the DcimPowerPortsList structure.
type DcimPowerPortsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerPortsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortsListOK creates a DcimPowerPortsListOK with default headers values
func NewDcimPowerPortsListOK() *DcimPowerPortsListOK {
	return &DcimPowerPortsListOK{}
}

/*DcimPowerPortsListOK handles this case with default header values.

DcimPowerPortsListOK dcim power ports list o k
*/
type DcimPowerPortsListOK struct {
	Payload *models.DcimPowerPortsListOKBody
}

func (o *DcimPowerPortsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/][%d] dcimPowerPortsListOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DcimPowerPortsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}