# Homework: Building Custom Kubernetes Controllers with Kubebuilder

## **Objective**
In this assignment, you will:
1. Install and configure **Kubebuilder** and its prerequisites.
2. Set up a local Kubernetes development environment.
3. Create a custom Kubernetes controller that manages a **ConfigMap** resource.
4. Implement a reconciliation loop with proper error handling.
5. Deploy and test your custom controller in a Kubernetes cluster.

---

## **Part 1: Installation & Configuration**

### **Step 1: Install Required Tools**
Ensure you have the following installed on your system:
- **Go (v1.19+)**
- **Docker**
- **kubectl**
- **kind** or **minikube** for local Kubernetes cluster
- **Kubebuilder**

#### **1. Install Go**
```sh
# For Linux/macOS - for Apple Silicon it should be
wget https://go.dev/dl/go1.20.linux-amd64.tar.gz `darwin-arm64` instead of `linux-amd64`

sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

#### **2. Install kubectl**
```sh
# For Linux/macOS - for Apple Silicon it should be
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```

#### **3. Install kind**
```sh
# For Linux/macOS - for Apple Silicon it should be
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

#### **4. Install Kubebuilder**
```sh
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder
sudo mv kubebuilder /usr/local/bin/
```

#### **5. Start a local Kubernetes cluster**
```sh
kind create cluster --name kubebuilder-demo
```

#### **6. Verify cluster is running**
```sh
kubectl cluster-info
```
You should see information about your Kubernetes cluster.

---

## **Part 2: Create the Kubebuilder Project**

### **Step 1: Initialize the Project**
```sh
mkdir config-operator && cd config-operator
go mod init example.com/config-operator
kubebuilder init --domain example.com --repo example.com/config-operator
```

### **Step 2: Create API and Controller**
```sh
kubebuilder create api --group apps --version v1 --kind ConfigSync
```
When prompted:
- Create Resource [y/n]: y
- Create Controller [y/n]: y

---

## **Part 3: Define Custom Resource**

### **Modify `api/v1/configsync_types.go`**
Replace or update the contents with:

```go
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
```

### **Step 3: Generate CRD Manifests**
```sh
make manifests
```

---

## **Part 4: Implement the Controller**

### **Modify `controller/configsync_controller.go`**
Replace the contents with:

```go
package controller

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "example.com/config-operator/api/v1"
)

// ConfigSyncReconciler reconciles a ConfigSync object
type ConfigSyncReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.example.com,resources=configsyncs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.example.com,resources=configsyncs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.example.com,resources=configsyncs/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop
func (r *ConfigSyncReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the ConfigSync instance
	configSync := &appsv1.ConfigSync{}
	err := r.Get(ctx, req.NamespacedName, configSync)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted
			return ctrl.Result{}, nil
		}
		// Error reading the object
		log.Error(err, "Failed to get ConfigSync")
		return ctrl.Result{}, err
	}

	// Default update interval to 60 seconds if not specified
	updateInterval := 60
	if configSync.Spec.UpdateInterval > 0 {
		updateInterval = configSync.Spec.UpdateInterval
	}

	// Check if ConfigMap exists
	targetConfigMap := &corev1.ConfigMap{}
	err = r.Get(ctx, types.NamespacedName{
		Name:      configSync.Spec.ConfigMapName,
		Namespace: configSync.Namespace,
	}, targetConfigMap)

	// Create or update ConfigMap
	if err != nil && errors.IsNotFound(err) {
		// ConfigMap doesn't exist, create a new one
		configMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      configSync.Spec.ConfigMapName,
				Namespace: configSync.Namespace,
			},
			Data: configSync.Spec.Data,
		}

		if err := r.Create(ctx, configMap); err != nil {
			log.Error(err, "Failed to create ConfigMap")

			// Update status with error
			configSync.Status.Status = "Error: Failed to create ConfigMap"
			if updateErr := r.Status().Update(ctx, configSync); updateErr != nil {
				log.Error(updateErr, "Failed to update ConfigSync status")
			}

			return ctrl.Result{}, err
		}

		log.Info("Created ConfigMap", "ConfigMap.Name", configMap.Name)
	} else if err == nil {
		// ConfigMap exists, update it
		targetConfigMap.Data = configSync.Spec.Data
		if err := r.Update(ctx, targetConfigMap); err != nil {
			log.Error(err, "Failed to update ConfigMap")

			// Update status with error
			configSync.Status.Status = "Error: Failed to update ConfigMap"
			if updateErr := r.Status().Update(ctx, configSync); updateErr != nil {
				log.Error(updateErr, "Failed to update ConfigSync status")
			}

			return ctrl.Result{}, err
		}
		log.Info("Updated ConfigMap", "ConfigMap.Name", targetConfigMap.Name)
	} else {
		// Error fetching the ConfigMap
		log.Error(err, "Failed to get ConfigMap")
		return ctrl.Result{}, err
	}

	// Update status
	configSync.Status.LastSyncTime = metav1.Now()
	configSync.Status.Status = "Synced"
	if err := r.Status().Update(ctx, configSync); err != nil {
		log.Error(err, "Failed to update ConfigSync status")
		return ctrl.Result{}, err
	}

	// Schedule next reconciliation based on updateInterval
	return ctrl.Result{RequeueAfter: time.Duration(updateInterval) * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConfigSyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.ConfigSync{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
```

