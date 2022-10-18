# Model
model:
  rest_name: mirrorrule
  resource_name: mirrorrules
  entity_name: MirrorRule
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: Represents an AWS mirroring filter rule.
  detached: true

# Attributes
attributes:
  v1:
  - name: action
    description: The action to take on the filtered traffic.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Accept
    - Reject
    default_value: Accept

  - name: description
    description: Description of the mirror rule.
    type: string
    exposed: true
    stored: true

  - name: destinationCIDR
    description: Destination CIDR block to assign to the mirror rule.
    type: string
    exposed: true
    stored: true
    default_value: 0.0.0.0/0
    validations:
    - $cidr

  - name: destinationPortRange
    description: Destination port range.
    type: string
    exposed: true
    stored: true
    validations:
    - $portrange

  - name: direction
    description: The direction of the traffic to be mirrored.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Ingress
    - Egress
    default_value: Ingress

  - name: number
    description: Number of a traffic mirror rule. Must be unique in each direction.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 100

  - name: protocol
    description: Protocol number to assign to the mirror rule.
    type: integer
    exposed: true
    stored: true
    max_value: 255

  - name: sourceCIDR
    description: Source CIDR block to assign to the mirror rule.
    type: string
    exposed: true
    stored: true
    default_value: 0.0.0.0/0
    validations:
    - $cidr

  - name: sourcePortRange
    description: Source port range.
    type: string
    exposed: true
    stored: true
    validations:
    - $portrange
