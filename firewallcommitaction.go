// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallCommitActionActionStatusValue represents the possible values for attribute "actionStatus".
type FirewallCommitActionActionStatusValue string

const (
	// FirewallCommitActionActionStatusFailure represents the value Failure.
	FirewallCommitActionActionStatusFailure FirewallCommitActionActionStatusValue = "Failure"

	// FirewallCommitActionActionStatusSuccess represents the value Success.
	FirewallCommitActionActionStatusSuccess FirewallCommitActionActionStatusValue = "Success"
)

// FirewallCommitActionObjectTypeValue represents the possible values for attribute "objectType".
type FirewallCommitActionObjectTypeValue string

const (
	// FirewallCommitActionObjectTypeAWSLogDefinition represents the value AWSLogDefinition.
	FirewallCommitActionObjectTypeAWSLogDefinition FirewallCommitActionObjectTypeValue = "AWSLogDefinition"

	// FirewallCommitActionObjectTypeFirewallRuleset represents the value FirewallRuleset.
	FirewallCommitActionObjectTypeFirewallRuleset FirewallCommitActionObjectTypeValue = "FirewallRuleset"

	// FirewallCommitActionObjectTypeFirewallSecurityProfile represents the value FirewallSecurityProfile.
	FirewallCommitActionObjectTypeFirewallSecurityProfile FirewallCommitActionObjectTypeValue = "FirewallSecurityProfile"

	// FirewallCommitActionObjectTypeFirewallTemplate represents the value FirewallTemplate.
	FirewallCommitActionObjectTypeFirewallTemplate FirewallCommitActionObjectTypeValue = "FirewallTemplate"
)

// FirewallCommitActionIdentity represents the Identity of the object.
var FirewallCommitActionIdentity = elemental.Identity{
	Name:     "firewallcommitaction",
	Category: "firewallcommitactions",
	Package:  "ngfw",
	Private:  false,
}

// FirewallCommitActionsList represents a list of FirewallCommitActions
type FirewallCommitActionsList []*FirewallCommitAction

// Identity returns the identity of the objects in the list.
func (o FirewallCommitActionsList) Identity() elemental.Identity {

	return FirewallCommitActionIdentity
}

// Copy returns a pointer to a copy the FirewallCommitActionsList.
func (o FirewallCommitActionsList) Copy() elemental.Identifiables {

	out := append(FirewallCommitActionsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the FirewallCommitActionsList.
func (o FirewallCommitActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FirewallCommitActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FirewallCommitAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FirewallCommitActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FirewallCommitActionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FirewallCommitActionsList converted to SparseFirewallCommitActionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FirewallCommitActionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFirewallCommitActionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFirewallCommitAction)
	}

	return out
}

// Version returns the version of the content.
func (o FirewallCommitActionsList) Version() int {

	return 1
}

