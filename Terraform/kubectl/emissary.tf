locals {

  project_id     = "sre-eng-cn-dev"
  cluster_name   = "sre-mgmt"
  cluster_region = "asia-east2"
  emissary_ns    = "emissary"
  chart_version  = "8.2.0"
}

//emissary crds
data "kubectl_file_documents" "emissary_crds" {
  content = file("yamls/emissary-crds.yaml")
}

resource "kubectl_manifest" "emissary_crds" {
  for_each  = data.kubectl_file_documents.emissary_crds.manifests
  yaml_body = each.value
}

# Install Emissary-ingress from Chart Repository
resource "helm_release" "emissary_ingress" {
  name             = "emissary-ingress"
  repository       = "https://app.getambassador.io"
  chart            = "emissary-ingress"
  version          = local.chart_version
  create_namespace = true
  namespace        = local.emissary_ns

  depends_on = [
    kubectl_manifest.emissary_crds
  ]
}
