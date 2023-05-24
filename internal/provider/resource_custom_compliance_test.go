package provider

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCustomComplianceConfig(t *testing.T) {
	var o policy.CustomCompliance
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomComplianceNetwork(t *testing.T) {
	var o policy.CustomCompliance
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomComplianceAuditEvent(t *testing.T) {
	var o policy.CustomCompliance
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomComplianceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomComplianceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomComplianceExists("prismacloudcompute_custom_compliance.test", &o),
					testAccCheckCustomComplianceAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func testAccCheckCustomComplianceExists(n string, o *policy.CustomCompliance) resource.TestCheckFunc {
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
		lo, err := policy.GetCustomComplianceByName(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = *lo

		return nil
	}
}

func testAccCheckCustomComplianceAttributes(o *policy.CustomCompliance, name string, description string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// if o.Name != name {
		// 	return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		// } else {
		// 	fmt.Printf("\n\nName is %s", o.Name)
		// }

		return nil
	}
}

func testAccCustomComplianceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_custom_compliance" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if _, err := policy.GetCustomComplianceByName(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}

	return nil
}

func testAccCustomComplianceConfig(name string) string {
	return fmt.Sprintf(`
	resource "prismacloudcompute_custom_compliance" "test" {
		name = %q
		title = "test compliance check"
		script = "if [ ! -f /tmp/foo.txt ]; then\n    echo \"File not found!\"\n    exit 1\nfi"
		severity = "high"
	}`, name)
}
