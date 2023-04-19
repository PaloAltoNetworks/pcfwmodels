# Model
model:
  rest_name: logqueryitem
  resource_name: logqueryitems
  entity_name: LogQueryItem
  package: logging
  group: core/log
  description: Represents a log line in a log query result.
  detached: true

# Attributes
attributes:
  v1:
  - name: fields
    description: The fields and values for the log line.
    type: external
    exposed: true
    subtype: map[string]string

  - name: timestamp
    description: The timestamp of the log line.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
