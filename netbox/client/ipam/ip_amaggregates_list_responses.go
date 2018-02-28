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

// IPAMAggregatesListReader is a Reader for the IPAMAggregatesList structure.
type IPAMAggregatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMAggregatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMAggregatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMAggregatesListOK creates a IPAMAggregatesListOK with default headers values
func NewIPAMAggregatesListOK() *IPAMAggregatesListOK {
	return &IPAMAggregatesListOK{}
}

/*IPAMAggregatesListOK handles this case with default header values.

IPAMAggregatesListOK ipam aggregates list o k
*/
type IPAMAggregatesListOK struct {
	Payload *models.IPAMAggregatesListOKBody
}

func (o *IPAMAggregatesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/aggregates/][%d] ipamAggregatesListOK  %+v", 200, o.Payload)
}

func (o *IPAMAggregatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMAggregatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}