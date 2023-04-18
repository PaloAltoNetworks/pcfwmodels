# Model
model:
  rest_name: mirrorsourceoption
  resource_name: mirrorsourceoptions
  entity_name: MirrorSourceOption
  package: discovery
  group: core/discovery
  description: Discovers eligible traffic mirror source instances and auto-scaling
    groups.

# Attributes
attributes:
  v1:
  - name: autoScalingGroupNames
    description: List of discovered auto-scaling groups.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true

  - name: firewallID
    description: The ID of the firewall whose VPCs/AZs should be used to search for
      instances.
    type: string
    exposed: true
    required: true
    example_value: 642ab7410fa464137cabba1a

  - name: instances
    description: List of discovered mirror source instances.
    type: refList
    exposed: true
    subtype: mirrorinstance
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer
