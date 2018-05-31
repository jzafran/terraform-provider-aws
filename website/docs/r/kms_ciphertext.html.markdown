---
layout: "aws"
page_title: "AWS: aws_kms_ciphertext"
sidebar_current: "docs-aws-resource-kms-ciphertext"
description: |-
    Provides ciphertext encrypted using a KMS key
---

# Data Source: aws_kms_ciphertext

The KMS ciphertext resource allows you to encrypt plaintext into ciphertext
by using an AWS KMS customer master key. This has similar functionality to the
aws_kms_ciphertext data source, except the ciphertext will only be regenerated/re-encrypted
only when `plaintext` changes.

~> **Note:** All arguments excluding the plaintext will be stored in the raw state as plain-text. A hash of the plaintext will be stored in the state. 
[Read more about sensitive data in state](/docs/state/sensitive-data.html).

## Example Usage

```hcl
resource "aws_kms_key" "oauth_config" {
  description = "oauth config"
  is_enabled = true
}

resource "aws_kms_ciphertext" "oauth" {
  key_id = "${aws_kms_key.oauth_config.key_id}"
  plaintext = <<EOF
{
  "client_id": "e587dbae22222f55da22",
  "client_secret": "8289575d00000ace55e1815ec13673955721b8a5"
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `plaintext` - (Required) Data to be encrypted. The value of this attribute when stored into the Terraform state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
* `key_id` - (Required) Globally unique key ID for the customer master key.
* `context` - (Optional) An optional mapping that makes up the encryption context.

## Attributes Reference

All of the argument attributes are also exported as result attributes.

* `ciphertext_blob` - Base64 encoded ciphertext
