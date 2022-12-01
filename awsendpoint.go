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

	// AWSEndpointStatusCreating represents the value Creating.
	AWSEndpointStatusCreating AWSEndpointStatusValue = "Creating"

	// AWSEndpointStatusDeleting represents the value Deleting.
	AWSEndpointStatusDeleting AWSEndpointStatusValue = "Deleting"

	// AWSEndpointStatusFailed represents the value Failed.
	AWSEndpointStatusFailed AWSEndpointStatusValue = "Failed"
)

// AWSEndpoint represents the model of a awsendpoint
type AWSEndpoint struct {
	// The AWS endpoint created by the NGFW.
	EndpointID string `json:"endpointID" msgpack:"endpointID" bson:"endpointid" mapstructure:"endpointID,omitempty"`

	// The previous endpoint Status.
	PreviousStatus string `json:"-" msgpack:"-" bson:"previousstatus" mapstructure:"-,omitempty"`

	// The status of the of endpoint.
	Status AWSEndpointStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// The status description of endpoint.
	StatusReason string `json:"statusReason" msgpack:"statusReason" bson:"statusreason" mapstructure:"statusReason,omitempty"`

	// The AWS subnet ID.
	SubnetID string `json:"subnetID" msgpack:"subnetID" bson:"subnetid" mapstructure:"subnetID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSEndpoint returns a new *AWSEndpoint
func NewAWSEndpoint() *AWSEndpoint {

	return &AWSEndpoint{
		ModelVersion: 1,
		Status:       AWSEndpointStatusCreating,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSEndpoint) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSEndpoint{}

	s.EndpointID = o.EndpointID
	s.PreviousStatus = o.PreviousStatus
	s.Status = o.Status
	s.StatusReason = o.StatusReason
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

	o.EndpointID = s.EndpointID
	o.PreviousStatus = s.PreviousStatus
	o.Status = s.Status
	o.StatusReason = s.StatusReason
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

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Creating", "Deleting", "Accepted", "Failed"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("subnetID", o.SubnetID); err != nil {
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
	case "endpointID":
		return o.EndpointID
	case "previousStatus":
		return o.PreviousStatus
	case "status":
		return o.Status
	case "statusReason":
		return o.StatusReason
	case "subnetID":
		return o.SubnetID
	}

	return nil
}

// AWSEndpointAttributesMap represents the map of attribute for AWSEndpoint.
var AWSEndpointAttributesMap = map[string]elemental.AttributeSpecification{
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

	"Status": {
		AllowedChoices: []string{"Creating", "Deleting", "Accepted", "Failed"},
		Autogenerated:  true,
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   AWSEndpointStatusCreating,
		Description:    `The status of the of endpoint.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"StatusReason": {
		AllowedChoices: []string{},
		BSONFieldName:  "statusreason",
		ConvertedName:  "StatusReason",
		Description:    `The status description of endpoint.`,
		Exposed:        true,
		Name:           "statusReason",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"SubnetID": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetid",
		ConvertedName:  "SubnetID",
		Description:    `The AWS subnet ID.`,
		Exposed:        true,
		Name:           "subnetID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// AWSEndpointLowerCaseAttributesMap represents the map of attribute for AWSEndpoint.
var AWSEndpointLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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

	"status": {
		AllowedChoices: []string{"Creating", "Deleting", "Accepted", "Failed"},
		Autogenerated:  true,
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   AWSEndpointStatusCreating,
		Description:    `The status of the of endpoint.`,
		Exposed:        true,
		Name:           "status",
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"statusreason": {
		AllowedChoices: []string{},
		BSONFieldName:  "statusreason",
		ConvertedName:  "StatusReason",
		Description:    `The status description of endpoint.`,
		Exposed:        true,
		Name:           "statusReason",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"subnetid": {
		AllowedChoices: []string{},
		BSONFieldName:  "subnetid",
		ConvertedName:  "SubnetID",
		Description:    `The AWS subnet ID.`,
		Exposed:        true,
		Name:           "subnetID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

type mongoAttributesAWSEndpoint struct {
	EndpointID     string                 `bson:"endpointid"`
	PreviousStatus string                 `bson:"previousstatus"`
	Status         AWSEndpointStatusValue `bson:"status"`
	StatusReason   string                 `bson:"statusreason"`
	SubnetID       string                 `bson:"subnetid"`
}
