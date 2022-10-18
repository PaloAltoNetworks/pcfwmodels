# Model
model:
  rest_name: firewallstatus
  resource_name: firewallstatuses
  entity_name: FirewallStatus
  package: ngfw
  group: core/ngfw
  description: Represents the status of a firewall.
  detached: true

# Attributes
attributes:
  v1:
  - name: firewallID
    description: The firewall ID.
    type: string
    exposed: true

  - name: firewallType
    description: The type of firewall.
    type: enum
    exposed: true
    allowed_choices:
    - AWS
    default_value: AWS

  - name: status
    description: The status of the of firewall.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Creating
    - Updating
    - Ready
    - Failed
    default_value: Creating

  - name: statusReason
    description: The status description of the firewall.
    type: string
    exposed: true
    read_only: true
