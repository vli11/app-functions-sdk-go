//
// Copyright (c) 2019 Intel Corporation
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

package appsdk

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/stretchr/testify/assert"
)

func TestConfigurableDeviceNameFilter(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}
	params := make(map[string]string)

	tests := []struct {
		name      string
		key       string
		value     string
		expectNil bool
	}{
		{"Non Existent Parameters", "", "", true},
		{"Empty Parameters", DeviceNames, "", false},
		{"Valid Parameters", DeviceNames, "GS1-AC-Drive01, GS1-AC-Drive02, GS1-AC-Drive03", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params[tt.key] = tt.value
			trx := configurable.DeviceNameFilter(params)
			if tt.expectNil {
				assert.Nil(t, trx, "return result from DeviceNameFilter should be nil")
			} else {
				assert.NotNil(t, trx, "return result from DeviceNameFilter should not be nil")
			}
		})
	}
}

func TestConfigurableValueDescriptorFilter(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}
	params := make(map[string]string)

	tests := []struct {
		name      string
		key       string
		value     string
		expectNil bool
	}{
		{"Non Existent Parameters", "", "", true},
		{"Empty Parameters", ValueDescriptors, "", false},
		{"Valid Parameters", ValueDescriptors, "GS1-AC-Drive01, GS1-AC-Drive02, GS1-AC-Drive03", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params[tt.key] = tt.value
			trx := configurable.ValueDescriptorFilter(params)
			if tt.expectNil {
				assert.Nil(t, trx, "return result from ValueDescriptorFilter should be nil")
			} else {
				assert.NotNil(t, trx, "return result from ValueDescriptorFilter should not be nil")
			}
		})
	}
}

func TestConfigurableXMLTransform(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{}

	trx := configurable.XMLTransform()
	assert.NotNil(t, trx, "return result from XMLTransform should not be nil")
}

func TestConfigurableJSONTransform(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{}

	trx := configurable.JSONTransform()
	assert.NotNil(t, trx, "return result from JSONTransform should not be nil")
}

func TestConfigurableHTTPPost(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}

	params := make(map[string]string)

	// no url in params
	params[""] = "fake"
	trx := configurable.HTTPPost(params)
	assert.Nil(t, trx, "return result from HTTPPost should be nil")

	// no mime type
	params[Url] = "http://url"
	trx = configurable.HTTPPost(params)
	assert.Nil(t, trx, "return result from HTTPPost should be nil")

	params[MimeType] = ""
	trx = configurable.HTTPPost(params)
	assert.NotNil(t, trx, "return result from HTTPPost should not be nil")
}

func TestConfigurableHTTPPostJSON(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}

	params := make(map[string]string)

	// no url in params
	params[""] = ""
	trx := configurable.HTTPPostJSON(params)
	assert.Nil(t, trx, "return result from HTTPPostJSON should be nil")

	params[Url] = "http://url"
	trx = configurable.HTTPPostJSON(params)
	assert.NotNil(t, trx, "return result from HTTPPostJSON should not be nil")
}

func TestConfigurableHTTPPostXML(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}

	params := make(map[string]string)

	// no url in params
	params[""] = ""
	trx := configurable.HTTPPostXML(params)
	assert.Nil(t, trx, "return result from HTTPPostXML should be nil")

	params[Url] = "http://url"
	trx = configurable.HTTPPostXML(params)
	assert.NotNil(t, trx, "return result from HTTPPostXML should not be nil")
}

func TestConfigurableMQTTSend(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{
		Sdk: &AppFunctionsSDK{
			LoggingClient: lc,
		},
	}

	params := make(map[string]string)
	addr := models.Addressable{}
	trx := configurable.MQTTSend(params, addr)
	assert.NotNil(t, trx, "return result from MQTTSend should not be nil")
}

func TestConfigurableSetOutputData(t *testing.T) {
	configurable := AppFunctionsSDKConfigurable{}

	trx := configurable.SetOutputData()
	assert.NotNil(t, trx, "return result from SetOutputData should not be nil")
}
