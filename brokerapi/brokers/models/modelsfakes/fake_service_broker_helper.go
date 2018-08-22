// Code generated by counterfeiter. DO NOT EDIT.
package modelsfakes

import (
	"sync"

	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceBrokerHelper struct {
	ProvisionStub        func(instanceId string, details brokerapi.ProvisionDetails, plan models.ServicePlan) (models.ServiceInstanceDetails, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		instanceId string
		details    brokerapi.ProvisionDetails
		plan       models.ServicePlan
	}
	provisionReturns struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}
	BindStub        func(instanceID, bindingID string, details brokerapi.BindDetails) (models.ServiceBindingCredentials, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}
	bindReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	BuildInstanceCredentialsStub        func(bindRecord models.ServiceBindingCredentials, instanceRecord models.ServiceInstanceDetails) (map[string]string, error)
	buildInstanceCredentialsMutex       sync.RWMutex
	buildInstanceCredentialsArgsForCall []struct {
		bindRecord     models.ServiceBindingCredentials
		instanceRecord models.ServiceInstanceDetails
	}
	buildInstanceCredentialsReturns struct {
		result1 map[string]string
		result2 error
	}
	buildInstanceCredentialsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	UnbindStub        func(details models.ServiceBindingCredentials) error
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		details models.ServiceBindingCredentials
	}
	unbindReturns struct {
		result1 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 error
	}
	DeprovisionStub        func(instanceID string, details brokerapi.DeprovisionDetails) error
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		instanceID string
		details    brokerapi.DeprovisionDetails
	}
	deprovisionReturns struct {
		result1 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 error
	}
	PollInstanceStub        func(instanceID string) (bool, error)
	pollInstanceMutex       sync.RWMutex
	pollInstanceArgsForCall []struct {
		instanceID string
	}
	pollInstanceReturns struct {
		result1 bool
		result2 error
	}
	pollInstanceReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	LastOperationWasDeleteStub        func(instanceID string) (bool, error)
	lastOperationWasDeleteMutex       sync.RWMutex
	lastOperationWasDeleteArgsForCall []struct {
		instanceID string
	}
	lastOperationWasDeleteReturns struct {
		result1 bool
		result2 error
	}
	lastOperationWasDeleteReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ProvisionsAsyncStub        func() bool
	provisionsAsyncMutex       sync.RWMutex
	provisionsAsyncArgsForCall []struct{}
	provisionsAsyncReturns     struct {
		result1 bool
	}
	provisionsAsyncReturnsOnCall map[int]struct {
		result1 bool
	}
	DeprovisionsAsyncStub        func() bool
	deprovisionsAsyncMutex       sync.RWMutex
	deprovisionsAsyncArgsForCall []struct{}
	deprovisionsAsyncReturns     struct {
		result1 bool
	}
	deprovisionsAsyncReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBrokerHelper) Provision(instanceId string, details brokerapi.ProvisionDetails, plan models.ServicePlan) (models.ServiceInstanceDetails, error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		instanceId string
		details    brokerapi.ProvisionDetails
		plan       models.ServicePlan
	}{instanceId, details, plan})
	fake.recordInvocation("Provision", []interface{}{instanceId, details, plan})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(instanceId, details, plan)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.provisionReturns.result1, fake.provisionReturns.result2
}

func (fake *FakeServiceBrokerHelper) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) ProvisionArgsForCall(i int) (string, brokerapi.ProvisionDetails, models.ServicePlan) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].instanceId, fake.provisionArgsForCall[i].details, fake.provisionArgsForCall[i].plan
}

func (fake *FakeServiceBrokerHelper) ProvisionReturns(result1 models.ServiceInstanceDetails, result2 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) ProvisionReturnsOnCall(i int, result1 models.ServiceInstanceDetails, result2 error) {
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 models.ServiceInstanceDetails
			result2 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 models.ServiceInstanceDetails
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Bind(instanceID string, bindingID string, details brokerapi.BindDetails) (models.ServiceBindingCredentials, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}{instanceID, bindingID, details})
	fake.recordInvocation("Bind", []interface{}{instanceID, bindingID, details})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(instanceID, bindingID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}

