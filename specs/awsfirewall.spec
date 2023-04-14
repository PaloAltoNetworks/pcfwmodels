# Model
model:
  rest_name: awsfirewall
  resource_name: awsfirewalls
  entity_name: AWSFirewall
  package: ngfw
  group: core/ngfw
  description: Represents an AWS firewall instance.
  get:
    description: Retrieves the awsfirewall with the given ID.
  update:
    description: Updates the awsfirewall with the given ID.
  delete:
    description: Deletes the awsfirewall with the given ID.
    parameters:
      entries:
      - name: purge
        description: Parameter to delete the awsfireall record without cleaning up
          the NGFW resources.
        type: boolean
        default_value: false
  extends:
  - '@identifiable-stored'
  - '@namedstrict'
  - '@described'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: NGFWExternalID
    description: The internal NGFW externalID for making API calls.
    type: string
    stored: true
    read_only: true

  - name: NGFWFirewall
    description: The internal name of the NGFW firewall.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: NGFWRuleStack
    description: The internal name of the NGFW rulestack associated with the firewall.
    type: string
    stored: true
    read_only: true

  - name: VPCIDs
    description: The list of VPC IDs.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value: vpc-23af3b89cd23, vpc-23af3b89cd24
    validations:
    - $vpcids

  - name: availabilityZones
    description: A list of availability zones.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value: us-east-1a

  - name: endpointServiceName
    description: The endpoint service name needed to create an AWS endpoint.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: endpoints
    description: The firewall endpoints.
    type: refList
    exposed: true
    subtype: awsendpoint
    stored: true
    read_only: true
    extensions:
      refMode: pointer

  - name: lastCommitTime
    description: The date when the fireall was last committed.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: licenseType
    description: The license type of the firewall.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - TAP
    - Basic
    - Advanced
    autogenerated: true
    default_value: TAP

  - name: mode
    description: The mode of the of firewall.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - TAP
    - NGFW
    default_value: TAP

  - name: region
    description: The AWS region of this Firewall.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: us-east-1

  - name: retryCount
    description: The number of times a Create/Update/Delete has been retried.
    type: integer
    stored: true

  - name: status
    description: The status of the of firewall.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Creating
    - Updating
    - Deleting
    - Ready
    - Failed
    - LogConfigFailed
    autogenerated: true
    default_value: Creating

  - name: statusReason
    description: The status description of the firewall.
    type: string
    exposed: true
    stored: true
    read_only: true
