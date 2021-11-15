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
package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "restandxorm/scheme"
	v1alpha2 "restandxorm/v1alpha2"
)

// DevicesGetter has a method to return a DeviceInterface.
// A group's client should implement this interface.
type DevicesGetter interface {
	Devices(namespace string) DeviceInterface
}

// DeviceInterface has methods to work with Device resources.
type DeviceInterface interface {
	Create(ctx context.Context, device *v1alpha2.Device, opts v1.CreateOptions) (*v1alpha2.Device, error)
	Update(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (*v1alpha2.Device, error)
	UpdateStatus(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (*v1alpha2.Device, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.Device, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.DeviceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Device, err error)
	DeviceExpansion
}

// devices implements DeviceInterface
type devices struct {
	client rest.Interface
	ns     string
}

// newDevices returns a Devices
func newDevices(c *DevicesV1alpha2Client, namespace string) *devices {
	return &devices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the device, and returns the corresponding device object, and an error if there is any.
func (c *devices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Device, err error) {
	result = &v1alpha2.Device{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Devices that match those selectors.
func (c *devices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.DeviceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.DeviceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devices.
func (c *devices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a device and creates it.  Returns the server's representation of the device, and an error, if there is any.
func (c *devices) Create(ctx context.Context, device *v1alpha2.Device, opts v1.CreateOptions) (result *v1alpha2.Device, err error) {
	result = &v1alpha2.Device{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(device).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a device and updates it. Returns the server's representation of the device, and an error, if there is any.
func (c *devices) Update(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (result *v1alpha2.Device, err error) {
	result = &v1alpha2.Device{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devices").
		Name(device.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(device).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *devices) UpdateStatus(ctx context.Context, device *v1alpha2.Device, opts v1.UpdateOptions) (result *v1alpha2.Device, err error) {
	result = &v1alpha2.Device{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devices").
		Name(device.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(device).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the device and deletes it. Returns an error if one occurs.
func (c *devices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched device.
func (c *devices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Device, err error) {
	result = &v1alpha2.Device{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
