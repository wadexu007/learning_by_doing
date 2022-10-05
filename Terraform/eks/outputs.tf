output "cluster_id" {
  description = "EKS cluster ID"
  value       = module.wade-eks.cluster_id
}

output "cluster_endpoint" {
  description = "Endpoint for EKS control plane"
  value       = module.wade-eks.cluster_endpoint
}

output "region" {
  description = "EKS region"
  value       = local.region
}

output "cluster_name" {
  description = "AWS Kubernetes Cluster Name"
  value       = local.cluster_name
}
