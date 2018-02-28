// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// ExtrasTopologyMapsCreateReader is a Reader for the ExtrasTopologyMapsCreate structure.
type ExtrasTopologyMapsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTopologyMapsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewExtrasTopologyMapsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasTopologyMapsCreateCreated creates a ExtrasTopologyMapsCreateCreated with default headers values
func NewExtrasTopologyMapsCreateCreated() *ExtrasTopologyMapsCreateCreated {
	return &ExtrasTopologyMapsCreateCreated{}
}

/*ExtrasTopologyMapsCreateCreated handles this case with default header values.

ExtrasTopologyMapsCreateCreated extras topology maps create created
*/
type ExtrasTopologyMapsCreateCreated struct {
	Payload *models.WritableTopologyMap
}

func (o *ExtrasTopologyMapsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/topology-maps/][%d] extrasTopologyMapsCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasTopologyMapsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableTopologyMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}