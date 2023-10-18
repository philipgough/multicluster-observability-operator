// Copyright (c) Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project
// Licensed under the Apache License 2.0

package util

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	addonv1alpha1 "open-cluster-management.io/api/addon/v1alpha1"
)

const (
	name      = "test"
	namespace = "test"
)

func TestManagedClusterAddon(t *testing.T) {
	s := scheme.Scheme
	addonv1alpha1.AddToScheme(s)
	c := fake.NewClientBuilder().Build()
	_, err := CreateManagedClusterAddonCR(c, namespace, "testKey", "value")
	if err != nil {
		t.Fatalf("Failed to create managedclusteraddon: (%v)", err)
	}
	addon := &addonv1alpha1.ManagedClusterAddOn{}
	err = c.Get(context.TODO(),
		types.NamespacedName{
			Name:      ManagedClusterAddonName,
			Namespace: namespace,
		},
		addon,
	)
	if err != nil {
		t.Fatalf("Failed to get managedclusteraddon: (%v)", err)
	}
}
