// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MirrorRuleActionValue represents the possible values for attribute "action".
type MirrorRuleActionValue string

const (
	// MirrorRuleActionAccept represents the value Accept.
	MirrorRuleActionAccept MirrorRuleActionValue = "Accept"

	// MirrorRuleActionReject represents the value Reject.
	MirrorRuleActionReject MirrorRuleActionValue = "Reject"
)

// MirrorRuleDirectionValue represents the possible values for attribute "direction".
type MirrorRuleDirectionValue string

const (
	// MirrorRuleDirectionEgress represents the value Egress.
	MirrorRuleDirectionEgress MirrorRuleDirectionValue = "Egress"

	// MirrorRuleDirectionIngress represents the value Ingress.
	MirrorRuleDirectionIngress MirrorRuleDirectionValue = "Ingress"
)

// MirrorRule represents the model of a mirrorrule
type MirrorRule struct {
	// The action to take on the filtered traffic.
	Action MirrorRuleActionValue `json:"action" msgpack:"action" bson:"action" mapstructure:"action,omitempty"`

	// Description of the mirror rule.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Destination CIDR block to assign to the mirror rule.
	DestinationCIDR string `json:"destinationCIDR" msgpack:"destinationCIDR" bson:"destinationcidr" mapstructure:"destinationCIDR,omitempty"`

	// Destination port range start.
	DestinationFromPort int `json:"destinationFromPort" msgpack:"destinationFromPort" bson:"destinationfromport" mapstructure:"destinationFromPort,omitempty"`

	// Destination port range end.
	DestinationToPort int `json:"destinationToPort" msgpack:"destinationToPort" bson:"destinationtoport" mapstructure:"destinationToPort,omitempty"`

	// The direction of the traffic to be mirrored.
	Direction MirrorRuleDirectionValue `json:"direction" msgpack:"direction" bson:"direction" mapstructure:"direction,omitempty"`

	// Number of a traffic mirror rule. Must be unique in each direction.
	Number int `json:"number" msgpack:"number" bson:"number" mapstructure:"number,omitempty"`

	// Protocol number to assign to the mirror rule.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Source CIDR block to assign to the mirror rule.
	SourceCIDR string `json:"sourceCIDR" msgpack:"sourceCIDR" bson:"sourcecidr" mapstructure:"sourceCIDR,omitempty"`

	// Source port range start.
	SourceFromPort int `json:"sourceFromPort" msgpack:"sourceFromPort" bson:"sourcefromport" mapstructure:"sourceFromPort,omitempty"`

	// Source port range end.
	SourceToPort int `json:"sourceToPort" msgpack:"sourceToPort" bson:"sourcetoport" mapstructure:"sourceToPort,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMirrorRule returns a new *MirrorRule
func NewMirrorRule() *MirrorRule {

	return &MirrorRule{
		ModelVersion:        1,
		DestinationCIDR:     "0.0.0.0/0",
		DestinationFromPort: 0,
		DestinationToPort:   65535,
		Protocol:            -1,
		SourceCIDR:          "0.0.0.0/0",
		SourceFromPort:      0,
		SourceToPort:        65535,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MirrorRule) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMirrorRule{}

	s.Action = o.Action
	s.Description = o.Description
	s.DestinationCIDR = o.DestinationCIDR
	s.DestinationFromPort = o.DestinationFromPort
	s.DestinationToPort = o.DestinationToPort
	s.Direction = o.Direction
	s.Number = o.Number
	s.Protocol = o.Protocol
	s.SourceCIDR = o.SourceCIDR
	s.SourceFromPort = o.SourceFromPort
	s.SourceToPort = o.SourceToPort

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MirrorRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMirrorRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Action = s.Action
	o.Description = s.Description
	o.DestinationCIDR = s.DestinationCIDR
	o.DestinationFromPort = s.DestinationFromPort
	o.DestinationToPort = s.DestinationToPort
	o.Direction = s.Direction
	o.Number = s.Number
	o.Protocol = s.Protocol
	o.SourceCIDR = s.SourceCIDR
	o.SourceFromPort = s.SourceFromPort
	o.SourceToPort = s.SourceToPort

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *MirrorRule) BleveType() string {

	return "mirrorrule"
}

// DeepCopy returns a deep copy if the MirrorRule.
func (o *MirrorRule) DeepCopy() *MirrorRule {

	if o == nil {
		return nil
	}

	out := &MirrorRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *MirrorRule.
func (o *MirrorRule) DeepCopyInto(out *MirrorRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy MirrorRule: %s", err))
	}

	*out = *target.(*MirrorRule)
}

// Validate valides the current information stored into the structure.
func (o *MirrorRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateCIDR("destinationCIDR", o.DestinationCIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePort("destinationFromPort", o.DestinationFromPort); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePort("destinationToPort", o.DestinationToPort); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("direction", string(o.Direction)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("direction", string(o.Direction), []string{"Ingress", "Egress"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("number", o.Number); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(255), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateCIDR("sourceCIDR", o.SourceCIDR); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePort("sourceFromPort", o.SourceFromPort); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidatePort("sourceToPort", o.SourceToPort); err != nil {
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
func (*MirrorRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MirrorRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MirrorRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MirrorRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MirrorRuleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *MirrorRule) ValueForAttribute(name string) any {

	switch name {
	case "action":
		return o.Action
	case "description":
		return o.Description
	case "destinationCIDR":
		return o.DestinationCIDR
	case "destinationFromPort":
		return o.DestinationFromPort
	case "destinationToPort":
		return o.DestinationToPort
	case "direction":
		return o.Direction
	case "number":
		return o.Number
	case "protocol":
		return o.Protocol
	case "sourceCIDR":
		return o.SourceCIDR
	case "sourceFromPort":
		return o.SourceFromPort
	case "sourceToPort":
		return o.SourceToPort
	}

	return nil
}

// MirrorRuleAttributesMap represents the map of attribute for MirrorRule.
var MirrorRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{"Accept", "Reject"},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		Description:    `The action to take on the filtered traffic.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the mirror rule.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	"DestinationCIDR": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationcidr",
		ConvertedName:  "DestinationCIDR",
		DefaultValue:   "0.0.0.0/0",
		Description:    `Destination CIDR block to assign to the mirror rule.`,
		Exposed:        true,
		Name:           "destinationCIDR",
		Stored:         true,
		Type:           "string",
	},
	"DestinationFromPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationfromport",
		ConvertedName:  "DestinationFromPort",
		Description:    `Destination port range start.`,
		Exposed:        true,
		Name:           "destinationFromPort",
		Stored:         true,
		Type:           "integer",
	},
	"DestinationToPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationtoport",
		ConvertedName:  "DestinationToPort",
		DefaultValue:   65535,
		Description:    `Destination port range end.`,
		Exposed:        true,
		Name:           "destinationToPort",
		Stored:         true,
		Type:           "integer",
	},
	"Direction": {
		AllowedChoices: []string{"Ingress", "Egress"},
		BSONFieldName:  "direction",
		ConvertedName:  "Direction",
		Description:    `The direction of the traffic to be mirrored.`,
		Exposed:        true,
		Name:           "direction",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Number": {
		AllowedChoices: []string{},
		BSONFieldName:  "number",
		ConvertedName:  "Number",
		Description:    `Number of a traffic mirror rule. Must be unique in each direction.`,
		Exposed:        true,
		Name:           "number",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Protocol": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocol",
		ConvertedName:  "Protocol",
		DefaultValue:   -1,
		Description:    `Protocol number to assign to the mirror rule.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"SourceCIDR": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcecidr",
		ConvertedName:  "SourceCIDR",
		DefaultValue:   "0.0.0.0/0",
		Description:    `Source CIDR block to assign to the mirror rule.`,
		Exposed:        true,
		Name:           "sourceCIDR",
		Stored:         true,
		Type:           "string",
	},
	"SourceFromPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcefromport",
		ConvertedName:  "SourceFromPort",
		Description:    `Source port range start.`,
		Exposed:        true,
		Name:           "sourceFromPort",
		Stored:         true,
		Type:           "integer",
	},
	"SourceToPort": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcetoport",
		ConvertedName:  "SourceToPort",
		DefaultValue:   65535,
		Description:    `Source port range end.`,
		Exposed:        true,
		Name:           "sourceToPort",
		Stored:         true,
		Type:           "integer",
	},
}

// MirrorRuleLowerCaseAttributesMap represents the map of attribute for MirrorRule.
var MirrorRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{"Accept", "Reject"},
		BSONFieldName:  "action",
		ConvertedName:  "Action",
		Description:    `The action to take on the filtered traffic.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the mirror rule.`,
		Exposed:        true,
		Name:           "description",
		Stored:         true,
		Type:           "string",
	},
	"destinationcidr": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationcidr",
		ConvertedName:  "DestinationCIDR",
		DefaultValue:   "0.0.0.0/0",
		Description:    `Destination CIDR block to assign to the mirror rule.`,
		Exposed:        true,
		Name:           "destinationCIDR",
		Stored:         true,
		Type:           "string",
	},
	"destinationfromport": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationfromport",
		ConvertedName:  "DestinationFromPort",
		Description:    `Destination port range start.`,
		Exposed:        true,
		Name:           "destinationFromPort",
		Stored:         true,
		Type:           "integer",
	},
	"destinationtoport": {
		AllowedChoices: []string{},
		BSONFieldName:  "destinationtoport",
		ConvertedName:  "DestinationToPort",
		DefaultValue:   65535,
		Description:    `Destination port range end.`,
		Exposed:        true,
		Name:           "destinationToPort",
		Stored:         true,
		Type:           "integer",
	},
	"direction": {
		AllowedChoices: []string{"Ingress", "Egress"},
		BSONFieldName:  "direction",
		ConvertedName:  "Direction",
		Description:    `The direction of the traffic to be mirrored.`,
		Exposed:        true,
		Name:           "direction",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"number": {
		AllowedChoices: []string{},
		BSONFieldName:  "number",
		ConvertedName:  "Number",
		Description:    `Number of a traffic mirror rule. Must be unique in each direction.`,
		Exposed:        true,
		Name:           "number",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"protocol": {
		AllowedChoices: []string{},
		BSONFieldName:  "protocol",
		ConvertedName:  "Protocol",
		DefaultValue:   -1,
		Description:    `Protocol number to assign to the mirror rule.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"sourcecidr": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcecidr",
		ConvertedName:  "SourceCIDR",
		DefaultValue:   "0.0.0.0/0",
		Description:    `Source CIDR block to assign to the mirror rule.`,
		Exposed:        true,
		Name:           "sourceCIDR",
		Stored:         true,
		Type:           "string",
	},
	"sourcefromport": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcefromport",
		ConvertedName:  "SourceFromPort",
		Description:    `Source port range start.`,
		Exposed:        true,
		Name:           "sourceFromPort",
		Stored:         true,
		Type:           "integer",
	},
	"sourcetoport": {
		AllowedChoices: []string{},
		BSONFieldName:  "sourcetoport",
		ConvertedName:  "SourceToPort",
		DefaultValue:   65535,
		Description:    `Source port range end.`,
		Exposed:        true,
		Name:           "sourceToPort",
		Stored:         true,
		Type:           "integer",
	},
}

type mongoAttributesMirrorRule struct {
	Action              MirrorRuleActionValue    `bson:"action"`
	Description         string                   `bson:"description"`
	DestinationCIDR     string                   `bson:"destinationcidr"`
	DestinationFromPort int                      `bson:"destinationfromport"`
	DestinationToPort   int                      `bson:"destinationtoport"`
	Direction           MirrorRuleDirectionValue `bson:"direction"`
	Number              int                      `bson:"number"`
	Protocol            int                      `bson:"protocol"`
	SourceCIDR          string                   `bson:"sourcecidr"`
	SourceFromPort      int                      `bson:"sourcefromport"`
	SourceToPort        int                      `bson:"sourcetoport"`
}
