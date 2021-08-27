package test

import (
	apps "k8s.io/api/apps/v1"
)

// DeploymentNamesUpdated converts a deploymentsUpdated slice into a string slice of deployment names
func DeploymentNamesUpdated(deploymentsUpdated *[]*apps.Deployment) []string {
	ns := []string{}
	for _, d := range *deploymentsUpdated {
		ns = append(ns, d.GetName())
	}
	return ns
}

// ClearDeploymentsUpdated clears the deploymentsUpdated list
func ClearDeploymentsUpdated(deploymentsUpdated *[]*apps.Deployment) {
	*deploymentsUpdated = []*apps.Deployment{}
}
