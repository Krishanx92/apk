package xds

import (
	"context"
	"errors"

	"github.com/wso2/apk/adapter/internal/loggers"
	logging "github.com/wso2/apk/adapter/internal/logging"
	cpv1alpha1 "github.com/wso2/apk/adapter/pkg/operator/apis/cp/v1alpha1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrlcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HandleApplicationEventsFromMgtServer handles the Application events
func HandleApplicationEventsFromMgtServer(c client.Client, cReader client.Reader) {
	for applicationEvent := range applicationChannel {
		switch applicationEvent.Type {
		case ApplicationCreate:
			if found, _, err := checkApplicationExists(applicationEvent.Application, c, cReader); err == nil && !found {
				if err := c.Create(context.Background(), *&applicationEvent.Application); err != nil {
					loggers.LoggerXds.ErrorC(logging.GetErrorByCode(logging.Error1707, err.Error()))
				} else {
					loggers.LoggerXds.Info("Application created: " + applicationEvent.Application.Name)
				}
			}
			break
		case ApplicationUpdate:
			if found, application, err := checkApplicationExists(applicationEvent.Application, c, cReader); err == nil && found {
				application.Spec = applicationEvent.Application.Spec
				err := c.Update(context.Background(), application)
				if err != nil {
					loggers.LoggerXds.ErrorC(logging.GetErrorByCode(logging.Error1709, err.Error()))
				} else {
					loggers.LoggerXds.Info("Application updated: " + applicationEvent.Application.Name)
				}
			}
			break
		case ApplicationDelete:
			err := c.Delete(context.Background(), *&applicationEvent.Application)
			if err != nil {
				loggers.LoggerXds.ErrorC(logging.GetErrorByCode(logging.Error1710, err.Error()))
			} else {
				loggers.LoggerXds.Info("Application deleted: " + applicationEvent.Application.Name)
			}
			break
		default:
			loggers.LoggerXds.Info("Unknown Application Event Type")
		}
	}
}

func checkApplicationExists(application *cpv1alpha1.Application, c client.Client, cReader client.Reader) (bool, *cpv1alpha1.Application, error) {
	var retrivedApplication = new(cpv1alpha1.Application)
	// Try reading from cache
	if err := c.Get(context.Background(), types.NamespacedName{
		Name:      application.Name,
		Namespace: application.Namespace}, retrivedApplication); err != nil {

		target := &ctrlcache.ErrCacheNotStarted{}
		if errors.As(err, &target) {
			// Try reading from api server directly
			if err := cReader.Get(context.Background(), types.NamespacedName{
				Name:      application.Name,
				Namespace: application.Namespace}, retrivedApplication); err != nil {

				if !apierrors.IsNotFound(err) {
					loggers.LoggerXds.ErrorC(logging.GetErrorByCode(logging.Error1711, err.Error()))
					return false, nil, err
				}
				return false, nil, nil
			}
		} else if !apierrors.IsNotFound(err) {
			loggers.LoggerXds.ErrorC(logging.GetErrorByCode(logging.Error1711, err.Error()))
			return false, nil, err
		} else {
			return false, nil, nil
		}
	}
	return true, retrivedApplication, nil
}
