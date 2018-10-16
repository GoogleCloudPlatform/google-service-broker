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

// Code generated by go generate; DO NOT EDIT.

package db_service

import (
	"context"

	"github.com/GoogleCloudPlatform/gcp-service-broker/brokerapi/brokers/models"
)


// CountServiceInstanceDetailsById gets the count of ServiceInstanceDetails by its key (id) in the datastore (0 or 1)
func CountServiceInstanceDetailsById(ctx context.Context, id string) (int, error) { return defaultDatastore().CountServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) CountServiceInstanceDetailsById(ctx context.Context, id string) (int, error) {
	var count int
	err := ds.db.Model(&models.ServiceInstanceDetails{}).Where("id = ?", id).Count(&count).Error
	return count, err
}

// CreateServiceInstanceDetails creates a new record in the database and assigns it a primary key.
func CreateServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error { return defaultDatastore().CreateServiceInstanceDetails(ctx, object) }
func (ds *SqlDatastore) CreateServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error {
	return ds.db.Create(object).Error
}

// SaveServiceInstanceDetails updates an existing record in the database.
func SaveServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error { return defaultDatastore().SaveServiceInstanceDetails(ctx, object) }
func (ds *SqlDatastore) SaveServiceInstanceDetails(ctx context.Context, object *models.ServiceInstanceDetails) error {
	return ds.db.Save(object).Error
}
// DeleteServiceInstanceDetailsById soft-deletes the record by its key (id).
func DeleteServiceInstanceDetailsById(ctx context.Context, id string) error { return defaultDatastore().DeleteServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) DeleteServiceInstanceDetailsById(ctx context.Context, id string) error {
	return ds.db.Where("id = ?", id).Delete(&models.ServiceInstanceDetails{}).Error
}



// DeleteServiceInstanceDetails soft-deletes the record.
func DeleteServiceInstanceDetails(ctx context.Context, record *models.ServiceInstanceDetails) error { return defaultDatastore().DeleteServiceInstanceDetails(ctx, record) }
func (ds *SqlDatastore) DeleteServiceInstanceDetails(ctx context.Context, record *models.ServiceInstanceDetails) error {
	return ds.db.Delete(record).Error
}
// GetServiceInstanceDetailsById gets an instance of ServiceInstanceDetails by its key (id).
func GetServiceInstanceDetailsById(ctx context.Context, id string) (*models.ServiceInstanceDetails, error) { return defaultDatastore().GetServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) GetServiceInstanceDetailsById(ctx context.Context, id string) (*models.ServiceInstanceDetails, error) {
	record := models.ServiceInstanceDetails{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedServiceInstanceDetailsById checks to see if an instance of ServiceInstanceDetails was soft deleted by its key (id).
func CheckDeletedServiceInstanceDetailsById(ctx context.Context, id string) (bool, error) { return defaultDatastore().CheckDeletedServiceInstanceDetailsById(ctx, id) }
func (ds *SqlDatastore) CheckDeletedServiceInstanceDetailsById(ctx context.Context, id string) (bool, error) {
	record := models.ServiceInstanceDetails{}
	if err := ds.db.Unscoped().Where("id = ?", id).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}




// CountServiceBindingCredentialsByServiceInstanceIdAndBindingId gets the count of ServiceBindingCredentials by its key (serviceInstanceId, bindingId) in the datastore (0 or 1)
func CountServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (int, error) { return defaultDatastore().CountServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) CountServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (int, error) {
	var count int
	err := ds.db.Model(&models.ServiceBindingCredentials{}).Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).Count(&count).Error
	return count, err
}


// CountServiceBindingCredentialsByBindingId gets the count of ServiceBindingCredentials by its key (bindingId) in the datastore (0 or 1)
func CountServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (int, error) { return defaultDatastore().CountServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) CountServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (int, error) {
	var count int
	err := ds.db.Model(&models.ServiceBindingCredentials{}).Where("binding_id = ?", bindingId).Count(&count).Error
	return count, err
}


// CountServiceBindingCredentialsById gets the count of ServiceBindingCredentials by its key (id) in the datastore (0 or 1)
func CountServiceBindingCredentialsById(ctx context.Context, id uint) (int, error) { return defaultDatastore().CountServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) CountServiceBindingCredentialsById(ctx context.Context, id uint) (int, error) {
	var count int
	err := ds.db.Model(&models.ServiceBindingCredentials{}).Where("id = ?", id).Count(&count).Error
	return count, err
}

// CreateServiceBindingCredentials creates a new record in the database and assigns it a primary key.
func CreateServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error { return defaultDatastore().CreateServiceBindingCredentials(ctx, object) }
func (ds *SqlDatastore) CreateServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error {
	return ds.db.Create(object).Error
}

// SaveServiceBindingCredentials updates an existing record in the database.
func SaveServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error { return defaultDatastore().SaveServiceBindingCredentials(ctx, object) }
func (ds *SqlDatastore) SaveServiceBindingCredentials(ctx context.Context, object *models.ServiceBindingCredentials) error {
	return ds.db.Save(object).Error
}
// DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId soft-deletes the record by its key (serviceInstanceId, bindingId).
func DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) error { return defaultDatastore().DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) error {
	return ds.db.Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).Delete(&models.ServiceBindingCredentials{}).Error
}

