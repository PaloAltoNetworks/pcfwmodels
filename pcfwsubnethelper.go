// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCFWSubnetHelperIdentity represents the Identity of the object.
var PCFWSubnetHelperIdentity = elemental.Identity{
	Name:     "pcfwsubnethelper",
	Category: "pcfwsubnethelpers",
	Package:  "deploymentadvisor",
	Private:  false,
}

// PCFWSubnetHelpersList represents a list of PCFWSubnetHelpers
type PCFWSubnetHelpersList []*PCFWSubnetHelper

// Identity returns the identity of the objects in the list.
func (o PCFWSubnetHelpersList) Identity() elemental.Identity {

	return PCFWSubnetHelperIdentity
}

// Copy returns a pointer to a copy the PCFWSubnetHelpersList.
func (o PCFWSubnetHelpersList) Copy() elemental.Identifiables {

	out := append(PCFWSubnetHelpersList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the PCFWSubnetHelpersList.
func (o PCFWSubnetHelpersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCFWSubnetHelpersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCFWSubnetHelper))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCFWSubnetHelpersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCFWSubnetHelpersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PCFWSubnetHelpersList converted to SparsePCFWSubnetHelpersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCFWSubnetHelpersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCFWSubnetHelpersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCFWSubnetHelper)
	}

	return out
}

// Version returns the version of the content.
func (o PCFWSubnetHelpersList) Version() int {

	return 1
}

