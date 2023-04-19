// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// VPCUsedSubnet represents the model of a vpcusedsubnet
type VPCUsedSubnet struct {
	// The CIDR block of the VPC.
	CIDRBlock string `json:"CIDRBlock" msgpack:"CIDRBlock" bson:"cidrblock" mapstructure:"CIDRBlock,omitempty"`

	// An AWS VPC ID.
	VPCID string `json:"VPCID" msgpack:"VPCID" bson:"vpcid" mapstructure:"VPCID,omitempty"`

	// The list of all availability zones associated with the VPC.
	AvailabilityZones []string `json:"availabilityZones" msgpack:"availabilityZones" bson:"availabilityzones" mapstructure:"availabilityZones,omitempty"`

	// The list of all currently used subnet CIDR blocks in this VPC.
	SubnetCIDRs []string `json:"subnetCIDRs" msgpack:"subnetCIDRs" bson:"subnetcidrs" mapstructure:"subnetCIDRs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewVPCUsedSubnet returns a new *VPCUsedSubnet
func NewVPCUsedSubnet() *VPCUsedSubnet {

	return &VPCUsedSubnet{
		ModelVersion:      1,
		AvailabilityZones: []string{},
		SubnetCIDRs:       []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCUsedSubnet) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesVPCUsedSubnet{}

	s.CIDRBlock = o.CIDRBlock
	s.VPCID = o.VPCID
	s.AvailabilityZones = o.AvailabilityZones
	s.SubnetCIDRs = o.SubnetCIDRs

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *VPCUsedSubnet) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesVPCUsedSubnet{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.CIDRBlock = s.CIDRBlock
	o.VPCID = s.VPCID
	o.AvailabilityZones = s.AvailabilityZones
	o.SubnetCIDRs = s.SubnetCIDRs

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *VPCUsedSubnet) BleveType() string {

	return "vpcusedsubnet"
}

// DeepCopy returns a deep copy if the VPCUsedSubnet.
func (o *VPCUsedSubnet) DeepCopy() *VPCUsedSubnet {

	if o == nil {
		return nil
	}

	out := &VPCUsedSubnet{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *VPCUsedSubnet.
func (o *VPCUsedSubnet) DeepCopyInto(out *VPCUsedSubnet) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy VPCUsedSubnet: %s", err))
	}

	*out = *target.(*VPCUsedSubnet)
}

// Validate valides the current information stored into the structure.
func (o *VPCUsedSubnet) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateCIDR("CIDRBlock", o.CIDRBlock); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateVPCID("VPCID", o.VPCID); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateCIDRList("subnetCIDRs", o.SubnetCIDRs); err != nil {
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
func (*VPCUsedSubnet) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := VPCUsedSubnetAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return VPCUsedSubnetLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*VPCUsedSubnet) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return VPCUsedSubnetAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *VPCUsedSubnet) ValueForAttribute(name string) any {

	switch name {
	case "CIDRBlock":
		return o.CIDRBlock
	case "VPCID":
		return o.VPCID
	case "availabilityZones":
		return o.AvailabilityZones
	case "subnetCIDRs":
		return o.SubnetCIDRs
	}

	return nil
}

// VPCUsedSubnetAttributesMap represents the map of attribute for VPCUsedSubnet.
var VPCUsedSubnetAttributesMap = map[string]elemental.AttributeSpecification{
	"CIDRBlock": {
		AllowedChoices: []string{},
		BSONFieldName:  "cidrblock",
		ConvertedName:  "CIDRBlock",
		Description:    `The CIDR block of the VPC.`,
		Exposed:        true,
		Name:           "CIDRBlock",
		Stored:         true,
		Type:           "string",
	},
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
	"AvailabilityZones": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzones",
		ConvertedName:  "AvailabilityZones",
		Description:    `The list of all availability zones associated with the VPC.`,
		Exposed:        true,
		Name:           "availabilityZones",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"SubnetCIDRs": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetcidrs",
		ConvertedName:  "SubnetCIDRs",
		Description:    `The list of all currently used subnet CIDR blocks in this VPC.`,
		Exposed:        true,
		Name:           "subnetCIDRs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

// VPCUsedSubnetLowerCaseAttributesMap represents the map of attribute for VPCUsedSubnet.
var VPCUsedSubnetLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cidrblock": {
		AllowedChoices: []string{},
		BSONFieldName:  "cidrblock",
		ConvertedName:  "CIDRBlock",
		Description:    `The CIDR block of the VPC.`,
		Exposed:        true,
		Name:           "CIDRBlock",
		Stored:         true,
		Type:           "string",
	},
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
	"availabilityzones": {
		AllowedChoices: []string{},
		BSONFieldName:  "availabilityzones",
		ConvertedName:  "AvailabilityZones",
		Description:    `The list of all availability zones associated with the VPC.`,
		Exposed:        true,
		Name:           "availabilityZones",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"subnetcidrs": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetcidrs",
		ConvertedName:  "SubnetCIDRs",
		Description:    `The list of all currently used subnet CIDR blocks in this VPC.`,
		Exposed:        true,
		Name:           "subnetCIDRs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
}

type mongoAttributesVPCUsedSubnet struct {
	CIDRBlock         string   `bson:"cidrblock"`
	VPCID             string   `bson:"vpcid"`
	AvailabilityZones []string `bson:"availabilityzones"`
	SubnetCIDRs       []string `bson:"subnetcidrs"`
}
