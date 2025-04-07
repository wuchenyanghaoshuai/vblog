package common

func NewPageRequest()*PageRequest {
	return &PageRequest{
		PageSize: 10,
		PageNum: 1,
	}
}


type PageRequest struct {
	//分页大小
	PageSize int `json:"page_size"`
	//第几页
	PageNum int `json:"page_num"`
}