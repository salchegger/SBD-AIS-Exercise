package rest

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"ordersystem/httptools"
	"ordersystem/model"
	"ordersystem/repository"
	"ordersystem/storage"
	"strings"

	"github.com/go-chi/render"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Failure     	500
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allDrinks, err := db.GetDrinks()
		if err != nil {
			slog.Error("Unable to load drinks", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to load drinks")
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, allDrinks)
	}
}

// GetOrders		godoc
// @tags 			Order
// @Description 	Returns all orders
// @Produce  		json
// @Success 		200 {array} model.Order
// @Failure     	500
// @Router 			/api/order/all [get]
func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allOrders, err := db.GetOrders()
		if err != nil {
			slog.Error("Unable to load orders", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to load order")
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, allOrders)
	}
}

// GetOrdersTotal		godoc
// @tags 				Order
// @Description 		Gets totalled orders
// @Produce  			json
// @Success 			200
// @Failure     		500
// @Router 				/api/order/totalled [get]
func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		totalledOrders, err := db.GetTotalledOrders()
		if err != nil {
			slog.Error("Unable to load order totals", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to load order totals")
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, totalledOrders)
	}
}

// GetReceiptFile		godoc
// @tags 				Order
// @Description 		Get receipt for order
// @Produce 			text/markdown
// @Success 			200 {file} markdown file
// @Param 				orderId path int true "Order ID"
// @Failure     		404
// @Failure     		500
// @Router 				/api/receipt/{orderId} [get]
func GetReceiptFile(db *repository.DatabaseHandler, s3 *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uintId, err := httptools.ParseUintUrlParam("orderId", r)
		if err != nil {
			slog.Error("No order id set")
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "No order id set")
			return
		}
		// get order from db
		order, err := db.GetOrder(uintId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, "This order does not exist")
				return
			}
			slog.Error("Unable to load order", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to load order")
			return
		}
		// read from s3
		// todo
		// Get the file from s3 using s3.GetObject(), the bucket name is defined in storage.OrdersBucket
		// dbOrder.Filename() can be used to get the filename.
		// handle any error!
		obj, err := s3.GetObject(
			r.Context(),
			storage.OrdersBucket,
			order.GetFilename(),
			minio.GetObjectOptions{})
		if err != nil {
			slog.Error("Unable to load the receipt from s3.", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to load the receipt.")
			return
		}

		// serve file
		// todo
		// set the correct header on w http.ResponseWriter ("Content-Type" and "Content-Disposition")
		// Use the correct filename for "Content-Disposition" (https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Content-Disposition)
		// io.Copy can be used to write the result of s3.GetObject() to w http.ResponseWriter
		w.Header().Set("Content-Type", "text/markdown")
		w.Header().Set("Content-Disposition", `attachment; filename="`+order.GetFilename()+`"`)

		if _, err := io.Copy(w, obj); err != nil {
			slog.Error("Unable to write receipt to response", slog.String("error", err.Error()))
		}
	}
}

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.Order true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Failure     	500
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler, s3 *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order model.Order
		// read body
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Unable to read body", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Unable to read body")
			return
		}
		err = json.Unmarshal(payload, &order)
		if err != nil {
			slog.Error("Unable to decode body", slog.String("error", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, "Unable to decode body")
			return
		}
		// store to db
		dbOrder, err := db.AddOrder(&order)
		if err != nil {
			slog.Error("Unable to add order to db", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to add order to db")
			return
		}
		// store to s3
		// todo
		// call dbOrder.ToMarkdown() --> use strings.NewReader to create a reader
		// Put the file into s3 using s3.PutObject(), the bucket name is defined in storage.OrdersBucket
		// dbOrder.Filename() can be used to get the filename.
		// Size of the file is determined by the string.
		// Use the following PutObjectOptions: minio.PutObjectOptions{ContentType: "text/markdown"}
		// Handle errors!
		md := dbOrder.ToMarkdown()
		reader := strings.NewReader(md)

		_, err = s3.PutObject(
			r.Context(),
			storage.OrdersBucket,
			dbOrder.GetFilename(),
			reader,
			int64(len(md)),
			minio.PutObjectOptions{ContentType: "text/markdown"},
		)
		if err != nil {
			slog.Error("Unable to upload receipt to S3", slog.String("error", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, "Unable to upload receipt to S3")
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, "ok")
	}
}
