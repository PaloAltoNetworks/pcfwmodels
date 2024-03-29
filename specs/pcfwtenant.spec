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
  update:
    description: Unsubscribe the tenant.
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

  - name: PCTenantIssuer
    description: |-
      Issuer for Prisma Cloud tenant (identifies the Prisma Cloud stack for IAM role
      usage).
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: athenaWorkgroup
    description: Athena workgroup to execute queries in.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: logs_workgroup
    max_length: 118
    validations:
    - $athenaworkgroup

  - name: decryptionRoleARN
    description: AWS rulestack decryption role ARN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::640574671726:role/DecryptionRole
    validations:
    - $rolearn

  - name: endpointRoleARN
    description: AWS endpoint role ARN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::640574671726:role/CustomerManagedEndpoint
    validations:
    - $rolearn

  - name: logDestination
    description: |-
      The log destination for logging. The value will be converted to lower case and
      have the AWS account ID and logging region appended to it (ie.
      logs-1234-us-east-2) when used to create an S3 bucket for logging.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: logs
    max_length: 30
    validations:
    - $logdestination

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

  - name: logPushRoleARN
    description: ARN of AWS role that allows the NGFW to push logs.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::1234567890:role/NGFWLogPushRole
    validations:
    - $rolearn

  - name: logQueryRoleARN
    description: ARN of AWS role that allows the PCFW to query logs.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: arn:aws:iam::1234567890:role/PCFWLogQueryRole
    validations:
    - $rolearn

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
    max_length: 8
    validations:
    - $logresourceprefix

  - name: offboardingTimestamp
    description: The timestamp when offboarding pending started.
    type: time
    stored: true

  - name: primaryAWSAccountID
    description: The primary AWS Account ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 640574671726
    validations:
    - $awsaccount

  - name: primaryAccountNamespace
    description: The namespace where the primary PCFW account will be created.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /tenant/account

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
