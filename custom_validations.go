package api

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"go.aporeto.io/elemental"
)

func makeErr(attribute string, message string) elemental.Error {

	err := elemental.NewError(
		"Validation Error",
		message,
		"pcfwbackend",
		http.StatusUnprocessableEntity,
	)

	if attribute != "" {
		err.Data = map[string]interface{}{"attribute": attribute}
	}

	return err
}

// ValidateWhyString validates the why field not empty
func ValidateWhyString(attribute string, why string) error {

	if why == "" {
		return makeErr(attribute, "why cannot be empty")
	}

	return nil
}

// Constants to validate tags.
const (
	prefixDynamicTag  = "$"
	prefixExpandedTag = "+"
	prefixMetadata    = "@"
)

// validateTagStrings validates the given string and check if it can be a valid value for a Tag.
func validateTagStrings(attribute string, acceptReservedPrefix bool, strs ...string) error {

	for _, s := range strs {

		if !acceptReservedPrefix && (strings.HasPrefix(s, prefixMetadata) || strings.HasPrefix(s, prefixDynamicTag) || strings.HasPrefix(s, prefixExpandedTag)) {
			return makeErr(attribute, fmt.Sprintf("%s starts with an @, a $ or a + that is reserved", s))
		}

		if err := ValidateTag(attribute, s); err != nil {
			return err
		}
	}

	return nil
}

// tagRegex is the regular expression to check the format of a tag.
var tagRegex = regexp.MustCompile(`^[^= ]+=.+`)

// ValidateTag validates a single tag.
func ValidateTag(attribute string, tag string) error {

	if !tagRegex.MatchString(tag) {
		return makeErr(attribute, fmt.Sprintf("'%s' must contain at least one '=' symbol separating two valid words", tag))
	}

	if len([]byte(tag)) >= 1024 {
		return makeErr(attribute, fmt.Sprintf("'%s' must be less than 1024 bytes", tag))
	}

	return nil
}

// ValidateTags validates a list of tags are valid. Accepts those with reserved prefix.
func ValidateTags(attribute string, tags []string) error {
	return validateTagStrings(attribute, true, tags...)
}

// ValidateTagsWithoutReservedPrefixes a list of tags are valid. Refuse those with reserved prefix.
func ValidateTagsWithoutReservedPrefixes(attribute string, tags []string) error {
	return validateTagStrings(attribute, false, tags...)
}

// ValidateCIDR validates a CIDR.
func ValidateCIDR(attribute string, network string) error {

	if _, _, err := net.ParseCIDR(network); err == nil {
		return nil
	}

	return makeErr(attribute, fmt.Sprintf("Attribute '%s' must be a CIDR", attribute))
}

// ValidateCIDRList validates a list of CIDRS.
// The list cannot be empty
func ValidateCIDRList(attribute string, networks []string) error {

	if len(networks) == 0 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	return ValidateOptionalCIDRList(attribute, networks)
}

// ValidateOptionalCIDRList validates a list of CIDRs.
// It can be empty.
func ValidateOptionalCIDRList(attribute string, networks []string) error {

	for _, network := range networks {
		if err := ValidateCIDR(attribute, network); err != nil {
			return err
		}
	}

	return nil
}

var rxDNSName = regexp.MustCompile(`^(\*\.){0,1}([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`)

// ValidateFQDN validates a FQDN.
func ValidateFQDN(attribute string, fqdn string) error {
	if !rxDNSName.MatchString(fqdn) {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must be a FQND", attribute))
	}
	return nil
}

// ValidateFQDNList validates a list of FQDNs.
// The list cannot be empty
func ValidateFQDNList(attribute string, fqdns []string) error {

	if len(fqdns) == 0 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	return ValidateOptionalFQDNList(attribute, fqdns)
}

// ValidateOptionalFQDNList validates a list of FQDNs.
// It can be empty.
func ValidateOptionalFQDNList(attribute string, fqdns []string) error {
	for _, fqdn := range fqdns {
		if err := ValidateFQDN(attribute, fqdn); err != nil {
			return err
		}
	}
	return nil
}

