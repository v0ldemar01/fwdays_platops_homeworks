## Reconciliation Loop Explanation
The reconciliation loop in the `ConfigSyncReconciler` ensures that the actual state of the cluster matches the desired state defined in the `ConfigSync` resource. It starts by retrieving the `ConfigSync` instance using the req.NamespacedName. If the resource exists, it checks for the associated ConfigMap. If the ConfigMap is missing, it creates a new one using the data from the `ConfigSync` spec. If the ConfigMap is present, it updates its data to match the spec. After successfully creating or updating the `ConfigMap`, the `ConfigSync` status is updated to reflect the last synchronization time and status. The loop also schedules the next reconciliation based on the `UpdateInterval` specified in the ConfigSync spec.

## Error Handling
Errors are managed at every critical step to ensure proper visibility and system stability. If the `ConfigSync` resource cannot be retrieved, the error is logged, and the reconciliation loop exits gracefully. If there is an issue creating or updating the `ConfigMap`, the error is logged, and the `ConfigSync` status is updated with a relevant error message. Additionally, if updating the ConfigSync status itself fails, the error is logged to ensure it is not overlooked. This approach ensures that all errors are surfaced and the system remains consistent.

## Challenges and Solutions
- `controllers` vs `controller`
- required go version(proposed 1.20 does not work for me), installed go1.21.7 darwin/arm64, but after executing `kubebuilder create api` - go1.23.0 darwin/arm64. Made some investigation and found this [issue](https://github.com/operator-framework/operator-sdk/issues/6681)
- version name for MacOS(Apple Silicon) for packages
