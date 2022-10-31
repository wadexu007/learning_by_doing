locals {
  cluster_name    = "test-eks-2022"
  cluster_version = "1.22"
  region          = "cn-north-1"

  vpc = {
    cidr = "10.0.0.0/16"
    private_subnets = ["10.0.1.0/24", "10.0.2.0/24"]
    public_subnets  = ["10.0.4.0/24", "10.0.5.0/24"]
  }

  master_authorized_networks =   [
    "4.14.xxx.xx/32",   # allow office 1, need to change to your whitelist
    "64.124.xxx.xx/32", # allow office 2, need to change to your whitelist
    "0.0.0.0/0"         # allow all access master node, not suggest
  ]

  # Extend cluster security group rules example
  cluster_security_group_additional_rules = {
    egress_nodes_ephemeral_ports_tcp = {
      description                = "To node 1025-65535"
      protocol                   = "tcp"
      from_port                  = 1025
      to_port                    = 65535
      type                       = "egress"
      source_node_security_group = true
    }
  }

  node_group_default = {
    ami_type     = "AL2_x86_64"
    min_size     = 1
    max_size     = 5
    desired_size = 1
  }

  dmz_group = {
  }

  app_group = {
    instance_types = ["t3.small"]
    disk_size    = 50

    # example rules added for app node group
    security_group_rules = {
        egress_1 = {
          description = "Hello CloudFlare"
          protocol    = "udp"
          from_port   = 53
          to_port     = 53
          type        = "egress"
          cidr_blocks = ["1.1.1.1/32"]
        }
      }
  }
}
