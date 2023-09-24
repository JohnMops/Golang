/*
Copyright 2023.

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

package controller

import (
	"context"

	johncomv1 "gcp/api/v1"
	"gcp/internal/gcp"

	// "github.com/k0kubun/pp"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var controllerLog = ctrl.Log.WithName("Controller")

const (
	FinalizerName = "serviceAccount.finalizer"
)

// ServiceAccountReconciler reconciles a ServiceAccount object
type ServiceAccountReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=john.com,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=john.com,resources=serviceaccounts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=john.com,resources=serviceaccounts/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ServiceAccount object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *ServiceAccountReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	serviceAccount := &johncomv1.ServiceAccount{}
	if err := r.Get(ctx, req.NamespacedName, serviceAccount); err != nil {
		// Handle the error
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if serviceAccount.ObjectMeta.DeletionTimestamp != nil {
		r.Delete(ctx, req)

	} else {
		controllerLog.Info("CR detected: " + serviceAccount.Spec.ServiceAccountName)
		check, _ := gcp.CheckServiceAccountExists(serviceAccount.Spec.ProjectID, serviceAccount.Spec.ServiceAccountName)

		if !check {
			controllerLog.Info(serviceAccount.Spec.ServiceAccountName + " does not exist and can be created")
			gcp.CreateServiceAccount(serviceAccount.Spec.ProjectID, serviceAccount.Spec.ServiceAccountName)
			controllerLog.Info("Created: " + serviceAccount.Spec.ServiceAccountName)
		} else {
			controllerLog.Info(serviceAccount.Spec.ServiceAccountName + " exists")
		}
	}

	return ctrl.Result{}, nil
}

func (r *ServiceAccountReconciler) Delete(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	controllerLog.Info("Firing delete")
	serviceAccount := &johncomv1.ServiceAccount{}
	if err := r.Get(ctx, req.NamespacedName, serviceAccount); err != nil {
		if errors.IsNotFound(err) {
			// Resource not found, nothing to delete
			controllerLog.Info("Resource not found, skipping delete reconciliation")
			return ctrl.Result{}, nil
		}
		// Handle other errors
		controllerLog.Error(err, "Failed to fetch resource for delete reconciliation")
		return ctrl.Result{}, err
	}

	if containsString(serviceAccount.ObjectMeta.Finalizers, FinalizerName) {
		// Run finalization logic here
		// For example, clean up resources associated with myResource
		controllerLog.Info(serviceAccount.Spec.ServiceAccountName + " CR is deleting, Will delete the SA")
		gcp.DeleteServiceAccount(serviceAccount.Spec.ProjectID, serviceAccount.Spec.ServiceAccountName)
		// Remove the finalizer
		serviceAccount.ObjectMeta.Finalizers = removeString(serviceAccount.ObjectMeta.Finalizers, FinalizerName)
		if err := r.Update(ctx, serviceAccount); err != nil {
			controllerLog.Error(err, "Failed to update resource while removing finalizer")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceAccountReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&johncomv1.ServiceAccount{}).
		Complete(r)
}
