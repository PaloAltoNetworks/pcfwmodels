# Model
model:
  rest_name: pcfwtenant
  resource_name: pcfwtenants
  entity_name: PCFWTenant
  package: ngfw
  group: onboarding
  description: Represents Prisma Cloud Firewall Tenant.
  get:
    description: Retrieve the tenant with the given PrismaID.
  delete:
    description: Delete the tenant with the given PrismaID.
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: AWSAccountID
    description: AWS Account ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 640574671726

  - name: NGFWExternalID
    description: NGFW external ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: NGFWTenantID
    description: NGFW tenant ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: decryptionRoleARN
    description: AWS rulestack decryption role ARN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::640574671726:role/DecryptionRole

  - name: endpointRoleARN
    description: AWS endpoint role ARN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::640574671726:role/CustomerManagedEndpoint

  - name: logDestination
    description: The log destination for logging.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: s3://640574671726-us-east-1-api-auth-bucket/Logging

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
    default_value: Prisma

  - name: logRegion
    description: The AWS region where logging data lives.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: us-east-1

  - name: loggingRoleARN
    description: AWS logging role ARN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::640574671726:role/LogMetricRole

  - name: offboardingTimestamp
    description: The timestamp when offboarding pending started.
    type: time
    stored: true

  - name: status
    description: status of tenant.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Creating
    - Success
    - Failed
    - Unsubscribe
    - UnsubscribePending
    - Deleting
    - FailedCreatingPrimaryAccount
    autogenerated: true

  - name: statusReason
    description: status failure reason.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
