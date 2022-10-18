# Model
model:
  rest_name: firewallexternalfeed
  resource_name: firewallexternalfeeds
  entity_name: FirewallExternalFeed
  package: ngfw
  group: core/ngfw
  description: Represents a NGFW external feed.
  detached: true

# Attributes
attributes:
  v1:
  - name: feedtype
    description: The type of external feed.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - IPList
    - URLList
    default_value: IPList

  - name: name
    description: External feed name.
    type: string
    exposed: true
    stored: true
    required: true
    read_only: true
    example_value: name