// PCFWSubnetHelper represents the model of a pcfwsubnethelper
type PCFWSubnetHelper struct {
	// Returns the list of AWS VPC IDs and information about available subnets for
	// every availability zones in a VPC.
	VPCAvailableSubnets []*VPCAvailableSubnet `json:"VPCAvailableSubnets" msgpack:"VPCAvailableSubnets" bson:"-" mapstructure:"VPCAvailableSubnets,omitempty"`

	// List of AWS VPC IDs with information about associated availability zones and
	// used subnets to check for available subnets.
	VPCUsedSubnets []*VPCUsedSubnet `json:"VPCUsedSubnets" msgpack:"VPCUsedSubnets" bson:"vpcusedsubnets" mapstructure:"VPCUsedSubnets,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPCFWSubnetHelper returns a new *PCFWSubnetHelper
func NewPCFWSubnetHelper() *PCFWSubnetHelper {

	return &PCFWSubnetHelper{
		ModelVersion:        1,
		VPCAvailableSubnets: []*VPCAvailableSubnet{},
		VPCUsedSubnets:      []*VPCUsedSubnet{},
	}
}

// Identity returns the Identity of the object.
func (o *PCFWSubnetHelper) Identity() elemental.Identity {

	return PCFWSubnetHelperIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCFWSubnetHelper) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCFWSubnetHelper) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCFWSubnetHelper) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCFWSubnetHelper{}

	s.VPCUsedSubnets = o.VPCUsedSubnets

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCFWSubnetHelper) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCFWSubnetHelper{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.VPCUsedSubnets = s.VPCUsedSubnets

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PCFWSubnetHelper) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCFWSubnetHelper) BleveType() string {

	return "pcfwsubnethelper"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCFWSubnetHelper) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PCFWSubnetHelper) Doc() string {

	return `Represents a PCFW service endpoint that returns available subnet for every
specified availability zones in every specified VPC.`
}

func (o *PCFWSubnetHelper) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCFWSubnetHelper) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCFWSubnetHelper{
			VPCAvailableSubnets: &o.VPCAvailableSubnets,
			VPCUsedSubnets:      &o.VPCUsedSubnets,
		}
	}

	sp := &SparsePCFWSubnetHelper{}
	for _, f := range fields {
		switch f {
		case "VPCAvailableSubnets":
			sp.VPCAvailableSubnets = &(o.VPCAvailableSubnets)
		case "VPCUsedSubnets":
			sp.VPCUsedSubnets = &(o.VPCUsedSubnets)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePCFWSubnetHelper to the object.
func (o *PCFWSubnetHelper) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCFWSubnetHelper)
	if so.VPCAvailableSubnets != nil {
		o.VPCAvailableSubnets = *so.VPCAvailableSubnets
	}
	if so.VPCUsedSubnets != nil {
		o.VPCUsedSubnets = *so.VPCUsedSubnets
	}
}

// DeepCopy returns a deep copy if the PCFWSubnetHelper.
func (o *PCFWSubnetHelper) DeepCopy() *PCFWSubnetHelper {

	if o == nil {
		return nil
	}

	out := &PCFWSubnetHelper{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCFWSubnetHelper.
func (o *PCFWSubnetHelper) DeepCopyInto(out *PCFWSubnetHelper) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCFWSubnetHelper: %s", err))
	}

	*out = *target.(*PCFWSubnetHelper)
}

// Validate valides the current information stored into the structure.
func (o *PCFWSubnetHelper) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.VPCAvailableSubnets {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.VPCUsedSubnets {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateVpcSubnetInfo("VPCUsedSubnets", o.VPCUsedSubnets); err != nil {
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
func (*PCFWSubnetHelper) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCFWSubnetHelperAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCFWSubnetHelperLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCFWSubnetHelper) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCFWSubnetHelperAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCFWSubnetHelper) ValueForAttribute(name string) any {

	switch name {
	case "VPCAvailableSubnets":
		return o.VPCAvailableSubnets
	case "VPCUsedSubnets":
		return o.VPCUsedSubnets
	}

	return nil
}

// PCFWSubnetHelperAttributesMap represents the map of attribute for PCFWSubnetHelper.
var PCFWSubnetHelperAttributesMap = map[string]elemental.AttributeSpecification{
	"VPCAvailableSubnets": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "VPCAvailableSubnets",
		Description: `Returns the list of AWS VPC IDs and information about available subnets for
every availability zones in a VPC.`,
		Exposed:  true,
		Name:     "VPCAvailableSubnets",
		ReadOnly: true,
		SubType:  "vpcavailablesubnet",
		Type:     "refList",
	},
	"VPCUsedSubnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcusedsubnets",
		ConvertedName:  "VPCUsedSubnets",
		Description: `List of AWS VPC IDs with information about associated availability zones and
used subnets to check for available subnets.`,
		Exposed: true,
		Name:    "VPCUsedSubnets",
		Stored:  true,
		SubType: "vpcusedsubnet",
		Type:    "refList",
	},
}

// PCFWSubnetHelperLowerCaseAttributesMap represents the map of attribute for PCFWSubnetHelper.
var PCFWSubnetHelperLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"vpcavailablesubnets": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "VPCAvailableSubnets",
		Description: `Returns the list of AWS VPC IDs and information about available subnets for
every availability zones in a VPC.`,
		Exposed:  true,
		Name:     "VPCAvailableSubnets",
		ReadOnly: true,
		SubType:  "vpcavailablesubnet",
		Type:     "refList",
	},
	"vpcusedsubnets": {
		AllowedChoices: []string{},
		BSONFieldName:  "vpcusedsubnets",
		ConvertedName:  "VPCUsedSubnets",
		Description: `List of AWS VPC IDs with information about associated availability zones and
used subnets to check for available subnets.`,
		Exposed: true,
		Name:    "VPCUsedSubnets",
		Stored:  true,
		SubType: "vpcusedsubnet",
		Type:    "refList",
	},
}

// SparsePCFWSubnetHelpersList represents a list of SparsePCFWSubnetHelpers
type SparsePCFWSubnetHelpersList []*SparsePCFWSubnetHelper

// Identity returns the identity of the objects in the list.
func (o SparsePCFWSubnetHelpersList) Identity() elemental.Identity {

	return PCFWSubnetHelperIdentity
}

// Copy returns a pointer to a copy the SparsePCFWSubnetHelpersList.
func (o SparsePCFWSubnetHelpersList) Copy() elemental.Identifiables {

	copy := append(SparsePCFWSubnetHelpersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCFWSubnetHelpersList.
func (o SparsePCFWSubnetHelpersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCFWSubnetHelpersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCFWSubnetHelper))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCFWSubnetHelpersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCFWSubnetHelpersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePCFWSubnetHelpersList converted to PCFWSubnetHelpersList.
func (o SparsePCFWSubnetHelpersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCFWSubnetHelpersList) Version() int {

	return 1
}

// SparsePCFWSubnetHelper represents the sparse version of a pcfwsubnethelper.
type SparsePCFWSubnetHelper struct {
	// Returns the list of AWS VPC IDs and information about available subnets for
	// every availability zones in a VPC.
	VPCAvailableSubnets *[]*VPCAvailableSubnet `json:"VPCAvailableSubnets,omitempty" msgpack:"VPCAvailableSubnets,omitempty" bson:"-" mapstructure:"VPCAvailableSubnets,omitempty"`

	// List of AWS VPC IDs with information about associated availability zones and
	// used subnets to check for available subnets.
	VPCUsedSubnets *[]*VPCUsedSubnet `json:"VPCUsedSubnets,omitempty" msgpack:"VPCUsedSubnets,omitempty" bson:"vpcusedsubnets,omitempty" mapstructure:"VPCUsedSubnets,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePCFWSubnetHelper returns a new  SparsePCFWSubnetHelper.
func NewSparsePCFWSubnetHelper() *SparsePCFWSubnetHelper {
	return &SparsePCFWSubnetHelper{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCFWSubnetHelper) Identity() elemental.Identity {

	return PCFWSubnetHelperIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCFWSubnetHelper) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCFWSubnetHelper) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCFWSubnetHelper) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCFWSubnetHelper{}

	if o.VPCUsedSubnets != nil {
		s.VPCUsedSubnets = o.VPCUsedSubnets
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCFWSubnetHelper) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCFWSubnetHelper{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.VPCUsedSubnets != nil {
		o.VPCUsedSubnets = s.VPCUsedSubnets
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePCFWSubnetHelper) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCFWSubnetHelper) ToPlain() elemental.PlainIdentifiable {

	out := NewPCFWSubnetHelper()
	if o.VPCAvailableSubnets != nil {
		out.VPCAvailableSubnets = *o.VPCAvailableSubnets
	}
	if o.VPCUsedSubnets != nil {
		out.VPCUsedSubnets = *o.VPCUsedSubnets
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePCFWSubnetHelper.
func (o *SparsePCFWSubnetHelper) DeepCopy() *SparsePCFWSubnetHelper {

	if o == nil {
		return nil
	}

	out := &SparsePCFWSubnetHelper{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCFWSubnetHelper.
func (o *SparsePCFWSubnetHelper) DeepCopyInto(out *SparsePCFWSubnetHelper) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCFWSubnetHelper: %s", err))
	}

	*out = *target.(*SparsePCFWSubnetHelper)
}

type mongoAttributesPCFWSubnetHelper struct {
	VPCUsedSubnets []*VPCUsedSubnet `bson:"vpcusedsubnets"`
}
type mongoAttributesSparsePCFWSubnetHelper struct {
	VPCUsedSubnets *[]*VPCUsedSubnet `bson:"vpcusedsubnets,omitempty"`
}
