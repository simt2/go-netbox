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

// DcimConsoleServerPortTemplatesListReader is a Reader for the DcimConsoleServerPortTemplatesList structure.
type DcimConsoleServerPortTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimConsoleServerPortTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesListOK creates a DcimConsoleServerPortTemplatesListOK with default headers values
func NewDcimConsoleServerPortTemplatesListOK() *DcimConsoleServerPortTemplatesListOK {
	return &DcimConsoleServerPortTemplatesListOK{}
}

/*DcimConsoleServerPortTemplatesListOK handles this case with default header values.

DcimConsoleServerPortTemplatesListOK dcim console server port templates list o k
*/
type DcimConsoleServerPortTemplatesListOK struct {
	Payload *models.DcimConsoleServerPortTemplatesListOKBody
}

func (o *DcimConsoleServerPortTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DcimConsoleServerPortTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}