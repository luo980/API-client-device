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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	v1alpha2 "restandxorm/v1alpha2"
)

type DevicesV1alpha2Interface interface {
	RESTClient() rest.Interface
	DevicesGetter
	DeviceModelsGetter
}

// DevicesV1alpha2Client is used to interact with features provided by the devices group.
type DevicesV1alpha2Client struct {
	RestClient rest.Interface
}

func (c *DevicesV1alpha2Client) Devices(namespace string) DeviceInterface {
	return newDevices(c, namespace)
}

func (c *DevicesV1alpha2Client) DeviceModels(namespace string) DeviceModelInterface {
	return newDeviceModels(c, namespace)
}

// NewForConfig creates a new DevicesV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*DevicesV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DevicesV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new DevicesV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DevicesV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DevicesV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *DevicesV1alpha2Client {
	return &DevicesV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.NewCodecFactory(runtime.NewScheme()).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DevicesV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.RestClient
}
