# Model
model:
  rest_name: vpcavailabilityzonesubnet
  resource_name: vpcvailabilityzonesubnets
  entity_name: VPCAvailabilityZoneSubnet
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: Represents association of a VPC ID to corresponding subnets, AZs and
    interfaces.
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

  - name: subnetInterfaces
    description: |-
      The list of all firewall subnets (one in each AZ) and associated network
      interfaces in that AZ.
    type: refList
    exposed: true
    subtype: availabilityzonesubnetinterface
    stored: true
    extensions:
      refMode: pointer
