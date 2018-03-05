package vm

import (
	"errors"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

// ClusterSet describes the build-out info for a set of cluster machines
type ClusterSet struct {
	Basename       string         `yaml:"basename"`
	NodeCount      int            `yaml:"nodeCount"`
	VMOptionsFile  string         `yaml:"vmOptionsFile,omitempty"`
	providerConfig ConfigRenderer `yaml:"-"`
}

// Cluster describes the build-out info for 2 cluster sets:
//		manager, worker
type Cluster struct {
	CloudProvider string     `yaml:"cloudProvider"`
	Manager       ClusterSet `yaml:"manager"`
	Worker        ClusterSet `yaml:"worker"`
}

// GetDefaultCluster returns a simple default cluster config
func GetDefaultCluster() *Cluster {
	providerName := "digitalocean"
	provider, _ := GetDefaultProviderConfig(providerName)
	return &Cluster{
		CloudProvider: providerName,
		Manager: ClusterSet{
			Basename:       "minion-mgr-%02d",
			NodeCount:      3,
			VMOptionsFile:  "",
			providerConfig: provider,
		},
		Worker: ClusterSet{
			Basename:       "minion-%02d",
			NodeCount:      2,
			VMOptionsFile:  "",
			providerConfig: provider,
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

	var basePath = path.Dir(clusterConfigFile)
	var cluster Cluster
	err = yaml.Unmarshal(buffer, &cluster)
	if err != nil {
		return nil, err
	}
	// TODO: add cloud provider factory
	if cluster.CloudProvider != "digitalocean" {
		return nil, errors.New("Currently, the sole supported cloud provider is 'digitalocean'")
	}
	cluster.Manager.providerConfig, err = cluster.loadOptions(basePath, cluster.Manager.VMOptionsFile)
	if err != nil {
		return nil, err
	}
	cluster.Worker.providerConfig, err = cluster.loadOptions(basePath, cluster.Worker.VMOptionsFile)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

func (c Cluster) loadOptions(basePath string, optionsFile string) (ConfigRenderer, error) {
	if optionsFile == "" {
		return GetDefaultProviderConfig(c.CloudProvider)
	}
	var fqnOptsFile string
	if path.IsAbs(optionsFile) {
		fqnOptsFile = optionsFile
	} else {
		fqnOptsFile = path.Join(basePath, optionsFile)
	}
	return LoadProviderConfig(c.CloudProvider, fqnOptsFile)
}
