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

// DcimPowerOutletsPartialUpdateReader is a Reader for the DcimPowerOutletsPartialUpdate structure.
type DcimPowerOutletsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPowerOutletsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletsPartialUpdateOK creates a DcimPowerOutletsPartialUpdateOK with default headers values
func NewDcimPowerOutletsPartialUpdateOK() *DcimPowerOutletsPartialUpdateOK {
	return &DcimPowerOutletsPartialUpdateOK{}
}

/*DcimPowerOutletsPartialUpdateOK handles this case with default header values.

DcimPowerOutletsPartialUpdateOK dcim power outlets partial update o k
*/
type DcimPowerOutletsPartialUpdateOK struct {
	Payload *models.WritablePowerOutlet
}

func (o *DcimPowerOutletsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/{id}/][%d] dcimPowerOutletsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritablePowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}