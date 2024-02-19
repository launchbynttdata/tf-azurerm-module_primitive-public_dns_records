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

locals {

  a_recod_names = {
    for k, v in var.a_records : "a_record_${k}" => {
      name       = "arecord${k}"
      max_length = 80
      region     = var.region
    }
  }
  cname_record_names = {
    for k, v in var.cname_records : "cname_record_${k}" => {
      name       = "canamerecord${k}"
      max_length = 80
      region     = var.region
    }
  }
  ns_record_names = {
    for k, v in var.ns_records : "ns_record_${k}" => {
      name       = "nsrecord${k}"
      max_length = 80
      region     = var.region
    }
  }
  txt_record_names = {
    for k, v in var.txt_records : "txt_record_${k}" => {
      name       = "txtrecord${k}"
      max_length = 80
      region     = var.region
    }
  }
  resource_names_map = merge(var.resource_names_map, local.a_recod_names, local.cname_record_names, local.ns_record_names, local.txt_record_names)


  resource_group_name = module.resource_names["resource_group"].standard
  dns_zone_name       = module.resource_names["public_dns_zone"].dns_compliant_standard

  // Records in the first dns zone
  a_records = {
    for k, v in var.a_records : "a_record_${k}" => {
      name                = module.resource_names["a_record_${k}"].standard
      resource_group_name = local.resource_group_name
      zone_name           = var.domain_names[0]
      ttl                 = v.ttl
      records             = v.records
      target_resource_id  = null
      tags                = v.tags
    }
  }

  cname_records = {
    for k, v in var.cname_records : "cname_record_${k}" => {
      name                = module.resource_names["cname_record_${k}"].standard
      resource_group_name = local.resource_group_name
      zone_name           = var.domain_names[0]
      ttl                 = v.ttl
      record              = v.record
      target_resource_id  = null
      tags                = v.tags
    }
  }

  ns_records = {
    for k, v in var.ns_records : "ns_record_${k}" => {
      name                = module.resource_names["ns_record_${k}"].standard
      resource_group_name = local.resource_group_name
      zone_name           = var.domain_names[0]
      ttl                 = v.ttl
      records             = v.records
      tags                = v.tags
    }
  }

  txt_records = {
    for k, v in var.txt_records : "txt_record_${k}" => {
      name                = module.resource_names["txt_record_${k}"].standard
      resource_group_name = local.resource_group_name
      zone_name           = var.domain_names[0]
      ttl                 = v.ttl
      records             = v.records
      tags                = v.tags
    }
  }
}
