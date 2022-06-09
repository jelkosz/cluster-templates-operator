/*
Copyright 2022.

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

type Quota struct {
	HelmRepositoryRef string `json:"helmRepositoryRef"`
	Count             int32  `json:"count"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClusterTemplateQuotaSpec defines the desired state of ClusterTemplateQuota
type ClusterTemplateQuotaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ClusterTemplateQuota. Edit clustertemplatequota_types.go to remove/update
	Quota map[string]Quota `json:"quota,omitempty"`
}

// ClusterTemplateQuotaStatus defines the observed state of ClusterTemplateQuota
type ClusterTemplateQuotaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Quota map[string]Quota `json:"quota,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterTemplateQuota is the Schema for the clustertemplatequota API
type ClusterTemplateQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterTemplateQuotaSpec   `json:"spec,omitempty"`
	Status ClusterTemplateQuotaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterTemplateQuotaList contains a list of ClusterTemplateQuota
type ClusterTemplateQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterTemplateQuota `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterTemplateQuota{}, &ClusterTemplateQuotaList{})
}