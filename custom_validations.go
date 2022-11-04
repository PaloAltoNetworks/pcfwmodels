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

// ValidatePort validates a port number.
func ValidatePort(attribute string, port int) error {
	if port < 0 || port > 65535 {
		return makeErr(attribute, fmt.Sprintf("Attribute '%s' is an invalid.  Invalid port value %d.  Must be between 0:65535", attribute, port))
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

// ValidateEnis validates the list containing eni strings.
// List can not be empty and each entry must start with "eni-" prefix
func ValidateEnis(attribute string, enis []string) error {

	if len(enis) == 0 {
		return makeErr(attribute, "eni list cannot be empty")
	}
	for _, eni := range enis {
		if !strings.HasPrefix(eni, "eni-") {
			return makeErr(attribute, fmt.Sprintf("eni %s does not start as 'eni-'", eni))
		}
	}
	return nil
}

// ValidateVpcInfo validates all the vpc information in VPCAvailabilityZoneSubnet list
func ValidateVpcInfo(attribute string, VPCAvailabilityZoneSubnets []*VPCAvailabilityZoneSubnet) error {

	// Empty list is invalid
	if len(VPCAvailabilityZoneSubnets) == 0 {
		return makeErr(attribute, "VPCAvailabilityZoneSubnet list cannot be empty")
	}
	// Walk through each entry and validate it
	for _, vpc := range VPCAvailabilityZoneSubnets {
		if len(vpc.SubnetInterfaces) == 0 {
			return makeErr(attribute, fmt.Sprintf("AvailabilityZoneSubnetInterface list for VPC %s cannot be empty", vpc.VPCID))
		}
	}
	return nil
}
