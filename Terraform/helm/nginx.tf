resource "helm_release" "my-nginx" {
  name       = "my-nginx-release"

  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx"

  set {
    name  = "service.type"
    value = "ClusterIP"
  }

  set {
    name  = "replicaCount"
    value = "2"
  }
}