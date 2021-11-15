/*
Copyright The KubeEdge Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "restandxorm/v1alpha2"
)

// FakeDevices implements DeviceInterface
type FakeDevices struct {
	Fake *FakeDevicesV1alpha2
	ns   string
}

var devicesResource = schema.GroupVersionResource{Group: "devices", Version: "v1alpha2", Resource: "devices"}

var devicesKind = schema.GroupVersionKind{Group: "devices", Version: "v1alpha2", Kind: "Device"}

// Get takes name of the device, and returns the corresponding device object, and an error if there is any.
func (c *FakeDevices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Device, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devicesResource, c.ns, name), &v1alpha2.Device{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Device), err
}

// List takes label and field selectors, and returns the list of Devices that match those selectors.
func (c *FakeDevices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.DeviceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devicesResource, devicesKind, c.ns, opts), &v1alpha2.DeviceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.DeviceList{ListMeta: obj.(*v1alpha2.DeviceList).ListMeta}
	for _, item := range obj.(*v1alpha2.DeviceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devices.
func (c *FakeDevices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devicesResource, c.ns, opts))

}

// Create takes the representation of a device and creates it.  Returns the server's representation of the device, and an error, if there is any.
func (c *FakeDevices) Create(ctx context.Context, device *v1alpha2.Device, opts v1.CreateOptions) (result *v1alpha2.Device, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devicesResource, c.ns, device), &v1alpha2.Device{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Device), err
}

// Update takes the representation of a device and updates it. Returns the server's representation of the device, and an error, if there is any.
func (c *FakeDevices) Update(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (result *v1alpha2.Device, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devicesResource, c.ns, device), &v1alpha2.Device{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Device), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevices) UpdateStatus(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (*v1alpha2.Device, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devicesResource, "status", c.ns, device), &v1alpha2.Device{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Device), err
}

// Delete takes name of the device and deletes it. Returns an error if one occurs.
func (c *FakeDevices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devicesResource, c.ns, name), &v1alpha2.Device{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.DeviceList{})
	return err
}

// Patch applies the patch and returns the patched device.
func (c *FakeDevices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Device, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devicesResource, c.ns, name, pt, data, subresources...), &v1alpha2.Device{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Device), err
}
