# Model
model:
  rest_name: firewalltemplate
  resource_name: firewalltemplates
  entity_name: FirewallTemplate
  package: ngfw
  group: core/ngfw
  description: |-
    Represents a list of filewall rules that have a priority from 1-100.
    This object can be created by the user and referenced by the Group object.
  get:
    description: Retrieves the firewalltemplate with the given ID.
  update:
    description: Updates the firewalltemplate with the given ID.
  delete:
    description: Deletes the firewalltemplate with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@tags'
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
    description: The date when the template was last committed.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: lastUpdatedTime
    description: The date when the template was last updated.
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
