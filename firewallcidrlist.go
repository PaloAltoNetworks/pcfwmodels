// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FirewallCIDRlistIdentity represents the Identity of the object.
var FirewallCIDRlistIdentity = elemental.Identity{
	Name:     "firewallcidrlist",
	Category: "firewallcidrlists",
	Package:  "ngfw",
	Private:  false,
}

// FirewallCIDRlistsList represents a list of FirewallCIDRlists
type FirewallCIDRlistsList []*FirewallCIDRlist

// Identity returns the identity of the objects in the list.
func (o FirewallCIDRlistsList) Identity() elemental.Identity {

	return FirewallCIDRlistIdentity
}

// Copy returns a pointer to a copy the FirewallCIDRlistsList.
func (o FirewallCIDRlistsList) Copy() elemental.Identifiables {

	copy := append(FirewallCIDRlistsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FirewallCIDRlistsList.
func (o FirewallCIDRlistsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FirewallCIDRlistsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FirewallCIDRlist))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FirewallCIDRlistsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FirewallCIDRlistsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the FirewallCIDRlistsList converted to SparseFirewallCIDRlistsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FirewallCIDRlistsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFirewallCIDRlistsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFirewallCIDRlist)
	}

	return out
}

// Version returns the version of the content.
func (o FirewallCIDRlistsList) Version() int {

	return 1
}

// FirewallCIDRlist represents the model of a firewallcidrlist
type FirewallCIDRlist struct {
	// List of CIDRs.
	CIDRs []string `json:"CIDRs" msgpack:"CIDRs" bson:"cidrs" mapstructure:"CIDRs,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// List of tags attached to an entity.
	Tags []string `json:"tags" msgpack:"tags" bson:"tags" mapstructure:"tags,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFirewallCIDRlist returns a new *FirewallCIDRlist
func NewFirewallCIDRlist() *FirewallCIDRlist {

	return &FirewallCIDRlist{
		ModelVersion: 1,
		CIDRs:        []string{},
		Tags:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *FirewallCIDRlist) Identity() elemental.Identity {

	return FirewallCIDRlistIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallCIDRlist) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallCIDRlist) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCIDRlist) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFirewallCIDRlist{}

	s.CIDRs = o.CIDRs
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.Tags = o.Tags
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FirewallCIDRlist) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFirewallCIDRlist{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.CIDRs = s.CIDRs
	o.ID = s.ID.Hex()
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.Tags = s.Tags
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FirewallCIDRlist) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FirewallCIDRlist) BleveType() string {

	return "firewallcidrlist"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FirewallCIDRlist) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *FirewallCIDRlist) Doc() string {

	return `Represents a list of CIDRs referenced by a TargetCriteria object.`
}

func (o *FirewallCIDRlist) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *FirewallCIDRlist) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *FirewallCIDRlist) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *FirewallCIDRlist) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *FirewallCIDRlist) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *FirewallCIDRlist) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *FirewallCIDRlist) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *FirewallCIDRlist) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *FirewallCIDRlist) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetTags returns the Tags of the receiver.
func (o *FirewallCIDRlist) GetTags() []string {

	return o.Tags
}

