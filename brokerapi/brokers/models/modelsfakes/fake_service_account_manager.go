// Code generated by counterfeiter. DO NOT EDIT.
package modelsfakes

import (
	"sync"

	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceAccountManager struct {
	CreateCredentialsStub        func(instanceID string, bindingID string, details brokerapi.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error)
	createCredentialsMutex       sync.RWMutex
	createCredentialsArgsForCall []struct {
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
		instance   models.ServiceInstanceDetails
	}
	createCredentialsReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	createCredentialsReturnsOnCall map[int]struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	DeleteCredentialsStub        func(creds models.ServiceBindingCredentials) error
	deleteCredentialsMutex       sync.RWMutex
	deleteCredentialsArgsForCall []struct {
		creds models.ServiceBindingCredentials
	}
	deleteCredentialsReturns struct {
		result1 error
	}
	deleteCredentialsReturnsOnCall map[int]struct {
		result1 error
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
	CreateAccountWithRolesStub        func(bindingID string, roles []string) (models.ServiceBindingCredentials, error)
	createAccountWithRolesMutex       sync.RWMutex
	createAccountWithRolesArgsForCall []struct {
		bindingID string
		roles     []string
	}
	createAccountWithRolesReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	createAccountWithRolesReturnsOnCall map[int]struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceAccountManager) CreateCredentials(instanceID string, bindingID string, details brokerapi.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error) {
	fake.createCredentialsMutex.Lock()
	ret, specificReturn := fake.createCredentialsReturnsOnCall[len(fake.createCredentialsArgsForCall)]
	fake.createCredentialsArgsForCall = append(fake.createCredentialsArgsForCall, struct {
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
		instance   models.ServiceInstanceDetails
	}{instanceID, bindingID, details, instance})
	fake.recordInvocation("CreateCredentials", []interface{}{instanceID, bindingID, details, instance})
	fake.createCredentialsMutex.Unlock()
	if fake.CreateCredentialsStub != nil {
		return fake.CreateCredentialsStub(instanceID, bindingID, details, instance)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCredentialsReturns.result1, fake.createCredentialsReturns.result2
}

func (fake *FakeServiceAccountManager) CreateCredentialsCallCount() int {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return len(fake.createCredentialsArgsForCall)
}

func (fake *FakeServiceAccountManager) CreateCredentialsArgsForCall(i int) (string, string, brokerapi.BindDetails, models.ServiceInstanceDetails) {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return fake.createCredentialsArgsForCall[i].instanceID, fake.createCredentialsArgsForCall[i].bindingID, fake.createCredentialsArgsForCall[i].details, fake.createCredentialsArgsForCall[i].instance
}

func (fake *FakeServiceAccountManager) CreateCredentialsReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateCredentialsStub = nil
	fake.createCredentialsReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) CreateCredentialsReturnsOnCall(i int, result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateCredentialsStub = nil
	if fake.createCredentialsReturnsOnCall == nil {
		fake.createCredentialsReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBindingCredentials
			result2 error
		})
	}
	fake.createCredentialsReturnsOnCall[i] = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) DeleteCredentials(creds models.ServiceBindingCredentials) error {
	fake.deleteCredentialsMutex.Lock()
	ret, specificReturn := fake.deleteCredentialsReturnsOnCall[len(fake.deleteCredentialsArgsForCall)]
	fake.deleteCredentialsArgsForCall = append(fake.deleteCredentialsArgsForCall, struct {
		creds models.ServiceBindingCredentials
	}{creds})
	fake.recordInvocation("DeleteCredentials", []interface{}{creds})
	fake.deleteCredentialsMutex.Unlock()
	if fake.DeleteCredentialsStub != nil {
		return fake.DeleteCredentialsStub(creds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteCredentialsReturns.result1
}

func (fake *FakeServiceAccountManager) DeleteCredentialsCallCount() int {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return len(fake.deleteCredentialsArgsForCall)
}

func (fake *FakeServiceAccountManager) DeleteCredentialsArgsForCall(i int) models.ServiceBindingCredentials {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return fake.deleteCredentialsArgsForCall[i].creds
}

func (fake *FakeServiceAccountManager) DeleteCredentialsReturns(result1 error) {
	fake.DeleteCredentialsStub = nil
	fake.deleteCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAccountManager) DeleteCredentialsReturnsOnCall(i int, result1 error) {
	fake.DeleteCredentialsStub = nil
	if fake.deleteCredentialsReturnsOnCall == nil {
		fake.deleteCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceAccountManager) BuildInstanceCredentials(bindRecord models.ServiceBindingCredentials, instanceRecord models.ServiceInstanceDetails) (map[string]string, error) {
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

func (fake *FakeServiceAccountManager) BuildInstanceCredentialsCallCount() int {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return len(fake.buildInstanceCredentialsArgsForCall)
}

func (fake *FakeServiceAccountManager) BuildInstanceCredentialsArgsForCall(i int) (models.ServiceBindingCredentials, models.ServiceInstanceDetails) {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.buildInstanceCredentialsArgsForCall[i].bindRecord, fake.buildInstanceCredentialsArgsForCall[i].instanceRecord
}

func (fake *FakeServiceAccountManager) BuildInstanceCredentialsReturns(result1 map[string]string, result2 error) {
	fake.BuildInstanceCredentialsStub = nil
	fake.buildInstanceCredentialsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) BuildInstanceCredentialsReturnsOnCall(i int, result1 map[string]string, result2 error) {
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

func (fake *FakeServiceAccountManager) CreateAccountWithRoles(bindingID string, roles []string) (models.ServiceBindingCredentials, error) {
	var rolesCopy []string
	if roles != nil {
		rolesCopy = make([]string, len(roles))
		copy(rolesCopy, roles)
	}
	fake.createAccountWithRolesMutex.Lock()
	ret, specificReturn := fake.createAccountWithRolesReturnsOnCall[len(fake.createAccountWithRolesArgsForCall)]
	fake.createAccountWithRolesArgsForCall = append(fake.createAccountWithRolesArgsForCall, struct {
		bindingID string
		roles     []string
	}{bindingID, rolesCopy})
	fake.recordInvocation("CreateAccountWithRoles", []interface{}{bindingID, rolesCopy})
	fake.createAccountWithRolesMutex.Unlock()
	if fake.CreateAccountWithRolesStub != nil {
		return fake.CreateAccountWithRolesStub(bindingID, roles)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAccountWithRolesReturns.result1, fake.createAccountWithRolesReturns.result2
}

func (fake *FakeServiceAccountManager) CreateAccountWithRolesCallCount() int {
	fake.createAccountWithRolesMutex.RLock()
	defer fake.createAccountWithRolesMutex.RUnlock()
	return len(fake.createAccountWithRolesArgsForCall)
}

func (fake *FakeServiceAccountManager) CreateAccountWithRolesArgsForCall(i int) (string, []string) {
	fake.createAccountWithRolesMutex.RLock()
	defer fake.createAccountWithRolesMutex.RUnlock()
	return fake.createAccountWithRolesArgsForCall[i].bindingID, fake.createAccountWithRolesArgsForCall[i].roles
}

func (fake *FakeServiceAccountManager) CreateAccountWithRolesReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateAccountWithRolesStub = nil
	fake.createAccountWithRolesReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) CreateAccountWithRolesReturnsOnCall(i int, result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateAccountWithRolesStub = nil
	if fake.createAccountWithRolesReturnsOnCall == nil {
		fake.createAccountWithRolesReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBindingCredentials
			result2 error
		})
	}
	fake.createAccountWithRolesReturnsOnCall[i] = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceAccountManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	fake.createAccountWithRolesMutex.RLock()
	defer fake.createAccountWithRolesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceAccountManager) recordInvocation(key string, args []interface{}) {
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

var _ models.ServiceAccountManager = new(FakeServiceAccountManager)
