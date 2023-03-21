# Model
model:
  rest_name: firewallcustomurlcategory
  resource_name: firewallcustomurlcategories
  entity_name: FirewallCustomURLCategory
  package: ngfw
  group: core/ngfw
  description: Represents a custom URL category that is referenced by a RuleURLCategory
    object.
  get:
    description: Retrieves the customurlcategory with the given ID.
  update:
    description: Updates the customurlcategory with the given ID.
  delete:
    description: Deletes the customurlcategory with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: URIs
    description: List of URIs.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - https://url1.com
    - https://url2.com
    validations:
    - $urls

  - name: action
    description: The action the firewall should take.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - None
    - Alert
    - Allow
    - Block
    default_value: None
