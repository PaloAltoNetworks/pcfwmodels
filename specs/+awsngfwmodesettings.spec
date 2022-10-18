# Model
model:
  rest_name: awsngfwmodesettings
  resource_name: awsngfwmodesettings
  entity_name: AWSNGFWModeSettings
  package: ngfw
  group: core/ngfw
  description: |-
    Represents settings specific to NGFW mode.
    This object is a sub-object of AWSFirewall.
  detached: true

# Attributes
attributes:
  v1:
  - name: firewallGroupID
    description: The firewallgroup ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxxxxxxx

  - name: firewallRulesetID
    description: The firewallruleset ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxxxxxxx
