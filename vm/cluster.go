package vm

import (
	"errors"
	"fmt"
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
	machines      []*Machine `yaml:"-"`
	runner        Runner     `yaml:"-"`
}

// GetDefaultCluster returns a simple default cluster config
func GetDefaultCluster(runnerOpt ...Runner) *Cluster {
	providerName := "digitalocean"
	provider, _ := GetDefaultProviderConfig(providerName)
	cluster := &Cluster{
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
	if len(runnerOpt) > 0 {
		cluster.runner = runnerOpt[0]
	}
	cluster.newClusterMachines()
	return cluster
}

// GetMachines returns list of machines and their configuration
func (c *Cluster) GetMachines() []*Machine {
	return c.machines
}

// LoadClusterConfig deserializes a yaml file with the cluster configuration
func LoadClusterConfig(clusterConfigFile string, runnerOpt ...Runner) (*Cluster, error) {
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
	if len(runnerOpt) > 0 {
		cluster.runner = runnerOpt[0]
	}

	cluster.Manager.providerConfig, err = cluster.loadOptions(basePath, cluster.Manager.VMOptionsFile)
	if err != nil {
		return nil, err
	}
	cluster.Worker.providerConfig, err = cluster.loadOptions(basePath, cluster.Worker.VMOptionsFile)
	if err != nil {
		return nil, err
	}
	err = cluster.newClusterMachines()
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

// CreateVMs creates a cluster of VMs on the cloud provider
func (c *Cluster) CreateVMs(cloudAPIToken string) error {
	for _, m := range c.machines {
		m.config.CloudAPIToken = cloudAPIToken
		err := m.CreateMachine()
		if err != nil {
			return err
		}
	}
	return nil
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

func (c *Cluster) newClusterMachines() error {
	vmList := make([]*Machine, 0, c.Manager.NodeCount+c.Worker.NodeCount)
	var vms []*Machine
	var err error

	vms, err = c.newMachines(c.Manager.Basename, c.Manager.NodeCount, c.Manager.providerConfig)
	if err != nil {
		return err
	}
	vmList = append(vmList, vms...)

	vms, err = c.newMachines(c.Worker.Basename, c.Worker.NodeCount, c.Worker.providerConfig)
	if err != nil {
		return err
	}
	vmList = append(vmList, vms...)

	c.machines = vmList
	return nil
}

func (c *Cluster) newMachines(basename string, count int, providerConfig ConfigRenderer) ([]*Machine, error) {
	var vmConfig *Config
	var err error

	vmList := make([]*Machine, count)

	for i := 0; i < count; i++ {
		vmConfig, err = NewConfigFrom(fmt.Sprintf(basename, i+1), c.CloudProvider, providerConfig)
		if err != nil {
			return nil, err
		}
		vmList[i] = NewMachine(vmConfig, c.runner)
	}
	return vmList, nil
}
