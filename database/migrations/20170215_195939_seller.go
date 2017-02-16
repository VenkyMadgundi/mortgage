package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Seller_20170215_195939 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Seller_20170215_195939{}
	m.Created = "20170215_195939"
	migration.Register("Seller_20170215_195939", m)
}

// Run the migrations
func (m *Seller_20170215_195939) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE seller(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Seller_20170215_195939) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `seller`")
}
