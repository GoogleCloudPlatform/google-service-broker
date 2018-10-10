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

package brokers

import (
	"context"
	"fmt"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/account_managers"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/api_service"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/bigquery"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/bigtable"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/broker_base"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/cloudsql"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/stackdriver_profiler"
	"github.com/pivotal-cf/brokerapi"
	"google.golang.org/api/googleapi"

	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/config"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/datastore"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/pubsub"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/spanner"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/stackdriver_debugger"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/stackdriver_trace"
	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/storage"
	"github.com/GoogleCloudPlatform/gcp-service-broker/db_service"
)

// GCPServiceBroker is a brokerapi.ServiceBroker that can be used to generate an OSB compatible service broker.
type GCPServiceBroker struct {
	Catalog          map[string]models.Service
	ServiceBrokerMap map[string]models.ServiceBrokerHelper

	Logger lager.Logger
}

// New creates a GCPServiceBroker.
// Exactly one of GCPServiceBroker or error will be nil when returned.
func New(cfg *config.BrokerConfig, logger lager.Logger) (*GCPServiceBroker, error) {

	self := GCPServiceBroker{}
	self.Logger = logger
	self.Catalog = cfg.Catalog

	saManager := &account_managers.ServiceAccountManager{
		HttpConfig: cfg.HttpConfig,
		ProjectId:  cfg.ProjectId,
	}

	sqlManager := &account_managers.SqlAccountManager{
		HttpConfig: cfg.HttpConfig,
		ProjectId:  cfg.ProjectId,
	}

	bb := broker_base.BrokerBase{
		AccountManager: saManager,
		HttpConfig:     cfg.HttpConfig,
		ProjectId:      cfg.ProjectId,
		Logger:         self.Logger,
	}

	// map service specific brokers to general broker
	self.ServiceBrokerMap = map[string]models.ServiceBrokerHelper{
		models.StorageName: &storage.StorageBroker{
			BrokerBase: bb,
		},
		models.PubsubName: &pubsub.PubSubBroker{
			BrokerBase: bb,
		},
		models.StackdriverDebuggerName: &stackdriver_debugger.StackdriverDebuggerBroker{
			BrokerBase: bb,
		},
		models.StackdriverProfilerName: &stackdriver_profiler.StackdriverProfilerBroker{
			BrokerBase: bb,
		},
		models.StackdriverTraceName: &stackdriver_trace.StackdriverTraceBroker{
			BrokerBase: bb,
		},
		models.BigqueryName: &bigquery.BigQueryBroker{
			BrokerBase: bb,
		},
		models.MlName: &api_service.ApiServiceBroker{
			BrokerBase: bb,
		},
		models.CloudsqlMySQLName: &cloudsql.CloudSQLBroker{
			HttpConfig:       cfg.HttpConfig,
			ProjectId:        cfg.ProjectId,
			Logger:           self.Logger,
			AccountManager:   sqlManager,
			SaAccountManager: saManager,
		},
		models.CloudsqlPostgresName: &cloudsql.CloudSQLBroker{
			HttpConfig:       cfg.HttpConfig,
			ProjectId:        cfg.ProjectId,
			Logger:           self.Logger,
			AccountManager:   sqlManager,
			SaAccountManager: saManager,
		},
		models.BigtableName: &bigtable.BigTableBroker{
			BrokerBase: bb,
		},
		models.SpannerName: &spanner.SpannerBroker{
			BrokerBase: bb,
		},
		models.DatastoreName: &datastore.DatastoreBroker{
			BrokerBase: bb,
		},
	}

	// replace the mapping from name to a mapping from id
	for _, service := range self.Catalog {
		self.ServiceBrokerMap[service.ID] = self.ServiceBrokerMap[service.Name]
		delete(self.ServiceBrokerMap, service.Name)
	}

	return &self, nil
}

// Services lists services in the broker's catalog.
// It is called through the `GET /v2/catalog` endpoint or the `cf marketplace` command.
func (gcpBroker *GCPServiceBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	svcs := []brokerapi.Service{}

	for _, svc := range gcpBroker.Catalog {
		svcs = append(svcs, svc.ToPlain())
	}

	return svcs, nil
}

func (gcpBroker *GCPServiceBroker) getPlanFromId(serviceId, planId string) (models.ServicePlan, error) {
	service, serviceOk := gcpBroker.Catalog[serviceId]
	if !serviceOk {
		return models.ServicePlan{}, fmt.Errorf("unknown service id: %q", serviceId)
	}

	for _, plan := range service.Plans {
		if plan.ID == planId {
			return plan, nil
		}
	}

	return models.ServicePlan{}, fmt.Errorf("unknown plan id: %q", planId)
}

