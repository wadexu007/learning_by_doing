module "example_sealed_secrets" {

  source  = "kbst.xyz/catalog/sealed-secrets/kustomization"
  version = "0.18.0-kbst.0"
  configuration_base_key = "default"
  configuration = {
    default = {
      namespace = "test" # namespace must exist

      additional_resources = [
        "${path.root}/manifests/namespace.yaml" # to create namespace
      ]

      common_labels = {
        "env" = "dev"
      }
    }

    ops = {}

    loc = {}
  }
}
