<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0, <= 1.5.5 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.77 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | git::https://github.com/launchbynttdata/tf-launch-module_library-resource_name.git | 1.0.0 |
| <a name="module_resource_group"></a> [resource\_group](#module\_resource\_group) | git::https://github.com/launchbynttdata/tf-azurerm-module_primitive-resource_group.git | 1.0.0 |
| <a name="module_public_dns_zone"></a> [public\_dns\_zone](#module\_public\_dns\_zone) | git::https://github.com/launchbynttdata/tf-azurerm-module_primitive-dns_zone.git | 1.0.0 |
| <a name="module_dns_records"></a> [dns\_records](#module\_dns\_records) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | A map of key to resource\_name that will be used by tf-launch-module\_library-resource\_name to generate resource names | <pre>map(object({<br>    name       = string<br>    max_length = optional(number, 60)<br>    region     = optional(string, "eastus2")<br>  }))</pre> | <pre>{<br>  "public_dns_zone": {<br>    "max_length": 80,<br>    "name": "zone",<br>    "region": "eastus"<br>  },<br>  "resource_group": {<br>    "max_length": 80,<br>    "name": "rg",<br>    "region": "eastus"<br>  }<br>}</pre> | no |
| <a name="input_environment"></a> [environment](#input\_environment) | Project environment | `string` | `"demo"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Number that represents the instance of the environment. | `number` | `0` | no |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Number that represents the instance of the resource. | `number` | `0` | no |
| <a name="input_region"></a> [region](#input\_region) | AWS Region in which the infra needs to be provisioned | `string` | `"eastus2"` | no |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | (Required) Name of the product family for which the resource is created.<br>    Example: org\_name, department\_name. | `string` | `"launch"` | no |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | (Required) Name of the product service for which the resource is created.<br>    For example, backend, frontend, middleware etc. | `string` | `"dns"` | no |
| <a name="input_a_records"></a> [a\_records](#input\_a\_records) | A map of A records to create | <pre>map(object({<br>    ttl                = number<br>    records            = list(string)<br>    target_resource_id = optional(string)<br>    tags               = optional(map(string))<br>  }))</pre> | `{}` | no |
| <a name="input_cname_records"></a> [cname\_records](#input\_cname\_records) | A map of CNAME records to create | <pre>map(object({<br>    ttl                = number<br>    record             = optional(string)<br>    target_resource_id = optional(string)<br>    tags               = optional(map(string))<br>  }))</pre> | `{}` | no |
| <a name="input_ns_records"></a> [ns\_records](#input\_ns\_records) | A map of NS records to create | <pre>map(object({<br>    ttl     = number<br>    records = list(string)<br>    tags    = optional(map(string))<br>  }))</pre> | `{}` | no |
| <a name="input_txt_records"></a> [txt\_records](#input\_txt\_records) | A map of TXT records to create | <pre>map(object({<br>    ttl     = number<br>    records = list(string)<br>    tags    = optional(map(string))<br>  }))</pre> | `{}` | no |
| <a name="input_domain_names"></a> [domain\_names](#input\_domain\_names) | A list of domain names to create | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to add to the DNS zone | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_a_record_ids"></a> [a\_record\_ids](#output\_a\_record\_ids) | The DNS A Record IDs. |
| <a name="output_a_record_fqdns"></a> [a\_record\_fqdns](#output\_a\_record\_fqdns) | The FQDNs of the DNS A Records. |
| <a name="output_cname_record_ids"></a> [cname\_record\_ids](#output\_cname\_record\_ids) | The DNS CNAME Record IDs. |
| <a name="output_cname_record_fqdns"></a> [cname\_record\_fqdns](#output\_cname\_record\_fqdns) | The FQDNs of the DNS CNAME Records. |
| <a name="output_ns_record_ids"></a> [ns\_record\_ids](#output\_ns\_record\_ids) | The DNS NS Record IDs. |
| <a name="output_ns_record_fqdns"></a> [ns\_record\_fqdns](#output\_ns\_record\_fqdns) | The FQDNs of the DNS NS Records. |
| <a name="output_txt_record_ids"></a> [txt\_record\_ids](#output\_txt\_record\_ids) | The DNS TXT Record IDs. |
| <a name="output_txt_record_fqdns"></a> [txt\_record\_fqdns](#output\_txt\_record\_fqdns) | The FQDNs of the DNS TXT Records. |
| <a name="output_resource_group_name"></a> [resource\_group\_name](#output\_resource\_group\_name) | The name of the resource group. |
| <a name="output_public_dns_zone_names"></a> [public\_dns\_zone\_names](#output\_public\_dns\_zone\_names) | The names of the public DNS zone. |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
