// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSEndpointStatusValue represents the possible values for attribute "status".
type AWSEndpointStatusValue string

const (
	// AWSEndpointStatusAccepted represents the value Accepted.
	AWSEndpointStatusAccepted AWSEndpointStatusValue = "Accepted"

	// AWSEndpointStatusPending represents the value Pending.
	AWSEndpointStatusPending AWSEndpointStatusValue = "Pending"

	// AWSEndpointStatusRejected represents the value Rejected.
	AWSEndpointStatusRejected AWSEndpointStatusValue = "Rejected"
)

// AWSEndpoint represents the model of a awsendpoint
type AWSEndpoint struct {
	// The AWS VPC ID.
	VPCID string `json:"VPCID" msgpack:"VPCID" bson:"vpcid" mapstructure:"VPCID,omitempty"`

	// The AWS endpoint created by the NGFW.
	EndpointID string `json:"endpointID" msgpack:"endpointID" bson:"endpointid" mapstructure:"endpointID,omitempty"`

	// The rejected reason.
	RejectedReason string `json:"rejectedReason" msgpack:"rejectedReason" bson:"rejectedreason" mapstructure:"rejectedReason,omitempty"`

	// The status of the of endpoint.
	Status AWSEndpointStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// The AWS subnet ID.
	SubnetID string `json:"subnetID" msgpack:"subnetID" bson:"subnetid" mapstructure:"subnetID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSEndpoint returns a new *AWSEndpoint
func NewAWSEndpoint() *AWSEndpoint {

	return &AWSEndpoint{
		ModelVersion: 1,
		Status:       AWSEndpointStatusPending,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSEndpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSEndpoint{}

	s.VPCID = o.VPCID
	s.EndpointID = o.EndpointID
	s.RejectedReason = o.RejectedReason
	s.Status = o.Status
	s.SubnetID = o.SubnetID

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSEndpoint) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAWSEndpoint{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCID = s.VPCID
	o.EndpointID = s.EndpointID
	o.RejectedReason = s.RejectedReason
	o.Status = s.Status
	o.SubnetID = s.SubnetID

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSEndpoint) BleveType() string {

	return "awsendpoint"
}

// DeepCopy returns a deep copy if the AWSEndpoint.
func (o *AWSEndpoint) DeepCopy() *AWSEndpoint {

	if o == nil {
		return nil
	}

	out := &AWSEndpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSEndpoint.
func (o *AWSEndpoint) DeepCopyInto(out *AWSEndpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSEndpoint: %s", err))
	}

	*out = *target.(*AWSEndpoint)
}

// Validate valides the current information stored into the structure.
func (o *AWSEndpoint) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Accepted", "Pending", "Rejected"}, true); err != nil {
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
func (*AWSEndpoint) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSEndpointAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSEndpointLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSEndpoint) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSEndpointAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSEndpoint) ValueForAttribute(name string) interface{} {

	switch name {
	case "VPCID":
		return o.VPCID
	case "endpointID":
		return o.EndpointID
	case "rejectedReason":
		return o.RejectedReason
	case "status":
		return o.Status
	case "subnetID":
		return o.SubnetID
	}

	return nil
}

// AWSEndpointAttributesMap represents the map of attribute for AWSEndpoint.
var AWSEndpointAttributesMap = map[string]elemental.AttributeSpecification{
	"VPCID": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `The AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"EndpointID": {
		AllowedChoices: []string{},
		BSONFieldName:  "endpointid",
		ConvertedName:  "EndpointID",
		Description:    `The AWS endpoint created by the NGFW.`,
		Exposed:        true,
		Name:           "endpointID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"RejectedReason": {
		AllowedChoices: []string{},
		BSONFieldName:  "rejectedreason",
		ConvertedName:  "RejectedReason",
		Description:    `The rejected reason.`,
		Exposed:        true,
		Name:           "rejectedReason",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"Accepted", "Pending", "Rejected"},
		Autogenerated:  true,
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   AWSEndpointStatusPending,
		Description:    `The status of the of endpoint.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"SubnetID": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetid",
		ConvertedName:  "SubnetID",
		Description:    `The AWS subnet ID.`,
		Exposed:        true,
		Name:           "subnetID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// AWSEndpointLowerCaseAttributesMap represents the map of attribute for AWSEndpoint.
var AWSEndpointLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpcid": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcid",
		ConvertedName:  "VPCID",
		Description:    `The AWS VPC ID.`,
		Exposed:        true,
		Name:           "VPCID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"endpointid": {
		AllowedChoices: []string{},
		BSONFieldName:  "endpointid",
		ConvertedName:  "EndpointID",
		Description:    `The AWS endpoint created by the NGFW.`,
		Exposed:        true,
		Name:           "endpointID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"rejectedreason": {
		AllowedChoices: []string{},
		BSONFieldName:  "rejectedreason",
		ConvertedName:  "RejectedReason",
		Description:    `The rejected reason.`,
		Exposed:        true,
		Name:           "rejectedReason",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"Accepted", "Pending", "Rejected"},
		Autogenerated:  true,
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   AWSEndpointStatusPending,
		Description:    `The status of the of endpoint.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"subnetid": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetid",
		ConvertedName:  "SubnetID",
		Description:    `The AWS subnet ID.`,
		Exposed:        true,
		Name:           "subnetID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesAWSEndpoint struct {
	VPCID          string                 `bson:"vpcid"`
	EndpointID     string                 `bson:"endpointid"`
	RejectedReason string                 `bson:"rejectedreason"`
	Status         AWSEndpointStatusValue `bson:"status"`
	SubnetID       string                 `bson:"subnetid"`
}
