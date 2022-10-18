# Model
model:
  rest_name: logquery
  resource_name: logqueries
  entity_name: LogQuery
  package: logserv
  group: core/log
  description: Answer general queries on firewall logs.

# Attributes
attributes:
  v1:
  - name: countHint
    description: Optional hint about the result set size, provided by the caller.
    type: integer
    exposed: true

  - name: logResult
    description: The result of the log query.
    type: refList
    exposed: true
    subtype: logqueryitem
    read_only: true
    autogenerated: true
    extensions:
      refMode: pointer

  - name: logType
    description: The type of firewall log to query.
    type: enum
    exposed: true
    allowed_choices:
    - Traffic
    - Threat
    - URLFiltering
    - Decryption