// ValidateURL validates a URL.
func ValidateURL(attribute string, address string) error {
	u, err := url.Parse(address)
	if err != nil {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must be a valid URL", attribute))
	}
	if u != nil && len(u.Host) == 0 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must have a host name", attribute))
	}
	return nil
}

// ValidateURLs validates a list of URLs.
// The list cannot be empty
func ValidateURLs(attribute string, addresses []string) error {
	if len(addresses) == 0 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}
	return ValidateOptionalURLs(attribute, addresses)
}

// ValidateOptionalURLs validates a list of URLs.
// It can be empty.
func ValidateOptionalURLs(attribute string, addresses []string) error {
	for _, address := range addresses {
		if err := ValidateURL(attribute, address); err != nil {
			return err
		}
	}
	return nil
}

// ValidateProtoPort validates a proto port.
func ValidateProtoPort(attribute string, protoport string) error {

	parts := strings.Split(protoport, ":")
	if len(parts) != 2 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid protoport", attribute))
	}

	proto := strings.ToUpper(parts[0])
	if proto != "TCP" && proto != "UDP" {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid protocol. Must be TCP or UDP", attribute))
	}

	ports := strings.Split(parts[1], ",")
	if len(ports) == 0 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  At least 1 port value must be specified", attribute))
	}

	for _, port := range ports {
		if val, err := strconv.Atoi(port); err != nil {
			return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Invalid port value %s", attribute, port))
		} else if val < 0 || val > 65535 {
			return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Invalid port value %s.  Must be between 0:65535", attribute, port))
		}
	}
	return nil
}

// ValidateOptionalProtoPorts validates a list of ProtoPorts.
// It can be empty.
func ValidateOptionalProtoPorts(attribute string, protoports []string) error {
	for _, protoport := range protoports {
		if err := ValidateProtoPort(attribute, protoport); err != nil {
			return err
		}
	}
	return nil
}

// ValidateMirrorRules validates traffic mirror filter rules.
func ValidateMirrorRules(attribute string, mirrorrules []*MirrorRule) error {
	// Walk through each rule and validate it
	for _, rule := range mirrorrules {
		for _, port := range []int{rule.DestinationFromPort, rule.DestinationToPort, rule.SourceFromPort, rule.SourceToPort} {
			if rule.Protocol == 6 || rule.Protocol == 17 {
				// For TCP/UDP ports should be within 0:65535 range
				if port < 0 || port > 65535 {
					return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Invalid port value %d.  Must be between 0:65535", attribute, port))
				}
			} else if port != -1 {
				// Ports shouldn't be specified if protocol is not TCP (6) or UDP (17)
				return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Port %d not valid for protocol %d.", attribute, port, rule.Protocol))
			}
		}
		if rule.DestinationFromPort > rule.DestinationToPort || rule.SourceFromPort > rule.SourceToPort {
			return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Start port is greater than end port in the range", attribute))
		}
		if (rule.DestinationFromPort == -1 && rule.DestinationToPort != -1) || (rule.DestinationFromPort != -1 && rule.DestinationToPort == -1) {
			return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Only one from destination start and end provided", attribute))
		}
		if (rule.SourceFromPort == -1 && rule.SourceToPort != -1) || (rule.SourceFromPort != -1 && rule.SourceToPort == -1) {
			return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Only one from source start and end provided", attribute))
		}
	}
	return nil
}

// ValidateAvailabilityZone validates the availability zone string.
// It can not be empty.
func ValidateAvailabilityZone(attribute string, availabilityzone string) error {

	if availabilityzone == "" {
		return makeErr(attribute, "availabilityzone cannot be empty")
	}
	// Validations against list of valid availability zones is going to fail
	// whenever a new zone is added, so avoid it for now.
	return nil
}

// ValidateAwsNetworkServices validates the networkServices list of strings that are
// used in mirror filter to filter aws services.
func ValidateAwsNetworkServices(attribute string, networkServices []string) error {

	// Empty list of network services is valid
	if len(networkServices) == 0 {
		return nil
	}
	// If not empty, then compare against the supported services.
	// Currently only "amazon-dns" is supported.
	supportedServices := []string{"amazon-dns"}
	if !reflect.DeepEqual(networkServices, supportedServices) {
		return makeErr(attribute, "Only amazon-dns is supported as network services in mirror filter")
	}
	return nil
}

