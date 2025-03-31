package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigSyncSpec defines the desired state of ConfigSync
type ConfigSyncSpec struct {
	// ConfigMapName is the name of the ConfigMap to create or update
	ConfigMapName string `json:"configMapName"`

	// Data contains key-value pairs that will be stored in the ConfigMap
	Data map[string]string `json:"data,omitempty"`

	// UpdateInterval specifies how often the controller should check for changes (in seconds)
	UpdateInterval int `json:"updateInterval,omitempty"`
}

// ConfigSyncStatus defines the observed state of ConfigSync
type ConfigSyncStatus struct {
	// LastSyncTime is the last time the ConfigMap was synced
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`

	// Status indicates the current sync status
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConfigSync is the Schema for the configsyncs API
type ConfigSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSyncSpec   `json:"spec,omitempty"`
	Status ConfigSyncStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigSyncList contains a list of ConfigSync
type ConfigSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigSync{}, &ConfigSyncList{})
}
