# Model
model:
  rest_name: firewallcountrycode
  resource_name: firewallcountrycodes
  entity_name: FirewallCountryCode
  package: ngfw
  group: core/ngfw
  description: A Country code.
  detached: true

# Attributes
attributes:
  v1:
  - name: description
    description: The Country description.
    type: string
    exposed: true
    read_only: true
    example_value: Canada

  - name: name
    description: The country code name.
    type: string
    exposed: true
    read_only: true
    example_value: CA
