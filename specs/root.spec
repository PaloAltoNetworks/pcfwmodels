# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  group: core
  description: root object.
  get:
    description: gets the object.
  root: true

# Relations
relations:
- rest_name: awsfirewall
  get:
    description: Retrieves the list of AWS firewalls.
    global_parameters:
    - $queryable
  create:
    description: Creates a new AWS firewall.

- rest_name: awslogdefinition
  get:
    description: Retrieves the list of AWS log definitions.
    global_parameters:
    - $queryable
  create:
    description: Creates a new AWS log definition.

- rest_name: firewallappidlist
  get:
    description: Retrieves a firewall AppID list.
    global_parameters:
    - $queryable

- rest_name: firewallapplicationgroup
  get:
    description: Retrieves the list of firewall application groups.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall application group.

- rest_name: firewallcertificate
  get:
    description: Retrieves the list of firewall certificates.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall certificate.

- rest_name: firewallcidrlist
  get:
    description: Retrieves the list of firewall CIDR lists.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall CIDR list.

- rest_name: firewallcommitaction
  create:
    description: Creates a new firewall commit action.

- rest_name: firewallcommittedruleset
  get:
    description: Retrieves the list of firewall committed rulesets.
    global_parameters:
    - $queryable

- rest_name: firewallcountrycodelist
  get:
    description: Retrieves a firewall country code list.
    global_parameters:
    - $queryable

- rest_name: firewallcustomexternalfeed
  get:
    description: Retrieves the list of firewall custom external feeds.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall custom external feed.

- rest_name: firewallcustomurlcategory
  get:
    description: Retrieves the list of firewall custom URL categories.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall custom URL category.

- rest_name: firewallexternalfeedlist
  get:
    description: Retrieves a firewall external feed list.
    global_parameters:
    - $queryable

- rest_name: firewallfqdnlist
  get:
    description: Retrieves a the list of firewall FQDN lists.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall FQDN list.

- rest_name: firewalllicensingcredits
  get:
    description: Retrieves the list of firewall licensing credits.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall licensing credits.

- rest_name: firewalllicensingmetrics
  get:
    description: Retrieves the list of firewall licensing metrics.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall licensing metrics.

- rest_name: firewalloption
  create:
    description: Discovers firewall deployment options.

- rest_name: firewallrollbackaction
  create:
    description: Creates a new firewall rollback action.

- rest_name: firewallrule
  get:
    description: Retrieves the list of firewall rules.
    global_parameters:
    - $queryable

- rest_name: firewallruleset
  get:
    description: Retrieves the list of firewall rulesets.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall ruleset.

- rest_name: firewallsecurityprofile
  get:
    description: Retrieves the list of firewall security profiles.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall security profile.

- rest_name: firewalltemplate
  get:
    description: Retrieves the list of firewall templates.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewall template.

- rest_name: firewallurlcategorylist
  get:
    description: Retrieves a firewall URL category list.
    global_parameters:
    - $queryable

- rest_name: logdebug
  create:
    description: Initiate a log debug command.
    global_parameters:
    - $timewindow

- rest_name: logfieldstat
  get:
    description: Retrieves statistics for a log field.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a statistic for a log field.

- rest_name: logincidentdetailsquery
  create:
    description: Creates a new log incident details query.
    global_parameters:
    - $timewindow

- rest_name: logincidentquery
  create:
    description: Creates a new log incident query.
    global_parameters:
    - $timewindow

- rest_name: logquery
  create:
    description: Creates a new log query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: mirrorsourceoption
  create:
    description: Discovers eligible traffic mirror source instances and auto-scaling
      groups.

- rest_name: pcfwaccount
  get:
    description: Retrieves the status of cloud account on Prisma firewall service.
    global_parameters:
    - $queryable
  create:
    description: Onboards Cloud Account on the Prisma Prisma Cloud service.

- rest_name: pcfwaccountmirrorsource
  get:
    description: Retrieves the mirror configuration of a firewall for a cloud account.
    global_parameters:
    - $queryable
  create:
    description: Retrieves the mirror configuration of a firewall for a cloud account.

- rest_name: pcfwaccountmirrorsourcestate
  get:
    description: Retrieves the mirror source status of a firewall for a cloud account.
    global_parameters:
    - $queryable
  create:
    description: Retrieves the mirror source status of a firewall for a cloud account.

- rest_name: pcfwfirewallconfigterraform
  create:
    description: Creates a firewall configuration terraform plan associated for a
      cloud account.
    global_parameters:
    - $queryable

- rest_name: pcfwlambdahealth
  get:
    description: Retrieves the health status of a lambda for a cloud account.
    global_parameters:
    - $queryable
  create:
    description: Creates the health status of a lambda for a cloud account.

- rest_name: pcfwsubnethelper
  create:
    description: Creates a list of free subnets per availability zone.
    global_parameters:
    - $queryable

- rest_name: pcfwtenant
  get:
    description: Retrieves the status of Prisma Cloud Tenant.
    global_parameters:
    - $queryable
  create:
    description: Onboards Prisma Cloud tenant on Prisma Cloud Firewall.

- rest_name: pcfwtenantaccountterraform
  create:
    description: Creates a new account setup terraform plan for a tenant.
    global_parameters:
    - $queryable

- rest_name: pcfwtenantterraform
  create:
    description: Creates a multi account setup terraform plan for a tenant.
    global_parameters:
    - $queryable

- rest_name: testmodel
  get:
    description: Retrieves the list of testmodels.
    global_parameters:
    - $queryable
  create:
    description: Creates a new testmodel.

- rest_name: topquery
  create:
    description: Creates a new top query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: topsnapshot
  get:
    description: Retrieves top snapshots.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a new top snapshot.

- rest_name: totalquery
  create:
    description: Creates a new total query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: trendquery
  create:
    description: Creates a new trend query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: trendsnapshot
  get:
    description: Retrieves trend snapshots.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a new trend snapshot.
