# Table: aws_wafv2_managed_rule_groups

This table shows data for Wafv2 Managed Rule Groups.

https://docs.aws.amazon.com/waf/latest/APIReference/API_ManagedRuleGroupSummary.html

The composite primary key for this table is (**account_id**, **region**, **scope**, **name**, **vendor_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|scope (PK)|String|
|properties|JSON|
|description|String|
|name (PK)|String|
|vendor_name (PK)|String|
|versioning_supported|Bool|