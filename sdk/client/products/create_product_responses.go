// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cesarvspr/golang-modules/sdk/models"
)

// CreateProductReader is a Reader for the CreateProduct structure.
type CreateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewCreateProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductOK creates a CreateProductOK with default headers values
func NewCreateProductOK() *CreateProductOK {
	return &CreateProductOK{}
}

/* CreateProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type CreateProductOK struct {
	Payload []*models.Product
}

func (o *CreateProductOK) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductOK  %+v", 200, o.Payload)
}
func (o *CreateProductOK) GetPayload() []*models.Product {
	return o.Payload
}

func (o *CreateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductUnprocessableEntity creates a CreateProductUnprocessableEntity with default headers values
func NewCreateProductUnprocessableEntity() *CreateProductUnprocessableEntity {
	return &CreateProductUnprocessableEntity{}
}

/* CreateProductUnprocessableEntity describes a response with status code 422, with default header values.

 Validation errors defined as an array of strings
ValidationError is a collection of validation error messages
*/
type CreateProductUnprocessableEntity struct {
	Payload *CreateProductUnprocessableEntityBody
}

func (o *CreateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateProductUnprocessableEntity) GetPayload() *CreateProductUnprocessableEntityBody {
	return o.Payload
}

func (o *CreateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateProductUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductNotImplemented creates a CreateProductNotImplemented with default headers values
func NewCreateProductNotImplemented() *CreateProductNotImplemented {
	return &CreateProductNotImplemented{}
}

/* CreateProductNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type CreateProductNotImplemented struct {
	Payload *CreateProductNotImplementedBody
}

func (o *CreateProductNotImplemented) Error() string {
	return fmt.Sprintf("[POST /products][%d] createProductNotImplemented  %+v", 501, o.Payload)
}
func (o *CreateProductNotImplemented) GetPayload() *CreateProductNotImplementedBody {
	return o.Payload
}

func (o *CreateProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateProductNotImplementedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateProductNotImplementedBody create product not implemented body
swagger:model CreateProductNotImplementedBody
*/
type CreateProductNotImplementedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create product not implemented body
func (o *CreateProductNotImplementedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create product not implemented body based on context it is used
func (o *CreateProductNotImplementedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateProductNotImplementedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateProductNotImplementedBody) UnmarshalBinary(b []byte) error {
	var res CreateProductNotImplementedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateProductUnprocessableEntityBody create product unprocessable entity body
swagger:model CreateProductUnprocessableEntityBody
*/
type CreateProductUnprocessableEntityBody struct {

	// messages
	Messages []string `json:"messages"`
}

// Validate validates this create product unprocessable entity body
func (o *CreateProductUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create product unprocessable entity body based on context it is used
func (o *CreateProductUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateProductUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateProductUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res CreateProductUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
