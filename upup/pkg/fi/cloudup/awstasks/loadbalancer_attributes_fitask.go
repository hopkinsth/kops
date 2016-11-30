/*
Copyright 2016 The Kubernetes Authors.

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

// Code generated by ""fitask" -type=LoadBalancer"; DO NOT EDIT

package awstasks

import "encoding/json"

// LoadBalancer

// JSON marshalling boilerplate
type realLoadBalancerAttributes LoadBalancerAttributes
type realLoadBalancerConnectionSettings LoadBalancerConnectionSettings

func (o *LoadBalancerAttributes) UnmarshalJSON(data []byte) error {
	var r realLoadBalancerAttributes
	var c realLoadBalancerConnectionSettings
	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	*o = LoadBalancerAttributes(r)
	*o.ConnectionSettings = LoadBalancerConnectionSettings(c)
	return nil
}
