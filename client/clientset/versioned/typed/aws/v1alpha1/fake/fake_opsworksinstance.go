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

// FakeOpsworksInstances implements OpsworksInstanceInterface
type FakeOpsworksInstances struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksinstancesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksinstances"}

var opsworksinstancesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksInstance"}

// Get takes name of the opsworksInstance, and returns the corresponding opsworksInstance object, and an error if there is any.
func (c *FakeOpsworksInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksinstancesResource, c.ns, name), &v1alpha1.OpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksInstance), err
}

// List takes label and field selectors, and returns the list of OpsworksInstances that match those selectors.
func (c *FakeOpsworksInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpsworksInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksinstancesResource, opsworksinstancesKind, c.ns, opts), &v1alpha1.OpsworksInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksInstanceList{ListMeta: obj.(*v1alpha1.OpsworksInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksInstances.
func (c *FakeOpsworksInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksinstancesResource, c.ns, opts))

}

// Create takes the representation of a opsworksInstance and creates it.  Returns the server's representation of the opsworksInstance, and an error, if there is any.
func (c *FakeOpsworksInstances) Create(ctx context.Context, opsworksInstance *v1alpha1.OpsworksInstance, opts v1.CreateOptions) (result *v1alpha1.OpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksinstancesResource, c.ns, opsworksInstance), &v1alpha1.OpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksInstance), err
}

// Update takes the representation of a opsworksInstance and updates it. Returns the server's representation of the opsworksInstance, and an error, if there is any.
func (c *FakeOpsworksInstances) Update(ctx context.Context, opsworksInstance *v1alpha1.OpsworksInstance, opts v1.UpdateOptions) (result *v1alpha1.OpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksinstancesResource, c.ns, opsworksInstance), &v1alpha1.OpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksInstances) UpdateStatus(ctx context.Context, opsworksInstance *v1alpha1.OpsworksInstance, opts v1.UpdateOptions) (*v1alpha1.OpsworksInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksinstancesResource, "status", c.ns, opsworksInstance), &v1alpha1.OpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksInstance), err
}

// Delete takes name of the opsworksInstance and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksinstancesResource, c.ns, name), &v1alpha1.OpsworksInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksInstanceList{})
	return err
}

// Patch applies the patch and returns the patched opsworksInstance.
func (c *FakeOpsworksInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpsworksInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksInstance), err
}
