package svc

import (
	"github.com/Nozomi9967/wmss-user-rpc/internal/config"
	"github.com/Nozomi9967/wmss-user-rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	AuthMiddleware rest.Middleware

	SysUserModel           model.SysUserModel
	SysRoleModel           model.SysRoleModel
	SysPermissionModel     model.SysPermissionModel
	SysRolePermissionModel model.SysRolePermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		// 初始化 Model（无缓存版本）
		SysUserModel:           model.NewSysUserModel(conn),
		SysRoleModel:           model.NewSysRoleModel(conn),
		SysPermissionModel:     model.NewSysPermissionModel(conn),
		SysRolePermissionModel: model.NewSysRolePermissionModel(conn),
	}
}
