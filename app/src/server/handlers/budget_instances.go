package handlers

import "github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"

var (
	// conections

	// database

	// dependencies
	useCaseCreateNewBudget = usecases.CreateNewBudget{}
	// handlers
	ImplementedNewBudgetHandler = &NewBudgetHandler{
		useCase: &useCaseCreateNewBudget,
	}
)
