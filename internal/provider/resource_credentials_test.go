package provider

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCredentialsConfig(t *testing.T) {
	var o auth.Credential
	id := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestCredentialsNetwork(t *testing.T) {
	var o auth.Credential
	id := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
		},
	})
}

func TestCredentialsAuditEvent(t *testing.T) {
	var o auth.Credential
	id := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
			{
				Config: testAccCredentialsConfig(id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialExists("prismacloudcompute_credential.test", &o),
					testAccCheckCredentialsAttributes(&o, id, true),
				),
			},
		},
	})
}

func testAccCheckCredentialExists(n string, o *auth.Credential) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Id is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		id := rs.Primary.ID
		lo, err := auth.GetCredential(*client, id)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}

		*o = *lo

		return nil
	}
}

func testAccCheckCredentialsAttributes(o *auth.Credential, id string, learningDisabled bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Id != id {
			return fmt.Errorf("\n\nId is %s, expected %s", o.Id, id)
		}
		// else {
		// 	fmt.Printf("\n\nId is %s", o.Id)
		// }

		return nil
	}
}

func testAccCredentialsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_credential" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := auth.DeleteCredential(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}

	return nil
}

func testAccCredentialsConfig(id string) string {
	return fmt.Sprintf(`
	resource "prismacloudcompute_credential" "test" {
		name			= "%s"
		type		= "basic"
		account_id	= "tf-test-credential"
		api_token {

		}
		secret {
			plain	  = "testpw"
		}
	}
	`, id)
}
