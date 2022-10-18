# Model
model:
  rest_name: firewallprotoportscriteria
  resource_name: firewallprotoportscriterias
  entity_name: FirewallProtoPortsCriteria
  package: ngfw
  group: core/ngfw
  description: |-
    Represents rule criteria for specifying protocols and ports.
    This is not an external object but a child object used by the Rule object.
  detached: true

# Attributes
attributes:
  v1:
  - name: applicationDefaults
    description: Use NGFW defined application protocols and ports.
    type: boolean
    exposed: true
    stored: true

  - name: protoportList
    description: A protocol:port list.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - TCP:80,443,8080
    - UDP:53,67,68
    validations:
    - $optionalprotoports
