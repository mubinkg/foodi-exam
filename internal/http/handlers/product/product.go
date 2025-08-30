package product

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mubinkg/foodi-exam/internal/storage"
	"github.com/mubinkg/foodi-exam/internal/types"
	"github.com/mubinkg/foodi-exam/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating products")

		var product types.Product
		err := json.NewDecoder(r.Body).Decode(&product)

		if errors.Is(err, io.EOF) {
			slog.Error("empty body")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			slog.Error("failed to decode request body", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(product); err != nil {
			slog.Error("validation failed", "error", err)
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		id, err := storage.CreateProduct(product.Title, product.Body, product.Price)
		if err != nil {
			slog.Error("failed to create product", "error", err)
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok", "id": fmt.Sprintf("%d", id)})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting product by ID")

		id := r.PathValue("id")
		if id == "" {
			slog.Error("missing id")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("missing id")))
			return
		}

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("invalid id", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		product, err := storage.GetProductById(intId)
		if err != nil {
			slog.Error("failed to get product", "error", err)
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, product)
	}
}


func GetAll(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting all products")

		products, err := storage.GetAllProducts()
		if err != nil {
			slog.Error("failed to get products", "error", err)
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, products)
	}
}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Updating product")

		id := r.PathValue("id")
		if id == "" {
			slog.Error("missing id")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("missing id")))
			return
		}

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("invalid id", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		var product types.Product
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			slog.Error("failed to decode request body", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(product); err != nil {
			slog.Error("validation failed", "error", err)
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		product.Id = intId
		err = storage.UpdateProduct(product.Id, product.Title, product.Body, product.Price)
		if err != nil {
			slog.Error("failed to update product", "error", err)
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"success": "ok"})
	}
}

func Search(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Searching products")

		query := r.URL.Query().Get("q")
		sort := r.URL.Query().Get("sort")

		if query == "" {
			slog.Error("missing query")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("missing query")))
			return
		}

		products, err := storage.SearchProducts(query, sort)
		if err != nil {
			slog.Error("failed to search products", "error", err)
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, products)
	}
}
