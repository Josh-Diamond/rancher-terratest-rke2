package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRke2DownSteamCluster(t *testing.T) {
	t.Parallel()

	expectedClusterName := "dadfish"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		TerraformDir: "../modules/terraform-rke2",
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
	actualClusterName := terraform.Output(t, terraformOptions, "my_rke2_name")

	assert.Equal(t, expectedClusterName, actualClusterName)

}
