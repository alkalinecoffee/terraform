package remote

import (
	"testing"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/state/remote"
)

func TestRemoteClient_impl(t *testing.T) {
	var _ remote.Client = new(remoteClient)
}

func TestRemoteClient(t *testing.T) {
	client := testRemoteClient(t)
	remote.TestClient(t, client)
}

func TestRemoteClient_stateLock(t *testing.T) {
	b := testBackendDefault(t)

	s1, err := b.State(backend.DefaultStateName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	s2, err := b.State(backend.DefaultStateName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	remote.TestRemoteLocks(t, s1.(*remote.State).Client, s2.(*remote.State).Client)
}
