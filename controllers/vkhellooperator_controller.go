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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	cachev1alpha1 "github.com/vikash21kumar/vk-hello-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VKHelloOperatorReconciler reconciles a VKHelloOperator object
type VKHelloOperatorReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cache.example.com,resources=vkhellooperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=vkhellooperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=vkhellooperators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VKHelloOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *VKHelloOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("vkhellooperator", req.NamespacedName)

	// your logic here
	ctx1 := context.Background()

	vKHelloOperator := &cachev1alpha1.VKHelloOperator{}
	err := r.Get(ctx1, req.NamespacedName, vKHelloOperator)
	/*
		if err != nil {
			if errors.IsNotFound(err) {
				// Request object not found, could have been deleted after reconcile request.
				// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
				// Return and don't requeue
				log.Info("Greeting resource not found. Ignoring since object must be deleted")
				return ctrl.Result{}, nil
			}
			// Error reading the object - requeue the request.
			log.Error(err, "Failed to get Greeting")
			return ctrl.Result{}, err
		}
	*/
	greeting := vKHelloOperator.Spec.Greetings

	log.Info("Saying", "Greeting", greeting)

	vKHelloOperator.Status.TheGreetings = greeting
	updateErr := r.Status().Update(ctx1, vKHelloOperator)
	if updateErr != nil {
		log.Error(err, "Failed to update Greeting status")
		return ctrl.Result{}, updateErr
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VKHelloOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.VKHelloOperator{}).
		Complete(r)
}
