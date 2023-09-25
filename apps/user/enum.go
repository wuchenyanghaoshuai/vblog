package user

type Role int

const (
	ROLE_MEMBER Role = iota
	ROLE_ADMIN
)
type DescribeBy int

const (
	DESCRIBE_BY_ID DescribeBy = iota
	DESCRIBE_BY_USERNAME
)