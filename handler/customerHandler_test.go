package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerById(t *testing.T) {

	test := []struct {
		Id      int
		Name    string
		Phone   int
		Address string
	}{
		{1, "Prashant", 8709131744, "Bng"},
		{3, "Hari", 1234567890, "Bng"},
	}
	for _, tc := range test {
		t.Run(tc.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/customer/", nil)
			request = mux.SetURLVars(request, map[string]string{"id": fmt.Sprint(tc.Id)})
			response := httptest.NewRecorder()

			GetCustomerById(response, request)
			//fmt.Println(response.Body)
			got := response.Body.String()
			want := fmt.Sprintf("{\"id\":%v,\"name\":\"%v\",\"phone\":%v,\"address\":\"%v\"}", tc.Id, tc.Name, tc.Phone, tc.Address)
			if got != want {
				t.Errorf("got %v -------------want %v", got, want)
			}
		})
	}
}

func TestAddCustomer(t *testing.T) {

	test := []struct {
		desc       string
		statusCode int
		customer   []byte
	}{
		{"One Customer Added", http.StatusOK, []byte(`{ "name":"rohan","phone":989784,"address": "jh"}`)},
	}

	for _, tc := range test {
		request := httptest.NewRequest("POST", "/customer/insert", bytes.NewReader(tc.customer))
		response := httptest.NewRecorder()

		AddCustomer(response, request)
		got := response.Result()

		if got.StatusCode != tc.statusCode {
			//t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, got.StatusCode, tc.statusCode)
			t.Errorf("got %v -------------want %v", got.StatusCode, tc.statusCode)
		}
	}
}

func TestUpdateCustomer(t *testing.T) {
	test := []struct {
		desc       string
		statusCode int
		customer   []byte
	}{
		{"Update successfully", http.StatusOK, []byte(`{"name":"utsav","phone":9876543,"address":"bhr","id":1}`)},
	}

	for i, tc := range test {
		request := httptest.NewRequest("PUT", "/customer/update", bytes.NewReader(tc.customer))
		response := httptest.NewRecorder()

		UpdateCustomer(response, request)
		got := response.Result()

		if got.StatusCode != tc.statusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, got.StatusCode, tc.statusCode)
		}
	}
}

func TestDeleteCustomerDetails(t *testing.T) {
	test := []struct {
		desc       string
		statusCode int
		id         int
	}{
		{"Deleted Successfully", http.StatusOK, 8},
	}

	for i, tc := range test {
		request := httptest.NewRequest("DELETE", "/customer/delete", nil)
		request = mux.SetURLVars(request, map[string]string{"id": fmt.Sprint(tc.id)})
		response := httptest.NewRecorder()

		DeleteCustomerDetails(response, request)

		got := response.Result()

		if got.StatusCode != tc.statusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, got.StatusCode, tc.statusCode)
		}
	}
}