---

## **Part 5: Build and Deploy the Controller**

### **Step 1: Build the Controller Image**
```sh
make docker-build IMG=config-operator:v1
```

### **Step 2: Load the Image into Kind Cluster**
```sh
kind load docker-image config-operator:v1 --name kubebuilder-demo
```

### **Step 3: Deploy the Controller to the Cluster**
```sh
make deploy IMG=config-operator:v1
```

---

## **Part 6: Test the Controller**

### **Step 1: Create a Sample ConfigSync Resource**
Create a file named `config/samples/apps_v1_configsync.yaml` with the following content:

```yaml
apiVersion: apps.example.com/v1
kind: ConfigSync
metadata:
  name: configsync-sample
spec:
  configMapName: my-config
  data:
    key1: "value1"
    key2: "value2"
    app.properties: |
      app.name=MyApp
      app.version=1.0.0
  updateInterval: 30
```

### **Step 2: Apply the Sample Resource**
```sh
kubectl apply -f config/samples/apps_v1_configsync.yaml
```

### **Step 3: Verify the Controller Created the ConfigMap**
```sh
kubectl get configmap my-config -o yaml
```

### **Step 4: Check the Status of the ConfigSync Resource**
```sh
kubectl get configsync configsync-sample -o yaml
```

You should see that the `status` field has been updated with `lastSyncTime` and `status: Synced`.

### **Step 5: Update the ConfigSync Resource and Observe the Changes**
Edit the ConfigSync resource to change some data:

```sh
kubectl edit configsync configsync-sample
```

Then verify the ConfigMap was updated:

```sh
kubectl get configmap my-config -o yaml
```

---

## **Part 7: Clean Up**

### **Step 1: Delete the ConfigSync Resource**
```sh
kubectl delete -f config/samples/apps_v1_configsync.yaml
```

### **Step 2: Uninstall the Controller**
```sh
make undeploy
```

### **Step 3: Delete the Kubernetes Cluster**
```sh
kind delete cluster --name kubebuilder-demo
```

---

## **Submission Requirements**
- Source code of your Kubebuilder project
- Screenshots showing:
  - Successful deployment of your controller
  - The created ConfigMap
  - Status updates of your ConfigSync resource
- A brief explanation (1-2 paragraphs) of:
  - How your reconciliation loop works
  - How you handled error conditions
  - What challenges you faced and how you overcame them

**Good luck! 🚀**
