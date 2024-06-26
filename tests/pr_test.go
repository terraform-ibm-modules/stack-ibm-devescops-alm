package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
	"testing"
)

func TestProjectsFullTest(t *testing.T) {

	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing:        t,
		ResourceGroup:  "default",
		Prefix:         "alm", // setting prefix here gets a random string appended to it
		ParallelDeploy: false,
	})

	options.StackInputs = map[string]interface{}{
		"prefix":                      options.Prefix,
		"resource_group_name":         options.Prefix,
		"sm_service_plan":             "trial",
		"scc_service_plan":            "security-compliance-center-standard-plan",
		"region":                      "us-south",
		"use_existing_resource_group": false,
		"ibmcloud_api_key":            options.RequiredEnvironmentVars["TF_VAR_ibmcloud_api_key"], // always required by the stack
	}

	err := options.RunProjectsTest()
	if assert.NoError(t, err) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}
