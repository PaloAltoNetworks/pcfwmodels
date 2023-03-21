# Model
model:
  rest_name: pcfwaccountmirrorsourcestate
  resource_name: pcfwaccountmirrorsourcestates
  entity_name: PCFWAccountMirrorSourceState
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: |-
    Represents state of traffic mirror session setup associated with an instance and
    network interface for a firewall associatd with a customer account.
  get:
    description: Retrieves the pcfwaccountmirrorsourcestate object with the given
      ID.
  update:
    description: Updates the pcfwaccountmirrorsourcestate object with the given ID.
  delete:
    description: Deletes the pcfwaccountmirrorsourcestate object with the given ID.
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
  - name: instanceID
    description: The instance ID that is the source of traffic mirroring.
    type: string
    exposed: true
    stored: true

  - name: networkInterface
    description: The ENI that is the source of traffic mirroring.
    type: string
    exposed: true
    stored: true

  - name: status
    description: The status of the traffic mirroring session.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Success
    - Failed

  - name: statusReason
    description: The status description of the traffic mirroring session.
    type: string
    exposed: true
    stored: true