func (fake *FakeServiceBrokerHelper) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BindArgsForCall(i int) (string, string, brokerapi.BindDetails) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].instanceID, fake.bindArgsForCall[i].bindingID, fake.bindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) BindReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BindReturnsOnCall(i int, result1 models.ServiceBindingCredentials, result2 error) {
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBindingCredentials
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentials(bindRecord models.ServiceBindingCredentials, instanceRecord models.ServiceInstanceDetails) (map[string]string, error) {
	fake.buildInstanceCredentialsMutex.Lock()
	ret, specificReturn := fake.buildInstanceCredentialsReturnsOnCall[len(fake.buildInstanceCredentialsArgsForCall)]
	fake.buildInstanceCredentialsArgsForCall = append(fake.buildInstanceCredentialsArgsForCall, struct {
		bindRecord     models.ServiceBindingCredentials
		instanceRecord models.ServiceInstanceDetails
	}{bindRecord, instanceRecord})
	fake.recordInvocation("BuildInstanceCredentials", []interface{}{bindRecord, instanceRecord})
	fake.buildInstanceCredentialsMutex.Unlock()
	if fake.BuildInstanceCredentialsStub != nil {
		return fake.BuildInstanceCredentialsStub(bindRecord, instanceRecord)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildInstanceCredentialsReturns.result1, fake.buildInstanceCredentialsReturns.result2
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsCallCount() int {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return len(fake.buildInstanceCredentialsArgsForCall)
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsArgsForCall(i int) (models.ServiceBindingCredentials, models.ServiceInstanceDetails) {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.buildInstanceCredentialsArgsForCall[i].bindRecord, fake.buildInstanceCredentialsArgsForCall[i].instanceRecord
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsReturns(result1 map[string]string, result2 error) {
	fake.BuildInstanceCredentialsStub = nil
	fake.buildInstanceCredentialsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) BuildInstanceCredentialsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.BuildInstanceCredentialsStub = nil
	if fake.buildInstanceCredentialsReturnsOnCall == nil {
		fake.buildInstanceCredentialsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.buildInstanceCredentialsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) Unbind(details models.ServiceBindingCredentials) error {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		details models.ServiceBindingCredentials
	}{details})
	fake.recordInvocation("Unbind", []interface{}{details})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(details)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unbindReturns.result1
}

func (fake *FakeServiceBrokerHelper) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeServiceBrokerHelper) UnbindArgsForCall(i int) models.ServiceBindingCredentials {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) UnbindReturnsOnCall(i int, result1 error) {
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Deprovision(instanceID string, details brokerapi.DeprovisionDetails) error {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		instanceID string
		details    brokerapi.DeprovisionDetails
	}{instanceID, details})
	fake.recordInvocation("Deprovision", []interface{}{instanceID, details})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(instanceID, details)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deprovisionReturns.result1
}

func (fake *FakeServiceBrokerHelper) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeServiceBrokerHelper) DeprovisionArgsForCall(i int) (string, brokerapi.DeprovisionDetails) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].instanceID, fake.deprovisionArgsForCall[i].details
}

