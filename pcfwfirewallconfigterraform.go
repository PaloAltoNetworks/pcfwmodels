// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCFWFirewallConfigTerraformIdentity represents the Identity of the object.
var PCFWFirewallConfigTerraformIdentity = elemental.Identity{
	Name:     "pcfwfirewallconfigterraform",
	Category: "pcfwfirewallconfigterraform",
	Package:  "deploymentadvisor",
	Private:  false,
}

// PCFWFirewallConfigTerraformsList represents a list of PCFWFirewallConfigTerraforms
type PCFWFirewallConfigTerraformsList []*PCFWFirewallConfigTerraform

// Identity returns the identity of the objects in the list.
func (o PCFWFirewallConfigTerraformsList) Identity() elemental.Identity {

	return PCFWFirewallConfigTerraformIdentity
}

// Copy returns a pointer to a copy the PCFWFirewallConfigTerraformsList.
func (o PCFWFirewallConfigTerraformsList) Copy() elemental.Identifiables {

	out := append(PCFWFirewallConfigTerraformsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the PCFWFirewallConfigTerraformsList.
func (o PCFWFirewallConfigTerraformsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCFWFirewallConfigTerraformsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCFWFirewallConfigTerraform))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCFWFirewallConfigTerraformsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCFWFirewallConfigTerraformsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the PCFWFirewallConfigTerraformsList converted to SparsePCFWFirewallConfigTerraformsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCFWFirewallConfigTerraformsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCFWFirewallConfigTerraformsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCFWFirewallConfigTerraform)
	}

	return out
}

// Version returns the version of the content.
func (o PCFWFirewallConfigTerraformsList) Version() int {

	return 1
}

// PCFWFirewallConfigTerraform represents the model of a pcfwfirewallconfigterraform
type PCFWFirewallConfigTerraform struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPCFWFirewallConfigTerraform returns a new *PCFWFirewallConfigTerraform
func NewPCFWFirewallConfigTerraform() *PCFWFirewallConfigTerraform {

	return &PCFWFirewallConfigTerraform{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PCFWFirewallConfigTerraform) Identity() elemental.Identity {

	return PCFWFirewallConfigTerraformIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCFWFirewallConfigTerraform) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCFWFirewallConfigTerraform) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCFWFirewallConfigTerraform) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCFWFirewallConfigTerraform{}

	s.Name = o.Name
	s.Namespace = o.Namespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCFWFirewallConfigTerraform) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCFWFirewallConfigTerraform{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Name = s.Name
	o.Namespace = s.Namespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PCFWFirewallConfigTerraform) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCFWFirewallConfigTerraform) BleveType() string {

	return "pcfwfirewallconfigterraform"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCFWFirewallConfigTerraform) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *PCFWFirewallConfigTerraform) Doc() string {

	return `Represents PCFW firewall configuration terraform generator.`
}

func (o *PCFWFirewallConfigTerraform) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *PCFWFirewallConfigTerraform) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *PCFWFirewallConfigTerraform) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *PCFWFirewallConfigTerraform) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *PCFWFirewallConfigTerraform) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCFWFirewallConfigTerraform) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCFWFirewallConfigTerraform{
			ID:        &o.ID,
			Name:      &o.Name,
			Namespace: &o.Namespace,
		}
	}

	sp := &SparsePCFWFirewallConfigTerraform{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePCFWFirewallConfigTerraform to the object.
func (o *PCFWFirewallConfigTerraform) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCFWFirewallConfigTerraform)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
}

// DeepCopy returns a deep copy if the PCFWFirewallConfigTerraform.
func (o *PCFWFirewallConfigTerraform) DeepCopy() *PCFWFirewallConfigTerraform {

	if o == nil {
		return nil
	}

	out := &PCFWFirewallConfigTerraform{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCFWFirewallConfigTerraform.
func (o *PCFWFirewallConfigTerraform) DeepCopyInto(out *PCFWFirewallConfigTerraform) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCFWFirewallConfigTerraform: %s", err))
	}

	*out = *target.(*PCFWFirewallConfigTerraform)
}