// SetTags sets the property Tags of the receiver using the given value.
func (o *FirewallCIDRlist) SetTags(tags []string) {

	o.Tags = tags
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *FirewallCIDRlist) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *FirewallCIDRlist) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FirewallCIDRlist) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFirewallCIDRlist{
			CIDRs:       &o.CIDRs,
			ID:          &o.ID,
			CreateTime:  &o.CreateTime,
			Description: &o.Description,
			Name:        &o.Name,
			Namespace:   &o.Namespace,
			Tags:        &o.Tags,
			UpdateTime:  &o.UpdateTime,
			ZHash:       &o.ZHash,
			Zone:        &o.Zone,
		}
	}

	sp := &SparseFirewallCIDRlist{}
	for _, f := range fields {
		switch f {
		case "CIDRs":
			sp.CIDRs = &(o.CIDRs)
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "tags":
			sp.Tags = &(o.Tags)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFirewallCIDRlist to the object.
func (o *FirewallCIDRlist) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFirewallCIDRlist)
	if so.CIDRs != nil {
		o.CIDRs = *so.CIDRs
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Tags != nil {
		o.Tags = *so.Tags
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the FirewallCIDRlist.
func (o *FirewallCIDRlist) DeepCopy() *FirewallCIDRlist {

	if o == nil {
		return nil
	}

	out := &FirewallCIDRlist{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FirewallCIDRlist.
func (o *FirewallCIDRlist) DeepCopyInto(out *FirewallCIDRlist) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FirewallCIDRlist: %s", err))
	}

	*out = *target.(*FirewallCIDRlist)
}

// Validate valides the current information stored into the structure.
func (o *FirewallCIDRlist) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("CIDRs", o.CIDRs); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateCIDRList("CIDRs", o.CIDRs); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("tags", o.Tags); err != nil {
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
func (*FirewallCIDRlist) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FirewallCIDRlistAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FirewallCIDRlistLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FirewallCIDRlist) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FirewallCIDRlistAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FirewallCIDRlist) ValueForAttribute(name string) interface{} {

	switch name {
	case "CIDRs":
		return o.CIDRs
	case "ID":
		return o.ID
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "tags":
		return o.Tags
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// FirewallCIDRlistAttributesMap represents the map of attribute for FirewallCIDRlist.
var FirewallCIDRlistAttributesMap = map[string]elemental.AttributeSpecification{
	"CIDRs": {
		AllowedChoices: []string{},
		BSONFieldName:  "cidrs",
		ConvertedName:  "CIDRs",
		Description:    `List of CIDRs.`,
		Exposed:        true,
		Name:           "CIDRs",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
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
	"Tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "tags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// FirewallCIDRlistLowerCaseAttributesMap represents the map of attribute for FirewallCIDRlist.
var FirewallCIDRlistLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cidrs": {
		AllowedChoices: []string{},
		BSONFieldName:  "cidrs",
		ConvertedName:  "CIDRs",
		Description:    `List of CIDRs.`,
		Exposed:        true,
		Name:           "CIDRs",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
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
	"tags": {
		AllowedChoices: []string{},
		BSONFieldName:  "tags",
		ConvertedName:  "Tags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "tags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseFirewallCIDRlistsList represents a list of SparseFirewallCIDRlists
type SparseFirewallCIDRlistsList []*SparseFirewallCIDRlist

// Identity returns the identity of the objects in the list.
func (o SparseFirewallCIDRlistsList) Identity() elemental.Identity {

	return FirewallCIDRlistIdentity
}

// Copy returns a pointer to a copy the SparseFirewallCIDRlistsList.
func (o SparseFirewallCIDRlistsList) Copy() elemental.Identifiables {

	copy := append(SparseFirewallCIDRlistsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFirewallCIDRlistsList.
func (o SparseFirewallCIDRlistsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFirewallCIDRlistsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFirewallCIDRlist))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFirewallCIDRlistsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFirewallCIDRlistsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseFirewallCIDRlistsList converted to FirewallCIDRlistsList.
func (o SparseFirewallCIDRlistsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFirewallCIDRlistsList) Version() int {

	return 1
}

// SparseFirewallCIDRlist represents the sparse version of a firewallcidrlist.
type SparseFirewallCIDRlist struct {
	// List of CIDRs.
	CIDRs *[]string `json:"CIDRs,omitempty" msgpack:"CIDRs,omitempty" bson:"cidrs,omitempty" mapstructure:"CIDRs,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// List of tags attached to an entity.
	Tags *[]string `json:"tags,omitempty" msgpack:"tags,omitempty" bson:"tags,omitempty" mapstructure:"tags,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFirewallCIDRlist returns a new  SparseFirewallCIDRlist.
func NewSparseFirewallCIDRlist() *SparseFirewallCIDRlist {
	return &SparseFirewallCIDRlist{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFirewallCIDRlist) Identity() elemental.Identity {

	return FirewallCIDRlistIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFirewallCIDRlist) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFirewallCIDRlist) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallCIDRlist) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFirewallCIDRlist{}

	if o.CIDRs != nil {
		s.CIDRs = o.CIDRs
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Tags != nil {
		s.Tags = o.Tags
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFirewallCIDRlist) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFirewallCIDRlist{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.CIDRs != nil {
		o.CIDRs = s.CIDRs
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Tags != nil {
		o.Tags = s.Tags
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFirewallCIDRlist) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFirewallCIDRlist) ToPlain() elemental.PlainIdentifiable {

	out := NewFirewallCIDRlist()
	if o.CIDRs != nil {
		out.CIDRs = *o.CIDRs
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Tags != nil {
		out.Tags = *o.Tags
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseFirewallCIDRlist) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseFirewallCIDRlist) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseFirewallCIDRlist) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseFirewallCIDRlist) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetTags returns the Tags of the receiver.
func (o *SparseFirewallCIDRlist) GetTags() (out []string) {

	if o.Tags == nil {
		return
	}

	return *o.Tags
}

// SetTags sets the property Tags of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetTags(tags []string) {

	o.Tags = &tags
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseFirewallCIDRlist) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseFirewallCIDRlist) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseFirewallCIDRlist.
func (o *SparseFirewallCIDRlist) DeepCopy() *SparseFirewallCIDRlist {

	if o == nil {
		return nil
	}

	out := &SparseFirewallCIDRlist{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFirewallCIDRlist.
func (o *SparseFirewallCIDRlist) DeepCopyInto(out *SparseFirewallCIDRlist) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFirewallCIDRlist: %s", err))
	}

	*out = *target.(*SparseFirewallCIDRlist)
}

type mongoAttributesFirewallCIDRlist struct {
	CIDRs       []string      `bson:"cidrs"`
	ID          bson.ObjectId `bson:"_id,omitempty"`
	CreateTime  time.Time     `bson:"createtime"`
	Description string        `bson:"description"`
	Name        string        `bson:"name"`
	Namespace   string        `bson:"namespace"`
	Tags        []string      `bson:"tags"`
	UpdateTime  time.Time     `bson:"updatetime"`
	ZHash       int           `bson:"zhash"`
	Zone        int           `bson:"zone"`
}
type mongoAttributesSparseFirewallCIDRlist struct {
	CIDRs       *[]string     `bson:"cidrs,omitempty"`
	ID          bson.ObjectId `bson:"_id,omitempty"`
	CreateTime  *time.Time    `bson:"createtime,omitempty"`
	Description *string       `bson:"description,omitempty"`
	Name        *string       `bson:"name,omitempty"`
	Namespace   *string       `bson:"namespace,omitempty"`
	Tags        *[]string     `bson:"tags,omitempty"`
	UpdateTime  *time.Time    `bson:"updatetime,omitempty"`
	ZHash       *int          `bson:"zhash,omitempty"`
	Zone        *int          `bson:"zone,omitempty"`
}
