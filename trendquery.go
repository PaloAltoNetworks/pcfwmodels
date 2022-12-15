// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/model.gotpl)

package api

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TrendQueryTrendTypeValue represents the possible values for attribute "trendType".
type TrendQueryTrendTypeValue string

const (
	// TrendQueryTrendTypeEgressBytes represents the value EgressBytes.
	TrendQueryTrendTypeEgressBytes TrendQueryTrendTypeValue = "EgressBytes"

	// TrendQueryTrendTypeIngressBytes represents the value IngressBytes.
	TrendQueryTrendTypeIngressBytes TrendQueryTrendTypeValue = "IngressBytes"

	// TrendQueryTrendTypeThreatsBlocked represents the value ThreatsBlocked.
	TrendQueryTrendTypeThreatsBlocked TrendQueryTrendTypeValue = "ThreatsBlocked"

	// TrendQueryTrendTypeThreatsDetected represents the value ThreatsDetected.
	TrendQueryTrendTypeThreatsDetected TrendQueryTrendTypeValue = "ThreatsDetected"

	// TrendQueryTrendTypeTraffic represents the value Traffic.
	TrendQueryTrendTypeTraffic TrendQueryTrendTypeValue = "Traffic"
)

// TrendQueryIdentity represents the Identity of the object.
var TrendQueryIdentity = elemental.Identity{
	Name:     "trendquery",
	Category: "trendqueries",
	Package:  "logserv",
	Private:  false,
}

// TrendQueriesList represents a list of TrendQueries
type TrendQueriesList []*TrendQuery

// Identity returns the identity of the objects in the list.
func (o TrendQueriesList) Identity() elemental.Identity {

	return TrendQueryIdentity
}

// Copy returns a pointer to a copy the TrendQueriesList.
func (o TrendQueriesList) Copy() elemental.Identifiables {

	copy := append(TrendQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TrendQueriesList.
func (o TrendQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TrendQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*TrendQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TrendQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TrendQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TrendQueriesList converted to SparseTrendQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TrendQueriesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTrendQueriesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTrendQuery)
	}

	return out
}

// Version returns the version of the content.
func (o TrendQueriesList) Version() int {

	return 1
}

