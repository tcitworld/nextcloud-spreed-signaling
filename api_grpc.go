/**
 * Standalone signaling server for the Nextcloud Spreed app.
 * Copyright (C) 2022 struktur AG
 *
 * @author Joachim Bauch <bauch@struktur.de>
 *
 * @license GNU AGPL version 3 or any later version
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package signaling

import (
	"fmt"
)

// Information on a GRPC target in the etcd cluster.

type GrpcTargetInformationEtcd struct {
	Address string `json:"address"`
}

func (p *GrpcTargetInformationEtcd) CheckValid() error {
	if l := len(p.Address); l == 0 {
		return fmt.Errorf("address missing")
	} else if p.Address[l-1] == '/' {
		p.Address = p.Address[:l-1]
	}
	return nil
}
