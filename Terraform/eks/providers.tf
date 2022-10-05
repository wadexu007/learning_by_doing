terraform {
  backend "s3" {
    bucket = "sre-dev-terraform"
    key    = "test/eks.tfstate"
    region = "cn-north-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.25.0"
    }
  }
}

provider "aws" {
  region     = local.region
}

# https://github.com/terraform-aws-modules/terraform-aws-eks/issues/2009
provider "kubernetes" {
  host                   = module.wade-eks.cluster_endpoint
  cluster_ca_certificate = base64decode(module.wade-eks.cluster_certificate_authority_data)

  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    # This requires the awscli to be installed locally where Terraform is executed
    args = ["eks", "get-token", "--cluster-name", module.wade-eks.cluster_id]
  }
}