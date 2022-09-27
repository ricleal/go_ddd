package main

import (
	"go-cleanarchitecture/internal/app"
	"go-cleanarchitecture/internal/app/presentation"
	"go-cleanarchitecture/internal/app/presentation/table"
	"go-cleanarchitecture/internal/app/service"
	"go-cleanarchitecture/internal/app/service/repository/sqlite"
)

func main() {
	// SQLite Repository
	db := sqlite.New(
		sqlite.WithFile("test.db"),
	)
	if err := db.Start(); err != nil {
		panic(err)
	}

	// Service
	service := service.NewPersonService(db)

	// Table
	t := table.NewTable()

	// Presentation
	presenter := presentation.NewPersonPresenter(t)

	// App
	app := app.New(service, presenter)
	app.Run()

}
