package routers

import (
	"github.com/astaxie/beego"
	"mortgage/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/buyer", &controllers.BuyerController{}, "get:GetAll")
	beego.Router("/lender", &controllers.LenderController{})
	beego.Router("/legal_representative", &controllers.Legal_representativeController{})
	beego.Router("/appraiser", &controllers.AppraiserController{})
	beego.Router("/tax_authorities", &controllers.Tax_authoritiesController{})
	beego.Router("/bank", &controllers.BankController{})
	beego.Router("/insurer", &controllers.InsurerController{})
	beego.Router("/seller", &controllers.SellerController{})
	beego.Router("/asset", &controllers.AssetController{})
}
