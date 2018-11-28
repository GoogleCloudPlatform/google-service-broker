// Copyright 2018 the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"log"

	"github.com/GoogleCloudPlatform/gcp-service-broker/pkg/broker"
	"github.com/GoogleCloudPlatform/gcp-service-broker/pkg/client"
	"github.com/GoogleCloudPlatform/gcp-service-broker/utils"
	"github.com/spf13/cobra"
)

var (
	serviceId      string
	planId         string
	instanceId     string
	bindingId      string
	parametersJson string

	serviceName string
)

func init() {

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "A CLI client for the service broker",
		Long: `A CLI client for the service broker.

The client commands use the same configuration values as the server and operate
on localhost using the HTTP protocol.

Configuration Params:

 - api.user
 - api.password
 - api.port

Environment Variables:

 - GSB_API_USER
 - GSB_API_PASSWORD
 - GSB_API_PORT

The client commands return formatted JSON when run if the exit code is 0:

	{
	    "url": "http://user:pass@localhost:8000/v2/catalog",
	    "http_method": "GET",
	    "status_code": 200,
	    "response": // Response Body as JSON
	}

Exit codes DO NOT correspond with status_code, if a request was made and the
response could be parsed then the exit code will be 0.
Non-zero exit codes indicate a failure in the executable.

Because of the format, you can use the client to do automated testing of your
user-defined plans.
`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(clientCmd)

	clientCatalogCmd := newClientCommand("catalog", "Show the service catalog", func(client *client.Client) *client.BrokerResponse {
		return client.Catalog()
	})

	provisionCmd := newClientCommand("provision", "Provision a service", func(client *client.Client) *client.BrokerResponse {
		return client.Provision(instanceId, serviceId, planId, json.RawMessage(parametersJson))
	})

	deprovisionCmd := newClientCommand("deprovision", "Derovision a service", func(client *client.Client) *client.BrokerResponse {
		return client.Deprovision(instanceId, serviceId, planId)
	})

	bindCmd := newClientCommand("bind", "Bind to a service", func(client *client.Client) *client.BrokerResponse {
		return client.Bind(instanceId, bindingId, serviceId, planId, json.RawMessage(parametersJson))
	})

	unbindCmd := newClientCommand("unbind", "Unbind a service", func(client *client.Client) *client.BrokerResponse {
		return client.Unbind(instanceId, bindingId, serviceId, planId)
	})

	lastCmd := newClientCommand("last", "Get the status of the last operation", func(client *client.Client) *client.BrokerResponse {
		return client.LastOperation(instanceId)
	})

	updateCmd := newClientCommand("update", "Update the instance details", func(client *client.Client) *client.BrokerResponse {
		return client.Update(instanceId, serviceId, planId, json.RawMessage(parametersJson))
	})

	runExamplesCmd := &cobra.Command{
		Use:   "run-examples",
		Short: "Run all examples in the use command.",
		Long: `Run all examples generated by the use command through a
	provision/bind/unbind/deprovision cycle.

	Exits with a 0 if all examples were successful, 1 otherwise.`,
		Run: func(cmd *cobra.Command, args []string) {
			apiClient, err := client.NewClientFromEnv()
			if err != nil {
				log.Fatalf("Error creating client: %v", err)
			}

			if err := client.RunExamplesForService(broker.DefaultRegistry, apiClient, serviceName); err != nil {
				log.Fatalf("Error executing examples: %v", err)
			}

			log.Println("Success")
		},
	}

	clientCmd.AddCommand(clientCatalogCmd, provisionCmd, deprovisionCmd, bindCmd, unbindCmd, lastCmd, runExamplesCmd, updateCmd)

	bindFlag := func(dest *string, name, description string, commands ...*cobra.Command) {
		for _, sc := range commands {
			sc.Flags().StringVarP(dest, name, "", "", description)
			sc.MarkFlagRequired(name)
		}
	}

	bindFlag(&instanceId, "instanceid", "id of the service instance to operate on (user defined)", provisionCmd, deprovisionCmd, bindCmd, unbindCmd, lastCmd, updateCmd)
	bindFlag(&serviceId, "serviceid", "GUID of the service instanceid references (see catalog)", provisionCmd, deprovisionCmd, bindCmd, unbindCmd, updateCmd)
	bindFlag(&planId, "planid", "GUID of the service instanceid references (see catalog entry for the associated serviceid)", provisionCmd, deprovisionCmd, bindCmd, unbindCmd, updateCmd)
	bindFlag(&bindingId, "bindingid", "GUID of the binding to work on (user defined)", bindCmd, unbindCmd)

	for _, sc := range []*cobra.Command{provisionCmd, bindCmd, updateCmd} {
		sc.Flags().StringVarP(&parametersJson, "params", "", "{}", "JSON string of user-defined parameters to pass to the request")
	}

	runExamplesCmd.Flags().StringVarP(&serviceName, "service-name", "", "", "name of the service to run tests for")
}

func newClientCommand(use, short string, run func(*client.Client) *client.BrokerResponse) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			apiClient, err := client.NewClientFromEnv()
			if err != nil {
				log.Fatalf("Could not create API client: %s", err)
			}

			results := run(apiClient)
			utils.PrettyPrintOrExit(results)
		},
	}
}
