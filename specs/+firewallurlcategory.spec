# Model
model:
  rest_name: firewallurlcategory
  resource_name: firewallurlcategories
  entity_name: FirewallURLCategory
  package: ngfw
  group: core/ngfw
  description: Represents a URL category.
  detached: true

# Attributes
attributes:
  v1:
  - name: action
    description: The action the firewall should take.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Alert
    - Allow
    - Block
    default_value: Alert

  - name: name
    description: The name of the URL Catgory.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: abused-drugs
