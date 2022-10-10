resource "helm_release" "my-redis" {
  name       = "my-redis-release"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "redis"
  version    = "17.3.4"

#   values = [
#     "${file("values.yaml")}"
#   ]

  set {
    name  = "architecture"
    value = "standalone" # for demo purpose, default is replication
  }

  set {
    name  = "auth.enabled"
    value = "false" # for demo purpose, default is true
  }

}