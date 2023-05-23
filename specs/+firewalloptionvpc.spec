# Model
model:
  rest_name: firewalloptionvpc
  resource_name: firewalloptionvpcs
  entity_name: FirewallOptionVPC
  package: discovery
  group: core/discovery
  description: Includes information about a VPC for use in firewall configuration.
  extends:
  - '@named'
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: An AWS VPC ID.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    validations:
    - $vpcid

  - name: availabilityZones
    description: The availability zones associated with the VPC.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true
