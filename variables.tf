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

variable "a_records" {
  description = "A list of A records to create"
  type = map(object({
    name                = string
    resource_group_name = string
    zone_name           = string
    ttl                 = number
    records             = list(string)
    target_resource_id  = optional(string)
    tags                = optional(map(string))
  }))
  default = {}
}

variable "cname_records" {
  description = "A list of A records to create"
  type = map(object({
    name                = string
    resource_group_name = string
    zone_name           = string
    ttl                 = number
    record              = optional(string)
    target_resource_id  = optional(string)
    tags                = optional(map(string))
  }))
  default = {}
}

variable "ns_records" {
  description = "A list of A records to create"
  type = map(object({
    name                = string
    resource_group_name = string
    zone_name           = string
    ttl                 = number
    records             = list(string)
    tags                = optional(map(string))
  }))
  default = {}
}

variable "txt_records" {
  description = "A list of A records to create"
  type = map(object({
    name                = string
    resource_group_name = string
    zone_name           = string
    ttl                 = number
    records             = list(string)
    tags                = optional(map(string))
  }))
  default = {}
}
