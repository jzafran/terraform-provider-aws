package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceAwsKmsCiphertext_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsKmsCiphertextConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttr(
						"aws_kms_ciphertext.foo", "plaintext",
						plaintextHashSum("Super secret data")),
				),
			},
		},
	})
}

func TestAccResourceAwsKmsCiphertext_validate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsKmsCiphertextConfig_validate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
		},
	})
}

func TestAccResourceAwsKmsCiphertext_validate_withContext(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsKmsCiphertextConfig_validate_withContext,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
		},
	})
}

func TestAccResourceAwsKmsCiphertext_validate_unchanged(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsKmsCiphertextConfig_validate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
			{
				Config:             testAccResourceAwsKmsCiphertextConfig_validate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
			{
				Config:             testAccResourceAwsKmsCiphertextConfig_validate,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
		},
	})
}

func TestAccResourceAwsKmsCiphertext_validate_plaintext(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAwsKmsCiphertextConfig_validate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
			{
				Config:             testAccResourceAwsKmsCiphertextConfig_validate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Super secret data"),
				),
			},
			{
				Config:             testAccResourceAwsKmsCiphertextConfig_validate_plaintext,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"aws_kms_ciphertext.foo", "ciphertext_blob"),
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_secret.foo", "plaintext"),
					resource.TestCheckResourceAttr(
						"data.aws_kms_secret.foo", "plaintext", "Other secret data"),
				),
			},
		},
	})
}

const testAccResourceAwsKmsCiphertextConfig_basic = `
provider "aws" {
 region = "us-west-2"
}

resource "aws_kms_key" "foo" {
  description = "tf-test-acc-resource-aws-kms-ciphertext-basic"
  is_enabled = true
  deletion_window_in_days = 7
}

resource "aws_kms_ciphertext" "foo" {
 key_id = "${aws_kms_key.foo.key_id}"
 plaintext = "Super secret data"
}
`

const testAccResourceAwsKmsCiphertextConfig_validate = `
provider "aws" {
 region = "us-west-2"
}

resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-validate"
  is_enabled = true
  deletion_window_in_days = 7
}

resource "aws_kms_ciphertext" "foo" {
 key_id = "${aws_kms_key.foo.key_id}"
 plaintext = "Super secret data"
}

data "aws_kms_secret" "foo" {
 secret {
   name = "plaintext"
   payload = "${aws_kms_ciphertext.foo.ciphertext_blob}"
 }
}
`

const testAccResourceAwsKmsCiphertextConfig_validate_plaintext = `
provider "aws" {
 region = "us-west-2"
}

resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-validate-plaintext"
  is_enabled = true
  deletion_window_in_days = 7
}

resource "aws_kms_ciphertext" "foo" {
 key_id = "${aws_kms_key.foo.key_id}"
 plaintext = "Other secret data"
}

data "aws_kms_secret" "foo" {
 secret {
   name = "plaintext"
   payload = "${aws_kms_ciphertext.foo.ciphertext_blob}"
 }
}
`

const testAccResourceAwsKmsCiphertextConfig_validate_withContext = `
provider "aws" {
 region = "us-west-2"
}

resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-validate-with-context"
  is_enabled = true
  deletion_window_in_days = 7
}

resource "aws_kms_ciphertext" "foo" {
 key_id = "${aws_kms_key.foo.key_id}"
 plaintext = "Super secret data"

 context {
	name = "value"
 }
}

data "aws_kms_secret" "foo" {
 secret {
   name = "plaintext"
   payload = "${aws_kms_ciphertext.foo.ciphertext_blob}"

   context {
	  name = "value"
   }
 }
}
`
