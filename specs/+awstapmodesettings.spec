# Model
model:
  rest_name: awstapmodesettings
  resource_name: awstapmodesettings
  entity_name: AWSTAPModeSettings
  package: ngfw
  group: core/ngfw
  description: |-
    Represents settings specific to TAP mode.
    This object is a sub-object of AWSFirewall.
  detached: true

# Attributes
attributes:
  v1:
  - name: logDefinitionID
    description: An awslogdefinition ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxxxxxxx
