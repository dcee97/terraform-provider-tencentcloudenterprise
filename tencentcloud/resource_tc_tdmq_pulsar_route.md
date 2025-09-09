Provide a resource to create a TDMQ Pulsar Route.

Example Usage
```hcl

	resource "cloud_tdmq_pulsar_route" "foo" {
	  remark = "this is description111."
	  cluster_id = 0
	  net_type = 2
	}

```

Import
Tdmq Route can be imported, e.g.

```
$ terraform import cloud_tdmq_pulsar_route.test tdmq_id
```