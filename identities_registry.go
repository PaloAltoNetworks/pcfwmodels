// Code generated by elegen. DO NOT EDIT.
// Source: go.aporeto.io/elemental (templates/identities_registry.gotpl)

package api

import "go.aporeto.io/elemental"

var (
	identityNamesMap = map[string]elemental.Identity{

		"awsfirewall":       AWSFirewallIdentity,
		"awslogdefinition":  AWSLogDefinitionIdentity,
		"firewallappidlist": FirewallAppIDListIdentity,

		"firewallapplicationgroup": FirewallApplicationGroupIdentity,

		"firewallcertificate":      FirewallCertificateIdentity,
		"firewallcidrlist":         FirewallCIDRlistIdentity,
		"firewallcommitaction":     FirewallCommitActionIdentity,
		"firewallcommittedruleset": FirewallCommittedRulesetIdentity,

		"firewallcountrycodelist":    FirewallCountryCodelistIdentity,
		"firewallcustomexternalfeed": FirewallCustomExternalFeedIdentity,
		"firewallcustomurlcategory":  FirewallCustomURLCategoryIdentity,

		"firewallexternalfeedlist": FirewallExternalFeedListIdentity,
		"firewallfqdnlist":         FirewallFQDNListIdentity,
		"firewallgroup":            FirewallGroupIdentity,
		"firewalllicensingcredits": FirewallLicensingCreditsIdentity,

		"firewalloption": FirewallOptionIdentity,

		"firewallrollbackaction":  FirewallRollbackActionIdentity,
		"firewallrule":            FirewallRuleIdentity,
		"firewallruleset":         FirewallRulesetIdentity,
		"firewallsecurityprofile": FirewallSecurityProfileIdentity,

		"firewalltemplate": FirewallTemplateIdentity,

		"firewallurlcategorylist": FirewallURLCategoryListIdentity,

		"logquery": LogQueryIdentity,

		"mirrorsourceoption":           MirrorSourceOptionIdentity,
		"pcfwaccount":                  PCFWAccountIdentity,
		"pcfwaccountmirrorsource":      PCFWAccountMirrorSourceIdentity,
		"pcfwaccountmirrorsourcestate": PCFWAccountMirrorSourceStateIdentity,
		"pcfwfirewallconfigterraform":  PCFWFirewallConfigTerraformIdentity,
		"pcfwsubnethelper":             PCFWSubnetHelperIdentity,
		"pcfwtenant":                   PCFWTenantIdentity,
		"pcfwtenantaccountterraform":   PCFWTenantAccountTerraformIdentity,
		"pcfwtenantterraform":          PCFWTenantTerraformIdentity,
		"root":                         RootIdentity,

		"topquery": TopQueryIdentity,

		"totalquery": TotalQueryIdentity,
		"trendquery": TrendQueryIdentity,
	}

	identitycategoriesMap = map[string]elemental.Identity{

		"awsfirewalls":      AWSFirewallIdentity,
		"awslogdefinitions": AWSLogDefinitionIdentity,
		"firewallappidlist": FirewallAppIDListIdentity,

		"firewallapplicationgroups": FirewallApplicationGroupIdentity,

		"firewallcertificates":      FirewallCertificateIdentity,
		"firewallcidrlists":         FirewallCIDRlistIdentity,
		"firewallcommitactions":     FirewallCommitActionIdentity,
		"firewallcommittedrulesets": FirewallCommittedRulesetIdentity,

		"firewallcountrycodelist":     FirewallCountryCodelistIdentity,
		"firewallcustomexternalfeeds": FirewallCustomExternalFeedIdentity,
		"firewallcustomurlcategories": FirewallCustomURLCategoryIdentity,

		"firewallexternalfeedlists": FirewallExternalFeedListIdentity,
		"firewallfqdnlists":         FirewallFQDNListIdentity,
		"firewallgroups":            FirewallGroupIdentity,
		"firewalllicensingcredits":  FirewallLicensingCreditsIdentity,

		"firewalloptions": FirewallOptionIdentity,

		"firewallrollbackactions":  FirewallRollbackActionIdentity,
		"firewallrules":            FirewallRuleIdentity,
		"firewallrulesets":         FirewallRulesetIdentity,
		"firewallsecurityprofiles": FirewallSecurityProfileIdentity,

		"firewalltemplates": FirewallTemplateIdentity,

		"firewallurlcategorylists": FirewallURLCategoryListIdentity,

		"logqueries": LogQueryIdentity,

		"mirrorsourceoptions":           MirrorSourceOptionIdentity,
		"pcfwaccounts":                  PCFWAccountIdentity,
		"pcfwaccountmirrorsources":      PCFWAccountMirrorSourceIdentity,
		"pcfwaccountmirrorsourcestates": PCFWAccountMirrorSourceStateIdentity,
		"pcfwfirewallconfigterraform":   PCFWFirewallConfigTerraformIdentity,
		"pcfwsubnethelpers":             PCFWSubnetHelperIdentity,
		"pcfwtenants":                   PCFWTenantIdentity,
		"pcfwtenantaccountterraform":    PCFWTenantAccountTerraformIdentity,
		"pcfwtenantterraform":           PCFWTenantTerraformIdentity,
		"root":                          RootIdentity,

		"topqueries": TopQueryIdentity,

		"totalqueries": TotalQueryIdentity,
		"trendqueries": TrendQueryIdentity,
	}

	aliasesMap = map[string]elemental.Identity{}

	indexesMap = map[string][][]string{
		"awsfirewall": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"awslogdefinition": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallappidlist": {
			{"namespace"},
		},
		"firewallapplicationgroup": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallcertificate": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallcidrlist": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallcommitaction": nil,
		"firewallcommittedruleset": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"namespace"},
			{"updateIdempotencyKey"},
		},
		"firewallcountrycodelist": {
			{"namespace"},
		},
		"firewallcustomexternalfeed": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallcustomurlcategory": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallexternalfeedlist": {
			{"namespace"},
		},
		"firewallfqdnlist": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallgroup": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewalllicensingcredits": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"namespace"},
			{"namespace", "firewallResource"},
			{"updateIdempotencyKey"},
		},
		"firewalloption":         nil,
		"firewallrollbackaction": nil,
		"firewallrule": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"namespace"},
			{"updateIdempotencyKey"},
		},
		"firewallruleset": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallsecurityprofile": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewalltemplate": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"firewallurlcategorylist": {
			{"namespace"},
		},
		"logquery":           nil,
		"mirrorsourceoption": nil,
		"pcfwaccount": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"namespace"},
			{"updateIdempotencyKey"},
		},
		"pcfwaccountmirrorsource": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"pcfwaccountmirrorsourcestate": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"name"},
			{"namespace"},
			{"namespace", "name"},
			{"updateIdempotencyKey"},
		},
		"pcfwfirewallconfigterraform": {
			{"name"},
			{"namespace"},
			{"namespace", "name"},
		},
		"pcfwsubnethelper": nil,
		"pcfwtenant": {
			{":shard", ":unique", "zone", "zHash"},
			{"createIdempotencyKey"},
			{"namespace"},
			{"updateIdempotencyKey"},
		},
		"pcfwtenantaccountterraform": {
			{"namespace"},
		},
		"pcfwtenantterraform": {
			{"namespace"},
		},
		"root":       nil,
		"topquery":   nil,
		"totalquery": nil,
		"trendquery": nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) elemental.Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) elemental.Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i elemental.Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity elemental.Identity) elemental.Identifiable {

	switch identity {

	case AWSFirewallIdentity:
		return NewAWSFirewall()
	case AWSLogDefinitionIdentity:
		return NewAWSLogDefinition()
	case FirewallAppIDListIdentity:
		return NewFirewallAppIDList()
	case FirewallApplicationGroupIdentity:
		return NewFirewallApplicationGroup()
	case FirewallCertificateIdentity:
		return NewFirewallCertificate()
	case FirewallCIDRlistIdentity:
		return NewFirewallCIDRlist()
	case FirewallCommitActionIdentity:
		return NewFirewallCommitAction()
	case FirewallCommittedRulesetIdentity:
		return NewFirewallCommittedRuleset()
	case FirewallCountryCodelistIdentity:
		return NewFirewallCountryCodelist()
	case FirewallCustomExternalFeedIdentity:
		return NewFirewallCustomExternalFeed()
	case FirewallCustomURLCategoryIdentity:
		return NewFirewallCustomURLCategory()
	case FirewallExternalFeedListIdentity:
		return NewFirewallExternalFeedList()
	case FirewallFQDNListIdentity:
		return NewFirewallFQDNList()
	case FirewallGroupIdentity:
		return NewFirewallGroup()
	case FirewallLicensingCreditsIdentity:
		return NewFirewallLicensingCredits()
	case FirewallOptionIdentity:
		return NewFirewallOption()
	case FirewallRollbackActionIdentity:
		return NewFirewallRollbackAction()
	case FirewallRuleIdentity:
		return NewFirewallRule()
	case FirewallRulesetIdentity:
		return NewFirewallRuleset()
	case FirewallSecurityProfileIdentity:
		return NewFirewallSecurityProfile()
	case FirewallTemplateIdentity:
		return NewFirewallTemplate()
	case FirewallURLCategoryListIdentity:
		return NewFirewallURLCategoryList()
	case LogQueryIdentity:
		return NewLogQuery()
	case MirrorSourceOptionIdentity:
		return NewMirrorSourceOption()
	case PCFWAccountIdentity:
		return NewPCFWAccount()
	case PCFWAccountMirrorSourceIdentity:
		return NewPCFWAccountMirrorSource()
	case PCFWAccountMirrorSourceStateIdentity:
		return NewPCFWAccountMirrorSourceState()
	case PCFWFirewallConfigTerraformIdentity:
		return NewPCFWFirewallConfigTerraform()
	case PCFWSubnetHelperIdentity:
		return NewPCFWSubnetHelper()
	case PCFWTenantIdentity:
		return NewPCFWTenant()
	case PCFWTenantAccountTerraformIdentity:
		return NewPCFWTenantAccountTerraform()
	case PCFWTenantTerraformIdentity:
		return NewPCFWTenantTerraform()
	case RootIdentity:
		return NewRoot()
	case TopQueryIdentity:
		return NewTopQuery()
	case TotalQueryIdentity:
		return NewTotalQuery()
	case TrendQueryIdentity:
		return NewTrendQuery()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity elemental.Identity) elemental.SparseIdentifiable {

	switch identity {

	case AWSFirewallIdentity:
		return NewSparseAWSFirewall()
	case AWSLogDefinitionIdentity:
		return NewSparseAWSLogDefinition()
	case FirewallAppIDListIdentity:
		return NewSparseFirewallAppIDList()
	case FirewallApplicationGroupIdentity:
		return NewSparseFirewallApplicationGroup()
	case FirewallCertificateIdentity:
		return NewSparseFirewallCertificate()
	case FirewallCIDRlistIdentity:
		return NewSparseFirewallCIDRlist()
	case FirewallCommitActionIdentity:
		return NewSparseFirewallCommitAction()
	case FirewallCommittedRulesetIdentity:
		return NewSparseFirewallCommittedRuleset()
	case FirewallCountryCodelistIdentity:
		return NewSparseFirewallCountryCodelist()
	case FirewallCustomExternalFeedIdentity:
		return NewSparseFirewallCustomExternalFeed()
	case FirewallCustomURLCategoryIdentity:
		return NewSparseFirewallCustomURLCategory()
	case FirewallExternalFeedListIdentity:
		return NewSparseFirewallExternalFeedList()
	case FirewallFQDNListIdentity:
		return NewSparseFirewallFQDNList()
	case FirewallGroupIdentity:
		return NewSparseFirewallGroup()
	case FirewallLicensingCreditsIdentity:
		return NewSparseFirewallLicensingCredits()
	case FirewallOptionIdentity:
		return NewSparseFirewallOption()
	case FirewallRollbackActionIdentity:
		return NewSparseFirewallRollbackAction()
	case FirewallRuleIdentity:
		return NewSparseFirewallRule()
	case FirewallRulesetIdentity:
		return NewSparseFirewallRuleset()
	case FirewallSecurityProfileIdentity:
		return NewSparseFirewallSecurityProfile()
	case FirewallTemplateIdentity:
		return NewSparseFirewallTemplate()
	case FirewallURLCategoryListIdentity:
		return NewSparseFirewallURLCategoryList()
	case LogQueryIdentity:
		return NewSparseLogQuery()
	case MirrorSourceOptionIdentity:
		return NewSparseMirrorSourceOption()
	case PCFWAccountIdentity:
		return NewSparsePCFWAccount()
	case PCFWAccountMirrorSourceIdentity:
		return NewSparsePCFWAccountMirrorSource()
	case PCFWAccountMirrorSourceStateIdentity:
		return NewSparsePCFWAccountMirrorSourceState()
	case PCFWFirewallConfigTerraformIdentity:
		return NewSparsePCFWFirewallConfigTerraform()
	case PCFWSubnetHelperIdentity:
		return NewSparsePCFWSubnetHelper()
	case PCFWTenantIdentity:
		return NewSparsePCFWTenant()
	case PCFWTenantAccountTerraformIdentity:
		return NewSparsePCFWTenantAccountTerraform()
	case PCFWTenantTerraformIdentity:
		return NewSparsePCFWTenantTerraform()
	case TopQueryIdentity:
		return NewSparseTopQuery()
	case TotalQueryIdentity:
		return NewSparseTotalQuery()
	case TrendQueryIdentity:
		return NewSparseTrendQuery()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity elemental.Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) elemental.Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity elemental.Identity) elemental.Identifiables {

	switch identity {

	case AWSFirewallIdentity:
		return &AWSFirewallsList{}
	case AWSLogDefinitionIdentity:
		return &AWSLogDefinitionsList{}
	case FirewallAppIDListIdentity:
		return &FirewallAppIDListsList{}
	case FirewallApplicationGroupIdentity:
		return &FirewallApplicationGroupsList{}
	case FirewallCertificateIdentity:
		return &FirewallCertificatesList{}
	case FirewallCIDRlistIdentity:
		return &FirewallCIDRlistsList{}
	case FirewallCommitActionIdentity:
		return &FirewallCommitActionsList{}
	case FirewallCommittedRulesetIdentity:
		return &FirewallCommittedRulesetsList{}
	case FirewallCountryCodelistIdentity:
		return &FirewallCountryCodelistsList{}
	case FirewallCustomExternalFeedIdentity:
		return &FirewallCustomExternalFeedsList{}
	case FirewallCustomURLCategoryIdentity:
		return &FirewallCustomURLCategoriesList{}
	case FirewallExternalFeedListIdentity:
		return &FirewallExternalFeedListsList{}
	case FirewallFQDNListIdentity:
		return &FirewallFQDNListsList{}
	case FirewallGroupIdentity:
		return &FirewallGroupsList{}
	case FirewallLicensingCreditsIdentity:
		return &FirewallLicensingCreditsList{}
	case FirewallOptionIdentity:
		return &FirewallOptionsList{}
	case FirewallRollbackActionIdentity:
		return &FirewallRollbackActionsList{}
	case FirewallRuleIdentity:
		return &FirewallRulesList{}
	case FirewallRulesetIdentity:
		return &FirewallRulesetsList{}
	case FirewallSecurityProfileIdentity:
		return &FirewallSecurityProfilesList{}
	case FirewallTemplateIdentity:
		return &FirewallTemplatesList{}
	case FirewallURLCategoryListIdentity:
		return &FirewallURLCategoryListsList{}
	case LogQueryIdentity:
		return &LogQueriesList{}
	case MirrorSourceOptionIdentity:
		return &MirrorSourceOptionsList{}
	case PCFWAccountIdentity:
		return &PCFWAccountsList{}
	case PCFWAccountMirrorSourceIdentity:
		return &PCFWAccountMirrorSourcesList{}
	case PCFWAccountMirrorSourceStateIdentity:
		return &PCFWAccountMirrorSourceStatesList{}
	case PCFWFirewallConfigTerraformIdentity:
		return &PCFWFirewallConfigTerraformsList{}
	case PCFWSubnetHelperIdentity:
		return &PCFWSubnetHelpersList{}
	case PCFWTenantIdentity:
		return &PCFWTenantsList{}
	case PCFWTenantAccountTerraformIdentity:
		return &PCFWTenantAccountTerraformsList{}
	case PCFWTenantTerraformIdentity:
		return &PCFWTenantTerraformsList{}
	case TopQueryIdentity:
		return &TopQueriesList{}
	case TotalQueryIdentity:
		return &TotalQueriesList{}
	case TrendQueryIdentity:
		return &TrendQueriesList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity elemental.Identity) elemental.SparseIdentifiables {

	switch identity {

	case AWSFirewallIdentity:
		return &SparseAWSFirewallsList{}
	case AWSLogDefinitionIdentity:
		return &SparseAWSLogDefinitionsList{}
	case FirewallAppIDListIdentity:
		return &SparseFirewallAppIDListsList{}
	case FirewallApplicationGroupIdentity:
		return &SparseFirewallApplicationGroupsList{}
	case FirewallCertificateIdentity:
		return &SparseFirewallCertificatesList{}
	case FirewallCIDRlistIdentity:
		return &SparseFirewallCIDRlistsList{}
	case FirewallCommitActionIdentity:
		return &SparseFirewallCommitActionsList{}
	case FirewallCommittedRulesetIdentity:
		return &SparseFirewallCommittedRulesetsList{}
	case FirewallCountryCodelistIdentity:
		return &SparseFirewallCountryCodelistsList{}
	case FirewallCustomExternalFeedIdentity:
		return &SparseFirewallCustomExternalFeedsList{}
	case FirewallCustomURLCategoryIdentity:
		return &SparseFirewallCustomURLCategoriesList{}
	case FirewallExternalFeedListIdentity:
		return &SparseFirewallExternalFeedListsList{}
	case FirewallFQDNListIdentity:
		return &SparseFirewallFQDNListsList{}
	case FirewallGroupIdentity:
		return &SparseFirewallGroupsList{}
	case FirewallLicensingCreditsIdentity:
		return &SparseFirewallLicensingCreditsList{}
	case FirewallOptionIdentity:
		return &SparseFirewallOptionsList{}
	case FirewallRollbackActionIdentity:
		return &SparseFirewallRollbackActionsList{}
	case FirewallRuleIdentity:
		return &SparseFirewallRulesList{}
	case FirewallRulesetIdentity:
		return &SparseFirewallRulesetsList{}
	case FirewallSecurityProfileIdentity:
		return &SparseFirewallSecurityProfilesList{}
	case FirewallTemplateIdentity:
		return &SparseFirewallTemplatesList{}
	case FirewallURLCategoryListIdentity:
		return &SparseFirewallURLCategoryListsList{}
	case LogQueryIdentity:
		return &SparseLogQueriesList{}
	case MirrorSourceOptionIdentity:
		return &SparseMirrorSourceOptionsList{}
	case PCFWAccountIdentity:
		return &SparsePCFWAccountsList{}
	case PCFWAccountMirrorSourceIdentity:
		return &SparsePCFWAccountMirrorSourcesList{}
	case PCFWAccountMirrorSourceStateIdentity:
		return &SparsePCFWAccountMirrorSourceStatesList{}
	case PCFWFirewallConfigTerraformIdentity:
		return &SparsePCFWFirewallConfigTerraformsList{}
	case PCFWSubnetHelperIdentity:
		return &SparsePCFWSubnetHelpersList{}
	case PCFWTenantIdentity:
		return &SparsePCFWTenantsList{}
	case PCFWTenantAccountTerraformIdentity:
		return &SparsePCFWTenantAccountTerraformsList{}
	case PCFWTenantTerraformIdentity:
		return &SparsePCFWTenantTerraformsList{}
	case TopQueryIdentity:
		return &SparseTopQueriesList{}
	case TotalQueryIdentity:
		return &SparseTotalQueriesList{}
	case TrendQueryIdentity:
		return &SparseTrendQueriesList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) elemental.Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func (f modelManager) AllIdentities() []elemental.Identity {
	return AllIdentities()
}

var manager = modelManager{}

// Manager returns the model elemental.ModelManager.
func Manager() elemental.ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AWSFirewallIdentity,
		AWSLogDefinitionIdentity,
		FirewallAppIDListIdentity,
		FirewallApplicationGroupIdentity,
		FirewallCertificateIdentity,
		FirewallCIDRlistIdentity,
		FirewallCommitActionIdentity,
		FirewallCommittedRulesetIdentity,
		FirewallCountryCodelistIdentity,
		FirewallCustomExternalFeedIdentity,
		FirewallCustomURLCategoryIdentity,
		FirewallExternalFeedListIdentity,
		FirewallFQDNListIdentity,
		FirewallGroupIdentity,
		FirewallLicensingCreditsIdentity,
		FirewallOptionIdentity,
		FirewallRollbackActionIdentity,
		FirewallRuleIdentity,
		FirewallRulesetIdentity,
		FirewallSecurityProfileIdentity,
		FirewallTemplateIdentity,
		FirewallURLCategoryListIdentity,
		LogQueryIdentity,
		MirrorSourceOptionIdentity,
		PCFWAccountIdentity,
		PCFWAccountMirrorSourceIdentity,
		PCFWAccountMirrorSourceStateIdentity,
		PCFWFirewallConfigTerraformIdentity,
		PCFWSubnetHelperIdentity,
		PCFWTenantIdentity,
		PCFWTenantAccountTerraformIdentity,
		PCFWTenantTerraformIdentity,
		RootIdentity,
		TopQueryIdentity,
		TotalQueryIdentity,
		TrendQueryIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AWSFirewallIdentity:
		return []string{}
	case AWSLogDefinitionIdentity:
		return []string{}
	case FirewallAppIDListIdentity:
		return []string{}
	case FirewallApplicationGroupIdentity:
		return []string{}
	case FirewallCertificateIdentity:
		return []string{}
	case FirewallCIDRlistIdentity:
		return []string{}
	case FirewallCommitActionIdentity:
		return []string{}
	case FirewallCommittedRulesetIdentity:
		return []string{}
	case FirewallCountryCodelistIdentity:
		return []string{}
	case FirewallCustomExternalFeedIdentity:
		return []string{}
	case FirewallCustomURLCategoryIdentity:
		return []string{}
	case FirewallExternalFeedListIdentity:
		return []string{}
	case FirewallFQDNListIdentity:
		return []string{}
	case FirewallGroupIdentity:
		return []string{}
	case FirewallLicensingCreditsIdentity:
		return []string{}
	case FirewallOptionIdentity:
		return []string{}
	case FirewallRollbackActionIdentity:
		return []string{}
	case FirewallRuleIdentity:
		return []string{}
	case FirewallRulesetIdentity:
		return []string{}
	case FirewallSecurityProfileIdentity:
		return []string{}
	case FirewallTemplateIdentity:
		return []string{}
	case FirewallURLCategoryListIdentity:
		return []string{}
	case LogQueryIdentity:
		return []string{}
	case MirrorSourceOptionIdentity:
		return []string{}
	case PCFWAccountIdentity:
		return []string{}
	case PCFWAccountMirrorSourceIdentity:
		return []string{}
	case PCFWAccountMirrorSourceStateIdentity:
		return []string{}
	case PCFWFirewallConfigTerraformIdentity:
		return []string{}
	case PCFWSubnetHelperIdentity:
		return []string{}
	case PCFWTenantIdentity:
		return []string{}
	case PCFWTenantAccountTerraformIdentity:
		return []string{}
	case PCFWTenantTerraformIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case TopQueryIdentity:
		return []string{}
	case TotalQueryIdentity:
		return []string{}
	case TrendQueryIdentity:
		return []string{}
	}

	return nil
}