// DeleteServiceBindingCredentialsByBindingId soft-deletes the record by its key (bindingId).
func DeleteServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) error { return defaultDatastore().DeleteServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) error {
	return ds.db.Where("binding_id = ?", bindingId).Delete(&models.ServiceBindingCredentials{}).Error
}

// DeleteServiceBindingCredentialsById soft-deletes the record by its key (id).
func DeleteServiceBindingCredentialsById(ctx context.Context, id uint) error { return defaultDatastore().DeleteServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) DeleteServiceBindingCredentialsById(ctx context.Context, id uint) error {
	return ds.db.Where("id = ?", id).Delete(&models.ServiceBindingCredentials{}).Error
}



// DeleteServiceBindingCredentials soft-deletes the record.
func DeleteServiceBindingCredentials(ctx context.Context, record *models.ServiceBindingCredentials) error { return defaultDatastore().DeleteServiceBindingCredentials(ctx, record) }
func (ds *SqlDatastore) DeleteServiceBindingCredentials(ctx context.Context, record *models.ServiceBindingCredentials) error {
	return ds.db.Delete(record).Error
}
// GetServiceBindingCredentialsByServiceInstanceIdAndBindingId gets an instance of ServiceBindingCredentials by its key (serviceInstanceId, bindingId).
func GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) GetServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedServiceBindingCredentialsByServiceInstanceIdAndBindingId checks to see if an instance of ServiceBindingCredentials was soft deleted by its key (serviceInstanceId, bindingId).
func CheckDeletedServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (bool, error) { return defaultDatastore().CheckDeletedServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx, serviceInstanceId, bindingId) }
func (ds *SqlDatastore) CheckDeletedServiceBindingCredentialsByServiceInstanceIdAndBindingId(ctx context.Context, serviceInstanceId string, bindingId string) (bool, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Unscoped().Where("service_instance_id = ? AND binding_id = ?", serviceInstanceId, bindingId).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}

// GetServiceBindingCredentialsByBindingId gets an instance of ServiceBindingCredentials by its key (bindingId).
func GetServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) GetServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("binding_id = ?", bindingId).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedServiceBindingCredentialsByBindingId checks to see if an instance of ServiceBindingCredentials was soft deleted by its key (bindingId).
func CheckDeletedServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (bool, error) { return defaultDatastore().CheckDeletedServiceBindingCredentialsByBindingId(ctx, bindingId) }
func (ds *SqlDatastore) CheckDeletedServiceBindingCredentialsByBindingId(ctx context.Context, bindingId string) (bool, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Unscoped().Where("binding_id = ?", bindingId).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}

