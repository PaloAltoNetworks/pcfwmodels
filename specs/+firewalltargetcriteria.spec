# Model
model:
  rest_name: firewalltargetcriteria
  resource_name: firewalltargetcriterias
  entity_name: FirewallTargetCriteria
  package: ngfw
  group: core/ngfw
  description: |-
    Represents Source or Destination rule matching criteria.
    This is not an external object but a child object used by the Rule object.
  detached: true

# Attributes
attributes:
  v1:
  - name: CIDRs
    description: List of CIDRs.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - “10.1.0.0/16“
    - “10.2.0.0/16“
    validations:
    - $optionalcidrs

  - name: FQDNs
    description: List of FQDNs.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - “www.google.com“
    - “redlock.atlassian.net“
    validations:
    - $optionalfqdns

  - name: countryCodes
    description: List of country codes.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - CA
    - “US“

  - name: customExternalFeedIDs
    description: List of firewall custom external feed IDs that are of type IPList.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: externalFeedIDs
    description: List of firewall external feed IDs that are of type IPList.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - PaloAltoNetworks-BulletproofIPAddresses
    - PaloAltoNetworks-HighRiskIPAddresses
    - PaloAltoNetworks-KnownMaliciousIPAddresses

  - name: firewallCIDRListIDs
    description: List of Firewall CIDR list IDs.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: firewallFQDNListIDs
    description: List of FQDN list IDs.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: negate
    description: Negates this match criteria.
    type: boolean
    exposed: true
    stored: true
