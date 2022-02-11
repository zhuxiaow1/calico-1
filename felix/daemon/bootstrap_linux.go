// Copyright (c) 2020-2021 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package daemon

import (
	"github.com/projectcalico/calico/felix/config"
	"github.com/projectcalico/calico/felix/netlinkshim"
	"github.com/projectcalico/calico/felix/wireguard"
	"github.com/projectcalico/calico/libcalico-go/lib/clientv3"

	log "github.com/sirupsen/logrus"
)

func bootstrapWireguard(configParams *config.Config, v3Client clientv3.Interface) error {
	log.Debug("bootstrapping wireguard host connectivity")
	return wireguard.BootstrapHostConnectivity(
		configParams,
		netlinkshim.NewRealWireguard,
		v3Client,
	)
}
