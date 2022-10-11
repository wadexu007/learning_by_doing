module "example_custom_manifests" {
  source  = "kbst.xyz/catalog/custom-manifests/kustomization"
  version = "0.3.0"

  configuration_base_key = "default"
  configuration = {
    default = {
      # namespace = "test" # namespace already specify in Kustomize/demo-manifests/services/demo-app/_common

      resources = [
        "${path.root}/../../Kustomize/demo-manifests/services/demo-app/dev"
      ]

      images = [{
        # Refers to the 'pod.spec.container.name' to modify the 'image' attribute of.
        # name     = "wadexu007/demo"
        
        # Customize the 'registry/name' part of the image. The part before the ':'
        # new_name = "wadexu007/demo"
        
        # Customize the 'tag' part of the image. The part after the ':'.
        new_tag  = "1.0.0"
      }]

      common_labels = {
        "env" = "dev"
      }
    }
  }
}
