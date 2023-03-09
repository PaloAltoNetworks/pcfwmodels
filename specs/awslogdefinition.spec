# Model
model:
  rest_name: awslogdefinition
  resource_name: awslogdefinitions
  entity_name: AWSLogDefinition
  package: ngfw
  group: core/ngfw
  description: Represents a Log Definition.
  get:
    description: Retrieves the logdefinition with the given ID.
  update:
    description: Updates the logdefinition with the given ID.
  delete:
    description: Deletes the logdefinition with the given ID.
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
  - name: decryptionEnabled
    description: The Decryption log is enabled.
    type: boolean
    exposed: true
    stored: true

  - name: logARN
    description: The ARN to access the log destination.
    type: string
    exposed: true
    stored: true
    example_value: arn:aws:iam::833962752675:role/SomeRole

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
