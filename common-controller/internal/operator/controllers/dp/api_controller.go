/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package dp

import (
	"context"

	"github.com/wso2/apk/adapter/pkg/logging"
	loggers "github.com/wso2/apk/common-controller/internal/loggers"
	dpv1alpha2 "github.com/wso2/apk/common-controller/internal/operator/apis/dp/v1alpha2"
	constants "github.com/wso2/apk/common-controller/internal/operator/constant"
	"github.com/wso2/apk/common-controller/internal/utils"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// APIReconciler reconciles a API object
type APIReconciler struct {
	client client.Client
	Scheme *runtime.Scheme
}

// NewApplicationController creates a new Application controller instance
func NewAPIController(mgr manager.Manager) error {
	r := &APIReconciler{
		client: mgr.GetClient(),
	}
	c, err := controller.New(constants.APIController, mgr, controller.Options{Reconciler: r})
	if err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2606, logging.BLOCKER, "Error creating API controller: %v", err.Error()))
		return err
	}

	if err := c.Watch(source.Kind(mgr.GetCache(), &dpv1alpha2.API{}), &handler.EnqueueRequestForObject{},
		predicate.NewPredicateFuncs(utils.FilterByNamespaces([]string{utils.GetOperatorPodNamespace()}))); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2607, logging.BLOCKER, "Error watching API resources: %v", err.Error()))
		return err
	}

	loggers.LoggerAPKOperator.Debug("API Controller successfully started. Watching API Objects...")
	return nil
}

//+kubebuilder:rbac:groups=dp.wso2.com,resources=apis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dp.wso2.com,resources=apis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dp.wso2.com,resources=apis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the API object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *APIReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *APIReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dpv1alpha2.API{}).
		Complete(r)
}
