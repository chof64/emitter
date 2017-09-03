/**********************************************************************************
* Copyright (c) 2009-2017 Misakai Ltd.
* This program is free software: you can redistribute it and/or modify it under the
* terms of the GNU Affero General Public License as published by the  Free Software
* Foundation, either version 3 of the License, or(at your option) any later version.
*
* This program is distributed  in the hope that it  will be useful, but WITHOUT ANY
* WARRANTY;  without even  the implied warranty of MERCHANTABILITY or FITNESS FOR A
* PARTICULAR PURPOSE.  See the GNU Affero General Public License  for  more details.
*
* You should have  received a copy  of the  GNU Affero General Public License along
* with this program. If not, see<http://www.gnu.org/licenses/>.
************************************************************************************/

package config

import (
	cfg "github.com/emitter-io/config"
)

// Constants used throughout the service.
const (
	ChannelSeparator = '/'   // The separator character.
	MaxMessageSize   = 65536 // Maximum message size allowed from/to the peer.
)

// NewDefault creates a default configuration.
func NewDefault() cfg.Config {
	return &Config{
		ListenAddr: ":8080",
		Cluster: &ClusterConfig{
			ListenAddr:    ":4000",
			AdvertiseAddr: "public:4000",
			Passphrase:    "emitter-io",
		},
	}
}

// Config represents main configuration.
type Config struct {
	ListenAddr string              `json:"listen"`             // The API port used for TCP & Websocket communication.'
	License    string              `json:"license"`            // The port used for gossip.'
	TLS        *cfg.TLSConfig      `json:"tls,omitempty"`      // The API port used for Secure TCP & Websocket communication.'
	Secrets    *cfg.VaultConfig    `json:"vault,omitempty"`    // The configuration for the Hashicorp Vault.
	Cluster    *ClusterConfig      `json:"cluster,omitempty"`  // The configuration for the clustering.
	Storage    *cfg.ProviderConfig `json:"storage,omitempty"`  // The configuration for the storage provider.
	Contract   *cfg.ProviderConfig `json:"contract,omitempty"` // The configuration for the contract provider.
	Metering   *cfg.ProviderConfig `json:"metering,omitempty"` // The configuration for the usage storage for metering.
}

// Vault returns a vault configuration.
func (c *Config) Vault() *cfg.VaultConfig {
	return c.Secrets
}

// ClusterConfig represents the configuration for the cluster.
type ClusterConfig struct {

	// The name of this node. This must be unique in the cluster. If this is not set, Emitter
	// will set it to the external IP address of the running machine.
	NodeName string `json:"name,omitempty"`

	// The IP address and port that is used to bind the inter-node communication network. This
	// is used for the actual binding of the port.
	ListenAddr string `json:"listen"`

	// The address and port to advertise inter-node communication network. This is used for nat
	// traversal.
	AdvertiseAddr string `json:"advertise"`

	// The seed address (or a domain name) for cluster join.
	Seed string `json:"seed"`

	// Passphrase is used to initialize the primary encryption key in a keyring. This key
	// is used for encrypting all the gossip messages (message-level encryption).
	Passphrase string `json:"passphrase,omitempty"`
}

// LoadProvider loads a provider from the configuration or panics if the configuration is
// specified, but the provider was not found or not able to configure. This uses the first
// provider as a default value.
var LoadProvider = cfg.LoadProvider
