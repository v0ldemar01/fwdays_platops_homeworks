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
