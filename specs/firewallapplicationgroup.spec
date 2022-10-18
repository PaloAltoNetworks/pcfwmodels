# Model
model:
  rest_name: firewallapplicationgroup
  resource_name: firewallapplicationgroups
  entity_name: FirewallApplicationGroup
  package: ngfw
  group: core/ngfw
  description: Represents a grouping of application IDs.
  get:
    description: Retrieves the firewallapplicationgroup with the given ID.
  update:
    description: Updates the firewallapplicationgroup with the given ID.
  delete:
    description: Deletes the firewallapplicationgroup with the given ID.
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
  - name: appIDs
    description: List of NGFW AppIDs.  See FirewallAppIDList.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - “ssh“
    - “ftp“