// TrendQuery represents the model of a trendquery
type TrendQuery struct {
	// The NGFW name.
	FirewallName string `json:"firewallName" msgpack:"firewallName" bson:"firewallname" mapstructure:"firewallName,omitempty"`

	// The result of the trend query.
	TrendResult []*TrendQueryItem `json:"trendResult" msgpack:"trendResult" bson:"-" mapstructure:"trendResult,omitempty"`

	// The type of field to get trends for.
	TrendType TrendQueryTrendTypeValue `json:"trendType" msgpack:"trendType" bson:"-" mapstructure:"trendType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTrendQuery returns a new *TrendQuery
func NewTrendQuery() *TrendQuery {

	return &TrendQuery{
		ModelVersion: 1,
		TrendResult:  []*TrendQueryItem{},
	}
}

// Identity returns the Identity of the object.
func (o *TrendQuery) Identity() elemental.Identity {

	return TrendQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TrendQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TrendQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TrendQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTrendQuery{}

	s.FirewallName = o.FirewallName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *TrendQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTrendQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.FirewallName = s.FirewallName

	return nil
}

// Version returns the hardcoded version of the model.
func (o *TrendQuery) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *TrendQuery) BleveType() string {

	return "trendquery"
}

// DefaultOrder returns the list of default ordering fields.
func (o *TrendQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TrendQuery) Doc() string {

	return `Answer trend queries on firewall logs.`
}

func (o *TrendQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *TrendQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTrendQuery{
			FirewallName: &o.FirewallName,
			TrendResult:  &o.TrendResult,
			TrendType:    &o.TrendType,
		}
	}

	sp := &SparseTrendQuery{}
	for _, f := range fields {
		switch f {
		case "firewallName":
			sp.FirewallName = &(o.FirewallName)
		case "trendResult":
			sp.TrendResult = &(o.TrendResult)
		case "trendType":
			sp.TrendType = &(o.TrendType)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTrendQuery to the object.
func (o *TrendQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTrendQuery)
	if so.FirewallName != nil {
		o.FirewallName = *so.FirewallName
	}
	if so.TrendResult != nil {
		o.TrendResult = *so.TrendResult
	}
	if so.TrendType != nil {
		o.TrendType = *so.TrendType
	}
}

// DeepCopy returns a deep copy if the TrendQuery.
func (o *TrendQuery) DeepCopy() *TrendQuery {

	if o == nil {
		return nil
	}

	out := &TrendQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TrendQuery.
func (o *TrendQuery) DeepCopyInto(out *TrendQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TrendQuery: %s", err))
	}

	*out = *target.(*TrendQuery)
}

// Validate valides the current information stored into the structure.
func (o *TrendQuery) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("firewallName", o.FirewallName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	for _, sub := range o.TrendResult {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateRequiredString("trendType", string(o.TrendType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("trendType", string(o.TrendType), []string{"Traffic", "IngressBytes", "EgressBytes", "ThreatsDetected", "ThreatsBlocked"}, false); err != nil {
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
func (*TrendQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TrendQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TrendQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TrendQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TrendQueryAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *TrendQuery) ValueForAttribute(name string) interface{} {

	switch name {
	case "firewallName":
		return o.FirewallName
	case "trendResult":
		return o.TrendResult
	case "trendType":
		return o.TrendType
	}

	return nil
}

// TrendQueryAttributesMap represents the map of attribute for TrendQuery.
var TrendQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"FirewallName": {
		AllowedChoices: []string{},
		BSONFieldName:  "firewallname",
		ConvertedName:  "FirewallName",
		Description:    `The NGFW name.`,
		Exposed:        true,
		Name:           "firewallName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"TrendResult": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TrendResult",
		Description:    `The result of the trend query.`,
		Exposed:        true,
		Name:           "trendResult",
		ReadOnly:       true,
		SubType:        "trendqueryitem",
		Type:           "refList",
	},
	"TrendType": {
		AllowedChoices: []string{"Traffic", "IngressBytes", "EgressBytes", "ThreatsDetected", "ThreatsBlocked"},
		ConvertedName:  "TrendType",
		Description:    `The type of field to get trends for.`,
		Exposed:        true,
		Name:           "trendType",
		Required:       true,
		Type:           "enum",
	},
}

// TrendQueryLowerCaseAttributesMap represents the map of attribute for TrendQuery.
var TrendQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"firewallname": {
		AllowedChoices: []string{},
		BSONFieldName:  "firewallname",
		ConvertedName:  "FirewallName",
		Description:    `The NGFW name.`,
		Exposed:        true,
		Name:           "firewallName",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"trendresult": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "TrendResult",
		Description:    `The result of the trend query.`,
		Exposed:        true,
		Name:           "trendResult",
		ReadOnly:       true,
		SubType:        "trendqueryitem",
		Type:           "refList",
	},
	"trendtype": {
		AllowedChoices: []string{"Traffic", "IngressBytes", "EgressBytes", "ThreatsDetected", "ThreatsBlocked"},
		ConvertedName:  "TrendType",
		Description:    `The type of field to get trends for.`,
		Exposed:        true,
		Name:           "trendType",
		Required:       true,
		Type:           "enum",
	},
}

// SparseTrendQueriesList represents a list of SparseTrendQueries
type SparseTrendQueriesList []*SparseTrendQuery

// Identity returns the identity of the objects in the list.
func (o SparseTrendQueriesList) Identity() elemental.Identity {

	return TrendQueryIdentity
}

// Copy returns a pointer to a copy the SparseTrendQueriesList.
func (o SparseTrendQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseTrendQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTrendQueriesList.
func (o SparseTrendQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTrendQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTrendQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTrendQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTrendQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTrendQueriesList converted to TrendQueriesList.
func (o SparseTrendQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTrendQueriesList) Version() int {

	return 1
}

// SparseTrendQuery represents the sparse version of a trendquery.
type SparseTrendQuery struct {
	// The NGFW name.
	FirewallName *string `json:"firewallName,omitempty" msgpack:"firewallName,omitempty" bson:"firewallname,omitempty" mapstructure:"firewallName,omitempty"`

	// The result of the trend query.
	TrendResult *[]*TrendQueryItem `json:"trendResult,omitempty" msgpack:"trendResult,omitempty" bson:"-" mapstructure:"trendResult,omitempty"`

	// The type of field to get trends for.
	TrendType *TrendQueryTrendTypeValue `json:"trendType,omitempty" msgpack:"trendType,omitempty" bson:"-" mapstructure:"trendType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTrendQuery returns a new  SparseTrendQuery.
func NewSparseTrendQuery() *SparseTrendQuery {
	return &SparseTrendQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTrendQuery) Identity() elemental.Identity {

	return TrendQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTrendQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTrendQuery) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrendQuery) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTrendQuery{}

	if o.FirewallName != nil {
		s.FirewallName = o.FirewallName
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTrendQuery) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTrendQuery{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.FirewallName != nil {
		o.FirewallName = s.FirewallName
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTrendQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTrendQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewTrendQuery()
	if o.FirewallName != nil {
		out.FirewallName = *o.FirewallName
	}
	if o.TrendResult != nil {
		out.TrendResult = *o.TrendResult
	}
	if o.TrendType != nil {
		out.TrendType = *o.TrendType
	}

	return out
}

// DeepCopy returns a deep copy if the SparseTrendQuery.
func (o *SparseTrendQuery) DeepCopy() *SparseTrendQuery {

	if o == nil {
		return nil
	}

	out := &SparseTrendQuery{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTrendQuery.
func (o *SparseTrendQuery) DeepCopyInto(out *SparseTrendQuery) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTrendQuery: %s", err))
	}

	*out = *target.(*SparseTrendQuery)
}

type mongoAttributesTrendQuery struct {
	FirewallName string `bson:"firewallname"`
}
type mongoAttributesSparseTrendQuery struct {
	FirewallName *string `bson:"firewallname,omitempty"`
}
