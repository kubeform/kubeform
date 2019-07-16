/*
Copyright 2019 The Kubeform Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// LicensemanagerLicenseConfigurationLister helps list LicensemanagerLicenseConfigurations.
type LicensemanagerLicenseConfigurationLister interface {
	// List lists all LicensemanagerLicenseConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerLicenseConfiguration, err error)
	// Get retrieves the LicensemanagerLicenseConfiguration from the index for a given name.
	Get(name string) (*v1alpha1.LicensemanagerLicenseConfiguration, error)
	LicensemanagerLicenseConfigurationListerExpansion
}

// licensemanagerLicenseConfigurationLister implements the LicensemanagerLicenseConfigurationLister interface.
type licensemanagerLicenseConfigurationLister struct {
	indexer cache.Indexer
}

// NewLicensemanagerLicenseConfigurationLister returns a new LicensemanagerLicenseConfigurationLister.
func NewLicensemanagerLicenseConfigurationLister(indexer cache.Indexer) LicensemanagerLicenseConfigurationLister {
	return &licensemanagerLicenseConfigurationLister{indexer: indexer}
}

// List lists all LicensemanagerLicenseConfigurations in the indexer.
func (s *licensemanagerLicenseConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.LicensemanagerLicenseConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LicensemanagerLicenseConfiguration))
	})
	return ret, err
}

// Get retrieves the LicensemanagerLicenseConfiguration from the index for a given name.
func (s *licensemanagerLicenseConfigurationLister) Get(name string) (*v1alpha1.LicensemanagerLicenseConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("licensemanagerlicenseconfiguration"), name)
	}
	return obj.(*v1alpha1.LicensemanagerLicenseConfiguration), nil
}
