# Model
model:
  rest_name: testmodel
  resource_name: testmodels
  entity_name: TestModel
  package: test
  group: test/testmodel
  description: Test model used for development.
  private: true
  get:
    description: Retrieves the testmodel with the given ID.
  update:
    description: Updates the testmodel with the given ID.
  delete:
    description: Deletes the testmodel with the given ID.
    global_parameters:
    - $queryable
  extends:
  - '@namespaced'
  - '@named'

# Indexes
indexes:
- - namespace

# Attributes
attributes:
  v1:
  - name: why
    description: Contains the reason why this object was created.
    type: string
    exposed: true
    stored: true
    validations:
    - $whystring
