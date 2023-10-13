package blog

type Status int

const (
	STATUS_DRAFT Status = iota
	STATUS_PUBLISHED
)

type UpdateMode int

const (
	//全量更新
	UPDATE_MODE_PUT UpdateMode = iota
	//部分更新(增量更新)
	UPDATE_MODE_PATCH
)
