# Model
model:
  rest_name: firewallgroup
  resource_name: firewallgroups
  entity_name: FirewallGroup
  package: ngfw
  group: core/ngfw
  description: Represents common configuration for 1 or more AWSFirewall objects.
  get:
    description: Retrieves the firewallgroup with the given ID.
  update:
    description: Updates the firewallgroup with the given ID.
  delete:
    description: Deletes the firewallgroup with the given ID.
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
  - name: autoIncludeProfile
    description: A profile used to attach newly created firewalls to this FirewallGroup.
    type: ref
    exposed: true
    subtype: firewallautoincludeprofile
    stored: true
    extensions:
      refMode: pointer

  - name: firewallSecurityProfileID
    description: A FirewallSecurityProfile ID.
    type: string
    exposed: true
    stored: true

  - name: firewallTemplateID
    description: A FirewallTemplate ID.
    type: string
    exposed: true
    stored: true

  - name: logDefinitionID
    description: Log Definition ID.
    type: string
    exposed: true
    stored: true
