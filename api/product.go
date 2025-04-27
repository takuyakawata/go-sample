package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sago-sample/feature/product/handler"
	"sago-sample/feature/product/infrastructure"
	usecase "sago-sample/feature/product/usecase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 1) DI: リポジトリ → ユースケース → ハンドラーを組み立て
	repo := infrastructure.NewProductRepository()
	ucGetAll := usecase.NewGetAllProductsUseCase(repo)
	ucGetByID := usecase.NewGetProductUseCase(repo)
	//ucCreate := usecase.NewCreateProductUseCase(repo)
	//ucUpdate := usecase.NewUpdateProductUseCase(repo)
	//ucDelete := usecase.NewDeleteProductUseCase(repo)
	//ucAddCat := usecase.NewAddCategoryToProductUseCase(repo)
	//ucRemCat := usecase.NewRemoveCategoryFromProductUseCase(repo)
	// 必要なら他のユースケースも…

	hGet := handler.NewGetProductHandler(ucGetByID, ucGetAll)
	//hCreate := handler.NewCreateProductHandler(ucCreate)
	//hUpdate := handler.NewUpdateProductHandler(ucUpdate)
	//hDelete := handler.NewDeleteProductHandler(ucDelete)
	//hCat := handler.NewCategoryHandler(ucAddCat, ucRemCat)
	// それぞれ internal/handler/*.go ファイルに定義しておきます

	// 2) chi ルーターにパスを定義
	rtr := chi.NewRouter()

	// Product
	//rtr.Get("/api/products", hGet.HandleGetAll)       // GET  /api/products
	rtr.Get("/api/products/{id}", hGet.HandleGetByID) // GET  /api/products/{id}
	//rtr.Post("/api/products", hCreate.Handle)         // POST /api/products
	//rtr.Put("/api/products/{id}", hUpdate.Handle)     // PUT  /api/products/{id}
	//rtr.Delete("/api/products/{id}", hDelete.Handle)  // DELETE /api/products/{id}

	// Category on Product
	//rtr.Post("/api/products/{id}/categories", hCat.HandleAdd)            // POST   /api/products/{id}/categories
	//rtr.Delete("/api/products/{id}/categories/{cid}", hCat.HandleRemove) // DELETE /api/products/{id}/categories/{cid}

	// Optional: List by category
	//rtr.Get("/api/categories/{id}/products", hGet.HandleByCategory)

	//hello world
	rtr.Get("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// 3) エントリポイントにリクエストを渡す
	rtr.ServeHTTP(w, r)
}
