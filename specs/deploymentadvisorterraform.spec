# Model
model:
  rest_name: deploymentadvisorterraform
  resource_name: deploymentadvisorterraforms
  entity_name: DeploymentAdvisorTerraform
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: Represents a Cloud NGFW Deployment Advisor core service.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: AMIIDs
    description: |-
      The list of all AMI IDs where dynamic updates are to be performed on associated
      instances.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: NGFWMode
    description: The mode of the Cloud NGFW instance.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - TAP
    - NGFW
    default_value: TAP

  - name: VPCAvailabilityZoneSubnets
    description: |-
      List of AWS VPC IDs and information about associated firewall subnets in their
      respective availability zones.
    type: refList
    exposed: true
    subtype: vpcavailabilityzonesubnet
    stored: true
    extensions:
      refMode: pointer

  - name: dynamicPolicyUpdateEnabled
    description: |-
      Whether or not perform dynamic updates in customer VPC e.g. for creating
      mirroring policies on new instances.
    type: boolean
    exposed: true
    stored: true

  - name: filter
    description: Traffic mirror filter to specify what traffic is to be mirrored.
    type: ref
    exposed: true
    subtype: mirrorfilter
    stored: true
    extensions:
      refMode: pointer

  - name: instanceIDs
    description: The list of all instance IDs where dynamic updates are to be performed.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: instanceTags
    description: The list of all instance tags where dynamic updates are to be performed.
    type: list
    exposed: true
    subtype: string
    stored: true
