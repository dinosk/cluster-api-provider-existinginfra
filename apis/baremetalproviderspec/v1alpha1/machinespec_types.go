/*
Copyright 2020 Weaveworks.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// MachineSpec is the Schema for the machinespecs API
type MachineSpec struct {
	metav1.TypeMeta `json:",inline"`
	// This ObjectMeta is not stored on encode, it is here just to provide
	// support for annotations that are required for comment preservation
	metav1.ObjectMeta `json:"-"`

	Address          string   `json:"address"`
	Port             uint16   `json:"port,omitempty"`
	PrivateAddress   string   `json:"privateAddress,omitempty"`
	PrivateInterface string   `json:"privateInterface,omitempty"`
	Private          EndPoint `json:"private,omitempty"`
	Public           EndPoint `json:"public,omitempty"`
}

// EndPoint groups the details required to establish a connection.
type EndPoint struct {
	Address string `json:"address"`
	Port    uint16 `json:"port"`
}

func init() {
	localSchemeBuilder.Register(addKnownMachineTypes)
}
