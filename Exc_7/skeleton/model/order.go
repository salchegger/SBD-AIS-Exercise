package model

import (
	"fmt"
	"time"
)

const (
	orderFilename = "order_%d.md"

	markdownTemplate = `
# Order: %d

| Created At | Drink ID | Amount |
|------------|----------|--------|
| %s         | %d       | %d     |

Thanks for drinking with us!
`
)

type Order struct {
	Base
	Amount uint64 `json:"amount"`
	// Relationships
	// foreign key
	DrinkID uint  `json:"drink_id" gorm:"not null"`
	Drink   Drink `json:"drink"`
}

func (o *Order) ToMarkdown() string {
	return fmt.Sprintf(markdownTemplate, o.ID, o.CreatedAt.Format(time.Stamp), o.DrinkID, o.Amount)
}

func (o *Order) GetFilename() string {
	return fmt.Sprintf(orderFilename, o.ID)
}
