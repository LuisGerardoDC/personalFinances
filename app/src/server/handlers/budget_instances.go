package handlers

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/LuisGerardoDC/personalFinances/app/src/utils"
)

var (
	// database
	dbConnection = utils.ConnectDB()

	// useCases
	useCaseCreateNewBudget = usecases.CreateNewBudget{
		DB: dbConnection,
	}
	useCaseAddRecord = usecases.AddRecord2Budget{
		DB: dbConnection,
	}

	useCaseGetBudget = usecases.GetBudget{
		DB: dbConnection,
	}

	useCaseDeleteRecord = usecases.RemoveRecord{
		DB: dbConnection,
	}

	useCaseGetUserBudgets = usecases.GetUserBudgets{
		DB: dbConnection,
	}

	// handlers
	ImplementedNewBudgetHandler = &NewBudgetHandler{
		useCase: &useCaseCreateNewBudget,
	}

	ImplementedAddRecordHandler = &AddRecordHandler{
		useCase: &useCaseAddRecord,
	}

	ImplementedGetBudget = &GetBudgetHandler{
		useCase: &useCaseGetBudget,
	}

	ImplementedDeleteRecordHandler = &DeleteRecordHandler{
		useCase: &useCaseDeleteRecord,
	}

	ImplementedGetBudgets = &GetUserBudgetsHandler{
		useCase: &useCaseGetUserBudgets,
	}
)
