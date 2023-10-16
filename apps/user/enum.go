package user

type Role int

const (
	//创建者
	ROLE_AUTHOR Role = iota
	//审核员
	ROLE_AUDITOR
	//系统管理员
	ROLE_ADMIN
)

type DescribeBy int

const (
	DESCRIBE_BY_ID DescribeBy = iota
	DESCRIBE_BY_USERNAME
)
