// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TopQueryItem represents the model of a topqueryitem
type TopQueryItem struct {
	// The count of the field.
	Count int `json:"count" msgpack:"count" bson:"-" mapstructure:"count,omitempty"`

	// The value of the field.
	Value string `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTopQueryItem returns a new *TopQueryItem
func NewTopQueryItem() *TopQueryItem {

	return &TopQueryItem{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TopQueryItem) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTopQueryItem{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TopQueryItem) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTopQueryItem{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *TopQueryItem) BleveType() string {

	return "topqueryitem"
}

// DeepCopy returns a deep copy if the TopQueryItem.
func (o *TopQueryItem) DeepCopy() *TopQueryItem {

	if o == nil {
		return nil
	}

	out := &TopQueryItem{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TopQueryItem.
func (o *TopQueryItem) DeepCopyInto(out *TopQueryItem) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TopQueryItem: %s", err))
	}

	*out = *target.(*TopQueryItem)
}

// Validate valides the current information stored into the structure.
func (o *TopQueryItem) Validate() error {

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
func (*TopQueryItem) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TopQueryItemAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TopQueryItemLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TopQueryItem) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TopQueryItemAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TopQueryItem) ValueForAttribute(name string) any {

	switch name {
	case "count":
		return o.Count
	case "value":
		return o.Value
	}

	return nil
}

// TopQueryItemAttributesMap represents the map of attribute for TopQueryItem.
var TopQueryItemAttributesMap = map[string]elemental.AttributeSpecification{
	"Count": {
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `The count of the field.`,
		Exposed:        true,
		Name:           "count",
		Type:           "integer",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `The value of the field.`,
		Exposed:        true,
		Name:           "value",
		Type:           "string",
	},
}

// TopQueryItemLowerCaseAttributesMap represents the map of attribute for TopQueryItem.
var TopQueryItemLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"count": {
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `The count of the field.`,
		Exposed:        true,
		Name:           "count",
		Type:           "integer",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `The value of the field.`,
		Exposed:        true,
		Name:           "value",
		Type:           "string",
	},
}

type mongoAttributesTopQueryItem struct {
}
