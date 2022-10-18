# Model
model:
  rest_name: availabilityzonesubnet
  resource_name: availabilityzonesubnets
  entity_name: AvailabilityZoneSubnet
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents association of an availability zone and an avaialble subnet in that
    zone.
  detached: true

# Attributes
attributes:
  v1:
  - name: availabilityZone
    description: The availability zone.
    type: string
    exposed: true
    validations:
    - $availabilityzone

  - name: subnetCIDR
    description: The AWS subnet CIDR block in this availability zone.
    type: string
    exposed: true
    validations:
    - $cidr
