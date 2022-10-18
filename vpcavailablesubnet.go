// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// VpcAvailableSubnet represents the model of a vpcavailablesubnet
type VpcAvailableSubnet struct {
	// An AWS VPC ID.
	VPCID string `json:"VPCID" msgpack:"VPCID" bson:"-" mapstructure:"VPCID,omitempty"`

	// The list of all availability zones and associated subnets for every VPC
	// specified.
	AvailabilityZoneSubnets []*AvailabilityZoneSubnet `json:"availabilityZoneSubnets" msgpack:"availabilityZoneSubnets" bson:"-" mapstructure:"availabilityZoneSubnets,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewVpcAvailableSubnet returns a new *VpcAvailableSubnet
func NewVpcAvailableSubnet() *VpcAvailableSubnet {

	return &VpcAvailableSubnet{
		ModelVersion:            1,
		AvailabilityZoneSubnets: []*AvailabilityZoneSubnet{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VpcAvailableSubnet) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesVpcAvailableSubnet{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VpcAvailableSubnet) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesVpcAvailableSubnet{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *VpcAvailableSubnet) BleveType() string {

	return "vpcavailablesubnet"
}

// DeepCopy returns a deep copy if the VpcAvailableSubnet.
func (o *VpcAvailableSubnet) DeepCopy() *VpcAvailableSubnet {

	if o == nil {
		return nil
	}

	out := &VpcAvailableSubnet{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *VpcAvailableSubnet.
func (o *VpcAvailableSubnet) DeepCopyInto(out *VpcAvailableSubnet) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy VpcAvailableSubnet: %s", err))
	}

	*out = *target.(*VpcAvailableSubnet)
}

// Validate valides the current information stored into the structure.
func (o *VpcAvailableSubnet) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.AvailabilityZoneSubnets {
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
func (*VpcAvailableSubnet) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := VpcAvailableSubnetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return VpcAvailableSubnetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*VpcAvailableSubnet) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return VpcAvailableSubnetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *VpcAvailableSubnet) ValueForAttribute(name string) interface{} {

	switch name {
	case "VPCID":
		return o.VPCID
	case "availabilityZoneSubnets":
		return o.AvailabilityZoneSubnets
	}

	return nil
}

// VpcAvailableSubnetAttributesMap represents the map of attribute for VpcAvailableSubnet.
var VpcAvailableSubnetAttributesMap = map[string]elemental.AttributeSpecification{
	"VPCID": {
		AllowedChoices: []string{},
		ConvertedName:  "VPCID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		Type:           "string",
	},
	"AvailabilityZoneSubnets": {
		AllowedChoices: []string{},
		ConvertedName:  "AvailabilityZoneSubnets",
		Description: `The list of all availability zones and associated subnets for every VPC
specified.`,
		Exposed: true,
		Name:    "availabilityZoneSubnets",
		SubType: "availabilityzonesubnet",
		Type:    "refList",
	},
}

// VpcAvailableSubnetLowerCaseAttributesMap represents the map of attribute for VpcAvailableSubnet.
var VpcAvailableSubnetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpcid": {
		AllowedChoices: []string{},
		ConvertedName:  "VPCID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		Type:           "string",
	},
	"availabilityzonesubnets": {
		AllowedChoices: []string{},
		ConvertedName:  "AvailabilityZoneSubnets",
		Description: `The list of all availability zones and associated subnets for every VPC
specified.`,
		Exposed: true,
		Name:    "availabilityZoneSubnets",
		SubType: "availabilityzonesubnet",
		Type:    "refList",
	},
}

type mongoAttributesVpcAvailableSubnet struct {
}
