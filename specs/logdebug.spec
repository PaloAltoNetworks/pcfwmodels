# Model
model:
  rest_name: logdebug
  resource_name: logdebugs
  entity_name: LogDebug
  package: logging
  group: core/log
  description: Get debug information or initiate debug commands.

# Attributes
attributes:
  v1:
  - name: command
    description: The command to perform.
    type: string
    exposed: true

  - name: firewallName
    description: The NGFW name.
    type: string
    exposed: true
    example_value: firewall1

  - name: information
    description: Additional information returned to caller.
    type: string
    exposed: true
    read_only: true
