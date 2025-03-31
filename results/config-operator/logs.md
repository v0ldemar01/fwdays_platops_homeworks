## Step 1.5. Start a local Kubernetes cluster
```
$ kind create cluster --name kubebuilder-demo

Creating cluster "kubebuilder-demo" ...
✓ Ensuring node image (kindest/node:v1.27.3) 🖼
✓ Preparing nodes 📦
✓ Writing configuration 📜
✓ Starting control-plane 🕹️
✓ Installing CNI 🔌
✓ Installing StorageClass 💾
Set kubectl context to "kind-kubebuilder-demo"
You can now use your cluster with:

kubectl cluster-info --context kind-kubebuilder-demo

Thanks for using kind! 😊
```
## Step 1.6. Verify cluster is running
```
$ kubectl cluster-info

Kubernetes control plane is running at https://127.0.0.1:55346
CoreDNS is running at https://127.0.0.1:55346/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```
## Step 2.1: Initialize the Project
```
$ mkdir config-operator && cd config-operator
go mod init example.com/config-operator
kubebuilder init --domain example.com --repo example.com/config-operator
go: creating new go.mod: module example.com/config-operator
Error: failed to initialize project: unable to run pre-scaffold tasks of "base.go.kubebuilder.io/v4": go version 'go1.21.7' is incompatible because 'plugin requires go1.23 <= version < go2.0alpha1'. You can skip this check using the --skip-go-version-check flag
Usage:
kubebuilder init [flags]

Examples:

# Initialize a new project with your domain and name in copyright

kubebuilder init --plugins go/v4 --domain example.org --owner "Your name"

# Initialize a new project defining a specific project version

kubebuilder init --plugins go/v4 --project-version 3

Flags:
--domain string domain for groups (default "my.domain")
--fetch-deps ensure dependencies are downloaded (default true)
-h, --help help for init
--license string license to use to boilerplate, may be one of 'apache2', 'none' (default "apache2")
--owner string owner to add to the copyright
--project-name string name of this project
--project-version string project version (default "3")
--repo string name to use for go module (e.g., github.com/user/repo), defaults to the go package of the current working directory.
--skip-go-version-check if specified, skip checking the Go version

Global Flags:
--plugins strings plugin keys to be used for this subcommand execution

FATA failed to initialize project: unable to run pre-scaffold tasks of "base.go.kubebuilder.io/v4": go version 'go1.21.7' is incompatible because 'plugin requires go1.23 <= version < go2.0alpha1'. You can skip this check using the --skip-go-version-check flag
FAIL: 1

NOTE: I've tried to install go1.23, but faced msg="failed to create API: unable to run post-scaffold tasks of \"base.go.kubebuilder.io/v4\": exit status 2". With reference https://github.com/operator-framework/operator-sdk/issues/6681, I decided to use --skip-go-version-check flag flag.

$ kubebuilder init --domain example.com --repo example.com/config-operator --skip-go-version-check
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.20.4
INFO Update dependencies:
$ go mod tidy
Next: define a resource with:
$ kubebuilder create api
```
## Step 2.2: Create API and Controller
```
$ kubebuilder create api --group apps --version v1 --kind ConfigSync

INFO Create Resource [y/n]
y
INFO Create Controller [y/n]
y
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...
INFO api/v1/configsync_types.go
INFO api/v1/groupversion_info.go
INFO internal/controller/suite_test.go
INFO internal/controller/configsync_controller.go
INFO internal/controller/configsync_controller_test.go
INFO Update dependencies:
$ go mod tidy
INFO Running make:
$ make generate
mkdir -p /Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.17.2
go: sigs.k8s.io/controller-tools@v0.17.2 requires go >= 1.23.0; switching to go1.23.7
/Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```
## Step 2.3: Generate CRD Manifests
```
$ make manifests

/Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```
## Step 5.1: Build the Controller Image
```
$ make docker-build IMG=config-operator:v1 [10:45:57]
docker build -t config-operator:v1 .
[+] Building 52.3s (18/18) FINISHED docker:desktop-linux
=> [internal] load build definition from Dockerfile 0.0s
=> => transferring dockerfile: 1.30kB 0.0s
=> [internal] load metadata for gcr.io/distroless/static:nonroot 1.7s
=> [internal] load metadata for docker.io/library/golang:1.23 1.3s
=> [auth] library/golang:pull token for registry-1.docker.io 0.0s
=> [internal] load .dockerignore 0.0s
=> => transferring context: 160B 0.0s
=> [builder 1/9] FROM docker.io/library/golang:1.23@sha256:cb45cf739cf6bc9eaeacf75d3cd7c157e7d39b757216d813d8115d026ee32e75 0.0s
=> [internal] load build context 0.0s
=> => transferring context: 53.31kB 0.0s
=> [stage-1 1/3] FROM gcr.io/distroless/static:nonroot@sha256:c0f429e16b13e583da7e5a6ec20dd656d325d88e6819cafe0adb0828976529dc 1.6s
=> => resolve gcr.io/distroless/static:nonroot@sha256:c0f429e16b13e583da7e5a6ec20dd656d325d88e6819cafe0adb0828976529dc 0.0s
=> => sha256:4eff9a62d888790350b2481ff4a4f38f9c94b3674d26b2f2c85ca39cdef43fd9 547.59kB / 547.59kB 0.7s
=> => sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212 188B / 188B 0.5s
=> => sha256:c0f429e16b13e583da7e5a6ec20dd656d325d88e6819cafe0adb0828976529dc 1.51kB / 1.51kB 0.0s
=> => sha256:a62778643d563b511190663ef9a77c30d46d282facfdce4f3a7aecc03423c1f3 67B / 67B 0.4s
=> => sha256:f5a49cb8447f5fdd665d0258421590dbff384bb8b6295ceafc0ab3e5a9ae35e6 1.95kB / 1.95kB 0.0s
=> => sha256:d7049f94161135bc1ad59e9f2776c2063b4dcb11339db98470a659278585d2a4 1.52kB / 1.52kB 0.0s
=> => sha256:3214acf345c0cc6bbdb56b698a41ccdefc624a09d6beb0d38b5de0b2303ecaf4 123B / 123B 0.8s
=> => sha256:5664b15f108bf9436ce3312090a767300800edbbfd4511aa1a6d64357024d5dd 168B / 168B 0.8s
=> => extracting sha256:4eff9a62d888790350b2481ff4a4f38f9c94b3674d26b2f2c85ca39cdef43fd9 0.6s
=> => sha256:0bab15eea81d0fe6ab56ebf5fba14e02c4c1775a7f7436fbddd3505add4e18fa 93B / 93B 1.0s
=> => sha256:4aa0ea1413d37a58615488592a0b827ea4b2e48fa5a77cf707d0e35f025e613f 385B / 385B 1.1s
=> => sha256:da7816fa955ea24533c388143c78804c28682eef99b4ee3723b548c70148bba6 321B / 321B 1.1s
=> => sha256:9aee425378d2c16cd44177dc54a274b312897f5860a8e78fdfda555a0d79dd71 130.50kB / 130.50kB 1.5s
=> => extracting sha256:a62778643d563b511190663ef9a77c30d46d282facfdce4f3a7aecc03423c1f3 0.0s
=> => extracting sha256:7c12895b777bcaa8ccae0605b4de635b68fc32d60fa08f421dc3818bf55ee212 0.0s
=> => extracting sha256:3214acf345c0cc6bbdb56b698a41ccdefc624a09d6beb0d38b5de0b2303ecaf4 0.0s
=> => extracting sha256:5664b15f108bf9436ce3312090a767300800edbbfd4511aa1a6d64357024d5dd 0.0s
=> => extracting sha256:0bab15eea81d0fe6ab56ebf5fba14e02c4c1775a7f7436fbddd3505add4e18fa 0.0s
=> => extracting sha256:4aa0ea1413d37a58615488592a0b827ea4b2e48fa5a77cf707d0e35f025e613f 0.0s
=> => extracting sha256:da7816fa955ea24533c388143c78804c28682eef99b4ee3723b548c70148bba6 0.0s
=> => extracting sha256:9aee425378d2c16cd44177dc54a274b312897f5860a8e78fdfda555a0d79dd71 0.0s
=> CACHED [builder 2/9] WORKDIR /workspace 0.0s
=> CACHED [builder 3/9] COPY go.mod go.mod 0.0s
=> CACHED [builder 4/9] COPY go.sum go.sum 0.0s
=> CACHED [builder 5/9] RUN go mod download 0.0s
=> CACHED [builder 6/9] COPY cmd/main.go cmd/main.go 0.0s
=> CACHED [builder 7/9] COPY api/ api/ 0.0s
=> [builder 8/9] COPY internal/ internal/ 0.1s
=> [builder 9/9] RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o manager cmd/main.go 49.9s
=> [stage-1 2/3] COPY --from=builder /workspace/manager . 0.3s
=> exporting to image 0.2s
=> => exporting layers 0.1s
=> => writing image sha256:2d377b25e219ac2c285785d3a692c5d8ac7ee5445e590d75262b4cc5edbec8f9 0.0s
=> => naming to docker.io/library/config-operator:v1 0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/k9s2idyofz7pdsgeeimc0j6cc

What's next:
View a summary of image vulnerabilities and recommendations → docker scout quickview
```
## Step 5.3: Deploy the Controller to the Cluster
```
$ kind load docker-image config-operator:v1 --name kubebuilder-demo [10:48:51]
Image: "config-operator:v1" with ID "sha256:2d377b25e219ac2c285785d3a692c5d8ac7ee5445e590d75262b4cc5edbec8f9" not yet present on node "kubebuilder-demo-control-plane", loading...
```
## Step 5.3: Deploy the Controller to the Cluster
```
$ make deploy IMG=config-operator:v1
/Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
Downloading sigs.k8s.io/kustomize/kustomize/v5@v5.6.0
go: sigs.k8s.io/kustomize/kustomize/v5@v5.6.0 requires go >= 1.22.7; switching to go1.23.7
cd config/manager && /Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/kustomize edit set image controller=config-operator:v1
/Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/kustomize build config/default | kubectl apply -f -
namespace/config-operator-system created
customresourcedefinition.apiextensions.k8s.io/configsyncs.apps.example.com created
serviceaccount/config-operator-controller-manager created
role.rbac.authorization.k8s.io/config-operator-leader-election-role created
clusterrole.rbac.authorization.k8s.io/config-operator-configsync-admin-role created
clusterrole.rbac.authorization.k8s.io/config-operator-configsync-editor-role created
clusterrole.rbac.authorization.k8s.io/config-operator-configsync-viewer-role created
clusterrole.rbac.authorization.k8s.io/config-operator-manager-role created
clusterrole.rbac.authorization.k8s.io/config-operator-metrics-auth-role created
clusterrole.rbac.authorization.k8s.io/config-operator-metrics-reader created
rolebinding.rbac.authorization.k8s.io/config-operator-leader-election-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/config-operator-manager-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/config-operator-metrics-auth-rolebinding created
service/config-operator-controller-manager-metrics-service created
deployment.apps/config-operator-controller-manager created
```
## Step 6.2: Apply the Sample Resource
```
$ kubectl apply -f config/samples/apps_v1_configsync.yaml [10:54:09]
configsync.apps.example.com/configsync-sample created
```
## Step 6.3: Verify the Controller Created the ConfigMap
```
$ kubectl get configmap my-config -o yaml
apiVersion: v1
data:
app.properties: |
app.name=MyApp
app.version=1.0.0
key1: value1
key2: value2
kind: ConfigMap
metadata:
creationTimestamp: "2025-03-31T07:55:35Z"
name: my-config
namespace: default
resourceVersion: "2213"
uid: 6262dfdd-083a-426c-9286-900350e3f485
```
## Step 6.4: Check the Status of the ConfigSync Resource
```
$ kubectl get configsync configsync-sample -o yaml
apiVersion: apps.example.com/v1
kind: ConfigSync
metadata:
annotations:
kubectl.kubernetes.io/last-applied-configuration: |
{"apiVersion":"apps.example.com/v1","kind":"ConfigSync","metadata":{"annotations":{},"name":"configsync-sample","namespace":"default"},"spec":{"configMapName":"my-config","data":{"app.properties":"app.name=MyApp\napp.version=1.0.0\n","key1":"value1","key2":"value2"},"updateInterval":30}}
creationTimestamp: "2025-03-31T07:55:35Z"
generation: 1
name: configsync-sample
namespace: default
resourceVersion: "2268"
uid: 71576be6-15ab-4dd2-8618-edada84acea3
spec:
configMapName: my-config
data:
app.properties: |
app.name=MyApp
app.version=1.0.0
key1: value1
key2: value2
updateInterval: 30
status:
lastSyncTime: "2025-03-31T07:56:05Z"
status: Synced
```
## Step 6.5: Update the ConfigSync Resource and Observe the Changes
```
$ kubectl edit configsync configsync-sample

configsync.apps.example.com/configsync-sample edited

$ kubectl get configmap my-config -o yaml
apiVersion: v1
data:
app.properties: |
app.name=MyApp
app.version=1.0.0
key1: value1
key2: value3 // changed
kind: ConfigMap
metadata:
creationTimestamp: "2025-03-31T07:55:35Z"
name: my-config
namespace: default
resourceVersion: "2442"
uid: 6262dfdd-083a-426c-9286-900350e3f485
```
## Step 7.1: Delete the ConfigSync Resource
```
$ kubectl delete -f config/samples/apps_v1_configsync.yaml

configsync.apps.example.com "configsync-sample" deleted
```
## Step 7.2: Uninstall the Controller
```
make undeploy

/Users/volodymyrminchenko/My-Education/PlatOps/fwdays_platops_homeworks/config-operator/bin/kustomize build config/default | kubectl delete --ignore-not-found=false -f -
namespace "config-operator-system" deleted
customresourcedefinition.apiextensions.k8s.io "configsyncs.apps.example.com" deleted
serviceaccount "config-operator-controller-manager" deleted
role.rbac.authorization.k8s.io "config-operator-leader-election-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-configsync-admin-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-configsync-editor-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-configsync-viewer-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-manager-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-metrics-auth-role" deleted
clusterrole.rbac.authorization.k8s.io "config-operator-metrics-reader" deleted
rolebinding.rbac.authorization.k8s.io "config-operator-leader-election-rolebinding" deleted
clusterrolebinding.rbac.authorization.k8s.io "config-operator-manager-rolebinding" deleted
clusterrolebinding.rbac.authorization.k8s.io "config-operator-metrics-auth-rolebinding" deleted
service "config-operator-controller-manager-metrics-service" deleted
deployment.apps "config-operator-controller-manager" deleted
```
## Step 7.3: Delete the Kubernetes Cluster
```
$ kind delete cluster --name kubebuilder-demo

Deleting cluster "kubebuilder-demo" ...
Deleted nodes: ["kubebuilder-demo-control-plane"]
```
