# Model
model:
  rest_name: firewallruleset
  resource_name: firewallrulesets
  entity_name: FirewallRuleset
  package: ngfw
  group: core/ngfw
  description: |-
    Represents a list of filewall rules that have a priority from 101+.
    This object can be created by the user and referenced by the AWSFirewall object.
  get:
    description: Retrieves the firewallruleset with the given ID.
  update:
    description: Updates the firewallruleset with the given ID.
  delete:
    description: Deletes the firewallruleset with the given ID.
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
  - name: committedRulesetID
    description: The ID of the last FirewallCommittedRuleset object.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: lastCommittedTime
    description: The date when the ruleset was last committed.
    type: time
    exposed: true
    stored: true
    read_only: true

# Relations
relations:
- rest_name: firewallrule
  get:
    description: Retrieves the list of firewall rules.
  create:
    description: Creates a firewall rule.