// Validate valides the current information stored into the structure.
func (o *PCFWFirewallConfigTerraform) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*PCFWFirewallConfigTerraform) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCFWFirewallConfigTerraformAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCFWFirewallConfigTerraformLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCFWFirewallConfigTerraform) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCFWFirewallConfigTerraformAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCFWFirewallConfigTerraform) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	}

	return nil
}

// PCFWFirewallConfigTerraformAttributesMap represents the map of attribute for PCFWFirewallConfigTerraform.
var PCFWFirewallConfigTerraformAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// PCFWFirewallConfigTerraformLowerCaseAttributesMap represents the map of attribute for PCFWFirewallConfigTerraform.
var PCFWFirewallConfigTerraformLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparsePCFWFirewallConfigTerraformsList represents a list of SparsePCFWFirewallConfigTerraforms
type SparsePCFWFirewallConfigTerraformsList []*SparsePCFWFirewallConfigTerraform

// Identity returns the identity of the objects in the list.
func (o SparsePCFWFirewallConfigTerraformsList) Identity() elemental.Identity {

	return PCFWFirewallConfigTerraformIdentity
}

// Copy returns a pointer to a copy the SparsePCFWFirewallConfigTerraformsList.
func (o SparsePCFWFirewallConfigTerraformsList) Copy() elemental.Identifiables {

	copy := append(SparsePCFWFirewallConfigTerraformsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCFWFirewallConfigTerraformsList.
func (o SparsePCFWFirewallConfigTerraformsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCFWFirewallConfigTerraformsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCFWFirewallConfigTerraform))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCFWFirewallConfigTerraformsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCFWFirewallConfigTerraformsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparsePCFWFirewallConfigTerraformsList converted to PCFWFirewallConfigTerraformsList.
func (o SparsePCFWFirewallConfigTerraformsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCFWFirewallConfigTerraformsList) Version() int {

	return 1
}

// SparsePCFWFirewallConfigTerraform represents the sparse version of a pcfwfirewallconfigterraform.
type SparsePCFWFirewallConfigTerraform struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePCFWFirewallConfigTerraform returns a new  SparsePCFWFirewallConfigTerraform.
func NewSparsePCFWFirewallConfigTerraform() *SparsePCFWFirewallConfigTerraform {
	return &SparsePCFWFirewallConfigTerraform{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCFWFirewallConfigTerraform) Identity() elemental.Identity {

	return PCFWFirewallConfigTerraformIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCFWFirewallConfigTerraform) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCFWFirewallConfigTerraform) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCFWFirewallConfigTerraform) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCFWFirewallConfigTerraform{}

	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCFWFirewallConfigTerraform) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCFWFirewallConfigTerraform{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePCFWFirewallConfigTerraform) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCFWFirewallConfigTerraform) ToPlain() elemental.PlainIdentifiable {

	out := NewPCFWFirewallConfigTerraform()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparsePCFWFirewallConfigTerraform) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparsePCFWFirewallConfigTerraform) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePCFWFirewallConfigTerraform) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePCFWFirewallConfigTerraform) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparsePCFWFirewallConfigTerraform.
func (o *SparsePCFWFirewallConfigTerraform) DeepCopy() *SparsePCFWFirewallConfigTerraform {

	if o == nil {
		return nil
	}

	out := &SparsePCFWFirewallConfigTerraform{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCFWFirewallConfigTerraform.
func (o *SparsePCFWFirewallConfigTerraform) DeepCopyInto(out *SparsePCFWFirewallConfigTerraform) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCFWFirewallConfigTerraform: %s", err))
	}

	*out = *target.(*SparsePCFWFirewallConfigTerraform)
}

type mongoAttributesPCFWFirewallConfigTerraform struct {
	Name      string `bson:"name"`
	Namespace string `bson:"namespace"`
}
type mongoAttributesSparsePCFWFirewallConfigTerraform struct {
	Name      *string `bson:"name,omitempty"`
	Namespace *string `bson:"namespace,omitempty"`
}