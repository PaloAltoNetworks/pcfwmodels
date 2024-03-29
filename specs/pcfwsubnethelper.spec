# Model
model:
  rest_name: pcfwsubnethelper
  resource_name: pcfwsubnethelpers
  entity_name: PCFWSubnetHelper
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents a PCFW service endpoint that returns available subnet for every
    specified availability zones in every specified VPC.

# Attributes
attributes:
  v1:
  - name: VPCAvailableSubnets
    description: |-
      Returns the list of AWS VPC IDs and information about available subnets for
      every availability zones in a VPC.
    type: refList
    exposed: true
    subtype: vpcavailablesubnet
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: VPCUsedSubnets
    description: |-
      List of AWS VPC IDs with information about associated availability zones and
      used subnets to check for available subnets.
    type: refList
    exposed: true
    subtype: vpcusedsubnet
    stored: true
    validations:
    - $vpcsubnetinfo
    extensions:
      refMode: pointer
