// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallOptionVPC represents the model of a firewalloptionvpc
type FirewallOptionVPC struct {
	// An AWS VPC ID.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The availability zones associated with the VPC.
	AvailabilityZones []string `json:"availabilityZones" msgpack:"availabilityZones" bson:"-" mapstructure:"availabilityZones,omitempty"`

	// Names of any firewalls already associated with the VPC.
	ExistingFirewalls []string `json:"existingFirewalls" msgpack:"existingFirewalls" bson:"-" mapstructure:"existingFirewalls,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallOptionVPC returns a new *FirewallOptionVPC
func NewFirewallOptionVPC() *FirewallOptionVPC {

	return &FirewallOptionVPC{
		ModelVersion:      1,
		AvailabilityZones: []string{},
		ExistingFirewalls: []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallOptionVPC) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallOptionVPC{}

	s.Name = o.Name

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallOptionVPC) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallOptionVPC{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Name = s.Name

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallOptionVPC) BleveType() string {

	return "firewalloptionvpc"
}

// GetName returns the Name of the receiver.
func (o *FirewallOptionVPC) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *FirewallOptionVPC) SetName(name string) {

	o.Name = name
}

// DeepCopy returns a deep copy if the FirewallOptionVPC.
func (o *FirewallOptionVPC) DeepCopy() *FirewallOptionVPC {

	if o == nil {
		return nil
	}

	out := &FirewallOptionVPC{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallOptionVPC.
func (o *FirewallOptionVPC) DeepCopyInto(out *FirewallOptionVPC) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallOptionVPC: %s", err))
	}

	*out = *target.(*FirewallOptionVPC)
}

// Validate valides the current information stored into the structure.
func (o *FirewallOptionVPC) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateVPCID("ID", o.ID); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
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
func (*FirewallOptionVPC) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallOptionVPCAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallOptionVPCLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallOptionVPC) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallOptionVPCAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallOptionVPC) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "availabilityZones":
		return o.AvailabilityZones
	case "existingFirewalls":
		return o.ExistingFirewalls
	case "name":
		return o.Name
	}

	return nil
}

// FirewallOptionVPCAttributesMap represents the map of attribute for FirewallOptionVPC.
var FirewallOptionVPCAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"AvailabilityZones": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AvailabilityZones",
		Description:    `The availability zones associated with the VPC.`,
		Exposed:        true,
		Name:           "availabilityZones",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"ExistingFirewalls": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExistingFirewalls",
		Description:    `Names of any firewalls already associated with the VPC.`,
		Exposed:        true,
		Name:           "existingFirewalls",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// FirewallOptionVPCLowerCaseAttributesMap represents the map of attribute for FirewallOptionVPC.
var FirewallOptionVPCLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"availabilityzones": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AvailabilityZones",
		Description:    `The availability zones associated with the VPC.`,
		Exposed:        true,
		Name:           "availabilityZones",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"existingfirewalls": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ExistingFirewalls",
		Description:    `Names of any firewalls already associated with the VPC.`,
		Exposed:        true,
		Name:           "existingFirewalls",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesFirewallOptionVPC struct {
	Name string `bson:"name"`
}
