/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package quota

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/openshift/rosa/pkg/arguments"
	"github.com/openshift/rosa/pkg/aws"
	"github.com/openshift/rosa/pkg/logging"
	"github.com/openshift/rosa/pkg/ocm"
	rprtr "github.com/openshift/rosa/pkg/reporter"
)

var Cmd = &cobra.Command{
	Use:   "quota",
	Short: "Verify AWS quota is ok for cluster install",
	Long:  "Verify AWS quota needed to create a cluster is configured as expected",
	Example: `  # Verify AWS quotas are configured correctly
  rosa verify quota

  # Verify AWS quotas in a different region
  rosa verify quota --region=us-west-2`,
	RunE: run,
}

func init() {
	flags := Cmd.Flags()

	arguments.AddRegionFlag(flags)
	arguments.AddProfileFlag(flags)
}

func run(cmd *cobra.Command, _ []string) (err error) {
	reporter := rprtr.CreateReporterOrExit()
	logger := logging.CreateLoggerOrExit(reporter)

	// Get AWS region
	region, err := aws.GetRegion(arguments.GetRegion())
	if err != nil {
		reporter.Errorf("Error getting region: %v", err)
		return err
	}

	// Create the client for the OCM API:
	ocmClient, err := ocm.NewClient().
		Logger(logger).
		Build()
	if err != nil {
		reporter.Errorf("Failed to create OCM connection: %v", err)
		return err
	}
	defer ocmClient.Close()

	// Create the AWS client:
	client, err := aws.NewClient().
		Logger(logger).
		Region(region).
		Build()
	if err != nil {
		// FIXME Hack to capture errors due to using STS accounts
		if strings.Contains(fmt.Sprintf("%s", err), "STS") {
			ocmClient.LogEvent("ROSAInitCredentialsSTS", nil)
		}
		reporter.Errorf("Error creating AWS client: %v", err)
		return err
	}

	if reporter.IsTerminal() {
		reporter.Infof("Validating AWS quota...")
	}
	_, err = client.ValidateQuota()
	if err != nil {
		ocmClient.LogEvent("ROSAVerifyQuotaInsufficient", nil)
		reporter.Errorf("Insufficient AWS quotas")
		reporter.Errorf("%v", err)
		return err
	}
	if reporter.IsTerminal() {
		reporter.Infof("AWS quota ok. " +
			"If cluster installation fails, validate actual AWS resource usage against " +
			"https://docs.openshift.com/rosa/rosa_getting_started/rosa-required-aws-service-quotas.html")
	}
	return nil
}
