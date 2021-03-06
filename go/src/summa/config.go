package summa

import (
	"encoding/json"
	"io/ioutil"
)

type AuthProvider func(username, password string) (*User, error)

type Config struct {
	Listen        string
	SSLEnable     bool
	SessionExpire int64
	AuthProvider  AuthProvider
	DirPaths      map[string]string
	FilePaths     map[string]string
}

var config *Config

// Load server configuration from a file containing a
// JSON object that maps to the config struct
func configLoad(path string) error {
	rawConfigData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	config = new(Config)
	if err := json.Unmarshal(rawConfigData, config); err != nil {
		return err
	}

	return nil
}

func (c *Config) SetAuthProvider(ap AuthProvider) {
	c.AuthProvider = ap
}

func (c *Config) WebRoot() string {
	return c.DirPaths["WebRoot"]
}

func (c *Config) GitRoot() string {
	return c.DirPaths["GitRoot"]
}

func (c *Config) LogFile() string {
	return c.FilePaths["LogFile"]
}

func (c *Config) DBFile() string {
	return c.FilePaths["DBFile"]
}

func (c *Config) SSLCertFile() string {
	return c.FilePaths["SSLCertFile"]
}

func (c *Config) SSLKeyFile() string {
	return c.FilePaths["SSLKeyFile"]
}
