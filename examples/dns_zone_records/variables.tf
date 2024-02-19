// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//variables required by resource names module
variable "resource_names_map" {
  description = "A map of key to resource_name that will be used by tf-module-resource_name to generate resource names"
  type = map(object({
    name       = string
    max_length = optional(number, 60)
    region     = optional(string, "eastus2")
  }))

  default = {
    resource_group = {
      name       = "rg"
      max_length = 80
      region     = "eastus"
    }
    public_dns_zone = {
      name       = "zone"
      max_length = 80
      region     = "eastus"
    }
  }
}

variable "environment" {
  description = "Project environment"
  type        = string
  default     = "demo"
}

variable "instance_env" {
  type        = number
  description = "Number that represents the instance of the environment."
  default     = 0

  validation {
    condition     = var.instance_env >= 0 && var.instance_env <= 999
    error_message = "Instance number should be between 1 to 999."
  }
}

variable "instance_resource" {
  type        = number
  description = "Number that represents the instance of the resource."
  default     = 0

  validation {
    condition     = var.instance_resource >= 0 && var.instance_resource <= 100
    error_message = "Instance number should be between 1 to 100."
  }
}

variable "region" {
  description = "AWS Region in which the infra needs to be provisioned"
  type        = string
  default     = "eastus2"
}

variable "logical_product_family" {
  type        = string
  description = <<EOF
    (Required) Name of the product family for which the resource is created.
    Example: org_name, department_name.
  EOF
  nullable    = false

  validation {
    condition     = can(regex("^[_\\-A-Za-z0-9]+$", var.logical_product_family))
    error_message = "The variable must contain letters, numbers, -, _, and .."
  }

  default = "launch"
}

variable "logical_product_service" {
  type        = string
  description = <<EOF
    (Required) Name of the product service for which the resource is created.
    For example, backend, frontend, middleware etc.
  EOF
  nullable    = false

  validation {
    condition     = can(regex("^[_\\-A-Za-z0-9]+$", var.logical_product_service))
    error_message = "The variable must contain letters, numbers, -, _, and .."
  }

  default = "dns"
}

# Variables required by the records module
variable "a_records" {
  description = "A map of A records to create"
  type = map(object({
    ttl                = number
    records            = list(string)
    target_resource_id = optional(string)
    tags               = optional(map(string))
  }))
  default = {}
}

variable "cname_records" {
  description = "A map of CNAME records to create"
  type = map(object({
    ttl                = number
    record             = optional(string)
    target_resource_id = optional(string)
    tags               = optional(map(string))
  }))
  default = {}
}

variable "ns_records" {
  description = "A map of NS records to create"
  type = map(object({
    ttl     = number
    records = list(string)
    tags    = optional(map(string))
  }))
  default = {}
}

variable "txt_records" {
  description = "A map of TXT records to create"
  type = map(object({
    ttl     = number
    records = list(string)
    tags    = optional(map(string))
  }))
  default = {}
}

# Variables required by the a dns zone module
variable "domain_names" {
  description = "A list of domain names to create"
  type        = list(string)
  default     = []
}

variable "tags" {
  description = "A map of tags to add to the DNS zone"
  type        = map(string)
  default     = {}
}
