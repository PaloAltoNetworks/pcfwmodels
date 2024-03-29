# Model
model:
  rest_name: firewallcidrlist
  resource_name: firewallcidrlists
  entity_name: FirewallCIDRlist
  package: ngfw
  group: core/ngfw
  description: Represents a list of CIDRs referenced by a target criteria object.
  get:
    description: Retrieves the firewall CIDR list with the given ID.
  update:
    description: Updates the firewall CIDR list with the given ID.
  delete:
    description: Deletes the firewall CIDR list with the given ID.
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
  - name: CIDRs
    description: List of CIDRs.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - “10.1.0.0/16“
    - “10.2.0.0/16“
    validations:
    - $cidrs
