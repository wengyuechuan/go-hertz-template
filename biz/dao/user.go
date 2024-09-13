package dao

import (
	"context"
	"github.com/RanFeng/ierror"
	"github.com/RanFeng/ilog"
	"uav/biz/consts"
	"uav/biz/infra"
	"uav/biz/model"
)

func QueryUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := infra.MysqlDB.WithContext(ctx).Debug().
		Where("user_name = ? OR email = ? OR tel = ? ", username).
		First(user).Error
	if err != nil {
		ilog.EventError(ctx, err, "dao_get_user_by_username_error", "username", username)
		return nil, ierror.NewIError(consts.DBError, err.Error())
	}

	return user, nil
}
