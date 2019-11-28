package config

type Container struct {
	name      string
	Image     string   `hcl:"image"`
	Command   []string `hcl:"command,optional"`
	Volumes   []Volume `hcl:"volume,block"`
	Network   string   `hcl:"network"`
	IPAddress string   `hcl:"ip_address,optional"`
}

type Volume struct {
	Source      string `hcl:"source"`
	Destination string `hcl:"destination"`
}