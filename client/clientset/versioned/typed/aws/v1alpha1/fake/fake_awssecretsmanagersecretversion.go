/*
Copyright 2019 The KubeDB Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeAwsSecretsmanagerSecretVersions implements AwsSecretsmanagerSecretVersionInterface
type FakeAwsSecretsmanagerSecretVersions struct {
	Fake *FakeAwsV1alpha1
}

var awssecretsmanagersecretversionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awssecretsmanagersecretversions"}

var awssecretsmanagersecretversionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsSecretsmanagerSecretVersion"}

// Get takes name of the awsSecretsmanagerSecretVersion, and returns the corresponding awsSecretsmanagerSecretVersion object, and an error if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awssecretsmanagersecretversionsResource, name), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecretsmanagerSecretVersion), err
}

// List takes label and field selectors, and returns the list of AwsSecretsmanagerSecretVersions that match those selectors.
func (c *FakeAwsSecretsmanagerSecretVersions) List(opts v1.ListOptions) (result *v1alpha1.AwsSecretsmanagerSecretVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awssecretsmanagersecretversionsResource, awssecretsmanagersecretversionsKind, opts), &v1alpha1.AwsSecretsmanagerSecretVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSecretsmanagerSecretVersionList{ListMeta: obj.(*v1alpha1.AwsSecretsmanagerSecretVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSecretsmanagerSecretVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSecretsmanagerSecretVersions.
func (c *FakeAwsSecretsmanagerSecretVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awssecretsmanagersecretversionsResource, opts))
}

// Create takes the representation of a awsSecretsmanagerSecretVersion and creates it.  Returns the server's representation of the awsSecretsmanagerSecretVersion, and an error, if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Create(awsSecretsmanagerSecretVersion *v1alpha1.AwsSecretsmanagerSecretVersion) (result *v1alpha1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awssecretsmanagersecretversionsResource, awsSecretsmanagerSecretVersion), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecretsmanagerSecretVersion), err
}

// Update takes the representation of a awsSecretsmanagerSecretVersion and updates it. Returns the server's representation of the awsSecretsmanagerSecretVersion, and an error, if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Update(awsSecretsmanagerSecretVersion *v1alpha1.AwsSecretsmanagerSecretVersion) (result *v1alpha1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awssecretsmanagersecretversionsResource, awsSecretsmanagerSecretVersion), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecretsmanagerSecretVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSecretsmanagerSecretVersions) UpdateStatus(awsSecretsmanagerSecretVersion *v1alpha1.AwsSecretsmanagerSecretVersion) (*v1alpha1.AwsSecretsmanagerSecretVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awssecretsmanagersecretversionsResource, "status", awsSecretsmanagerSecretVersion), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecretsmanagerSecretVersion), err
}

// Delete takes name of the awsSecretsmanagerSecretVersion and deletes it. Returns an error if one occurs.
func (c *FakeAwsSecretsmanagerSecretVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awssecretsmanagersecretversionsResource, name), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSecretsmanagerSecretVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awssecretsmanagersecretversionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSecretsmanagerSecretVersionList{})
	return err
}

// Patch applies the patch and returns the patched awsSecretsmanagerSecretVersion.
func (c *FakeAwsSecretsmanagerSecretVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awssecretsmanagersecretversionsResource, name, pt, data, subresources...), &v1alpha1.AwsSecretsmanagerSecretVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSecretsmanagerSecretVersion), err
}
