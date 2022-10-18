// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallCountryCode represents the model of a firewallcountrycode
type FirewallCountryCode struct {
	// The Country description.
	Description string `json:"description" msgpack:"description" bson:"-" mapstructure:"description,omitempty"`

	// The country code name.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallCountryCode returns a new *FirewallCountryCode
func NewFirewallCountryCode() *FirewallCountryCode {

	return &FirewallCountryCode{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCountryCode) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallCountryCode{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCountryCode) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallCountryCode{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallCountryCode) BleveType() string {

	return "firewallcountrycode"
}

// DeepCopy returns a deep copy if the FirewallCountryCode.
func (o *FirewallCountryCode) DeepCopy() *FirewallCountryCode {

	if o == nil {
		return nil
	}

	out := &FirewallCountryCode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallCountryCode.
func (o *FirewallCountryCode) DeepCopyInto(out *FirewallCountryCode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallCountryCode: %s", err))
	}

	*out = *target.(*FirewallCountryCode)
}

// Validate valides the current information stored into the structure.
func (o *FirewallCountryCode) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*FirewallCountryCode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallCountryCodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallCountryCodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallCountryCode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallCountryCodeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallCountryCode) ValueForAttribute(name string) interface{} {

	switch name {
	case "description":
		return o.Description
	case "name":
		return o.Name
	}

	return nil
}

// FirewallCountryCodeAttributesMap represents the map of attribute for FirewallCountryCode.
var FirewallCountryCodeAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The Country description.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The country code name.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
}

// FirewallCountryCodeLowerCaseAttributesMap represents the map of attribute for FirewallCountryCode.
var FirewallCountryCodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The Country description.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The country code name.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Type:           "string",
	},
}

type mongoAttributesFirewallCountryCode struct {
}
