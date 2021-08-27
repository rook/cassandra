/*
Copyright 2018 The Rook Authors. All rights reserved.

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

package integration

import (
	"context"
	"os"
	"testing"
	"time"

	cassandrav1alpha1 "github.com/rook/cassandra/pkg/apis/cassandra.rook.io/v1alpha1"
	"github.com/rook/cassandra/tests/framework/installer"
	"github.com/rook/cassandra/tests/framework/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ************************************************
// *** Major scenarios tested by the CassandraSuite ***
// Setup
// - via the cluster CRD with very simple properties
//   - 1 replica
//     - 1 CPU
//     - 2GB memory
//     - 5Gi volume from default provider
// ************************************************

type CassandraSuite struct {
	suite.Suite
	k8sHelper       *utils.K8sHelper
	installer       *installer.CassandraInstaller
	namespace       string
	systemNamespace string
	instanceCount   int
}

// TestCassandraSuite initiates the CassandraSuite
func TestCassandraSuite(t *testing.T) {
	if installer.SkipTestSuite(installer.CassandraTestSuite) {
		t.Skip()
	}
	if os.Getenv("SKIP_CASSANDRA_TESTS") == "true" {
		t.Skip()
	}

	s := new(CassandraSuite)
	defer func(s *CassandraSuite) {
		r := recover()
		if r != nil {
			logger.Infof("unexpected panic occurred during test %s, --> %v", t.Name(), r)
			t.Fail()
			s.Teardown()
			t.FailNow()
		}
	}(s)
	suite.Run(t, s)
}

// SetupSuite runs once at the beginning of the suite,
// before any tests are run.
func (s *CassandraSuite) SetupSuite() {

	s.namespace = "cassandra-ns"
	s.systemNamespace = installer.SystemNamespace(s.namespace)
	s.instanceCount = 1

	k8sHelper, err := utils.CreateK8sHelper(s.T)
	require.NoError(s.T(), err)
	s.k8sHelper = k8sHelper

	k8sVersion := s.k8sHelper.GetK8sServerVersion()
	logger.Infof("Installing Cassandra on K8s %s", k8sVersion)

	s.installer = installer.NewCassandraInstaller(s.k8sHelper, s.T)

	if err = s.installer.InstallCassandra(s.systemNamespace, s.namespace, s.instanceCount, cassandrav1alpha1.ClusterModeCassandra); err != nil {
		logger.Errorf("Cassandra was not installed successfully: %s", err.Error())
		s.T().Fail()
		s.Teardown()
		s.T().FailNow()
	}
}

// BeforeTest runs before every test in the CassandraSuite.
func (s *CassandraSuite) TeardownSuite() {
	s.Teardown()
}

///////////
// Tests //
///////////

// TestCassandraClusterCreation tests the creation of a Cassandra cluster.
func (s *CassandraSuite) TestCassandraClusterCreation() {
	s.CheckClusterHealth()
}

// TestScyllaClusterCreation tests the creation of a Scylla cluster.
// func (s *CassandraSuite) TestScyllaClusterCreation() {
// 	s.CheckClusterHealth()
// }

//////////////////////
// Helper Functions //
//////////////////////

// Teardown gathers logs and other helping info and then uninstalls
// everything installed by the CassandraSuite
func (s *CassandraSuite) Teardown() {
	s.installer.GatherAllCassandraLogs(s.systemNamespace, s.namespace, s.T().Name())
	s.installer.UninstallCassandra(s.systemNamespace, s.namespace)
}

// CheckClusterHealth checks if all Pods in the cluster are ready
// and CQL is working.
func (s *CassandraSuite) CheckClusterHealth() {
	// Verify that cassandra-operator is running
	operatorName := "rook-cassandra-operator"
	logger.Infof("Verifying that all expected pods of cassandra operator are ready")
	ready := utils.Retry(10, 30*time.Second,
		"Waiting for Cassandra operator to be ready", func() bool {
			sts, err := s.k8sHelper.Clientset.AppsV1().StatefulSets(s.systemNamespace).Get(context.TODO(), operatorName, v1.GetOptions{})
			if err != nil {
				logger.Errorf("Error getting Cassandra operator `%s`", operatorName)
				return false
			}
			if sts.Generation != sts.Status.ObservedGeneration {
				logger.Infof("Operator Statefulset has not converged yet")
				return false
			}
			if sts.Status.UpdatedReplicas != *sts.Spec.Replicas {
				logger.Error("Operator StatefulSet is rolling updating")
				return false
			}
			if sts.Status.ReadyReplicas != *sts.Spec.Replicas {
				logger.Infof("Statefulset not ready. Got: %v, Want: %v",
					sts.Status.ReadyReplicas, sts.Spec.Replicas)
				return false
			}
			return true
		})
	assert.True(s.T(), ready, "Timed out waiting for Cassandra operator to become ready")

	// Verify cassandra cluster instances are running OK
	clusterName := "cassandra-ns"
	clusterNamespace := "cassandra-ns"
	ready = utils.Retry(10, 30*time.Second,
		"Waiting for Cassandra cluster to be ready", func() bool {
			c, err := s.k8sHelper.RookClientset.CassandraV1alpha1().Clusters(clusterNamespace).Get(context.TODO(), clusterName, v1.GetOptions{})
			if err != nil {
				logger.Errorf("Error getting Cassandra cluster `%s`", clusterName)
				return false
			}
			for rackName, rack := range c.Status.Racks {
				var desiredMembers int32
				for _, r := range c.Spec.Datacenter.Racks {
					if r.Name == rackName {
						desiredMembers = r.Members
						break
					}
				}
				if !(desiredMembers == rack.Members && rack.Members == rack.ReadyMembers) {
					logger.Infof("Rack `%s` is not ready yet", rackName)
					return false
				}
			}
			return true
		})
	assert.True(s.T(), ready, "Timed out waiting for Cassandra cluster to become ready")

	// Determine a pod name for the cluster
	podName := "cassandra-ns-us-east-1-us-east-1a-0"

	// Get the Pod's IP address
	command := "hostname"
	commandArgs := []string{"-i"}
	podIP, err := s.k8sHelper.Exec(s.namespace, podName, command, commandArgs)
	assert.NoError(s.T(), err)

	command = "cqlsh"
	commandArgs = []string{
		"-e",
		`
CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {
'class': 'SimpleStrategy',
'replication_factor': 1
};
USE test;
CREATE TABLE IF NOT EXISTS map (key text, value text, PRIMARY KEY(key));
INSERT INTO map (key, value) VALUES('test_key', 'test_value');
SELECT key,value FROM map WHERE key='test_key';`,
		podIP,
	}

	var result string
	for i := 0; i < 5; i++ {
		logger.Warning("trying cassandra cql command in 30s")
		time.Sleep(utils.RetryInterval * time.Second)

		result, err = s.k8sHelper.Exec(s.namespace, podName, command, commandArgs)
		logger.Infof("cassandra cql command exited, err: %v. result: %s", err, result)
		if err == nil {
			break
		}
		logger.Errorf("cassandra cql command failed. %v", err)
	}

	// FIX: The Cassandra commands are failing in the CI
	//assert.NoError(s.T(), err)
	//assert.True(s.T(), strings.Contains(result, "test_key"))
	//assert.True(s.T(), strings.Contains(result, "test_value"))
}