func (fake *FakeServiceBrokerHelper) DeprovisionReturns(result1 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) DeprovisionReturnsOnCall(i int, result1 error) {
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerHelper) PollInstance(instanceID string) (bool, error) {
	fake.pollInstanceMutex.Lock()
	ret, specificReturn := fake.pollInstanceReturnsOnCall[len(fake.pollInstanceArgsForCall)]
	fake.pollInstanceArgsForCall = append(fake.pollInstanceArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("PollInstance", []interface{}{instanceID})
	fake.pollInstanceMutex.Unlock()
	if fake.PollInstanceStub != nil {
		return fake.PollInstanceStub(instanceID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.pollInstanceReturns.result1, fake.pollInstanceReturns.result2
}

func (fake *FakeServiceBrokerHelper) PollInstanceCallCount() int {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return len(fake.pollInstanceArgsForCall)
}

func (fake *FakeServiceBrokerHelper) PollInstanceArgsForCall(i int) string {
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	return fake.pollInstanceArgsForCall[i].instanceID
}

func (fake *FakeServiceBrokerHelper) PollInstanceReturns(result1 bool, result2 error) {
	fake.PollInstanceStub = nil
	fake.pollInstanceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) PollInstanceReturnsOnCall(i int, result1 bool, result2 error) {
	fake.PollInstanceStub = nil
	if fake.pollInstanceReturnsOnCall == nil {
		fake.pollInstanceReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.pollInstanceReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) LastOperationWasDelete(instanceID string) (bool, error) {
	fake.lastOperationWasDeleteMutex.Lock()
	ret, specificReturn := fake.lastOperationWasDeleteReturnsOnCall[len(fake.lastOperationWasDeleteArgsForCall)]
	fake.lastOperationWasDeleteArgsForCall = append(fake.lastOperationWasDeleteArgsForCall, struct {
		instanceID string
	}{instanceID})
	fake.recordInvocation("LastOperationWasDelete", []interface{}{instanceID})
	fake.lastOperationWasDeleteMutex.Unlock()
	if fake.LastOperationWasDeleteStub != nil {
		return fake.LastOperationWasDeleteStub(instanceID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastOperationWasDeleteReturns.result1, fake.lastOperationWasDeleteReturns.result2
}

func (fake *FakeServiceBrokerHelper) LastOperationWasDeleteCallCount() int {
	fake.lastOperationWasDeleteMutex.RLock()
	defer fake.lastOperationWasDeleteMutex.RUnlock()
	return len(fake.lastOperationWasDeleteArgsForCall)
}

func (fake *FakeServiceBrokerHelper) LastOperationWasDeleteArgsForCall(i int) string {
	fake.lastOperationWasDeleteMutex.RLock()
	defer fake.lastOperationWasDeleteMutex.RUnlock()
	return fake.lastOperationWasDeleteArgsForCall[i].instanceID
}

func (fake *FakeServiceBrokerHelper) LastOperationWasDeleteReturns(result1 bool, result2 error) {
	fake.LastOperationWasDeleteStub = nil
	fake.lastOperationWasDeleteReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) LastOperationWasDeleteReturnsOnCall(i int, result1 bool, result2 error) {
	fake.LastOperationWasDeleteStub = nil
	if fake.lastOperationWasDeleteReturnsOnCall == nil {
		fake.lastOperationWasDeleteReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.lastOperationWasDeleteReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsync() bool {
	fake.provisionsAsyncMutex.Lock()
	ret, specificReturn := fake.provisionsAsyncReturnsOnCall[len(fake.provisionsAsyncArgsForCall)]
	fake.provisionsAsyncArgsForCall = append(fake.provisionsAsyncArgsForCall, struct{}{})
	fake.recordInvocation("ProvisionsAsync", []interface{}{})
	fake.provisionsAsyncMutex.Unlock()
	if fake.ProvisionsAsyncStub != nil {
		return fake.ProvisionsAsyncStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.provisionsAsyncReturns.result1
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncCallCount() int {
	fake.provisionsAsyncMutex.RLock()
	defer fake.provisionsAsyncMutex.RUnlock()
	return len(fake.provisionsAsyncArgsForCall)
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncReturns(result1 bool) {
	fake.ProvisionsAsyncStub = nil
	fake.provisionsAsyncReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) ProvisionsAsyncReturnsOnCall(i int, result1 bool) {
	fake.ProvisionsAsyncStub = nil
	if fake.provisionsAsyncReturnsOnCall == nil {
		fake.provisionsAsyncReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.provisionsAsyncReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsync() bool {
	fake.deprovisionsAsyncMutex.Lock()
	ret, specificReturn := fake.deprovisionsAsyncReturnsOnCall[len(fake.deprovisionsAsyncArgsForCall)]
	fake.deprovisionsAsyncArgsForCall = append(fake.deprovisionsAsyncArgsForCall, struct{}{})
	fake.recordInvocation("DeprovisionsAsync", []interface{}{})
	fake.deprovisionsAsyncMutex.Unlock()
	if fake.DeprovisionsAsyncStub != nil {
		return fake.DeprovisionsAsyncStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deprovisionsAsyncReturns.result1
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncCallCount() int {
	fake.deprovisionsAsyncMutex.RLock()
	defer fake.deprovisionsAsyncMutex.RUnlock()
	return len(fake.deprovisionsAsyncArgsForCall)
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncReturns(result1 bool) {
	fake.DeprovisionsAsyncStub = nil
	fake.deprovisionsAsyncReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) DeprovisionsAsyncReturnsOnCall(i int, result1 bool) {
	fake.DeprovisionsAsyncStub = nil
	if fake.deprovisionsAsyncReturnsOnCall == nil {
		fake.deprovisionsAsyncReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.deprovisionsAsyncReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeServiceBrokerHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.pollInstanceMutex.RLock()
	defer fake.pollInstanceMutex.RUnlock()
	fake.lastOperationWasDeleteMutex.RLock()
	defer fake.lastOperationWasDeleteMutex.RUnlock()
	fake.provisionsAsyncMutex.RLock()
	defer fake.provisionsAsyncMutex.RUnlock()
	fake.deprovisionsAsyncMutex.RLock()
	defer fake.deprovisionsAsyncMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceBrokerHelper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ models.ServiceBrokerHelper = new(FakeServiceBrokerHelper)
