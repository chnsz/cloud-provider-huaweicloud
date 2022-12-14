package huaweicloud

import (
	"context"

	"k8s.io/apimachinery/pkg/types"

	cloudprovider "k8s.io/cloud-provider"
)

type Zones struct {
	*HuaweiCloud
}

// GetZone returns the Zone containing the current failure zone and locality region that the program is running in
// In most cases, this method is called from the kubelet querying a local metadata service to acquire its zone.
// For the case of external cloud providers, use GetZoneByProviderID or GetZoneByNodeName since GetZone
// can no longer be called from the kubelets.
func (z *Zones) GetZone(ctx context.Context) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{}, nil
}

// GetZoneByProviderID returns the Zone containing the current zone and locality region of the node specified by providerID
// This method is particularly used in the context of external cloud providers where node initialization must be done
// outside the kubelets.
func (z *Zones) GetZoneByProviderID(ctx context.Context, providerID string) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{}, nil
}

// GetZoneByNodeName returns the Zone containing the current zone and locality region of the node specified by node name
// This method is particularly used in the context of external cloud providers where node initialization must be done
// outside the kubelets.
func (z *Zones) GetZoneByNodeName(ctx context.Context, nodeName types.NodeName) (cloudprovider.Zone, error) {
	return cloudprovider.Zone{}, nil
}
