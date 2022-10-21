# Model
model:
  rest_name: awsattachment
  resource_name: awsattachments
  entity_name: AWSAttachment
  package: ngfw
  group: core/ngfw
  description: |-
    Represents an AWS VPC ID and subnets where the NGFW attaches.
    This object is a sub-object of AWSFirewall.
  detached: true

# Attributes
attributes:
  v1:
  - name: VPC
    description: AWS VPC ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: vpc-23af3b89cd23

  - name: endpoints
    description: A list of awsendpoint objects.
    type: refList
    exposed: true
    subtype: awsendpoint
    stored: true
    required: true
    example_value: awsendpoint list
    extensions:
      refMode: pointer

  - name: region
    description: The AWS region of this VPC.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: us-east-1