// Provision creates a new instance of a service.
// It is bound to the `PUT /v2/service_instances/:instance_id` endpoint and can be called using the `cf create-service` command.
func (gcpBroker *GCPServiceBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, clientSupportsAsync bool) (brokerapi.ProvisionedServiceSpec, error) {
	gcpBroker.Logger.Info("Provisioning", lager.Data{
		"instanceId":         instanceID,
		"accepts_incomplete": clientSupportsAsync,
		"details":            details,
	})

	// make sure that instance hasn't already been provisioned
	count, err := db_service.CountServiceInstanceDetailsById(ctx, instanceID)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, fmt.Errorf("Database error checking for existing instance: %s", err)
	}
	if count > 0 {
		return brokerapi.ProvisionedServiceSpec{}, brokerapi.ErrInstanceAlreadyExists
	}

	serviceId := details.ServiceID

	// verify the service exists and
	plan, err := gcpBroker.getPlanFromId(serviceId, details.PlanID)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}

	// verify async provisioning is allowed if it is required
	service := gcpBroker.ServiceBrokerMap[serviceId]
	shouldProvisionAsync := service.ProvisionsAsync()
	if shouldProvisionAsync && !clientSupportsAsync {
		return brokerapi.ProvisionedServiceSpec{}, brokerapi.ErrAsyncRequired
	}

	// get instance details
	instanceDetails, err := service.Provision(ctx, instanceID, details, plan)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}

	// save instance details
	instanceDetails.ServiceId = serviceId
	instanceDetails.ID = instanceID
	instanceDetails.PlanId = details.PlanID
	instanceDetails.SpaceGuid = details.SpaceGUID
	instanceDetails.OrganizationGuid = details.OrganizationGUID

	err = db_service.CreateServiceInstanceDetails(ctx, &instanceDetails)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, fmt.Errorf("Error saving instance details to database: %s. WARNING: this instance cannot be deprovisioned through cf. Contact your operator for cleanup", err)
	}

	// save provision request details
	pr := models.ProvisionRequestDetails{
		ServiceInstanceId: instanceID,
		RequestDetails:    string(details.RawParameters),
	}
	if err = db_service.CreateProvisionRequestDetails(ctx, &pr); err != nil {
		return brokerapi.ProvisionedServiceSpec{}, fmt.Errorf("Error saving provision request details to database: %s. Services relying on async provisioning will not be able to complete provisioning", err)
	}

	return brokerapi.ProvisionedServiceSpec{IsAsync: shouldProvisionAsync, DashboardURL: "", OperationData: instanceDetails.OperationId}, nil
}

// Deprovision destroys an existing instance of a service.
// It is bound to the `DELETE /v2/service_instances/:instance_id` endpoint and can be called using the `cf delete-service` command.
func (gcpBroker *GCPServiceBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, clientSupportsAsync bool) (brokerapi.DeprovisionServiceSpec, error) {
	gcpBroker.Logger.Info("Deprovisioning", lager.Data{
		"instance_id":        instanceID,
		"accepts_incomplete": clientSupportsAsync,
		"details":            details,
	})

	service := gcpBroker.ServiceBrokerMap[details.ServiceID]
	deprovisionIsAsync := service.DeprovisionsAsync()
	response := brokerapi.DeprovisionServiceSpec{IsAsync: deprovisionIsAsync}

	// make sure that instance actually exists
	count, err := db_service.CountServiceInstanceDetailsById(ctx, instanceID)
	if err != nil {
		return response, fmt.Errorf("Database error checking for existing instance: %s", err)
	}
	if count == 0 {
		return response, brokerapi.ErrInstanceDoesNotExist
	}

	// if async deprovisioning isn't allowed but this service needs it, throw an error
	if deprovisionIsAsync && !clientSupportsAsync {
		return brokerapi.DeprovisionServiceSpec{}, brokerapi.ErrAsyncRequired
	}

	// deprovision
	instance, err := db_service.GetServiceInstanceDetailsById(ctx, instanceID)
	if err != nil {
		return response, brokerapi.NewFailureResponseBuilder(err, http.StatusInternalServerError, "fetching instance from database").Build()
	}

	if err := service.Deprovision(ctx, *instance, details); err != nil {
		return response, err
	}

	// soft-delete instance details from the db if this is a synchronous operation
	// if it's an async operation we can't delete from the db until we're sure delete succeeded, so this is
	// handled internally to LastOperation
	if !deprovisionIsAsync {
		err = db_service.DeleteServiceInstanceDetailsById(ctx, instanceID)
		if err != nil {
			return response, fmt.Errorf("Error deleting instance details from database: %s. WARNING: this instance will remain visible in cf. Contact your operator for cleanup", err)
		}
	}

	return response, nil
}

