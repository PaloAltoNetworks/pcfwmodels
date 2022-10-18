# Model
model:
  rest_name: firewallapplicationcriteria
  resource_name: firewallapplicationcriterias
  entity_name: FirewallApplicationCriteria
  package: ngfw
  group: core/ngfw
  description: |-
    Represents rule criteria for specifying AppIds.
    This is not an external object but a child object used by the Rule object.
  detached: true

# Attributes
attributes:
  v1:
  - name: appIDs
    description: List of NGFW App IDs.  See AppIDList.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - ftp
    - http-proxy
    - ssh

  - name: applicationGroupIDs
    description: A list of FirewallApplicationGroup IDs.
    type: list
    exposed: true
    subtype: string
    stored: true
