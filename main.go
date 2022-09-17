package main

import (
	"go-cleanarchitecture/infrastructure"
	"go-cleanarchitecture/interfaces"
	"go-cleanarchitecture/usecases"
	"net/http"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("/tmp/db.sqlite")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	usersRepo := interfaces.NewDbUserRepo(handlers)
	ordersRepo := interfaces.NewDbOrderRepo(handlers)
	itemsRepo := interfaces.NewDbItemRepo(handlers)
	logger := infrastructure.NewLogger()
	orderInteractor := usecases.NewOrderInteractor(
		usersRepo,
		ordersRepo,
		itemsRepo,
		logger,
	)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
