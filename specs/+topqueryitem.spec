# Model
model:
  rest_name: topqueryitem
  resource_name: topqueryitems
  entity_name: TopQueryItem
  package: logging
  group: core/log
  description: Represents an item in a topquery result.
  detached: true

# Attributes
attributes:
  v1:
  - name: count
    description: The count of the field.
    type: integer
    exposed: true

  - name: value
    description: The value of the field.
    type: string
    exposed: true
