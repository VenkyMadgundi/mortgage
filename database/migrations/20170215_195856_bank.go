package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Bank_20170215_195856 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Bank_20170215_195856{}
	m.Created = "20170215_195856"
	migration.Register("Bank_20170215_195856", m)
}

// Run the migrations
func (m *Bank_20170215_195856) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE bank(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Bank_20170215_195856) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `bank`")
}
