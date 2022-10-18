# Model
model:
  rest_name: firewallcommitaction
  resource_name: firewallcommitactions
  entity_name: FirewallCommitAction
  package: ngfw
  group: core/ngfw
  description: "Represents committing FilewallRulesets, FilewallTemplates,\nFilewallSecurityProfiles,
    \nand AWSLogDefinitions to a firewall."

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
    description: The firewalls effected by the commit action.
    type: refList
    exposed: true
    subtype: firewallstatus
    stored: true
    extensions:
      refMode: pointer

  - name: objectID
    description: The ID of the object being committed.
    type: string
    exposed: true
    required: true
    example_value: xxxxxxxxx

  - name: objectType
    description: The type of object being committed.
    type: enum
    exposed: true
    allowed_choices:
    - FirewallRuleset
    - FirewallTemplate
    - FirewallSecurityProfile
    - AWSLogDefinition
    default_value: FirewallRuleset
