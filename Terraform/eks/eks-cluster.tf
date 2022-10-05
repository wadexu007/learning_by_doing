data "aws_availability_zones" "available" {

  # Cannot create cluster because cn-north-1d, 
  # the targeted availability zone, does not currently have sufficient capacity to support the cluster.
  exclude_names = ["cn-north-1d"]
}

module "wade-eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "18.27.1"

  cluster_name    = local.cluster_name
  cluster_version = local.cluster_version

  cluster_endpoint_private_access = true
  cluster_endpoint_public_access  = true

  # api server authorized network list
  cluster_endpoint_public_access_cidrs = local.master_authorized_networks

  cluster_addons = {
    coredns = {
      resolve_conflicts = "OVERWRITE"
    }
    kube-proxy = {}
    vpc-cni = {
      resolve_conflicts = "OVERWRITE"
    }
  }

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets

  # Extend cluster security group rules
  cluster_security_group_additional_rules = local.cluster_security_group_additional_rules

  eks_managed_node_group_defaults = {
    ami_type     = local.node_group_default.ami_type
    min_size     = local.node_group_default.min_size
    max_size     = local.node_group_default.max_size
    desired_size = local.node_group_default.desired_size
  }

  eks_managed_node_groups = {
    # dmz = {
    #   name = "dmz-pool"
    # }
    app = {
      name                           = "app-pool"
      instance_types                 = local.app_group.instance_types
      create_launch_template         = false
      launch_template_name           = ""
      disk_size                      = local.app_group.disk_size

      create_security_group          = true
      security_group_name            = "app-node-group-sg"
      security_group_use_name_prefix = false
      security_group_description     = "EKS managed app node group security group"
      security_group_rules            = local.app_group.security_group_rules

      update_config = {
        max_unavailable_percentage = 50
      }
    }
  }

  # aws-auth configmap
  # create_aws_auth_configmap = true
  manage_aws_auth_configmap = true

  aws_auth_roles = [
    {
      rolearn  = "arn:aws-cn:iam::9935108xxxxx:role/CN-SRE" # replace me
      username = "sre"
      groups   = ["system:masters"]
    },
  ]

  aws_auth_users = [
    {
      userarn  = "arn:aws-cn:iam::9935108xxxxx:user/wadexu" # replace me
      username = "wadexu"
      groups   = ["system:masters"]
    },
  ]

  tags = {
    Environment = "dev"
    Terraform   = "true"
  }

  # aws china only because https://github.com/terraform-aws-modules/terraform-aws-eks/pull/1905
  cluster_iam_role_dns_suffix = "amazonaws.com"
}
