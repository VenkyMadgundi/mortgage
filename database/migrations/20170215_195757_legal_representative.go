package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LegalRepresentative_20170215_195757 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LegalRepresentative_20170215_195757{}
	m.Created = "20170215_195757"
	migration.Register("LegalRepresentative_20170215_195757", m)
}

// Run the migrations
func (m *LegalRepresentative_20170215_195757) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE legal_representative(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`representative_type` int(11) DEFAULT NULL,`wallet_id` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *LegalRepresentative_20170215_195757) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `legal_representative`")
}
