// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MirrorSourceOptionIdentity represents the Identity of the object.
var MirrorSourceOptionIdentity = elemental.Identity{
	Name:     "mirrorsourceoption",
	Category: "mirrorsourceoptions",
	Package:  "discovery",
	Private:  false,
}

// MirrorSourceOptionsList represents a list of MirrorSourceOptions
type MirrorSourceOptionsList []*MirrorSourceOption

// Identity returns the identity of the objects in the list.
func (o MirrorSourceOptionsList) Identity() elemental.Identity {

	return MirrorSourceOptionIdentity
}

// Copy returns a pointer to a copy the MirrorSourceOptionsList.
func (o MirrorSourceOptionsList) Copy() elemental.Identifiables {

	out := append(MirrorSourceOptionsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the MirrorSourceOptionsList.
func (o MirrorSourceOptionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(MirrorSourceOptionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*MirrorSourceOption))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o MirrorSourceOptionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o MirrorSourceOptionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the MirrorSourceOptionsList converted to SparseMirrorSourceOptionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o MirrorSourceOptionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseMirrorSourceOptionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseMirrorSourceOption)
	}

	return out
}

// Version returns the version of the content.
func (o MirrorSourceOptionsList) Version() int {

	return 1
}

// MirrorSourceOption represents the model of a mirrorsourceoption
type MirrorSourceOption struct {
	// List of discovered auto-scaling groups.
	AutoScalingGroupNames []string `json:"autoScalingGroupNames" msgpack:"autoScalingGroupNames" bson:"-" mapstructure:"autoScalingGroupNames,omitempty"`

	// The firewall name whose VPCs/AZs should be used to search for instances.
	FirewallName string `json:"firewallName" msgpack:"firewallName" bson:"-" mapstructure:"firewallName,omitempty"`

	// List of discovered mirror source instances.
	Instances []*MirrorInstance `json:"instances" msgpack:"instances" bson:"-" mapstructure:"instances,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMirrorSourceOption returns a new *MirrorSourceOption
func NewMirrorSourceOption() *MirrorSourceOption {

	return &MirrorSourceOption{
		ModelVersion:          1,
		AutoScalingGroupNames: []string{},
		Instances:             []*MirrorInstance{},
	}
}

// Identity returns the Identity of the object.
func (o *MirrorSourceOption) Identity() elemental.Identity {

	return MirrorSourceOptionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MirrorSourceOption) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MirrorSourceOption) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MirrorSourceOption) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMirrorSourceOption{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *MirrorSourceOption) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMirrorSourceOption{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *MirrorSourceOption) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *MirrorSourceOption) BleveType() string {

	return "mirrorsourceoption"
}

// DefaultOrder returns the list of default ordering fields.
func (o *MirrorSourceOption) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *MirrorSourceOption) Doc() string {

	return `Discovers eligible traffic mirror source instances and auto-scaling groups.`
}

func (o *MirrorSourceOption) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *MirrorSourceOption) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseMirrorSourceOption{
			AutoScalingGroupNames: &o.AutoScalingGroupNames,
			FirewallName:          &o.FirewallName,
			Instances:             &o.Instances,
		}
	}

	sp := &SparseMirrorSourceOption{}
	for _, f := range fields {
		switch f {
		case "autoScalingGroupNames":
			sp.AutoScalingGroupNames = &(o.AutoScalingGroupNames)
		case "firewallName":
			sp.FirewallName = &(o.FirewallName)
		case "instances":
			sp.Instances = &(o.Instances)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseMirrorSourceOption to the object.
func (o *MirrorSourceOption) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseMirrorSourceOption)
	if so.AutoScalingGroupNames != nil {
		o.AutoScalingGroupNames = *so.AutoScalingGroupNames
	}
	if so.FirewallName != nil {
		o.FirewallName = *so.FirewallName
	}
	if so.Instances != nil {
		o.Instances = *so.Instances
	}
}

// DeepCopy returns a deep copy if the MirrorSourceOption.
func (o *MirrorSourceOption) DeepCopy() *MirrorSourceOption {

	if o == nil {
		return nil
	}

	out := &MirrorSourceOption{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *MirrorSourceOption.
func (o *MirrorSourceOption) DeepCopyInto(out *MirrorSourceOption) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy MirrorSourceOption: %s", err))
	}

	*out = *target.(*MirrorSourceOption)
}

// Validate valides the current information stored into the structure.
func (o *MirrorSourceOption) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("firewallName", o.FirewallName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.Instances {
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
func (*MirrorSourceOption) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MirrorSourceOptionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MirrorSourceOptionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*MirrorSourceOption) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MirrorSourceOptionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *MirrorSourceOption) ValueForAttribute(name string) any {

	switch name {
	case "autoScalingGroupNames":
		return o.AutoScalingGroupNames
	case "firewallName":
		return o.FirewallName
	case "instances":
		return o.Instances
	}

	return nil
}

// MirrorSourceOptionAttributesMap represents the map of attribute for MirrorSourceOption.
var MirrorSourceOptionAttributesMap = map[string]elemental.AttributeSpecification{
	"AutoScalingGroupNames": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AutoScalingGroupNames",
		Description:    `List of discovered auto-scaling groups.`,
		Exposed:        true,
		Name:           "autoScalingGroupNames",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"FirewallName": {
		AllowedChoices: []string{},
		ConvertedName:  "FirewallName",
		Description:    `The firewall name whose VPCs/AZs should be used to search for instances.`,
		Exposed:        true,
		Name:           "firewallName",
		Required:       true,
		Type:           "string",
	},
	"Instances": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Instances",
		Description:    `List of discovered mirror source instances.`,
		Exposed:        true,
		Name:           "instances",
		ReadOnly:       true,
		SubType:        "mirrorinstance",
		Type:           "refList",
	},
}

// MirrorSourceOptionLowerCaseAttributesMap represents the map of attribute for MirrorSourceOption.
var MirrorSourceOptionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"autoscalinggroupnames": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AutoScalingGroupNames",
		Description:    `List of discovered auto-scaling groups.`,
		Exposed:        true,
		Name:           "autoScalingGroupNames",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"firewallname": {
		AllowedChoices: []string{},
		ConvertedName:  "FirewallName",
		Description:    `The firewall name whose VPCs/AZs should be used to search for instances.`,
		Exposed:        true,
		Name:           "firewallName",
		Required:       true,
		Type:           "string",
	},
	"instances": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Instances",
		Description:    `List of discovered mirror source instances.`,
		Exposed:        true,
		Name:           "instances",
		ReadOnly:       true,
		SubType:        "mirrorinstance",
		Type:           "refList",
	},
}

// SparseMirrorSourceOptionsList represents a list of SparseMirrorSourceOptions
type SparseMirrorSourceOptionsList []*SparseMirrorSourceOption

// Identity returns the identity of the objects in the list.
func (o SparseMirrorSourceOptionsList) Identity() elemental.Identity {

	return MirrorSourceOptionIdentity
}

// Copy returns a pointer to a copy the SparseMirrorSourceOptionsList.
func (o SparseMirrorSourceOptionsList) Copy() elemental.Identifiables {

	copy := append(SparseMirrorSourceOptionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseMirrorSourceOptionsList.
func (o SparseMirrorSourceOptionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseMirrorSourceOptionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseMirrorSourceOption))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseMirrorSourceOptionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseMirrorSourceOptionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseMirrorSourceOptionsList converted to MirrorSourceOptionsList.
func (o SparseMirrorSourceOptionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseMirrorSourceOptionsList) Version() int {

	return 1
}

// SparseMirrorSourceOption represents the sparse version of a mirrorsourceoption.
type SparseMirrorSourceOption struct {
	// List of discovered auto-scaling groups.
	AutoScalingGroupNames *[]string `json:"autoScalingGroupNames,omitempty" msgpack:"autoScalingGroupNames,omitempty" bson:"-" mapstructure:"autoScalingGroupNames,omitempty"`

	// The firewall name whose VPCs/AZs should be used to search for instances.
	FirewallName *string `json:"firewallName,omitempty" msgpack:"firewallName,omitempty" bson:"-" mapstructure:"firewallName,omitempty"`

	// List of discovered mirror source instances.
	Instances *[]*MirrorInstance `json:"instances,omitempty" msgpack:"instances,omitempty" bson:"-" mapstructure:"instances,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseMirrorSourceOption returns a new  SparseMirrorSourceOption.
func NewSparseMirrorSourceOption() *SparseMirrorSourceOption {
	return &SparseMirrorSourceOption{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseMirrorSourceOption) Identity() elemental.Identity {

	return MirrorSourceOptionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseMirrorSourceOption) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseMirrorSourceOption) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMirrorSourceOption) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseMirrorSourceOption{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMirrorSourceOption) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseMirrorSourceOption{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseMirrorSourceOption) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseMirrorSourceOption) ToPlain() elemental.PlainIdentifiable {

	out := NewMirrorSourceOption()
	if o.AutoScalingGroupNames != nil {
		out.AutoScalingGroupNames = *o.AutoScalingGroupNames
	}
	if o.FirewallName != nil {
		out.FirewallName = *o.FirewallName
	}
	if o.Instances != nil {
		out.Instances = *o.Instances
	}

	return out
}

// DeepCopy returns a deep copy if the SparseMirrorSourceOption.
func (o *SparseMirrorSourceOption) DeepCopy() *SparseMirrorSourceOption {

	if o == nil {
		return nil
	}

	out := &SparseMirrorSourceOption{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseMirrorSourceOption.
func (o *SparseMirrorSourceOption) DeepCopyInto(out *SparseMirrorSourceOption) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseMirrorSourceOption: %s", err))
	}

	*out = *target.(*SparseMirrorSourceOption)
}

type mongoAttributesMirrorSourceOption struct {
}
type mongoAttributesSparseMirrorSourceOption struct {
}