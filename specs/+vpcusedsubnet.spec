# Model
model:
  rest_name: vpcusedsubnet
  resource_name: vpcusedsubnets
  entity_name: VpcUsedSubnet
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents association of a VPC ID to corresponding availability zones and used
    subnet CIDR blocks in the VPC.
  detached: true

# Attributes
attributes:
  v1:
  - name: VPCID
    description: An AWS VPC ID.
    type: string
    exposed: true
    stored: true
    validations:
    - $vpcid

  - name: availabilityZones
    description: The list of all availability zones associated with the VPC.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: subnetCIDRs
    description: The list of all currently used subnet CIDR blocks in this VPC.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $cidrs
