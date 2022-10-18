# Model
model:
  rest_name: firewallcountrycodelist
  resource_name: firewallcountrycodelist
  entity_name: FirewallCountryCodelist
  package: ngfw
  group: core/ngfw
  description: A list of NGFW Country codes.
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: countrycodes
    description: A list of NGFW Country codes.
    type: refList
    exposed: true
    subtype: firewallcountrycode
    stored: true
    read_only: true
    extensions:
      refMode: pointer
