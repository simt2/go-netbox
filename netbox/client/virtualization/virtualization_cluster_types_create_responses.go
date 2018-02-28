// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// VirtualizationClusterTypesCreateReader is a Reader for the VirtualizationClusterTypesCreate structure.
type VirtualizationClusterTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewVirtualizationClusterTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterTypesCreateCreated creates a VirtualizationClusterTypesCreateCreated with default headers values
func NewVirtualizationClusterTypesCreateCreated() *VirtualizationClusterTypesCreateCreated {
	return &VirtualizationClusterTypesCreateCreated{}
}

/*VirtualizationClusterTypesCreateCreated handles this case with default header values.

VirtualizationClusterTypesCreateCreated virtualization cluster types create created
*/
type VirtualizationClusterTypesCreateCreated struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-types/][%d] virtualizationClusterTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *VirtualizationClusterTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}