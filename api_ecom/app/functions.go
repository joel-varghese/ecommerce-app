package app

import (
	"api_ecom/handler"
	"net/http"
)

func (a *App) AddToCart(w http.ResponseWriter, r *http.Request) {
	handler.AddToCart(a.DB, w, r)
}

func (a *App) GetCartPrice(w http.ResponseWriter, r *http.Request) {
	handler.GetCartPrice(a.DB, w, r)
}

func (a *App) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	handler.DeleteFromCart(a.DB, w, r)
}

func (a *App) getAllProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProducts(a.DB, w, r)
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	handler.CreateProduct(a.DB, w, r)
}

func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	handler.GetProduct(a.DB, w, r)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	handler.GetProducts(a.DB, w, r)
}

func (a *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProduct(a.DB, w, r)
}

func (a *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProduct(a.DB, w, r)
}

func (a *App) GetProductByCategory(w http.ResponseWriter, r *http.Request) {
	handler.GetProductByCategory(a.DB, w, r)
}

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	handler.RegisterUser(a.DB, w, r)
}

func (a *App) findUser(w http.ResponseWriter, r *http.Request) {
	handler.FindUser(a.DB, w, r)
}

func (a *App) findUsers(w http.ResponseWriter, r *http.Request) {
	handler.FindUsers(a.DB, w, r)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	handler.UpdateUser(a.DB, w, r)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}
/*
func (a *App) loginUser(w http.ResponseWriter, r *http.Request) {
	handler.LoginUser(a.DB, w, r)
}

func (a *App) SendSmsVerfication(w http.ResponseWriter, r *http.Request) {
	handler.SendSmsVerfication(a.DB, w, r)
}

func (a *App) GetOtpFromUser(w http.ResponseWriter, r *http.Request) {
	handler.GetOtpFromUser(a.DB, w, r)
}

func (a *App) RefreshToken(w http.ResponseWriter, r *http.Request) {
	handler.RefreshToken(w, r)
}
*/