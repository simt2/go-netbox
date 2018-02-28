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

// DcimInventoryItemsPartialUpdateReader is a Reader for the DcimInventoryItemsPartialUpdate structure.
type DcimInventoryItemsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimInventoryItemsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInventoryItemsPartialUpdateOK creates a DcimInventoryItemsPartialUpdateOK with default headers values
func NewDcimInventoryItemsPartialUpdateOK() *DcimInventoryItemsPartialUpdateOK {
	return &DcimInventoryItemsPartialUpdateOK{}
}

/*DcimInventoryItemsPartialUpdateOK handles this case with default header values.

DcimInventoryItemsPartialUpdateOK dcim inventory items partial update o k
*/
type DcimInventoryItemsPartialUpdateOK struct {
	Payload *models.WritableInventoryItem
}

func (o *DcimInventoryItemsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-items/{id}/][%d] dcimInventoryItemsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableInventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}