// GetServiceBindingCredentialsById gets an instance of ServiceBindingCredentials by its key (id).
func GetServiceBindingCredentialsById(ctx context.Context, id uint) (*models.ServiceBindingCredentials, error) { return defaultDatastore().GetServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) GetServiceBindingCredentialsById(ctx context.Context, id uint) (*models.ServiceBindingCredentials, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedServiceBindingCredentialsById checks to see if an instance of ServiceBindingCredentials was soft deleted by its key (id).
func CheckDeletedServiceBindingCredentialsById(ctx context.Context, id uint) (bool, error) { return defaultDatastore().CheckDeletedServiceBindingCredentialsById(ctx, id) }
func (ds *SqlDatastore) CheckDeletedServiceBindingCredentialsById(ctx context.Context, id uint) (bool, error) {
	record := models.ServiceBindingCredentials{}
	if err := ds.db.Unscoped().Where("id = ?", id).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}




// CountProvisionRequestDetailsById gets the count of ProvisionRequestDetails by its key (id) in the datastore (0 or 1)
func CountProvisionRequestDetailsById(ctx context.Context, id uint) (int, error) { return defaultDatastore().CountProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) CountProvisionRequestDetailsById(ctx context.Context, id uint) (int, error) {
	var count int
	err := ds.db.Model(&models.ProvisionRequestDetails{}).Where("id = ?", id).Count(&count).Error
	return count, err
}

// CreateProvisionRequestDetails creates a new record in the database and assigns it a primary key.
func CreateProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error { return defaultDatastore().CreateProvisionRequestDetails(ctx, object) }
func (ds *SqlDatastore) CreateProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error {
	return ds.db.Create(object).Error
}

// SaveProvisionRequestDetails updates an existing record in the database.
func SaveProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error { return defaultDatastore().SaveProvisionRequestDetails(ctx, object) }
func (ds *SqlDatastore) SaveProvisionRequestDetails(ctx context.Context, object *models.ProvisionRequestDetails) error {
	return ds.db.Save(object).Error
}
// DeleteProvisionRequestDetailsById soft-deletes the record by its key (id).
func DeleteProvisionRequestDetailsById(ctx context.Context, id uint) error { return defaultDatastore().DeleteProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) DeleteProvisionRequestDetailsById(ctx context.Context, id uint) error {
	return ds.db.Where("id = ?", id).Delete(&models.ProvisionRequestDetails{}).Error
}



// DeleteProvisionRequestDetails soft-deletes the record.
func DeleteProvisionRequestDetails(ctx context.Context, record *models.ProvisionRequestDetails) error { return defaultDatastore().DeleteProvisionRequestDetails(ctx, record) }
func (ds *SqlDatastore) DeleteProvisionRequestDetails(ctx context.Context, record *models.ProvisionRequestDetails) error {
	return ds.db.Delete(record).Error
}
// GetProvisionRequestDetailsById gets an instance of ProvisionRequestDetails by its key (id).
func GetProvisionRequestDetailsById(ctx context.Context, id uint) (*models.ProvisionRequestDetails, error) { return defaultDatastore().GetProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) GetProvisionRequestDetailsById(ctx context.Context, id uint) (*models.ProvisionRequestDetails, error) {
	record := models.ProvisionRequestDetails{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedProvisionRequestDetailsById checks to see if an instance of ProvisionRequestDetails was soft deleted by its key (id).
func CheckDeletedProvisionRequestDetailsById(ctx context.Context, id uint) (bool, error) { return defaultDatastore().CheckDeletedProvisionRequestDetailsById(ctx, id) }
func (ds *SqlDatastore) CheckDeletedProvisionRequestDetailsById(ctx context.Context, id uint) (bool, error) {
	record := models.ProvisionRequestDetails{}
	if err := ds.db.Unscoped().Where("id = ?", id).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}




// CountPlanDetailsV1ByServiceIdAndName gets the count of PlanDetailsV1 by its key (serviceId, name) in the datastore (0 or 1)
func CountPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (int, error) { return defaultDatastore().CountPlanDetailsV1ByServiceIdAndName(ctx, serviceId, name) }
func (ds *SqlDatastore) CountPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (int, error) {
	var count int
	err := ds.db.Model(&models.PlanDetailsV1{}).Where("service_id = ? AND name = ?", serviceId, name).Count(&count).Error
	return count, err
}


// CountPlanDetailsV1ById gets the count of PlanDetailsV1 by its key (id) in the datastore (0 or 1)
func CountPlanDetailsV1ById(ctx context.Context, id string) (int, error) { return defaultDatastore().CountPlanDetailsV1ById(ctx, id) }
func (ds *SqlDatastore) CountPlanDetailsV1ById(ctx context.Context, id string) (int, error) {
	var count int
	err := ds.db.Model(&models.PlanDetailsV1{}).Where("id = ?", id).Count(&count).Error
	return count, err
}

// CreatePlanDetailsV1 creates a new record in the database and assigns it a primary key.
func CreatePlanDetailsV1(ctx context.Context, object *models.PlanDetailsV1) error { return defaultDatastore().CreatePlanDetailsV1(ctx, object) }
func (ds *SqlDatastore) CreatePlanDetailsV1(ctx context.Context, object *models.PlanDetailsV1) error {
	return ds.db.Create(object).Error
}

// SavePlanDetailsV1 updates an existing record in the database.
func SavePlanDetailsV1(ctx context.Context, object *models.PlanDetailsV1) error { return defaultDatastore().SavePlanDetailsV1(ctx, object) }
func (ds *SqlDatastore) SavePlanDetailsV1(ctx context.Context, object *models.PlanDetailsV1) error {
	return ds.db.Save(object).Error
}
// DeletePlanDetailsV1ByServiceIdAndName soft-deletes the record by its key (serviceId, name).
func DeletePlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) error { return defaultDatastore().DeletePlanDetailsV1ByServiceIdAndName(ctx, serviceId, name) }
func (ds *SqlDatastore) DeletePlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) error {
	return ds.db.Where("service_id = ? AND name = ?", serviceId, name).Delete(&models.PlanDetailsV1{}).Error
}

// DeletePlanDetailsV1ById soft-deletes the record by its key (id).
func DeletePlanDetailsV1ById(ctx context.Context, id string) error { return defaultDatastore().DeletePlanDetailsV1ById(ctx, id) }
func (ds *SqlDatastore) DeletePlanDetailsV1ById(ctx context.Context, id string) error {
	return ds.db.Where("id = ?", id).Delete(&models.PlanDetailsV1{}).Error
}



// DeletePlanDetailsV1 soft-deletes the record.
func DeletePlanDetailsV1(ctx context.Context, record *models.PlanDetailsV1) error { return defaultDatastore().DeletePlanDetailsV1(ctx, record) }
func (ds *SqlDatastore) DeletePlanDetailsV1(ctx context.Context, record *models.PlanDetailsV1) error {
	return ds.db.Delete(record).Error
}
// GetPlanDetailsV1ByServiceIdAndName gets an instance of PlanDetailsV1 by its key (serviceId, name).
func GetPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (*models.PlanDetailsV1, error) { return defaultDatastore().GetPlanDetailsV1ByServiceIdAndName(ctx, serviceId, name) }
func (ds *SqlDatastore) GetPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (*models.PlanDetailsV1, error) {
	record := models.PlanDetailsV1{}
	if err := ds.db.Where("service_id = ? AND name = ?", serviceId, name).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedPlanDetailsV1ByServiceIdAndName checks to see if an instance of PlanDetailsV1 was soft deleted by its key (serviceId, name).
func CheckDeletedPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (bool, error) { return defaultDatastore().CheckDeletedPlanDetailsV1ByServiceIdAndName(ctx, serviceId, name) }
func (ds *SqlDatastore) CheckDeletedPlanDetailsV1ByServiceIdAndName(ctx context.Context, serviceId string, name string) (bool, error) {
	record := models.PlanDetailsV1{}
	if err := ds.db.Unscoped().Where("service_id = ? AND name = ?", serviceId, name).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}

// GetPlanDetailsV1ById gets an instance of PlanDetailsV1 by its key (id).
func GetPlanDetailsV1ById(ctx context.Context, id string) (*models.PlanDetailsV1, error) { return defaultDatastore().GetPlanDetailsV1ById(ctx, id) }
func (ds *SqlDatastore) GetPlanDetailsV1ById(ctx context.Context, id string) (*models.PlanDetailsV1, error) {
	record := models.PlanDetailsV1{}
	if err := ds.db.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}

// CheckDeletedPlanDetailsV1ById checks to see if an instance of PlanDetailsV1 was soft deleted by its key (id).
func CheckDeletedPlanDetailsV1ById(ctx context.Context, id string) (bool, error) { return defaultDatastore().CheckDeletedPlanDetailsV1ById(ctx, id) }
func (ds *SqlDatastore) CheckDeletedPlanDetailsV1ById(ctx context.Context, id string) (bool, error) {
	record := models.PlanDetailsV1{}
	if err := ds.db.Unscoped().Where("id = ?", id).First(&record).Error; err != nil {
		return false, err
	}

	return record.DeletedAt != nil, nil
}


