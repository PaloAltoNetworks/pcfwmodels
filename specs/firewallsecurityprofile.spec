# Model
model:
  rest_name: firewallsecurityprofile
  resource_name: firewallsecurityprofiles
  entity_name: FirewallSecurityProfile
  package: ngfw
  group: core/ngfw
  description: "Represents firewall security setting.  \nThis object can be created
    by the user and referenced by a Group object."
  get:
    description: Retrieves the firewallsecurityprofile with the given ID.
  update:
    description: Updates the firewallsecurityprofile with the given ID.
  delete:
    description: Deletes the firewallsecurityprofile with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@tags'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: IPSEnabled
    description: Enables IPS.
    type: boolean
    exposed: true
    stored: true

  - name: URLCategoriesEnabled
    description: Enables best practice URL Categories.
    type: boolean
    exposed: true
    stored: true

  - name: antiSpywareEnabled
    description: Enables AntiSpyware.
    type: boolean
    exposed: true
    stored: true

  - name: antiVirusEnabled
    description: Enables AntiVirus.
    type: boolean
    exposed: true
    stored: true

  - name: fileBlockingEnabled
    description: Enables file blocking.
    type: boolean
    exposed: true
    stored: true

  - name: lastCommittedTime
    description: The date when the securityprofile was last committed.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: lastUpdatedTime
    description: The date when the securityprofile was last updated.
    type: time
    exposed: true
    stored: true
    read_only: true

  - name: trustedCertificateID
    description: The trusted Certificate ID for Outbound Decryption.
    type: string
    exposed: true
    stored: true

  - name: untrustedCertificateID
    description: The untrusted Certificate ID for Outbound Decryption.
    type: string
    exposed: true
    stored: true
