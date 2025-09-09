Use this data source to query detailed information of tdmq pulsar environments

Example Usage

```hcl

	data "cloud_tdmq_pulsar_environments" "example" {
	    environment_id = "keep-ns"
	    cluster_id     = "pulsar-9n95ax58b9vn"
	}

```