package provider

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/rule"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCustomRuleConfig(t *testing.T) {
	var o rule.CustomRule
	name := fmt.Sprintf("test-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomRuleNetwork(t *testing.T) {
	var o rule.CustomRule
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func TestAccCustomRuleAuditEvent(t *testing.T) {
	var o rule.CustomRule
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCustomRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
			{
				Config: testAccCustomRuleConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCustomRuleExists("prismacloudcompute_custom_rule.test", &o),
					testAccCheckCustomRuleAttributes(&o, name, "description", "#000000"),
				),
			},
		},
	})
}

func testAccCheckCustomRuleExists(n string, o *rule.CustomRule) resource.TestCheckFunc {
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
		lo, err := rule.GetCustomRuleByName(*client, name)
		if err != nil {
			return fmt.Errorf("Error in get: %s", err)
		}
		*o = *lo

		return nil
	}
}

func testAccCheckCustomRuleAttributes(o *rule.CustomRule, name string, description string, color string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("\n\nName is %s, expected %s", o.Name, name)
		}
		// else {
		// 	fmt.Printf("\n\nName is %s", o.Name)
		// }

		if o.Description != description {
			return fmt.Errorf("Description is %s, expected %s", o.Description, description)
		}

		return nil
	}
}

func testAccCustomRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*api.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type != "prismacloudcompute_custom_rule" {
			continue
		}

		if rs.Primary.ID != "" {
			name := rs.Primary.ID
			if _, err := rule.GetCustomRuleByName(*client, name); err == nil {
				return fmt.Errorf("Object %q still exists", name)
			}
		}
		return nil
	}

	return nil
}

func testAccCustomRuleConfig(name string) string {
	return fmt.Sprintf(`
	resource "prismacloudcompute_custom_rule" "test" {
		name = "%s"
		type = "waas-request"
		script = "req.headers[\"Range\"] contains /(?i)bytes[\\S\\s]*-[\\S\\s]*\\d{16,}/"
		description = "description"
		message = "Created by Terraform"
		min_version = "21.08.525"
		vuln_ids = [
			"CVE-2015-1635"
		]
	}`, name)
}
