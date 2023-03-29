// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallOptionIdentity represents the Identity of the object.
var FirewallOptionIdentity = elemental.Identity{
	Name:     "firewalloption",
	Category: "firewalloptions",
	Package:  "discovery",
	Private:  false,
}

// FirewallOptionsList represents a list of FirewallOptions
type FirewallOptionsList []*FirewallOption

// Identity returns the identity of the objects in the list.
func (o FirewallOptionsList) Identity() elemental.Identity {

	return FirewallOptionIdentity
}

// Copy returns a pointer to a copy the FirewallOptionsList.
func (o FirewallOptionsList) Copy() elemental.Identifiables {

	out := append(FirewallOptionsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the FirewallOptionsList.
func (o FirewallOptionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FirewallOptionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FirewallOption))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FirewallOptionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FirewallOptionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FirewallOptionsList converted to SparseFirewallOptionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FirewallOptionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFirewallOptionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFirewallOption)
	}

	return out
}

// Version returns the version of the content.
func (o FirewallOptionsList) Version() int {

	return 1
}

// FirewallOption represents the model of a firewalloption
type FirewallOption struct {
	// Optionally restrict results to this AWS region.
	RegionFilter string `json:"regionFilter" msgpack:"regionFilter" bson:"-" mapstructure:"regionFilter,omitempty"`

	// List of VPC IDs and their region and associated availability zones.
	RegionVPCs []*FirewallOptionRegion `json:"regionVPCs" msgpack:"regionVPCs" bson:"-" mapstructure:"regionVPCs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallOption returns a new *FirewallOption
func NewFirewallOption() *FirewallOption {

	return &FirewallOption{
		ModelVersion: 1,
		RegionVPCs:   []*FirewallOptionRegion{},
	}
}

// Identity returns the Identity of the object.
func (o *FirewallOption) Identity() elemental.Identity {

	return FirewallOptionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallOption) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallOption) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallOption) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallOption{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallOption) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallOption{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FirewallOption) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallOption) BleveType() string {

	return "firewalloption"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FirewallOption) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FirewallOption) Doc() string {

	return `Discovers firewall deployment options.`
}

func (o *FirewallOption) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FirewallOption) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFirewallOption{
			RegionFilter: &o.RegionFilter,
			RegionVPCs:   &o.RegionVPCs,
		}
	}

	sp := &SparseFirewallOption{}
	for _, f := range fields {
		switch f {
		case "regionFilter":
			sp.RegionFilter = &(o.RegionFilter)
		case "regionVPCs":
			sp.RegionVPCs = &(o.RegionVPCs)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFirewallOption to the object.
func (o *FirewallOption) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFirewallOption)
	if so.RegionFilter != nil {
		o.RegionFilter = *so.RegionFilter
	}
	if so.RegionVPCs != nil {
		o.RegionVPCs = *so.RegionVPCs
	}
}

// DeepCopy returns a deep copy if the FirewallOption.
func (o *FirewallOption) DeepCopy() *FirewallOption {

	if o == nil {
		return nil
	}

	out := &FirewallOption{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallOption.
func (o *FirewallOption) DeepCopyInto(out *FirewallOption) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallOption: %s", err))
	}

	*out = *target.(*FirewallOption)
}

// Validate valides the current information stored into the structure.
func (o *FirewallOption) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.RegionVPCs {
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
func (*FirewallOption) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallOptionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallOptionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallOption) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallOptionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallOption) ValueForAttribute(name string) any {

	switch name {
	case "regionFilter":
		return o.RegionFilter
	case "regionVPCs":
		return o.RegionVPCs
	}

	return nil
}

// FirewallOptionAttributesMap represents the map of attribute for FirewallOption.
var FirewallOptionAttributesMap = map[string]elemental.AttributeSpecification{
	"RegionFilter": {
		AllowedChoices: []string{},
		ConvertedName:  "RegionFilter",
		CreationOnly:   true,
		Description:    `Optionally restrict results to this AWS region.`,
		Exposed:        true,
		Name:           "regionFilter",
		Type:           "string",
	},
	"RegionVPCs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RegionVPCs",
		Description:    `List of VPC IDs and their region and associated availability zones.`,
		Exposed:        true,
		Name:           "regionVPCs",
		ReadOnly:       true,
		SubType:        "firewalloptionregion",
		Type:           "refList",
	},
}

// FirewallOptionLowerCaseAttributesMap represents the map of attribute for FirewallOption.
var FirewallOptionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"regionfilter": {
		AllowedChoices: []string{},
		ConvertedName:  "RegionFilter",
		CreationOnly:   true,
		Description:    `Optionally restrict results to this AWS region.`,
		Exposed:        true,
		Name:           "regionFilter",
		Type:           "string",
	},
	"regionvpcs": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RegionVPCs",
		Description:    `List of VPC IDs and their region and associated availability zones.`,
		Exposed:        true,
		Name:           "regionVPCs",
		ReadOnly:       true,
		SubType:        "firewalloptionregion",
		Type:           "refList",
	},
}

// SparseFirewallOptionsList represents a list of SparseFirewallOptions
type SparseFirewallOptionsList []*SparseFirewallOption

// Identity returns the identity of the objects in the list.
func (o SparseFirewallOptionsList) Identity() elemental.Identity {

	return FirewallOptionIdentity
}

// Copy returns a pointer to a copy the SparseFirewallOptionsList.
func (o SparseFirewallOptionsList) Copy() elemental.Identifiables {

	copy := append(SparseFirewallOptionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFirewallOptionsList.
func (o SparseFirewallOptionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFirewallOptionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFirewallOption))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFirewallOptionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFirewallOptionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFirewallOptionsList converted to FirewallOptionsList.
func (o SparseFirewallOptionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFirewallOptionsList) Version() int {

	return 1
}

// SparseFirewallOption represents the sparse version of a firewalloption.
type SparseFirewallOption struct {
	// Optionally restrict results to this AWS region.
	RegionFilter *string `json:"regionFilter,omitempty" msgpack:"regionFilter,omitempty" bson:"-" mapstructure:"regionFilter,omitempty"`

	// List of VPC IDs and their region and associated availability zones.
	RegionVPCs *[]*FirewallOptionRegion `json:"regionVPCs,omitempty" msgpack:"regionVPCs,omitempty" bson:"-" mapstructure:"regionVPCs,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFirewallOption returns a new  SparseFirewallOption.
func NewSparseFirewallOption() *SparseFirewallOption {
	return &SparseFirewallOption{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFirewallOption) Identity() elemental.Identity {

	return FirewallOptionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFirewallOption) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFirewallOption) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallOption) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFirewallOption{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallOption) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFirewallOption{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFirewallOption) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFirewallOption) ToPlain() elemental.PlainIdentifiable {

	out := NewFirewallOption()
	if o.RegionFilter != nil {
		out.RegionFilter = *o.RegionFilter
	}
	if o.RegionVPCs != nil {
		out.RegionVPCs = *o.RegionVPCs
	}

	return out
}

// DeepCopy returns a deep copy if the SparseFirewallOption.
func (o *SparseFirewallOption) DeepCopy() *SparseFirewallOption {

	if o == nil {
		return nil
	}

	out := &SparseFirewallOption{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFirewallOption.
func (o *SparseFirewallOption) DeepCopyInto(out *SparseFirewallOption) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFirewallOption: %s", err))
	}

	*out = *target.(*SparseFirewallOption)
}

type mongoAttributesFirewallOption struct {
}
type mongoAttributesSparseFirewallOption struct {
}