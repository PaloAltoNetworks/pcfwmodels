// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// VPCAvailabilityZoneSubnet represents the model of a vpcavailabilityzonesubnet
type VPCAvailabilityZoneSubnet struct {
	// An AWS VPC ID.
	VPCID string `json:"VPCID" msgpack:"VPCID" bson:"vpcid" mapstructure:"VPCID,omitempty"`

	// The list of all firewall subnets (one in each AZ) and associated network
	// interfaces in that AZ.
	SubnetInterfaces []*AvailabilityZoneSubnetInterface `json:"subnetInterfaces" msgpack:"subnetInterfaces" bson:"subnetinterfaces" mapstructure:"subnetInterfaces,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewVPCAvailabilityZoneSubnet returns a new *VPCAvailabilityZoneSubnet
func NewVPCAvailabilityZoneSubnet() *VPCAvailabilityZoneSubnet {

	return &VPCAvailabilityZoneSubnet{
		ModelVersion:     1,
		SubnetInterfaces: []*AvailabilityZoneSubnetInterface{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCAvailabilityZoneSubnet) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesVPCAvailabilityZoneSubnet{}

	s.VPCID = o.VPCID
	s.SubnetInterfaces = o.SubnetInterfaces

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCAvailabilityZoneSubnet) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesVPCAvailabilityZoneSubnet{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCID = s.VPCID
	o.SubnetInterfaces = s.SubnetInterfaces

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *VPCAvailabilityZoneSubnet) BleveType() string {

	return "vpcavailabilityzonesubnet"
}

// DeepCopy returns a deep copy if the VPCAvailabilityZoneSubnet.
func (o *VPCAvailabilityZoneSubnet) DeepCopy() *VPCAvailabilityZoneSubnet {

	if o == nil {
		return nil
	}

	out := &VPCAvailabilityZoneSubnet{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *VPCAvailabilityZoneSubnet.
func (o *VPCAvailabilityZoneSubnet) DeepCopyInto(out *VPCAvailabilityZoneSubnet) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy VPCAvailabilityZoneSubnet: %s", err))
	}

	*out = *target.(*VPCAvailabilityZoneSubnet)
}

// Validate valides the current information stored into the structure.
func (o *VPCAvailabilityZoneSubnet) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateVPCID("VPCID", o.VPCID); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.SubnetInterfaces {
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
func (*VPCAvailabilityZoneSubnet) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := VPCAvailabilityZoneSubnetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return VPCAvailabilityZoneSubnetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*VPCAvailabilityZoneSubnet) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return VPCAvailabilityZoneSubnetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *VPCAvailabilityZoneSubnet) ValueForAttribute(name string) interface{} {

	switch name {
	case "VPCID":
		return o.VPCID
	case "subnetInterfaces":
		return o.SubnetInterfaces
	}

	return nil
}

// VPCAvailabilityZoneSubnetAttributesMap represents the map of attribute for VPCAvailabilityZoneSubnet.
var VPCAvailabilityZoneSubnetAttributesMap = map[string]elemental.AttributeSpecification{
	"VPCID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		Stored:         true,
		Type:           "string",
	},
	"SubnetInterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetinterfaces",
		ConvertedName:  "SubnetInterfaces",
		Description: `The list of all firewall subnets (one in each AZ) and associated network
interfaces in that AZ.`,
		Exposed: true,
		Name:    "subnetInterfaces",
		Stored:  true,
		SubType: "availabilityzonesubnetinterface",
		Type:    "refList",
	},
}

// VPCAvailabilityZoneSubnetLowerCaseAttributesMap represents the map of attribute for VPCAvailabilityZoneSubnet.
var VPCAvailabilityZoneSubnetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `An AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		Stored:         true,
		Type:           "string",
	},
	"subnetinterfaces": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetinterfaces",
		ConvertedName:  "SubnetInterfaces",
		Description: `The list of all firewall subnets (one in each AZ) and associated network
interfaces in that AZ.`,
		Exposed: true,
		Name:    "subnetInterfaces",
		Stored:  true,
		SubType: "availabilityzonesubnetinterface",
		Type:    "refList",
	},
}

type mongoAttributesVPCAvailabilityZoneSubnet struct {
	VPCID            string                             `bson:"vpcid"`
	SubnetInterfaces []*AvailabilityZoneSubnetInterface `bson:"subnetinterfaces"`
}
