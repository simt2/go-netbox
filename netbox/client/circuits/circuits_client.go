// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new circuits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for circuits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CircuitsChoicesList circuits choices list API
*/
func (a *Client) CircuitsChoicesList(params *CircuitsChoicesListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsChoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsChoicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits__choices_list",
		Method:             "GET",
		PathPattern:        "/circuits/_choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsChoicesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsChoicesListOK), nil

}

/*
CircuitsChoicesRead circuits choices read API
*/
func (a *Client) CircuitsChoicesRead(params *CircuitsChoicesReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsChoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsChoicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits__choices_read",
		Method:             "GET",
		PathPattern:        "/circuits/_choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsChoicesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsChoicesReadOK), nil

}

/*
CircuitsCircuitTerminationsCreate circuits circuit terminations create API
*/
func (a *Client) CircuitsCircuitTerminationsCreate(params *CircuitsCircuitTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsCreateCreated), nil

}

/*
CircuitsCircuitTerminationsDelete circuits circuit terminations delete API
*/
func (a *Client) CircuitsCircuitTerminationsDelete(params *CircuitsCircuitTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsDeleteNoContent), nil

}

/*
CircuitsCircuitTerminationsList circuits circuit terminations list API
*/
func (a *Client) CircuitsCircuitTerminationsList(params *CircuitsCircuitTerminationsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsListOK), nil

}

/*
CircuitsCircuitTerminationsPartialUpdate circuits circuit terminations partial update API
*/
func (a *Client) CircuitsCircuitTerminationsPartialUpdate(params *CircuitsCircuitTerminationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsPartialUpdateOK), nil

}

/*
CircuitsCircuitTerminationsRead circuits circuit terminations read API
*/
func (a *Client) CircuitsCircuitTerminationsRead(params *CircuitsCircuitTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsReadOK), nil

}

/*
CircuitsCircuitTerminationsUpdate circuits circuit terminations update API
*/
func (a *Client) CircuitsCircuitTerminationsUpdate(params *CircuitsCircuitTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTerminationsUpdateOK), nil

}

/*
CircuitsCircuitTypesCreate circuits circuit types create API
*/
func (a *Client) CircuitsCircuitTypesCreate(params *CircuitsCircuitTypesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesCreateCreated), nil

}

/*
CircuitsCircuitTypesDelete circuits circuit types delete API
*/
func (a *Client) CircuitsCircuitTypesDelete(params *CircuitsCircuitTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesDeleteNoContent), nil

}

/*
CircuitsCircuitTypesList circuits circuit types list API
*/
func (a *Client) CircuitsCircuitTypesList(params *CircuitsCircuitTypesListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesListOK), nil

}

/*
CircuitsCircuitTypesPartialUpdate circuits circuit types partial update API
*/
func (a *Client) CircuitsCircuitTypesPartialUpdate(params *CircuitsCircuitTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesPartialUpdateOK), nil

}

/*
CircuitsCircuitTypesRead circuits circuit types read API
*/
func (a *Client) CircuitsCircuitTypesRead(params *CircuitsCircuitTypesReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesReadOK), nil

}

/*
CircuitsCircuitTypesUpdate circuits circuit types update API
*/
func (a *Client) CircuitsCircuitTypesUpdate(params *CircuitsCircuitTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitTypesUpdateOK), nil

}

/*
CircuitsCircuitsCreate circuits circuits create API
*/
func (a *Client) CircuitsCircuitsCreate(params *CircuitsCircuitsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsCreateCreated), nil

}

/*
CircuitsCircuitsDelete circuits circuits delete API
*/
func (a *Client) CircuitsCircuitsDelete(params *CircuitsCircuitsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsDeleteNoContent), nil

}

/*
CircuitsCircuitsList circuits circuits list API
*/
func (a *Client) CircuitsCircuitsList(params *CircuitsCircuitsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsListOK), nil

}

/*
CircuitsCircuitsPartialUpdate circuits circuits partial update API
*/
func (a *Client) CircuitsCircuitsPartialUpdate(params *CircuitsCircuitsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsPartialUpdateOK), nil

}

/*
CircuitsCircuitsRead circuits circuits read API
*/
func (a *Client) CircuitsCircuitsRead(params *CircuitsCircuitsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsReadOK), nil

}

/*
CircuitsCircuitsUpdate circuits circuits update API
*/
func (a *Client) CircuitsCircuitsUpdate(params *CircuitsCircuitsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsCircuitsUpdateOK), nil

}

/*
CircuitsProvidersCreate circuits providers create API
*/
func (a *Client) CircuitsProvidersCreate(params *CircuitsProvidersCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_create",
		Method:             "POST",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersCreateCreated), nil

}

/*
CircuitsProvidersDelete circuits providers delete API
*/
func (a *Client) CircuitsProvidersDelete(params *CircuitsProvidersDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersDeleteNoContent), nil

}

/*
CircuitsProvidersGraphs A convenience method for rendering graphs for a particular provider.
*/
func (a *Client) CircuitsProvidersGraphs(params *CircuitsProvidersGraphsParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersGraphsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersGraphsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_graphs",
		Method:             "GET",
		PathPattern:        "/circuits/providers/{id}/graphs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersGraphsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersGraphsOK), nil

}

/*
CircuitsProvidersList circuits providers list API
*/
func (a *Client) CircuitsProvidersList(params *CircuitsProvidersListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_list",
		Method:             "GET",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersListOK), nil

}

/*
CircuitsProvidersPartialUpdate circuits providers partial update API
*/
func (a *Client) CircuitsProvidersPartialUpdate(params *CircuitsProvidersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersPartialUpdateOK), nil

}

/*
CircuitsProvidersRead circuits providers read API
*/
func (a *Client) CircuitsProvidersRead(params *CircuitsProvidersReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_read",
		Method:             "GET",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersReadOK), nil

}

/*
CircuitsProvidersUpdate circuits providers update API
*/
func (a *Client) CircuitsProvidersUpdate(params *CircuitsProvidersUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_update",
		Method:             "PUT",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CircuitsProvidersUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}