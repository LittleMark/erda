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

package dao

import (
	"github.com/jinzhu/gorm"

	"github.com/erda-project/erda/apistructs"
)

// LibReference 库引用
type LibReference struct {
	BaseModel
	AppID          uint64
	LibID          uint64
	LibName        string
	LibDesc        string
	ApprovalID     uint64
	ApprovalStatus apistructs.ApprovalStatus
	Creator        string
}

// TableName 表名
func (LibReference) TableName() string {
	return "dice_library_references"
}

// CreateLibReference 创建库引用
func (client *DBClient) CreateLibReference(libReference *LibReference) error {
	return client.Create(libReference).Error
}

// UpdateApprovalStatusByApprovalID 更新审批流状态
func (client *DBClient) UpdateApprovalStatusByApprovalID(approvalID uint64, status apistructs.ApprovalStatus) error {
	return client.Model(LibReference{}).Where("approval_id = ?", approvalID).
		Updates(LibReference{ApprovalStatus: status}).Error
}

// DeleteLibReference 删除库引用
func (client *DBClient) DeleteLibReference(libReferenceID uint64) error {
	return client.Where("id = ?", libReferenceID).Delete(LibReference{}).Error
}

// GetLibReference 库引用详情
func (client *DBClient) GetLibReference(libReferenceID uint64) (*LibReference, error) {
	var libReference LibReference
	if err := client.Where("id = ?", libReferenceID).Find(&libReference).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &libReference, nil
}

// ListLibReference 库引用列表
func (client *DBClient) ListLibReference(req *apistructs.LibReferenceListRequest) (uint64, []LibReference, error) {
	var (
		total         uint64
		libReferences []LibReference
	)
	cond := LibReference{}
	if req.AppID > 0 {
		cond.AppID = req.AppID
	}
	if req.LibID > 0 {
		cond.LibID = req.LibID
	}
	if req.ApprovalStatus != "" {
		cond.ApprovalStatus = req.ApprovalStatus
	}
	if err := client.Where(cond).Offset((req.PageNo - 1) * req.PageSize).Limit(req.PageSize).
		Order("updated_at DESC").Find(&libReferences).
		Offset(0).Limit(-1).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	return total, libReferences, nil
}
