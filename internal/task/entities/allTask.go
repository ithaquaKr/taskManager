package entities

import "github.com/ithaquaKr/taskManager/pkg/utils"

type AllTask struct {
	Paginate utils.PaginationResponse
	Result   []*Task `json:"result"`
}
