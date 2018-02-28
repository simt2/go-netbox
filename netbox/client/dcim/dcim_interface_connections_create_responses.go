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

// DcimInterfaceConnectionsCreateReader is a Reader for the DcimInterfaceConnectionsCreate structure.
type DcimInterfaceConnectionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceConnectionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimInterfaceConnectionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfaceConnectionsCreateCreated creates a DcimInterfaceConnectionsCreateCreated with default headers values
func NewDcimInterfaceConnectionsCreateCreated() *DcimInterfaceConnectionsCreateCreated {
	return &DcimInterfaceConnectionsCreateCreated{}
}

/*DcimInterfaceConnectionsCreateCreated handles this case with default header values.

DcimInterfaceConnectionsCreateCreated dcim interface connections create created
*/
type DcimInterfaceConnectionsCreateCreated struct {
	Payload *models.WritableInterfaceConnection
}

func (o *DcimInterfaceConnectionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-connections/][%d] dcimInterfaceConnectionsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfaceConnectionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableInterfaceConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}