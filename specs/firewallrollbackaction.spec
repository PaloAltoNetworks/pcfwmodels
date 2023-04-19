# Model
model:
  rest_name: firewallrollbackaction
  resource_name: firewallrollbackactions
  entity_name: FirewallRollbackAction
  package: ngfw
  group: core/ngfw
  description: Represents rolling back a committed Rulesets.

# Attributes
attributes:
  v1:
  - name: actionStatus
    description: The status of action.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Success
    - Failure
    example_value: Success

  - name: actionStatusReason
    description: The action status description.
    type: string
    exposed: true
    read_only: true

  - name: affectedFirewalls
    description: The firewalls affected by the rollback action.
    type: refList
    exposed: true
    subtype: firewallstatus
    stored: true
    extensions:
      refMode: pointer

  - name: committedFirewallRulesetID
    description: The ID of the committed firewall ruleset that we are rolling back
      to.
    type: string
    exposed: true
    required: true
    example_value: xxxxxxxxx
