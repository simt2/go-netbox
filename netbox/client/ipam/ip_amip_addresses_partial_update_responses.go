// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// IPAMIPAddressesPartialUpdateReader is a Reader for the IPAMIPAddressesPartialUpdate structure.
type IPAMIPAddressesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMIPAddressesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMIPAddressesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMIPAddressesPartialUpdateOK creates a IPAMIPAddressesPartialUpdateOK with default headers values
func NewIPAMIPAddressesPartialUpdateOK() *IPAMIPAddressesPartialUpdateOK {
	return &IPAMIPAddressesPartialUpdateOK{}
}

/*IPAMIPAddressesPartialUpdateOK handles this case with default header values.

IPAMIPAddressesPartialUpdateOK ipam Ip addresses partial update o k
*/
type IPAMIPAddressesPartialUpdateOK struct {
	Payload *models.WritableIPAddress
}

func (o *IPAMIPAddressesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMIPAddressesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableIPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}