package impl

import (
	"context"

	"vblog/apps/common"
	"vblog/apps/conf"
	"vblog/apps/ioc"
	"vblog/apps/user"

	"gorm.io/gorm"
)

// func NewUserServiceImpl() *UserServiceImpl {
// 	return &UserServiceImpl{
// 		db: conf.C().MySQL.GetDB(),
// 	}

// }
//
func init(){
	ioc.Controller.Registry(user.AppName,&UserServiceImpl{})
}


type UserServiceImpl struct {
	db *gorm.DB
}
func (i *UserServiceImpl) Init() error {
	// 1. 获取数据库连接
	i.db = conf.C().MySQL.GetDB()
	return nil
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	// 1. 校验请求的合法性
	if err := common.Validate(in); err != nil {
		return nil, err
	}
	// hash password不存储明文密码
	if err := in.HashPassword(); err!= nil {
		return nil, err
	}
	

	// 2. 创建user对象(资源)
	ins := user.NewUser(in)

	// 3. user 对象保持入库
	/*
	  读取数据库配置
	  获取数据库连接
	  操作连接 保证数据
	*/
	// INSERT INTO `users` (`created_at`,`updated_at`,`username`,`password`,`role`,`label`) VALUES (1716623778,1716623778,'admin','123456',0,'{}')
	if err := i.db.WithContext(ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	// 4. 返回保持后的user对象
	return ins, nil

}

//怎么查询用户”： 根据过滤条件去数据库查询
// where 以及 limit
func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	set := &user.UserSet{}
	// 构造一个查询语句
	query := i.db.Model(&user.User{}).WithContext(ctx)
	// 构造where语句
	if in.Username!= "" {
		//返回一个新的query对象,并没有在原有的query对象上进行修改
		query = query.Where("username = ?", in.Username)
	}
	// 构造limit语句
	if err := query.Count(&set.Total).Error; err!= nil {
		return nil, err
	}


	if err := query.Offset(in.Offset()).Limit(in.PageSize).Find(&set.Items).Error;err!= nil {
		return nil, err
	}
	return set, nil
}

