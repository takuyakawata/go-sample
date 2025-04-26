package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sago-sample/internal/handler"
	"sago-sample/internal/product/infrastructure"
	usecase "sago-sample/internal/product/usecase"
)

package main

import (
"net/http"
"github.com/go-chi/chi/v5"

"myapp/internal/handler"
"myapp/internal/product/infrastructure"
"myapp/internal/product/usecase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// DI
	repo            := infrastructure.NewProductRepository()
	ucGetAll        := usecase.NewGetAllProductsUseCase(repo)
	ucGetByID       := usecase.NewGetProductByIDUseCase(repo)
	ucCreate        := usecase.NewCreateProductUseCase(repo)
	//ucUpdate        := usecase.NewUpdateProductUseCase(repo)
	//ucDelete        := usecase.NewDeleteProductUseCase(repo)
	//ucAddCategory   := usecase.NewAddCategoryToProductUseCase(repo)
	//ucRemoveCategory:= usecase.NewRemoveCategoryFromProductUseCase(repo)

	// ハンドラー生成
	hGet    := handler.NewGetProductHandler(ucGetAll, ucGetByID)
	hCreate := handler.NewCreateProductHandler(ucCreate)
	// …他のハンドラーも同様…

	// ルーター
	rtr := chi.NewRouter()
	rtr.Get("/api/products",            hGet.HandleGetAll)
	rtr.Get("/api/products/{id}",       hGet.HandleGetByID)
	rtr.Post("/api/products",           hCreate.Handle)
	// …PUT, DELETE, category 周り…

	rtr.ServeHTTP(w, r)
}
