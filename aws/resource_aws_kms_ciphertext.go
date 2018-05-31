package aws

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceAwsKmsCiphertext() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsKmsCiphertextCreate,
		Read:   resourceAwsKmsCiphertextRead,
		Delete: resourceAwsKmsCiphertextDelete,

		Schema: map[string]*schema.Schema{
			"plaintext": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
				StateFunc: plaintextHashSum,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"context": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"ciphertext_blob": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func plaintextHashSum(v interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v.(string))))
}

func resourceAwsKmsCiphertextCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).kmsconn

	req := &kms.EncryptInput{
		KeyId:     aws.String(d.Get("key_id").(string)),
		Plaintext: []byte(d.Get("plaintext").(string)),
	}

	if v, exists := d.GetOk("context"); exists {
		req.EncryptionContext = stringMapToPointers(v.(map[string]interface{}))
	}

	log.Printf("[DEBUG] KMS encrypt for key: %s", d.Get("key_id").(string))

	resp, err := conn.Encrypt(req)
	if err != nil {
		return err
	}

	d.Set("ciphertext_blob", base64.StdEncoding.EncodeToString(resp.CiphertextBlob))
	d.SetId(resource.UniqueId())

	return nil
}

func resourceAwsKmsCiphertextRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsKmsCiphertextDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
