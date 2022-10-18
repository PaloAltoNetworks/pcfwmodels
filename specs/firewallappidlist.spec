# Model
model:
  rest_name: firewallappidlist
  resource_name: firewallappidlist
  entity_name: FirewallAppIDList
  package: ngfw
  group: core/ngfw
  description: |-
    This a readonly list that returns all of the NGFW AppIDs that the user can
    select.
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: appIDs
    description: A list of NGFW AppIDs.
    type: list
    exposed: true
    subtype: string
    read_only: true
