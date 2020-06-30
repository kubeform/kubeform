/*
Copyright The Kubeform Authors.

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
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDbInstances implements DbInstanceInterface
type FakeDbInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dbinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dbinstances"}

var dbinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DbInstance"}

// Get takes name of the dbInstance, and returns the corresponding dbInstance object, and an error if there is any.
func (c *FakeDbInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbinstancesResource, c.ns, name), &v1alpha1.DbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbInstance), err
}

// List takes label and field selectors, and returns the list of DbInstances that match those selectors.
func (c *FakeDbInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DbInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbinstancesResource, dbinstancesKind, c.ns, opts), &v1alpha1.DbInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DbInstanceList{ListMeta: obj.(*v1alpha1.DbInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DbInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbInstances.
func (c *FakeDbInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbinstancesResource, c.ns, opts))

}

// Create takes the representation of a dbInstance and creates it.  Returns the server's representation of the dbInstance, and an error, if there is any.
func (c *FakeDbInstances) Create(ctx context.Context, dbInstance *v1alpha1.DbInstance, opts v1.CreateOptions) (result *v1alpha1.DbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbinstancesResource, c.ns, dbInstance), &v1alpha1.DbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbInstance), err
}

// Update takes the representation of a dbInstance and updates it. Returns the server's representation of the dbInstance, and an error, if there is any.
func (c *FakeDbInstances) Update(ctx context.Context, dbInstance *v1alpha1.DbInstance, opts v1.UpdateOptions) (result *v1alpha1.DbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbinstancesResource, c.ns, dbInstance), &v1alpha1.DbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbInstances) UpdateStatus(ctx context.Context, dbInstance *v1alpha1.DbInstance, opts v1.UpdateOptions) (*v1alpha1.DbInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbinstancesResource, "status", c.ns, dbInstance), &v1alpha1.DbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbInstance), err
}

// Delete takes name of the dbInstance and deletes it. Returns an error if one occurs.
func (c *FakeDbInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dbinstancesResource, c.ns, name), &v1alpha1.DbInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DbInstanceList{})
	return err
}

// Patch applies the patch and returns the patched dbInstance.
func (c *FakeDbInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DbInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DbInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DbInstance), err
}
