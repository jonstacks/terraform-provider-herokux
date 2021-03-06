package herokux

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strings"
	"testing"
)

func TestAccHerokuxPostgresDataLink_Basic(t *testing.T) {
	localID := testAccConfig.GetAddonIDorSkip(t)
	remoteName := testAccConfig.GetDBNameorSkip(t)
	name := fmt.Sprintf("tftest_%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHerokuxPostgresDataLink_WithCustomName(localID, remoteName, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "local_db_id", localID),
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "remote_db_name", remoteName),
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "name", name),
					resource.TestCheckResourceAttrSet(
						"herokux_postgres_data_link.foobar", "link_id"),
					resource.TestCheckResourceAttrSet(
						"herokux_postgres_data_link.foobar", "remote_attachment_name"),
				),
			},
		},
	})
}

func TestAccHerokuxPostgresDataLink_BasicNoCustomName(t *testing.T) {
	localID := testAccConfig.GetAddonIDorSkip(t)
	remoteName := testAccConfig.GetDBNameorSkip(t)
	expectedLinkName := strings.ReplaceAll(remoteName, "-", "_")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHerokuxPostgresDataLink_NoCustomName(localID, remoteName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "local_db_id", localID),
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "remote_db_name", remoteName),
					resource.TestCheckResourceAttr(
						"herokux_postgres_data_link.foobar", "name", expectedLinkName),
					resource.TestCheckResourceAttrSet(
						"herokux_postgres_data_link.foobar", "link_id"),
					resource.TestCheckResourceAttrSet(
						"herokux_postgres_data_link.foobar", "remote_attachment_name"),
				),
			},
		},
	})
}

func testAccCheckHerokuxPostgresDataLink_WithCustomName(localID, remoteName, name string) string {
	return fmt.Sprintf(`
resource "herokux_postgres_data_link" "foobar" {
	local_db_id = "%s"
	remote_db_name = "%s"
	name = "%s"
}
`, localID, remoteName, name)
}

func testAccCheckHerokuxPostgresDataLink_NoCustomName(localID, remoteName string) string {
	return fmt.Sprintf(`
resource "herokux_postgres_data_link" "foobar" {
	local_db_id = "%s"
	remote_db_name = "%s"
}
`, localID, remoteName)
}
