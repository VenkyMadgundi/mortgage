package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Buyer_20170215_195426 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Buyer_20170215_195426{}
	m.Created = "20170215_195426"
	migration.Register("Buyer_20170215_195426", m)
}

// Run the migrations
func (m *Buyer_20170215_195426) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE buyer(`id` int(11) NOT NULL AUTO_INCREMENT,`buyer_id` varchar(128) NOT NULL,`buyer_name` varchar(128) NOT NULL,`buyer_creation` datetime,`passport` varchar(128),`driving_licence` varchar(128),`national_id` varchar(128) ,`employment_status` varchar(128) ,`annual_employment_income` float ,`annual_income_from_other_source` float ,`life_and_sickness_insurance_cover` float ,`mo_food` float ,`mo_child_care` float ,`mo_dependents` float ,`mo_utility_bills` float ,`mo_phone_broadband_tv` float ,`mo_travel_entertainment` float ,`mo_building_and_contents_insurance` float ,`net__income` float ,`wallet_balance` float ,`wallet_id` varchar(128) ,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Buyer_20170215_195426) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `buyer`")
}
