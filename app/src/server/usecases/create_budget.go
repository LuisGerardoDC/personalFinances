package usecases

import (
	"fmt"

	models "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
)

type CreateNewBudget struct{}

func (c *CreateNewBudget) CreateNewBudget(b models.Budget) (models.Budget, error) {
	fmt.Println("Use case Budget Created")
	return b, nil
}
