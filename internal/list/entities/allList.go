package entities

import (
	"github.com/ithaquaKr/taskManager/pkg/utils"
)

type AllList struct {
	Paginate utils.PaginationResponse
	Result   []*List `json:"result"`
}
