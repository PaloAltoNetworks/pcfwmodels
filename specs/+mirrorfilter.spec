# Model
model:
  rest_name: mirrorfilter
  resource_name: mirrorfilters
  entity_name: MirrorFilter
  package: deploymentadvisor
  group: core/deploymentadvisor
  description: Represents an AWS mirroring filter.
  detached: true

# Attributes
attributes:
  v1:
  - name: networkServices
    description: AWS network services to be mirrored.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $awsNetworkServicesFilter

  - name: rules
    description: A list of mirror filter rules.
    type: refList
    exposed: true
    subtype: mirrorrule
    stored: true
    extensions:
      refMode: pointer
