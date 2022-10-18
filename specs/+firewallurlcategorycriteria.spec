# Model
model:
  rest_name: firewallurlcategorycriteria
  resource_name: firewallurlcategorycriterias
  entity_name: FirewallURLCategoryCriteria
  package: ngfw
  group: core/ngfw
  description: |-
    Represents rule URL category criteria.
    This is not an external object but a child object used by the Rule object.
  detached: true

# Attributes
attributes:
  v1:
  - name: URLCategories
    description: NGFW URL categories.  See FirewallURLCategoryList for the list.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - adult
    - high-risk
    - hacking

  - name: customExternalFeedIDs
    description: List FirewallCustomExternalFeed IDs that are of type URLList.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: customURLCategoryIDs
    description: List of FirewallCustomURLCategory IDs.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: externalFeedIDs
    description: List of URL intelligence feeds.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - PaloAltoNetworks-AuthenticationPortalExcluedList
