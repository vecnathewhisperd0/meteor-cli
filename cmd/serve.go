package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/odpf/meteor/api"
	"github.com/odpf/meteor/config"
	"github.com/odpf/meteor/extractors"
	"github.com/odpf/meteor/processors"
	"github.com/odpf/meteor/recipes"
	"github.com/odpf/meteor/sinks"
	"github.com/odpf/meteor/stores"
)

func Serve() {
	var err error

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	recipeStore := initRecipeStore(config.RecipeStorageURL)
	extractorStore := initExtractorStore()
	processorStore := initProcessorStore()
	sinkStore := initSinkStore()
	recipeService := recipes.NewService(
		recipeStore,
		extractorStore,
		processorStore,
		sinkStore,
	)

	recipeHandler := api.NewRecipeHandler(recipeService)
	router := api.NewRouter()
	api.SetupRoutes(router, recipeHandler)

	fmt.Println("Listening on port :" + config.Port)
	err = http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		fmt.Println(err)
	}
}
func initRecipeStore(recipeStorageURL string) recipes.Store {
	store, err := stores.NewRecipeStore(recipeStorageURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	return store
}
func initExtractorStore() *extractors.Store {
	store := extractors.NewStore()
	extractors.PopulateStore(store)
	return store
}
func initProcessorStore() *processors.Store {
	store := processors.NewStore()
	processors.PopulateStore(store)
	return store
}
func initSinkStore() *sinks.Store {
	store := sinks.NewStore()
	sinks.PopulateStore(store)
	return store
}
