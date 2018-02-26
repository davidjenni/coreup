package vm

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ClusterSet describes the build-out info for a set of cluster machines
type ClusterSet struct {
	Basename      string `yaml:"basename"`
	NodeCount     int    `yaml:"nodeCount"`
	VMOptionsFile string `yaml:"vmOptionsFile,omitempty"`
	VMOptions     Config `yaml:"-"`
}

// Cluster describes the build-out info for 2 cluster sets:
//		manager, worker
type Cluster struct {
	Manager ClusterSet `yaml:"manager"`
	Worker  ClusterSet `yaml:"worker"`
}

// GetDefaultCluster returns a simple default cluster config
func GetDefaultCluster() *Cluster {
	return &Cluster{
		Manager: ClusterSet{
			Basename:      "minion-mgr-%02d",
			NodeCount:     3,
			VMOptionsFile: "",
		},
		Worker: ClusterSet{
			Basename:      "minion-%02d",
			NodeCount:     2,
			VMOptionsFile: "",
		},
	}
}

// LoadClusterConfig deserializes a yaml file with the cluster configuration
func LoadClusterConfig(clusterConfigFile string) (*Cluster, error) {
	var err error
	buffer, err := ioutil.ReadFile(clusterConfigFile)
	if err != nil {
		return nil, err
	}

	var cluster Cluster
	err = yaml.Unmarshal(buffer, &cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}
