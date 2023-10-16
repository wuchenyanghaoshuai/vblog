package impl

import (
	"context"

	"github.com/wuchenyanghaoshuai/vblog/apps/user"
	"github.com/wuchenyanghaoshuai/vblog/exception"
	"github.com/wuchenyanghaoshuai/vblog/ioc"

	"github.com/wuchenyanghaoshuai/vblog/conf"
	"gorm.io/gorm"
)

// 倒入这个包的时候，直接把这个对象 UserServiceImpl注册给ioc
//注册user业务模块的控制器
func init() {
	ioc.Controller().Registry(&UserServiceImpl{})
}

//var _ user.Service = &UserServiceImpl{}

// 装逼写法，等价于上面的写法
var _ user.Service = (*UserServiceImpl)(nil)

func (i *UserServiceImpl) Init() error {
	i.db = conf.C().Mysql.GetConn().Debug()
	return nil
}

//定义到托管Ioc里面的名称
func (i *UserServiceImpl) Name() string {
	return user.AppName
}

//他是user service 服务的控制器
type UserServiceImpl struct {
	db *gorm.DB
}

// 为什么这块是使用*user.CreateUserRequest，使用user是因为impl.go引用了上层的文件要加目录，*user 是因为函数要求传递指针
func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	//第一步 校验用户参数 判断用户传递过来的是否为空,抽象一个validate方法来专门校验
	if err := req.Validate(); err != nil {
		return nil, err

	}
	//t := time.Now()
	//第二部 生成一个user对象(orm对象)
	ins := user.NewUser(req)
	//第三部 将user对象插入到数据库中
	//	i.db.Migrator().AutoMigrate(&user.User{})
	//create方法会检测有没有tablename方法，这个是orm提供的功能， 也可以写成i.db.table("xxx").Create(ins)
	// gorm:"column:username" 通过struct tag 定义对象的映射关系
	//ctx 取消了，这个数据会保存吗，数据库相应慢，或者用户取消了操作
	//现在存储在数据库里面的密码是明文的，怎么办，需要加密 sha256，hash一下再存储
	//1.加密  通过密钥可以解密 : 1. 对称加密  2. 非对称加密
	//2.hash(不可逆) 
	//3 关于password的使用， 不是 password 解密 以后 再对比
	// 可以比对hash结果。 password.hash_code = req.password.hash_code
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	//fmt.Println(res.Error)
	return ins, nil
}

//删除用户
func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {

	//直接删除
	return i.db.WithContext(ctx).Where("id=?", req.Id).Delete(&user.User{}).Error
	//少了一步检查， 如果没有这个id 也没有报错

}

func (i *UserServiceImpl) DescribeUserRequest(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {

	// 1. 构造查询条件
	// id = ? or username=?
	query := i.db.WithContext(ctx)
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		//通过返回值来修改原来的对象
		query = query.Where("id=?", req.DescribeValue)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("username=?", req.DescribeValue)
	}

	ins := user.NewUser(user.NewCreateUserRequest())

	if err := query.First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user %s not found", req.DescribeValue)
		}
	}
	ins.SetIsHashed()
	return ins, nil
}
