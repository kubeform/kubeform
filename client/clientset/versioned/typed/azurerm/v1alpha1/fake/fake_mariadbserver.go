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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMariadbServers implements MariadbServerInterface
type FakeMariadbServers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mariadbserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mariadbservers"}

var mariadbserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MariadbServer"}

// Get takes name of the mariadbServer, and returns the corresponding mariadbServer object, and an error if there is any.
func (c *FakeMariadbServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mariadbserversResource, c.ns, name), &v1alpha1.MariadbServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbServer), err
}

// List takes label and field selectors, and returns the list of MariadbServers that match those selectors.
func (c *FakeMariadbServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MariadbServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mariadbserversResource, mariadbserversKind, c.ns, opts), &v1alpha1.MariadbServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MariadbServerList{ListMeta: obj.(*v1alpha1.MariadbServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.MariadbServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mariadbServers.
func (c *FakeMariadbServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mariadbserversResource, c.ns, opts))

}

// Create takes the representation of a mariadbServer and creates it.  Returns the server's representation of the mariadbServer, and an error, if there is any.
func (c *FakeMariadbServers) Create(ctx context.Context, mariadbServer *v1alpha1.MariadbServer, opts v1.CreateOptions) (result *v1alpha1.MariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mariadbserversResource, c.ns, mariadbServer), &v1alpha1.MariadbServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbServer), err
}

// Update takes the representation of a mariadbServer and updates it. Returns the server's representation of the mariadbServer, and an error, if there is any.
func (c *FakeMariadbServers) Update(ctx context.Context, mariadbServer *v1alpha1.MariadbServer, opts v1.UpdateOptions) (result *v1alpha1.MariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mariadbserversResource, c.ns, mariadbServer), &v1alpha1.MariadbServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMariadbServers) UpdateStatus(ctx context.Context, mariadbServer *v1alpha1.MariadbServer, opts v1.UpdateOptions) (*v1alpha1.MariadbServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mariadbserversResource, "status", c.ns, mariadbServer), &v1alpha1.MariadbServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbServer), err
}

// Delete takes name of the mariadbServer and deletes it. Returns an error if one occurs.
func (c *FakeMariadbServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mariadbserversResource, c.ns, name), &v1alpha1.MariadbServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMariadbServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mariadbserversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MariadbServerList{})
	return err
}

// Patch applies the patch and returns the patched mariadbServer.
func (c *FakeMariadbServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MariadbServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mariadbserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.MariadbServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbServer), err
}
