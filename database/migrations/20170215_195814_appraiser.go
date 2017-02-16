package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Appraiser_20170215_195814 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Appraiser_20170215_195814{}
	m.Created = "20170215_195814"
	migration.Register("Appraiser_20170215_195814", m)
}

// Run the migrations
func (m *Appraiser_20170215_195814) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE appraiser(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Appraiser_20170215_195814) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `appraiser`")
}
