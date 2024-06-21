package handlers

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/LuisGerardoDC/personalFinances/app/src/utils"
)

var (
	// database
	dbConnection = utils.ConnectDB()

	// dependencies
	useCaseCreateNewBudget = usecases.CreateNewBudget{
		DB: dbConnection,
	}

	// handlers
	ImplementedNewBudgetHandler = &NewBudgetHandler{
		useCase: &useCaseCreateNewBudget,
	}
)
