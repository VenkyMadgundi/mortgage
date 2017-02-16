package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Lender_20170215_195654 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Lender_20170215_195654{}
	m.Created = "20170215_195654"
	migration.Register("Lender_20170215_195654", m)
}

// Run the migrations
func (m *Lender_20170215_195654) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE lender(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Lender_20170215_195654) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `lender`")
}
