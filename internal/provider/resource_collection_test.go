package provider

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCollectionConfig(t *testing.T) {
	var o collection.Collection
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCollectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
		},
	})
}

func TestAccCollectionNetwork(t *testing.T) {
	var o collection.Collection
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCollectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
		},
	})
}

func TestAccCollectionAuditEvent(t *testing.T) {
	var o collection.Collection
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCollectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
			{
				Config: testAccCollectionConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectionExists("prismacloudcompute_collection.test", &o),
					testAccCheckCollectionAttributes(&o, name, "collection created with Terraform", "#FF0000"),
				),
			},
		},
	})
}

func testAccCheckCollectionExists(n string, o *collection.Collection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label Name is not set")
		}

		client := testAccProvider.Meta().(*api.Client)
		name := rs.Primary.ID
		lo, err := collection.GetCollection(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = *lo

		return nil
	}
}

func testAccCheckCollectionAttributes(o *collection.Collection, name string, description string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// if o.Name != name {
		// 	return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		// } else {
		// 	fmt.Printf("\n\nName is %s", o.Name)
		// }

		if o.Description != description {
			return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		}

		if o.Color != color {
			return fmt.Errorf("Color type is %q, expected %q", o.Color, color)
		}

		return nil
	}
}

func testAccCollectionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_collection" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if err := collection.DeleteCollection(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}

	return nil
}

func testAccCollectionConfig(name string) string {
	return fmt.Sprintf(`
	resource "prismacloudcompute_collection" "test" {
		name              = "%s"
		description       = "collection created with Terraform"
		color             = "#FF0000"
		application_ids   = ["*"]
		account_ids		  = ["*"]
		clusters		  = ["*"]
		containers	  	  = ["*"]
		code_repositories = ["*"]
		functions		  = ["*"]
		hosts			  = ["*"]
		images            = ["*"]
		labels            = ["*"]
		namespaces        = ["*"]
	  }	  
	`, name)
}
