// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// LogQueryItem represents the model of a logqueryitem
type LogQueryItem struct {
	// The fields and values for the log line.
	Fields map[string]string `json:"fields" msgpack:"fields" bson:"-" mapstructure:"fields,omitempty"`

	// The timestamp of the log line.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewLogQueryItem returns a new *LogQueryItem
func NewLogQueryItem() *LogQueryItem {

	return &LogQueryItem{
		ModelVersion: 1,
		Fields:       map[string]string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LogQueryItem) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesLogQueryItem{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *LogQueryItem) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesLogQueryItem{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *LogQueryItem) BleveType() string {

	return "logqueryitem"
}

// DeepCopy returns a deep copy if the LogQueryItem.
func (o *LogQueryItem) DeepCopy() *LogQueryItem {

	if o == nil {
		return nil
	}

	out := &LogQueryItem{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *LogQueryItem.
func (o *LogQueryItem) DeepCopyInto(out *LogQueryItem) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy LogQueryItem: %s", err))
	}

	*out = *target.(*LogQueryItem)
}

// Validate valides the current information stored into the structure.
func (o *LogQueryItem) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*LogQueryItem) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := LogQueryItemAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return LogQueryItemLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*LogQueryItem) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return LogQueryItemAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *LogQueryItem) ValueForAttribute(name string) interface{} {

	switch name {
	case "fields":
		return o.Fields
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// LogQueryItemAttributesMap represents the map of attribute for LogQueryItem.
var LogQueryItemAttributesMap = map[string]elemental.AttributeSpecification{
	"Fields": {
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `The fields and values for the log line.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the log line.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// LogQueryItemLowerCaseAttributesMap represents the map of attribute for LogQueryItem.
var LogQueryItemLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fields": {
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `The fields and values for the log line.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "map[string]string",
		Type:           "external",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the log line.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

type mongoAttributesLogQueryItem struct {
}
