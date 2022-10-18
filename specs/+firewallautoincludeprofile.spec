# Model
model:
  rest_name: firewallautoincludeprofile
  resource_name: firewallautoincludeprofiles
  entity_name: FirewallAutoIncludeProfile
  package: ngfw
  group: core/ngfw
  description: |-
    Represents match criteria for selecting which AWSInstance objects belong to a
    Group.
  detached: true

# Attributes
attributes:
  v1:
  - name: VPC
    description: An AWS VPC ID.
    type: string
    exposed: true
    stored: true

  - name: tags
    description: A tag list.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - name=value
    - app=prod
    validations:
    - $tagsWithoutReservedPrefixes
