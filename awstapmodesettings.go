// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSTAPModeSettings represents the model of a awstapmodesettings
type AWSTAPModeSettings struct {
	// An awslogdefinition ID.
	LogDefinitionID string `json:"logDefinitionID" msgpack:"logDefinitionID" bson:"logdefinitionid" mapstructure:"logDefinitionID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSTAPModeSettings returns a new *AWSTAPModeSettings
func NewAWSTAPModeSettings() *AWSTAPModeSettings {

	return &AWSTAPModeSettings{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSTAPModeSettings) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSTAPModeSettings{}

	s.LogDefinitionID = o.LogDefinitionID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSTAPModeSettings) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAWSTAPModeSettings{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.LogDefinitionID = s.LogDefinitionID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSTAPModeSettings) BleveType() string {

	return "awstapmodesettings"
}

// DeepCopy returns a deep copy if the AWSTAPModeSettings.
func (o *AWSTAPModeSettings) DeepCopy() *AWSTAPModeSettings {

	if o == nil {
		return nil
	}

	out := &AWSTAPModeSettings{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSTAPModeSettings.
func (o *AWSTAPModeSettings) DeepCopyInto(out *AWSTAPModeSettings) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSTAPModeSettings: %s", err))
	}

	*out = *target.(*AWSTAPModeSettings)
}

// Validate valides the current information stored into the structure.
func (o *AWSTAPModeSettings) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("logDefinitionID", o.LogDefinitionID); err != nil {
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
func (*AWSTAPModeSettings) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSTAPModeSettingsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSTAPModeSettingsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSTAPModeSettings) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSTAPModeSettingsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSTAPModeSettings) ValueForAttribute(name string) interface{} {

	switch name {
	case "logDefinitionID":
		return o.LogDefinitionID
	}

	return nil
}

// AWSTAPModeSettingsAttributesMap represents the map of attribute for AWSTAPModeSettings.
var AWSTAPModeSettingsAttributesMap = map[string]elemental.AttributeSpecification{
	"LogDefinitionID": {
		AllowedChoices: []string{},
		BSONFieldName:  "logdefinitionid",
		ConvertedName:  "LogDefinitionID",
		Description:    `An awslogdefinition ID.`,
		Exposed:        true,
		Name:           "logDefinitionID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// AWSTAPModeSettingsLowerCaseAttributesMap represents the map of attribute for AWSTAPModeSettings.
var AWSTAPModeSettingsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"logdefinitionid": {
		AllowedChoices: []string{},
		BSONFieldName:  "logdefinitionid",
		ConvertedName:  "LogDefinitionID",
		Description:    `An awslogdefinition ID.`,
		Exposed:        true,
		Name:           "logDefinitionID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesAWSTAPModeSettings struct {
	LogDefinitionID string `bson:"logdefinitionid"`
}
