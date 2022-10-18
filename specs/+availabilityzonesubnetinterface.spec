# Model
model:
  rest_name: availabilityzonesubnetinterface
  resource_name: availabilityzonesubnetinterfaces
  entity_name: AvailabilityZoneSubnetInterface
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents an association of a subnet in an availability zone where a firewall
    endpoint needs to be placed, and network interfaces in that availability zone
    from where traffic needs to be mirrored.
  detached: true

# Attributes
attributes:
  v1:
  - name: availabilityZone
    description: The availability zone in the VPC.
    type: string
    exposed: true
    stored: true
    validations:
    - $availabilityzone

  - name: sourceNetworkInterfaces
    description: |-
      Source network interfaces in this availability zone from where traffic is to be
      mirrored.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: subnetCIDR
    description: The AWS subnet CIDR block in this availability zone.
    type: string
    exposed: true
    stored: true
    validations:
    - $cidr