// ValidateVPCID validates the VPC ID string.
// It should always start with "vpc-" prefix
func ValidateVPCID(attribute string, vpcid string) error {

	if !strings.HasPrefix(vpcid, "vpc-") {
		return makeErr(attribute, fmt.Sprintf("vpcid %s does not start as 'vpc-'", vpcid))
	}
	return nil
}

// ValidateVPCIDs validates a list of VPC IDs.
func ValidateVPCIDs(attribute string, vpcids []string) error {

	for _, vpcid := range vpcids {
		if err := ValidateVPCID(attribute, vpcid); err != nil {
			return err
		}
	}
	return nil
}

// ValidateVpcSubnetInfo validates all the information in VPCUsedSubnets list
func ValidateVpcSubnetInfo(attribute string, VpcUsedSubnets []*VpcUsedSubnet) error {

	// Empty list is invalid
	if len(VpcUsedSubnets) == 0 {
		return makeErr(attribute, "VpcUsedSubnets list cannot be empty")
	}
	// Walk through each entry and validate it
	for _, vpc := range VpcUsedSubnets {
		if len(vpc.AvailabilityZones) == 0 {
			return makeErr(attribute, fmt.Sprintf("AvailabilityZones list for VPC %s cannot be empty", vpc.VPCID))
		}
	}
	return nil
}

// athenaWorkGroupRegex is the regular expression to check the format of the athenaWorkGroup.
var athenaWorkGroupRegex = regexp.MustCompile(`[0-9a-zA-Z_@\-]`)

// ValidateAthenaWorkGroup validates an athena workgroup
func ValidateAthenaWorkGroup(attribute string, athenaWorkGroup string) error {

	if !athenaWorkGroupRegex.MatchString(athenaWorkGroup) {
		return makeErr(attribute, fmt.Sprintf("'%s' must only contain a-z, A-Z, 0-9, _(underscore), @(at sign) and -(hyphen)", athenaWorkGroup))
	}

	if len(athenaWorkGroup) > 127 {
		return makeErr(attribute, fmt.Sprintf("'%s' must be less than or equal to 127 characters", athenaWorkGroup))
	}

	return nil
}

// logResourcePrefixRegex is the regular expression to check the format of the logResourcePrefix.
var logResourcePrefixRegex = regexp.MustCompile(`[a-z]`)

// ValidateLogResourcePrefix validates a log resource prefix
func ValidateLogResourcePrefix(attribute string, logResourcePrefix string) error {

	if !logResourcePrefixRegex.MatchString(logResourcePrefix) {
		return makeErr(attribute, fmt.Sprintf("'%s' must only contain a-z", logResourcePrefix))
	}

	if len(logResourcePrefix) > 8 {
		return makeErr(attribute, fmt.Sprintf("'%s' must be less than or equal to 8 characters", logResourcePrefix))
	}

	return nil
}

// roleARNRegex is the regular expression to check the format of a roleARN.
var roleARNRegex = regexp.MustCompile(`arn:aws:iam::[0-9]{12}:role\/.+`)

// ValidateRoleARN validates a role ARN
func ValidateRoleARN(attribute string, roleARN string) error {

	if !roleARNRegex.MatchString(roleARN) {
		return makeErr(attribute, fmt.Sprintf("'%s' is an invalid role ARN", roleARN))
	}
	return nil
}

// AWSAccountIDRegex is the regular expression to check the format of an AWS Account ID.
var AWSAccountIDRegex = regexp.MustCompile(`^[0-9]{12}$`)

// ValidateAWSAccountID validates an AWS account ID
func ValidateAWSAccountID(attribute string, AWSAccountID string) error {

	if !AWSAccountIDRegex.MatchString(AWSAccountID) {
		return makeErr(attribute, fmt.Sprintf("'%s' is an invalid AWS Account ID", AWSAccountID))
	}
	return nil
}
