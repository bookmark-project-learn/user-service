package fixture

import (
	"time"

	baseEntity "github.com/bookmark-project-learn/user-service/internal/models/base"
	"github.com/google/uuid"
)

var BaseTime = time.Date(2025, 5, 5, 5, 40, 0, 0, time.UTC)

func GetBaseEntity(id string) baseEntity.Base {
	if id == "" {
		id = uuid.NewString()
	}
	return baseEntity.Base{
		Id:        id,
		CreatedAt: BaseTime,
		UpdatedAt: BaseTime,
	}
}
