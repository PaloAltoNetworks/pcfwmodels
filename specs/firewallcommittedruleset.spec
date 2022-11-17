# Model
model:
  rest_name: firewallcommittedruleset
  resource_name: firewallcommittedrulesets
  entity_name: FirewallCommittedRuleset
  package: ngfw
  group: core/ngfw
  description: Represents a list of committed filewall rules.
  get:
    description: Retrieves the firewallcommittedruleset with the given ID.
  update:
    description: Updates the firewallcommittedruleset with the given ID.
  delete:
    description: Deletes the firewallcommittedruleset with the given ID.
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: committedTime
    description: The date when the ruleset or template was committed.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: objectID
    description: The ID of the ruleset or template that was committed.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: objectType
    description: The type of object committed.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - FirewallRuleset
    - FirewallTemplate
    default_value: FirewallRuleset

# Relations
relations:
- rest_name: firewallrule
  get:
    description: Retrieves the list of firewall rules.
  create:
    description: Creates a firewall rule.
