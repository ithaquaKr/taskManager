package entities

import "github.com/ithaquaKr/taskManager/pkg/utils"

type AllNote struct {
	Paginate utils.PaginationResponse
	Result   []*Note `json:"result"`
}
