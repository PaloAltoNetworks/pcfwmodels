// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallURLCategoryListIdentity represents the Identity of the object.
var FirewallURLCategoryListIdentity = elemental.Identity{
	Name:     "firewallurlcategorylist",
	Category: "firewallurlcategorylists",
	Package:  "ngfw",
	Private:  false,
}

// FirewallURLCategoryListsList represents a list of FirewallURLCategoryLists
type FirewallURLCategoryListsList []*FirewallURLCategoryList

// Identity returns the identity of the objects in the list.
func (o FirewallURLCategoryListsList) Identity() elemental.Identity {

	return FirewallURLCategoryListIdentity
}

// Copy returns a pointer to a copy the FirewallURLCategoryListsList.
func (o FirewallURLCategoryListsList) Copy() elemental.Identifiables {

	out := append(FirewallURLCategoryListsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the FirewallURLCategoryListsList.
func (o FirewallURLCategoryListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FirewallURLCategoryListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FirewallURLCategoryList))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FirewallURLCategoryListsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FirewallURLCategoryListsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FirewallURLCategoryListsList converted to SparseFirewallURLCategoryListsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FirewallURLCategoryListsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFirewallURLCategoryListsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFirewallURLCategoryList)
	}

	return out
}

// Version returns the version of the content.
func (o FirewallURLCategoryListsList) Version() int {

	return 1
}

// FirewallURLCategoryList represents the model of a firewallurlcategorylist
type FirewallURLCategoryList struct {
	// A list of NGFW URL category objects.
	Categories []*FirewallURLCategory `json:"categories" msgpack:"categories" bson:"-" mapstructure:"categories,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallURLCategoryList returns a new *FirewallURLCategoryList
func NewFirewallURLCategoryList() *FirewallURLCategoryList {

	return &FirewallURLCategoryList{
		ModelVersion: 1,
		Categories:   []*FirewallURLCategory{},
	}
}

// Identity returns the Identity of the object.
func (o *FirewallURLCategoryList) Identity() elemental.Identity {

	return FirewallURLCategoryListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallURLCategoryList) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallURLCategoryList) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallURLCategoryList) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallURLCategoryList{}

	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallURLCategoryList) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallURLCategoryList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FirewallURLCategoryList) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallURLCategoryList) BleveType() string {

	return "firewallurlcategorylist"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FirewallURLCategoryList) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FirewallURLCategoryList) Doc() string {

	return `This a read-only list that returns the default NGFW URL categories.`
}

func (o *FirewallURLCategoryList) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *FirewallURLCategoryList) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *FirewallURLCategoryList) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FirewallURLCategoryList) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFirewallURLCategoryList{
			Categories: &o.Categories,
			Namespace:  &o.Namespace,
		}
	}

	sp := &SparseFirewallURLCategoryList{}
	for _, f := range fields {
		switch f {
		case "categories":
			sp.Categories = &(o.Categories)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFirewallURLCategoryList to the object.
func (o *FirewallURLCategoryList) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFirewallURLCategoryList)
	if so.Categories != nil {
		o.Categories = *so.Categories
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the FirewallURLCategoryList.
func (o *FirewallURLCategoryList) DeepCopy() *FirewallURLCategoryList {

	if o == nil {
		return nil
	}

	out := &FirewallURLCategoryList{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallURLCategoryList.
func (o *FirewallURLCategoryList) DeepCopyInto(out *FirewallURLCategoryList) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallURLCategoryList: %s", err))
	}

	*out = *target.(*FirewallURLCategoryList)
}

// Validate valides the current information stored into the structure.
func (o *FirewallURLCategoryList) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Categories {
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
func (*FirewallURLCategoryList) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallURLCategoryListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallURLCategoryListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallURLCategoryList) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallURLCategoryListAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallURLCategoryList) ValueForAttribute(name string) any {

	switch name {
	case "categories":
		return o.Categories
	case "namespace":
		return o.Namespace
	}

	return nil
}

// FirewallURLCategoryListAttributesMap represents the map of attribute for FirewallURLCategoryList.
var FirewallURLCategoryListAttributesMap = map[string]elemental.AttributeSpecification{
	"Categories": {
		AllowedChoices: []string{},
		ConvertedName:  "Categories",
		Description:    `A list of NGFW URL category objects.`,
		Exposed:        true,
		Name:           "categories",
		ReadOnly:       true,
		SubType:        "firewallurlcategory",
		Type:           "refList",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// FirewallURLCategoryListLowerCaseAttributesMap represents the map of attribute for FirewallURLCategoryList.
var FirewallURLCategoryListLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"categories": {
		AllowedChoices: []string{},
		ConvertedName:  "Categories",
		Description:    `A list of NGFW URL category objects.`,
		Exposed:        true,
		Name:           "categories",
		ReadOnly:       true,
		SubType:        "firewallurlcategory",
		Type:           "refList",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseFirewallURLCategoryListsList represents a list of SparseFirewallURLCategoryLists
type SparseFirewallURLCategoryListsList []*SparseFirewallURLCategoryList

// Identity returns the identity of the objects in the list.
func (o SparseFirewallURLCategoryListsList) Identity() elemental.Identity {

	return FirewallURLCategoryListIdentity
}

// Copy returns a pointer to a copy the SparseFirewallURLCategoryListsList.
func (o SparseFirewallURLCategoryListsList) Copy() elemental.Identifiables {

	copy := append(SparseFirewallURLCategoryListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFirewallURLCategoryListsList.
func (o SparseFirewallURLCategoryListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFirewallURLCategoryListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFirewallURLCategoryList))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFirewallURLCategoryListsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFirewallURLCategoryListsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFirewallURLCategoryListsList converted to FirewallURLCategoryListsList.
func (o SparseFirewallURLCategoryListsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFirewallURLCategoryListsList) Version() int {

	return 1
}

// SparseFirewallURLCategoryList represents the sparse version of a firewallurlcategorylist.
type SparseFirewallURLCategoryList struct {
	// A list of NGFW URL category objects.
	Categories *[]*FirewallURLCategory `json:"categories,omitempty" msgpack:"categories,omitempty" bson:"-" mapstructure:"categories,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFirewallURLCategoryList returns a new  SparseFirewallURLCategoryList.
func NewSparseFirewallURLCategoryList() *SparseFirewallURLCategoryList {
	return &SparseFirewallURLCategoryList{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFirewallURLCategoryList) Identity() elemental.Identity {

	return FirewallURLCategoryListIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFirewallURLCategoryList) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFirewallURLCategoryList) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallURLCategoryList) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFirewallURLCategoryList{}

	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallURLCategoryList) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFirewallURLCategoryList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFirewallURLCategoryList) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFirewallURLCategoryList) ToPlain() elemental.PlainIdentifiable {

	out := NewFirewallURLCategoryList()
	if o.Categories != nil {
		out.Categories = *o.Categories
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseFirewallURLCategoryList) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseFirewallURLCategoryList) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseFirewallURLCategoryList.
func (o *SparseFirewallURLCategoryList) DeepCopy() *SparseFirewallURLCategoryList {

	if o == nil {
		return nil
	}

	out := &SparseFirewallURLCategoryList{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFirewallURLCategoryList.
func (o *SparseFirewallURLCategoryList) DeepCopyInto(out *SparseFirewallURLCategoryList) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFirewallURLCategoryList: %s", err))
	}

	*out = *target.(*SparseFirewallURLCategoryList)
}

type mongoAttributesFirewallURLCategoryList struct {
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparseFirewallURLCategoryList struct {
	Namespace *string `bson:"namespace,omitempty"`
}
