locals {

  project_id     = "global-sre-dev"
  cluster_name   = "sre-gke"
  cluster_region = "us-central1"
  emissary_ns    = "emissary"
  chart_version  = "8.2.0"
  common_yaml_d  = "../common/helm/yamls"
  ambassador_id  = "ambassador"

  emissary_ingress_map = {
    ambassadorID          = local.ambassador_id
    loadBalancerIP        = "35.232.98.249" # Prepare a Static IP first instead to use Ephemeral
    replicaCount          = 2
    minReplicas           = 2
    maxReplicas           = 3
    canaryEnabled         = false # set to true in Prod
    logLevel              = "error" # valid log levels are error, warn/warning, info, debug, and trace
    endpointEnable        = true
    endpointName          = "my-resolver"
    diagnosticsEnable     = false
    clusterRequestTimeout = 120000 # milliseconds
  }

  emissary_listeners_map = {
    ambassadorID          = local.ambassador_id
    listenersEnabled      = true # custom listeners
  }
}

resource "helm_release" "emissary_crds" {
  name             = "emissary-crds"
  create_namespace = true # create `emissary-system` namespace, this is CRDs default namespace
  namespace        = "emissary-system"
  chart            = "../common/helm/repos/emissary-crds-8.2.0.tgz"
}

# Install Emissary-ingress from Chart Repository
resource "helm_release" "emissary_ingress" {
  name             = "emissary-ingress"
  repository       = "https://app.getambassador.io"
  chart            = "emissary-ingress"
  version          = local.chart_version
  create_namespace = true
  namespace        = local.emissary_ns

  values = [
    templatefile("${local.common_yaml_d}/emissary-ingress-template.yaml", local.emissary_ingress_map)
  ]

  depends_on = [
    helm_release.emissary_crds
  ]
}

# This is for install Host/Listener/Mapping/TLSContext from a local custom chart
# also can upload chart to a bucket or a public github for install from a url
# e.g. [Publish to a GCS bucket](https://github.com/hayorov/helm-gcs)
resource "helm_release" "emissary_config" {
  name      = "emissary-config"
  namespace = local.emissary_ns
  chart     = "../common/helm/repos/emissary-config-8.2.0.tgz"

  values = [
    templatefile("${local.common_yaml_d}/emissary-listeners-template.yaml", local.emissary_listeners_map),
    local.emissary_config_yaml
  ]

  depends_on = [
    helm_release.emissary_ingress
  ]
}