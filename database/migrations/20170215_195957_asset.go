package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Asset_20170215_195957 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Asset_20170215_195957{}
	m.Created = "20170215_195957"
	migration.Register("Asset_20170215_195957", m)
}

// Run the migrations
func (m *Asset_20170215_195957) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE asset(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`address2` varchar(128) ,`city` varchar(128) ,`state` varchar(128) ,`country` varchar(128) ,`postal_code` varchar(128) ,`current_owner` varchar(128) ,`real_estate_value` float ,`seller_id` varchar(128) ,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Asset_20170215_195957) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `asset`")
}
