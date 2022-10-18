# Model
model:
  rest_name: vpcavailablesubnet
  resource_name: vpcavailablesubnets
  entity_name: VpcAvailableSubnet
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents association of a VPC ID to corresponding availability zones and one
    avaialble subnet every zone.
  detached: true

# Attributes
attributes:
  v1:
  - name: VPCID
    description: An AWS VPC ID.
    type: string
    exposed: true

  - name: availabilityZoneSubnets
    description: |-
      The list of all availability zones and associated subnets for every VPC
      specified.
    type: refList
    exposed: true
    subtype: availabilityzonesubnet
    extensions:
      refMode: pointer
