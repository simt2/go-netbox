// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimPowerOutletsDeleteReader is a Reader for the DcimPowerOutletsDelete structure.
type DcimPowerOutletsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimPowerOutletsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletsDeleteNoContent creates a DcimPowerOutletsDeleteNoContent with default headers values
func NewDcimPowerOutletsDeleteNoContent() *DcimPowerOutletsDeleteNoContent {
	return &DcimPowerOutletsDeleteNoContent{}
}

/*DcimPowerOutletsDeleteNoContent handles this case with default header values.

DcimPowerOutletsDeleteNoContent dcim power outlets delete no content
*/
type DcimPowerOutletsDeleteNoContent struct {
}

func (o *DcimPowerOutletsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlets/{id}/][%d] dcimPowerOutletsDeleteNoContent ", 204)
}

func (o *DcimPowerOutletsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}