// Bind creates an account with credentials to access an instance of a service.
// It is bound to the `PUT /v2/service_instances/:instance_id/service_bindings/:binding_id` endpoint and can be called using the `cf bind-service` command.
func (gcpBroker *GCPServiceBroker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	gcpBroker.Logger.Info("Binding", lager.Data{
		"instance_id": instanceID,
		"binding_id":  bindingID,
		"details":     details,
	})

	service := gcpBroker.ServiceBrokerMap[details.ServiceID]

	// check for existing binding

	count, err := db_service.CountServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, instanceID, bindingID)
	if err != nil {
		return brokerapi.Binding{}, fmt.Errorf("Error checking for existing binding: %s", err)
	}
	if count > 0 {
		return brokerapi.Binding{}, brokerapi.ErrBindingAlreadyExists
	}

	// create binding
	newCreds, err := service.Bind(ctx, instanceID, bindingID, details)
	if err != nil {
		return brokerapi.Binding{}, err
	}

	// save binding to database
	newCreds.ServiceInstanceId = instanceID
	newCreds.BindingId = bindingID
	newCreds.ServiceId = details.ServiceID

	if err := db_service.CreateServiceBindingCredentials(ctx, &newCreds); err != nil {
		return brokerapi.Binding{}, fmt.Errorf("Error saving credentials to database: %s. WARNING: these credentials cannot be unbound through cf. Please contact your operator for cleanup",
			err)
	}

	// get existing service instance details
	instanceRecord, err := db_service.GetServiceInstanceDetailsById(ctx, instanceID)
	if err != nil {
		return brokerapi.Binding{}, fmt.Errorf("Error retrieving service instance details: %s", err)
	}

	updatedCreds, err := service.BuildInstanceCredentials(ctx, newCreds, *instanceRecord)
	if err != nil {
		return brokerapi.Binding{}, err
	}

	return brokerapi.Binding{Credentials: updatedCreds}, nil
}

// Unbind destroys an account and credentials with access to an instance of a service.
// It is bound to the `DELETE /v2/service_instances/:instance_id/service_bindings/:binding_id` endpoint and can be called using the `cf unbind-service` command.
func (gcpBroker *GCPServiceBroker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	gcpBroker.Logger.Info("Unbinding", lager.Data{
		"instance_id": instanceID,
		"binding_id":  bindingID,
		"details":     details,
	})

	service := gcpBroker.ServiceBrokerMap[details.ServiceID]

	// validate existence of binding
	existingBinding, err := db_service.GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, instanceID, bindingID)
	if err != nil {
		return brokerapi.ErrBindingDoesNotExist
	}

	// remove binding from Google
	if err := service.Unbind(ctx, *existingBinding); err != nil {
		return err
	}

	// remove binding from database
	if err := db_service.DeleteServiceBindingCredentials(ctx, existingBinding); err != nil {
		return fmt.Errorf("Error soft-deleting credentials from database: %s. WARNING: these credentials will remain visible in cf. Contact your operator for cleanup", err)
	}

	return nil
}

// Unbind destroys an account and credentials with access to an instance of a service.
// It is bound to the `GET /v2/service_instances/:instance_id/last_operation` endpoint.
// It is called by `cf create-service` or `cf delete-service` if the operation was asynchronous.
func (gcpBroker *GCPServiceBroker) LastOperation(ctx context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	gcpBroker.Logger.Info("Last Operation", lager.Data{
		"instance_id":    instanceID,
		"operation_data": operationData,
	})

	instance, err := db_service.GetServiceInstanceDetailsById(ctx, instanceID)
	if err != nil {
		return brokerapi.LastOperation{}, brokerapi.ErrInstanceDoesNotExist
	}

	service := gcpBroker.ServiceBrokerMap[instance.ServiceId]
	isAsyncService := service.ProvisionsAsync() || service.DeprovisionsAsync()

	if isAsyncService {
		return gcpBroker.lastOperationAsync(ctx, instanceID, service)
	}

	return brokerapi.LastOperation{}, brokerapi.ErrAsyncRequired
}

func (gcpBroker *GCPServiceBroker) lastOperationAsync(ctx context.Context, instanceId string, service models.ServiceBrokerHelper) (brokerapi.LastOperation, error) {
	done, err := service.PollInstance(ctx, instanceId)
	if err != nil {
		// this is a retryable error
		if gerr, ok := err.(*googleapi.Error); ok {
			if gerr.Code == 503 {
				return brokerapi.LastOperation{State: brokerapi.InProgress}, err
			}
		}
		// This is not a retryable error. Return fail
		return brokerapi.LastOperation{State: brokerapi.Failed}, err
	}

	if !done {
		return brokerapi.LastOperation{State: brokerapi.InProgress}, nil
	}

	// no error and we're done! Delete from the SB database if this was a delete flow and return success
	deleteFlow, err := service.LastOperationWasDelete(ctx, instanceId)
	if err != nil {
		return brokerapi.LastOperation{State: brokerapi.Succeeded}, fmt.Errorf("Couldn't determine if provision or deprovision flow, this may leave orphaned resources, contact your operator for cleanup: %s", err)
	}
	if deleteFlow {
		err = db_service.DeleteServiceInstanceDetailsById(ctx, instanceId)
		if err != nil {
			return brokerapi.LastOperation{State: brokerapi.Succeeded}, fmt.Errorf("Error deleting instance details from database: %s. WARNING: this instance will remain visible in cf. Contact your operator for cleanup", err)
		}
	}
	return brokerapi.LastOperation{State: brokerapi.Succeeded}, nil
}

// Update a service instance plan.
// This functionality is not implemented and will return an error indicating that plan changes are not supported.
func (gcpBroker *GCPServiceBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, brokerapi.ErrPlanChangeNotSupported
}