// FirewallCommitAction represents the model of a firewallcommitaction
type FirewallCommitAction struct {
	// The status of action.
	ActionStatus FirewallCommitActionActionStatusValue `json:"actionStatus" msgpack:"actionStatus" bson:"-" mapstructure:"actionStatus,omitempty"`

	// The action status description.
	ActionStatusReason string `json:"actionStatusReason" msgpack:"actionStatusReason" bson:"-" mapstructure:"actionStatusReason,omitempty"`

	// The firewalls effected by the commit action.
	AffectedFirewalls []*FirewallStatus `json:"affectedFirewalls" msgpack:"affectedFirewalls" bson:"affectedfirewalls" mapstructure:"affectedFirewalls,omitempty"`

	// The ID of the object being committed.
	ObjectID string `json:"objectID" msgpack:"objectID" bson:"-" mapstructure:"objectID,omitempty"`

	// The type of object being committed.
	ObjectType FirewallCommitActionObjectTypeValue `json:"objectType" msgpack:"objectType" bson:"-" mapstructure:"objectType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallCommitAction returns a new *FirewallCommitAction
func NewFirewallCommitAction() *FirewallCommitAction {

	return &FirewallCommitAction{
		ModelVersion:      1,
		AffectedFirewalls: []*FirewallStatus{},
		ObjectType:        FirewallCommitActionObjectTypeFirewallRuleset,
	}
}

// Identity returns the Identity of the object.
func (o *FirewallCommitAction) Identity() elemental.Identity {

	return FirewallCommitActionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallCommitAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallCommitAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCommitAction) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallCommitAction{}

	s.AffectedFirewalls = o.AffectedFirewalls

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCommitAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallCommitAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.AffectedFirewalls = s.AffectedFirewalls

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FirewallCommitAction) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallCommitAction) BleveType() string {

	return "firewallcommitaction"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FirewallCommitAction) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FirewallCommitAction) Doc() string {

	return `Represents committing FilewallRulesets, FilewallTemplates,
FilewallSecurityProfiles, 
and AWSLogDefinitions to a firewall.`
}

func (o *FirewallCommitAction) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FirewallCommitAction) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFirewallCommitAction{
			ActionStatus:       &o.ActionStatus,
			ActionStatusReason: &o.ActionStatusReason,
			AffectedFirewalls:  &o.AffectedFirewalls,
			ObjectID:           &o.ObjectID,
			ObjectType:         &o.ObjectType,
		}
	}

	sp := &SparseFirewallCommitAction{}
	for _, f := range fields {
		switch f {
		case "actionStatus":
			sp.ActionStatus = &(o.ActionStatus)
		case "actionStatusReason":
			sp.ActionStatusReason = &(o.ActionStatusReason)
		case "affectedFirewalls":
			sp.AffectedFirewalls = &(o.AffectedFirewalls)
		case "objectID":
			sp.ObjectID = &(o.ObjectID)
		case "objectType":
			sp.ObjectType = &(o.ObjectType)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFirewallCommitAction to the object.
func (o *FirewallCommitAction) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFirewallCommitAction)
	if so.ActionStatus != nil {
		o.ActionStatus = *so.ActionStatus
	}
	if so.ActionStatusReason != nil {
		o.ActionStatusReason = *so.ActionStatusReason
	}
	if so.AffectedFirewalls != nil {
		o.AffectedFirewalls = *so.AffectedFirewalls
	}
	if so.ObjectID != nil {
		o.ObjectID = *so.ObjectID
	}
	if so.ObjectType != nil {
		o.ObjectType = *so.ObjectType
	}
}

// DeepCopy returns a deep copy if the FirewallCommitAction.
func (o *FirewallCommitAction) DeepCopy() *FirewallCommitAction {

	if o == nil {
		return nil
	}

	out := &FirewallCommitAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallCommitAction.
func (o *FirewallCommitAction) DeepCopyInto(out *FirewallCommitAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallCommitAction: %s", err))
	}

	*out = *target.(*FirewallCommitAction)
}

// Validate valides the current information stored into the structure.
func (o *FirewallCommitAction) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("actionStatus", string(o.ActionStatus), []string{"Success", "Failure"}, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.AffectedFirewalls {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("objectID", o.ObjectID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("objectType", string(o.ObjectType), []string{"FirewallRuleset", "FirewallTemplate", "FirewallSecurityProfile", "AWSLogDefinition"}, false); err != nil {
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
func (*FirewallCommitAction) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallCommitActionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallCommitActionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallCommitAction) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallCommitActionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallCommitAction) ValueForAttribute(name string) any {

	switch name {
	case "actionStatus":
		return o.ActionStatus
	case "actionStatusReason":
		return o.ActionStatusReason
	case "affectedFirewalls":
		return o.AffectedFirewalls
	case "objectID":
		return o.ObjectID
	case "objectType":
		return o.ObjectType
	}

	return nil
}

// FirewallCommitActionAttributesMap represents the map of attribute for FirewallCommitAction.
var FirewallCommitActionAttributesMap = map[string]elemental.AttributeSpecification{
	"ActionStatus": {
		AllowedChoices: []string{"Success", "Failure"},
		ConvertedName:  "ActionStatus",
		Description:    `The status of action.`,
		Exposed:        true,
		Name:           "actionStatus",
		ReadOnly:       true,
		Type:           "enum",
	},
	"ActionStatusReason": {
		AllowedChoices: []string{},
		ConvertedName:  "ActionStatusReason",
		Description:    `The action status description.`,
		Exposed:        true,
		Name:           "actionStatusReason",
		ReadOnly:       true,
		Type:           "string",
	},
	"AffectedFirewalls": {
		AllowedChoices: []string{},
		BSONFieldName:  "affectedfirewalls",
		ConvertedName:  "AffectedFirewalls",
		Description:    `The firewalls effected by the commit action.`,
		Exposed:        true,
		Name:           "affectedFirewalls",
		Stored:         true,
		SubType:        "firewallstatus",
		Type:           "refList",
	},
	"ObjectID": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `The ID of the object being committed.`,
		Exposed:        true,
		Name:           "objectID",
		Required:       true,
		Type:           "string",
	},
	"ObjectType": {
		AllowedChoices: []string{"FirewallRuleset", "FirewallTemplate", "FirewallSecurityProfile", "AWSLogDefinition"},
		ConvertedName:  "ObjectType",
		DefaultValue:   FirewallCommitActionObjectTypeFirewallRuleset,
		Description:    `The type of object being committed.`,
		Exposed:        true,
		Name:           "objectType",
		Type:           "enum",
	},
}

// FirewallCommitActionLowerCaseAttributesMap represents the map of attribute for FirewallCommitAction.
var FirewallCommitActionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"actionstatus": {
		AllowedChoices: []string{"Success", "Failure"},
		ConvertedName:  "ActionStatus",
		Description:    `The status of action.`,
		Exposed:        true,
		Name:           "actionStatus",
		ReadOnly:       true,
		Type:           "enum",
	},
	"actionstatusreason": {
		AllowedChoices: []string{},
		ConvertedName:  "ActionStatusReason",
		Description:    `The action status description.`,
		Exposed:        true,
		Name:           "actionStatusReason",
		ReadOnly:       true,
		Type:           "string",
	},
	"affectedfirewalls": {
		AllowedChoices: []string{},
		BSONFieldName:  "affectedfirewalls",
		ConvertedName:  "AffectedFirewalls",
		Description:    `The firewalls effected by the commit action.`,
		Exposed:        true,
		Name:           "affectedFirewalls",
		Stored:         true,
		SubType:        "firewallstatus",
		Type:           "refList",
	},
	"objectid": {
		AllowedChoices: []string{},
		ConvertedName:  "ObjectID",
		Description:    `The ID of the object being committed.`,
		Exposed:        true,
		Name:           "objectID",
		Required:       true,
		Type:           "string",
	},
	"objecttype": {
		AllowedChoices: []string{"FirewallRuleset", "FirewallTemplate", "FirewallSecurityProfile", "AWSLogDefinition"},
		ConvertedName:  "ObjectType",
		DefaultValue:   FirewallCommitActionObjectTypeFirewallRuleset,
		Description:    `The type of object being committed.`,
		Exposed:        true,
		Name:           "objectType",
		Type:           "enum",
	},
}

// SparseFirewallCommitActionsList represents a list of SparseFirewallCommitActions
type SparseFirewallCommitActionsList []*SparseFirewallCommitAction

// Identity returns the identity of the objects in the list.
func (o SparseFirewallCommitActionsList) Identity() elemental.Identity {

	return FirewallCommitActionIdentity
}

// Copy returns a pointer to a copy the SparseFirewallCommitActionsList.
func (o SparseFirewallCommitActionsList) Copy() elemental.Identifiables {

	copy := append(SparseFirewallCommitActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFirewallCommitActionsList.
func (o SparseFirewallCommitActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFirewallCommitActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFirewallCommitAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFirewallCommitActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFirewallCommitActionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFirewallCommitActionsList converted to FirewallCommitActionsList.
func (o SparseFirewallCommitActionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFirewallCommitActionsList) Version() int {

	return 1
}

// SparseFirewallCommitAction represents the sparse version of a firewallcommitaction.
type SparseFirewallCommitAction struct {
	// The status of action.
	ActionStatus *FirewallCommitActionActionStatusValue `json:"actionStatus,omitempty" msgpack:"actionStatus,omitempty" bson:"-" mapstructure:"actionStatus,omitempty"`

	// The action status description.
	ActionStatusReason *string `json:"actionStatusReason,omitempty" msgpack:"actionStatusReason,omitempty" bson:"-" mapstructure:"actionStatusReason,omitempty"`

	// The firewalls effected by the commit action.
	AffectedFirewalls *[]*FirewallStatus `json:"affectedFirewalls,omitempty" msgpack:"affectedFirewalls,omitempty" bson:"affectedfirewalls,omitempty" mapstructure:"affectedFirewalls,omitempty"`

	// The ID of the object being committed.
	ObjectID *string `json:"objectID,omitempty" msgpack:"objectID,omitempty" bson:"-" mapstructure:"objectID,omitempty"`

	// The type of object being committed.
	ObjectType *FirewallCommitActionObjectTypeValue `json:"objectType,omitempty" msgpack:"objectType,omitempty" bson:"-" mapstructure:"objectType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFirewallCommitAction returns a new  SparseFirewallCommitAction.
func NewSparseFirewallCommitAction() *SparseFirewallCommitAction {
	return &SparseFirewallCommitAction{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFirewallCommitAction) Identity() elemental.Identity {

	return FirewallCommitActionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFirewallCommitAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFirewallCommitAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallCommitAction) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFirewallCommitAction{}

	if o.AffectedFirewalls != nil {
		s.AffectedFirewalls = o.AffectedFirewalls
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallCommitAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFirewallCommitAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.AffectedFirewalls != nil {
		o.AffectedFirewalls = s.AffectedFirewalls
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFirewallCommitAction) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFirewallCommitAction) ToPlain() elemental.PlainIdentifiable {

	out := NewFirewallCommitAction()
	if o.ActionStatus != nil {
		out.ActionStatus = *o.ActionStatus
	}
	if o.ActionStatusReason != nil {
		out.ActionStatusReason = *o.ActionStatusReason
	}
	if o.AffectedFirewalls != nil {
		out.AffectedFirewalls = *o.AffectedFirewalls
	}
	if o.ObjectID != nil {
		out.ObjectID = *o.ObjectID
	}
	if o.ObjectType != nil {
		out.ObjectType = *o.ObjectType
	}

	return out
}

// DeepCopy returns a deep copy if the SparseFirewallCommitAction.
func (o *SparseFirewallCommitAction) DeepCopy() *SparseFirewallCommitAction {

	if o == nil {
		return nil
	}

	out := &SparseFirewallCommitAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFirewallCommitAction.
func (o *SparseFirewallCommitAction) DeepCopyInto(out *SparseFirewallCommitAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFirewallCommitAction: %s", err))
	}

	*out = *target.(*SparseFirewallCommitAction)
}

type mongoAttributesFirewallCommitAction struct {
	AffectedFirewalls []*FirewallStatus `bson:"affectedfirewalls"`
}
type mongoAttributesSparseFirewallCommitAction struct {
	AffectedFirewalls *[]*FirewallStatus `bson:"affectedfirewalls,omitempty"`
}
