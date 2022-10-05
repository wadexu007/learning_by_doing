variable "network_name" {
  description = "The name of the vpc"
  type        = string
}

variable "project" {
  description = "The name of the project"
  type        = string
}

variable "target_tags" {
  description = "default no target tags means firewall apply to all instances"
  type    = list
  default = []
}