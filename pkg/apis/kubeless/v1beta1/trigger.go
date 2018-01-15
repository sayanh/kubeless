/*
Copyright (c) 2016-2017 Bitnami

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

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Trigger object
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TriggerSpec `json:"spec"`
}

// TriggerSpec contains Trigger specification
type TriggerSpec struct {
	Type                    string                          `json:"type"`                  // Trigger type
	Topic                   string                          `json:"topic"`                 // Trigger topic (for PubSub type)
	Schedule                string                          `json:"schedule"`              // Scheduled time (for Schedule type)
	ServiceSpec             v1.ServiceSpec                  `json:"service"`	           // Specification on how service to be exposed in case HTTP trigger type
	FunctionName            string                          `json:"function-name"`         // Function name associated with the trigger
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TriggerList contains map of triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of third party objects
	Items []*Trigger `json:"items"`
}