# Model
model:
  rest_name: firewallfqdnlist
  resource_name: firewallfqdnlists
  entity_name: FirewallFQDNList
  package: ngfw
  group: core/ngfw
  description: Represents a list of FQDNs referenced by a rule match criteria object.
  get:
    description: Retrieves the FQDN list with the given ID.
  update:
    description: Updates the FQDN list with the given ID.
  delete:
    description: Deletes the FQDN list with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: FQDNs
    description: List of FQDNs.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - “www.google.com“
    - “redlock.atlassian.net“
    validations:
    - $fqdns
