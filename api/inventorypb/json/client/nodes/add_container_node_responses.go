// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddContainerNodeReader is a Reader for the AddContainerNode structure.
type AddContainerNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddContainerNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddContainerNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddContainerNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddContainerNodeOK creates a AddContainerNodeOK with default headers values
func NewAddContainerNodeOK() *AddContainerNodeOK {
	return &AddContainerNodeOK{}
}

/*AddContainerNodeOK handles this case with default header values.

A successful response.
*/
type AddContainerNodeOK struct {
	Payload *AddContainerNodeOKBody
}

func (o *AddContainerNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddContainer][%d] addContainerNodeOk  %+v", 200, o.Payload)
}

func (o *AddContainerNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddContainerNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddContainerNodeDefault creates a AddContainerNodeDefault with default headers values
func NewAddContainerNodeDefault(code int) *AddContainerNodeDefault {
	return &AddContainerNodeDefault{
		_statusCode: code,
	}
}

/*AddContainerNodeDefault handles this case with default header values.

An error response.
*/
type AddContainerNodeDefault struct {
	_statusCode int

	Payload *AddContainerNodeDefaultBody
}

// Code gets the status code for the add container node default response
func (o *AddContainerNodeDefault) Code() int {
	return o._statusCode
}

func (o *AddContainerNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddContainer][%d] AddContainerNode default  %+v", o._statusCode, o.Payload)
}

func (o *AddContainerNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddContainerNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddContainerNodeBody add container node body
swagger:model AddContainerNodeBody
*/
type AddContainerNodeBody struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. Auto-detected and auto-updated.
	// If defined, Generic Node with that machine_id must exist.
	MachineID string `json:"machine_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this add container node body
func (o *AddContainerNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeBody) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddContainerNodeDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddContainerNodeDefaultBody
*/
type AddContainerNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add container node default body
func (o *AddContainerNodeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddContainerNodeOKBody add container node OK body
swagger:model AddContainerNodeOKBody
*/
type AddContainerNodeOKBody struct {

	// container
	Container *AddContainerNodeOKBodyContainer `json:"container,omitempty"`
}

// Validate validates this add container node OK body
func (o *AddContainerNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddContainerNodeOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addContainerNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddContainerNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model AddContainerNodeOKBodyContainer
*/
type AddContainerNodeOKBodyContainer struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	// Auto-detected and auto-updated.
	ContainerID string `json:"container_id,omitempty"`

	// Container name. Auto-detected and auto-updated.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. Auto-detected and auto-updated.
	// If defined, Generic Node with that machine_id must exist.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this add container node OK body container
func (o *AddContainerNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}