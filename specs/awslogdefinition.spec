# Model
model:
  rest_name: awslogdefinition
  resource_name: awslogdefinitions
  entity_name: AWSLogDefinition
  package: ngfw
  group: core/ngfw
  description: Represents a Log Definition.
  get:
    description: Retrieves the AWS log definition with the given ID.
  update:
    description: Updates the AWS log definition with the given ID.
  delete:
    description: Deletes the AWS log definition with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: decryptionEnabled
    description: The Decryption log is enabled.
    type: boolean
    exposed: true
    stored: true

  - name: logDestination
    description: Destination for log output.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: name

  - name: logDestinationType
    description: Destination type for log output.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Prisma
    - S3
    - Cloudwatch
    - KinesisFirehose
    example_value: Prisma

  - name: logPushRoleARN
    description: The ARN for NGFW to write to the destination.
    type: string
    exposed: true
    stored: true
    example_value: arn:aws:iam::833962752675:role/SomeRole

  - name: logQueryRoleARN
    description: The ARN for PCFW to query the destination.
    type: string
    exposed: true
    stored: true
    example_value: arn:aws:iam::833962752675:role/SomeRole

  - name: logRegion
    description: The AWS region where logging data lives.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: us-east-1

  - name: logResourcePrefix
    description: Prefix to use for logging resources.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: pcfw

  - name: threatEnabled
    description: The Threat log is enabled.
    type: boolean
    exposed: true
    stored: true

  - name: trafficEnabled
    description: The Traffic log is enabled.
    type: boolean
    exposed: true
    stored: true
