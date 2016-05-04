// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package emr

import (
	"github.com/qProust/aws-sdk-go/private/waiter"
)

func (c *EMR) WaitUntilClusterRunning(input *DescribeClusterInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeCluster",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Cluster.Status.State",
				Expected: "RUNNING",
			},
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Cluster.Status.State",
				Expected: "WAITING",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "Cluster.Status.State",
				Expected: "TERMINATING",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "Cluster.Status.State",
				Expected: "TERMINATED",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "Cluster.Status.State",
				Expected: "TERMINATED_WITH_ERRORS",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EMR) WaitUntilStepComplete(input *DescribeStepInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeStep",
		Delay:       30,
		MaxAttempts: 60,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "Step.Status.State",
				Expected: "COMPLETED",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "Step.Status.State",
				Expected: "FAILED",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "Step.Status.State",
				Expected: "CANCELLED",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
