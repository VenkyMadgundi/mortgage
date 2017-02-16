package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Insurer_20170215_195918 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Insurer_20170215_195918{}
	m.Created = "20170215_195918"
	migration.Register("Insurer_20170215_195918", m)
}

// Run the migrations
func (m *Insurer_20170215_195918) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE insurer(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Insurer_20170215_195918) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `insurer`")
}
