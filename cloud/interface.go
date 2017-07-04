package cloud

import (
	"github.com/kris-nova/kubicorn/apis/cluster"
)

type Reconciler interface {
	Init() error
	GetActual() (*cluster.Cluster, error)
	GetExpected() (*cluster.Cluster, error)
	Reconcile(actualCluster, expectedCluster *cluster.Cluster) (*cluster.Cluster, error)
	Destroy() error
}

type Resource interface {
	Actual(known *cluster.Cluster) (Resource, error)
	Expected(known *cluster.Cluster) (Resource, error)
	Apply(actual, expected Resource) (Resource, error)
	Delete(actual Resource) error
	Render(renderResource Resource, renderCluster *cluster.Cluster) (*cluster.Cluster, error)
	Tag(tags map[string]string) error
}
