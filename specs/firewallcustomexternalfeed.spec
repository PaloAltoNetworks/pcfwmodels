# Model
model:
  rest_name: firewallcustomexternalfeed
  resource_name: firewallcustomexternalfeeds
  entity_name: FirewallCustomExternalFeed
  package: ngfw
  group: core/ngfw
  description: Represents an external intelligence feed referenced by TargetCriteria.
  get:
    description: Retrieves the externalfeed with the given ID.
  update:
    description: Updates the externalfeed with the given ID.
  delete:
    description: Deletes the externalfeed with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@tags'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: certificateProfile
    description: TODO figure out what this is.
    type: string
    exposed: true
    stored: true

  - name: feedtype
    description: The type of external feed.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - IPList
    - URLList
    default_value: IPList

  - name: sourceURL
    description: The URL of the external feed.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://saasedl.paloaltonetworks.com/feeds/m365/china/any/all/ipv4
    validations:
    - $url

  - name: updateDailyTime
    description: Specifies the time of day when the updateFrequency is Daily.
    type: time
    exposed: true
    stored: true

  - name: updateFrequency
    description: How often is the external feed updated.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Daily
    - Hourly
    default_value: Daily
