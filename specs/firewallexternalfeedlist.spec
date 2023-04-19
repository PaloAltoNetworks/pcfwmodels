# Model
model:
  rest_name: firewallexternalfeedlist
  resource_name: firewallexternalfeedlists
  entity_name: FirewallExternalFeedList
  package: ngfw
  group: core/ngfw
  description: This a read-only list that returns all of the NGFW external feeds.
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: feeds
    description: A list of external feed objects.
    type: refList
    exposed: true
    subtype: firewallexternalfeed
    stored: true
    read_only: true
    extensions:
      refMode: pointer
