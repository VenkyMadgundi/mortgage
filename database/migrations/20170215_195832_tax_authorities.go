package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TaxAuthorities_20170215_195832 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TaxAuthorities_20170215_195832{}
	m.Created = "20170215_195832"
	migration.Register("TaxAuthorities_20170215_195832", m)
}

// Run the migrations
func (m *TaxAuthorities_20170215_195832) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE tax_authorities(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *TaxAuthorities_20170215_195832) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `tax_authorities`")
}
