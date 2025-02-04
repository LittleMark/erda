// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmdb

import (
	"net/http"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var CMDB_ORG_NEXUS_SHOW_PASSWORD = apis.ApiSpec{
	Path:         "/api/orgs/<orgID>/show-nexus-password",
	BackendPath:  "/api/orgs/<orgID>/show-nexus-password",
	Host:         "cmdb.marathon.l4lb.thisdcos.directory:9093",
	Scheme:       "http",
	Method:       http.MethodGet,
	CheckLogin:   true,
	IsOpenAPI:    true,
	RequestType:  apistructs.OrgNexusShowPasswordRequest{},
	ResponseType: apistructs.OrgNexusShowPasswordResponse{},
	Doc:          "summary: 获取企业 nexus 密码",
}
