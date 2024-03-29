// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSAttachment represents the model of a awsattachment
type AWSAttachment struct {
	// AWS VPC ID.
	VPC string `json:"VPC" msgpack:"VPC" bson:"vpc" mapstructure:"VPC,omitempty"`

	// A list of AWSEndpoint objects.
	Endpoints []*AWSEndpoint `json:"endpoints" msgpack:"endpoints" bson:"endpoints" mapstructure:"endpoints,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSAttachment returns a new *AWSAttachment
func NewAWSAttachment() *AWSAttachment {

	return &AWSAttachment{
		ModelVersion: 1,
		Endpoints:    []*AWSEndpoint{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSAttachment) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSAttachment{}

	s.VPC = o.VPC
	s.Endpoints = o.Endpoints

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSAttachment) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAWSAttachment{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPC = s.VPC
	o.Endpoints = s.Endpoints

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSAttachment) BleveType() string {

	return "awsattachment"
}

// DeepCopy returns a deep copy if the AWSAttachment.
func (o *AWSAttachment) DeepCopy() *AWSAttachment {

	if o == nil {
		return nil
	}

	out := &AWSAttachment{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSAttachment.
func (o *AWSAttachment) DeepCopyInto(out *AWSAttachment) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSAttachment: %s", err))
	}

	*out = *target.(*AWSAttachment)
}

// Validate valides the current information stored into the structure.
func (o *AWSAttachment) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("VPC", o.VPC); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.Endpoints {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*AWSAttachment) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSAttachmentAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSAttachmentLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSAttachment) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSAttachmentAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSAttachment) ValueForAttribute(name string) any {

	switch name {
	case "VPC":
		return o.VPC
	case "endpoints":
		return o.Endpoints
	}

	return nil
}

// AWSAttachmentAttributesMap represents the map of attribute for AWSAttachment.
var AWSAttachmentAttributesMap = map[string]elemental.AttributeSpecification{
	"VPC": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpc",
		ConvertedName:  "VPC",
		Description:    `AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPC",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Endpoints": {
		AllowedChoices: []string{},
		BSONFieldName:  "endpoints",
		ConvertedName:  "Endpoints",
		Description:    `A list of AWSEndpoint objects.`,
		Exposed:        true,
		Name:           "endpoints",
		Required:       true,
		Stored:         true,
		SubType:        "awsendpoint",
		Type:           "refList",
	},
}

// AWSAttachmentLowerCaseAttributesMap represents the map of attribute for AWSAttachment.
var AWSAttachmentLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpc": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpc",
		ConvertedName:  "VPC",
		Description:    `AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPC",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"endpoints": {
		AllowedChoices: []string{},
		BSONFieldName:  "endpoints",
		ConvertedName:  "Endpoints",
		Description:    `A list of AWSEndpoint objects.`,
		Exposed:        true,
		Name:           "endpoints",
		Required:       true,
		Stored:         true,
		SubType:        "awsendpoint",
		Type:           "refList",
	},
}

type mongoAttributesAWSAttachment struct {
	VPC       string         `bson:"vpc"`
	Endpoints []*AWSEndpoint `bson:"endpoints"`
}
