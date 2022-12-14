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
  - '@named'
  - '@described'
  - '@namespaced'
  - '@tags'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: AWSNGFWModeSettings
    description: Settings specific to NGFW mode.
    type: ref
    exposed: true
    subtype: awsngfwmodesettings
    stored: true
    extensions:
      noInit: true
      refMode: pointer

  - name: AWSTAPModeSettings
    description: Settings specific to TAP mode.
    type: ref
    exposed: true
    subtype: awstapmodesettings
    stored: true
    extensions:
      noInit: true
      refMode: pointer

  - name: NGFWFirewall
    description: The name of the NGFW fireall.
    type: string
    stored: true
    read_only: true

  - name: NGFWRuleStack
    description: The name of the NGFW rulestack associated with the firewall.
    type: string
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

  - name: VPCID
    description: An AWS VPC ID.
    type: string
    required: true
    exposed: true
    stored: true
    example_value: vpc-23af3b89cd23
    validations:
    - $vpcid

  - name: availabilityZones
    description: A list of availability zones.
    required: true
    type: list
    exposed: true
    subtype: string
    example_value: us-east-1a
    stored: true

  - name: endpointServiceName
    description: The endpoint service name needed to create an AWS endpoint.
    type: string
    exposed: true
    stored: true
    read_only: true
