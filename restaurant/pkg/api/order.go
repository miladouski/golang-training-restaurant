package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/miladouski/golang-training-restaurant/restaurant/pkg/data"
)

type orderAPI struct {
	data *data.OrderData
}

func ServeOrderResource(r *mux.Router, data data.OrderData) {
	api := &orderAPI{data: &data}
	r.HandleFunc("/orders", api.getAllOrders).Methods("GET")
	r.HandleFunc("/orders", api.createOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", api.FindOrder).Methods("GET")
	r.HandleFunc("/orders/{id}", api.updateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", api.deleteOrder).Methods("DELETE")
}

func (o orderAPI) getAllOrders(writer http.ResponseWriter, request *http.Request) {
	orders, err := o.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get orders"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
	err = json.NewEncoder(writer).Encode(orders)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (o orderAPI) createOrder(writer http.ResponseWriter, request *http.Request) {
	order := new(data.Order)
	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if order == nil {
		log.Printf("failed empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = o.data.Create(*order)
	if err != nil {
		log.Println("order hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (o orderAPI) updateOrder(writer http.ResponseWriter, request *http.Request) {
	order := new(data.Order)
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	order.Id = int(id)
	err = o.data.Update(order.Id, order.Payment)
	if err != nil {
		log.Println("order hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (o orderAPI) deleteOrder(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = o.data.Delete(int(id))
	if err != nil {
		log.Println("order hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (o orderAPI) FindOrder(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := o.data.Read(int(id))
	if err != nil {
		log.Println("order not found")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(order)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
