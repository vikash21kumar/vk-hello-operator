/*
Copyright 2021.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VKHelloOperatorSpec defines the desired state of VKHelloOperator
type VKHelloOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of VKHelloOperator. Edit vkhellooperator_types.go to remove/update
	Greetings string `json:"foo,omitempty"`
}

// VKHelloOperatorStatus defines the observed state of VKHelloOperator
type VKHelloOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	TheGreetings string `json:"foo,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VKHelloOperator is the Schema for the vkhellooperators API
type VKHelloOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VKHelloOperatorSpec   `json:"spec,omitempty"`
	Status VKHelloOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VKHelloOperatorList contains a list of VKHelloOperator
type VKHelloOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VKHelloOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VKHelloOperator{}, &VKHelloOperatorList{})
}
