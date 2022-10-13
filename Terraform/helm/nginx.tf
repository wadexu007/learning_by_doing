resource "helm_release" "my-nginx" {
  name       = "my-nginx-release"

  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx"

  # values = [
  #   "${file("additional_values.yaml")}"
  # ]

  set {
    name  = "service.type"
    value = "ClusterIP"
  }

  set {
    name  = "replicaCount"
    value = "2"
  }
}