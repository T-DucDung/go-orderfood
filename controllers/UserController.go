package controllers

import (
	"encoding/json"
	"go-orderfood/models"
	"go-orderfood/requests"
	"go-orderfood/responses"
	"log"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// Define a struct to return when user is authorized
// Define a struct to return user
//@Title Create Customer
//@Description Create Customer
//@Summary "CreateCustomer"
// @Params auth header string true "token"
// @Params username body requests.CreateCustomer true "username"
//@Success 200
//@Failure 404
//@router /creuser [post]
func (this *UserController) CreateCustomer() {
	defer this.ServeJSON()
	var cus models.Customers
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.CreateCustomer{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	log.Println(err)
	log.Println(req.Type)
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	_, err = cus.CreateCustomer(req)
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	log.Println(err)
	this.Data["json"] = responses.ResCustomer{
		Data:  []models.Customers{},
		Error: responses.NewErr(responses.Success),
	}
	return
}

// Define a struct to return store
//@Title Create Store
//@Description Create Store
//@Summary "CreateStore"
// @Params auth header string true "token"
// @Params username body requests.CreateStore true "username"
//@Success 200
//@Failure 404
//@router /crestore [post]
func (this *UserController) CreateStore() {
	defer this.ServeJSON()
	var store models.Store
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResStore{
			Data:  []models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.CreateStore{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	log.Println(req)
	if err != nil {
		this.Data["json"] = ""
		return
	}
	_, err = store.CreateStore(req)
	if err != nil {
		this.Data["json"] = responses.ResStore{
			Data:  []models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = ""
	return
}

// Define a struct to return shipper
//@Title Create Shipper
//@Description Create Shipper
//@Summary "CreateShipper"
// @Params auth header string true "token"
// @Params username body requests.CreateShipper true "username"
//@Success 200
//@Failure 404
//@router /creshipper [post]
func (this *UserController) CreateShipper() {
	defer this.ServeJSON()
	var ship models.Shipper
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResShipper{
			Data:  []models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	req := requests.CreateShipper{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = ""
		return
	}
	_, err = ship.CreateShipper(req)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responses.ResShipper{
			Data:  []models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = ""
	return
}

// Define a struct to return all Customer
//@Title Get All Customer
//@Description Get All Customer
//@Summary "Get toàn bộ thông tin khách hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResCustomer
//@Failure 404 responses.ResCustomer
//@router /customer [get]
func (this *UserController) GetCustomer() {
	defer this.ServeJSON()
	var customer models.Customers
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	customers, err := customer.GetTotalCustomer()
	if err != nil {
		this.Data["json"] = responses.ResCustomer{
			Data:  []models.Customers{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResCustomer{
		Data:  customers,
		Error: responses.NewErr(responses.Success),
	}
}

// Define a struct to return all Store
//@Title Get All Store
//@Description Get All Store
//@Summary "Get toàn bộ thông tin Cửa hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResStore
//@Failure 404 responses.ResStore
//@router /store [get]
func (this *UserController) GetStore() {
	defer this.ServeJSON()
	var store models.Store
	typeid := this.Ctx.Request.Header.Get("Type")
	log.Println(typeid)
	if typeid != "0" {
		this.Data["json"] = responses.ResStore{
			Data:  []models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	stores, err := store.GetTotalStore()
	if err != nil {
		this.Data["json"] = responses.ResStore{
			Data:  []models.Store{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResStore{
		Data:  stores,
		Error: responses.NewErr(responses.Success),
	}
}

// Define a struct to return all Shipper
//@Title Get All Shipper
//@Description Get All Shipper
//@Summary "Get toàn bộ thông tin người giao hàng"
// @Params auth header string true "token"
//@Success 200 {object} responses.ResShipper
//@Failure 404 responses.ResShipper
//@router /shipper [get]
func (this *UserController) GetShipper() {
	defer this.ServeJSON()
	var shipper models.Shipper
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.ResShipper{
			Data:  []models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	shippers, err := shipper.GetTotalShipper()
	if err != nil {
		this.Data["json"] = responses.ResShipper{
			Data:  []models.Shipper{},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}

	this.Data["json"] = responses.ResShipper{
		Data:  shippers,
		Error: responses.NewErr(responses.Success),
	}
}

// Define a struct
//@Title Update Status Customer
//@Description Update Status Customer
//@Summary "Kích hoat tài khoản khách hàng"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /activeuser [put]
func (this *UserController) ActivateCustomerAPI() {
	defer this.ServeJSON()
	var cus models.Customers
	req := requests.DeactivateAcc{}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}
	
	err = cus.ActivateCustomer(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error at cus.ActivateCustomer(req)"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Activate User Account Success!"}}
	return
}

// Define a struct
//@Title Update Status Customer
//@Description Update Status Customer
//@Summary "Vô hiệu hoá tài khoản khách hàng"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /deactiveuser [put]
func (this *UserController) DeactivateCustomerAPI() {
	defer this.ServeJSON()
	var cus models.Customers
	req := requests.DeactivateAcc{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}
	typeid := this.Ctx.Request.Header.Get("Type")

	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}
	
	err = cus.DeactivateCustomer(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Deactivate User Account Success!"}}
	return
}


// Define a struct
//@Title Update Status Shipper
//@Description Update Status Shipper
//@Summary "Kích tài khoản shipper"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /activeshipper [put]
func (this *UserController) ActivateShipperAPI() {
	defer this.ServeJSON()
	var shipper models.Shipper
	req := requests.DeactivateAcc{}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}
	
	err = shipper.ActivateShipper(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error at shipper.ActivateShipper(req)"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Activate Shipper Account Success!"}}
	return
}

// Define a struct
//@Title Update Status Shipper
//@Description Update Status Shipper
//@Summary "Vô hiệu hoá tài khoản Shipper"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /deactiveshipper [put]
func (this *UserController) DeactivateShipperAPI() {
	defer this.ServeJSON()
	var shipper models.Shipper
	req := requests.DeactivateAcc{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}

	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}
	
	err = shipper.DeactivateShipper(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Deactivate Shipper Account Success!"}}
	return
}


// Define a struct
//@Title Update Status Store
//@Description Update Status Store
//@Summary "Kích tài khoản Store"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /activestore [put]
func (this *UserController) ActivateStoreAPI() {
	defer this.ServeJSON()
	var store models.Store
	req := requests.DeactivateAcc{}
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}
	
	err = store.ActivateStore(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error at store.ActivateStore(req)"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Activate Store Account Success!"}}
	return
}


// Define a struct
//@Title Update Status Store
//@Description Update Status Store
//@Summary "Vô hiệu hoá tài khoản Store"
// @Params auth header string true "token"
// @Param data body requests.DeactivateAcc true "body"
//@Success 200 
//@Failure 404
//@router /deactivestore [put]
func (this *UserController) DeactivateStoreAPI() {
	defer this.ServeJSON()
	var store models.Store
	req := requests.DeactivateAcc{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 401, Message: "Error at json.Unmarshal"}}
		return
	}
	
	typeid := this.Ctx.Request.Header.Get("Type")
	if typeid != "0" {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 403, Message: "Permission Not Allowed"}}
		return
	}
	
	err = store.DeactivateStore(req)
	// log.Println(err)
	if err != nil {
		this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 400, Message: "Error at store.DeactivateStore(req)"}}
		return
	}
	this.Data["json"] = responses.Response{Data: "", Error: &responses.Err{Code: 200, Message: "Deactivate Store Account Success!"}}
	return
}
