# Model
model:
  rest_name: pcfwlambdahealth
  resource_name: pcfwlambdahealths
  entity_name: PCFWLambdaHealth
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: Represents the health state of a deployed Lambda function.
  get:
    description: Retrieves the PCFW Lambda health object with the given ID.
  update:
    description: Updates the PCFW Lambda health object with the given ID.
  delete:
    description: Deletes the PCFW Lambda health object with the given ID.
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: lambdaARN
    description: The ARN of the lambda function resource.
    type: string
    exposed: true
    stored: true

  - name: lambdaName
    description: The name of the lambda function.
    type: string
    exposed: true
    stored: true

  - name: region
    description: The AWS region where the lambda function is running.
    type: string
    exposed: true
    stored: true

  - name: status
    description: The status of the lambda function.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Success
    - Failed
    - InProgress

  - name: statusMessage
    description: The status description of the lambda function.
    type: string
    exposed: true
    stored: true
