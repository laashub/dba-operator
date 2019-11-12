/*

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DatabaseMigrationSchemaHintColumn contains information about one column
// being added to a table, or in a table definition
type DatabaseMigrationSchemaHintColumn struct {
	Name             string `json:"name,omitempty"`
	NotNullable      bool   `json:"notNullable,omitempty"`
	HasServerDefault bool   `json:"hasServerDefault,omitempty"`
}

// DatabaseMigrationSchemaHint approximately describes what the migration is going to change
type DatabaseMigrationSchemaHint struct {
	TableReference `json:",omitempty"`

	// +kubebuilder:validation:Enum=addColumn;createTable;createIndex;alterColumn
	Operation string `json:"operation,omitempty"`

	Columns []DatabaseMigrationSchemaHintColumn `json:"columns,omitempty"`

	// +kubebuilder:validation:Enum=index;unique;fulltext
	IndexType string `json:"indexType,omitempty"`
}

// DatabaseMigrationSpec defines the desired state of DatabaseMigration
type DatabaseMigrationSpec struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=12
	Previous               string                        `json:"previous,omitempty"`
	MigrationContainerSpec corev1.Container              `json:"migrationContainerSpec,omitempty"`
	Scalable               bool                          `json:"scalable,omitempty"`
	SchemaHints            []DatabaseMigrationSchemaHint `json:"schemaHints,omitempty"`
}

// DatabaseMigrationStatus defines the observed state of DatabaseMigration
type DatabaseMigrationStatus struct {
}

// +kubebuilder:object:root=true

// DatabaseMigration is the Schema for the databasemigrations API
type DatabaseMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseMigrationSpec   `json:"spec,omitempty"`
	Status DatabaseMigrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseMigrationList contains a list of DatabaseMigration
type DatabaseMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseMigration{}, &DatabaseMigrationList{})
}
