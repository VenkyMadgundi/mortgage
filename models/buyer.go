package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Buyer struct {
	Id                             int64     `orm:"auto"`
	BuyerId                        string    `orm:"size(128)"`
	BuyerName                      string    `orm:"size(128)"`
	BuyerCreation                  time.Time `orm:"type(datetime)"`
	Passport                       string    `orm:"size(128)"`
	DrivingLicence                 string    `orm:"size(128)"`
	NationalId                     string    `orm:"size(128)"`
	EmploymentStatus               string    `orm:"size(128)"`
	AnnualEmploymentIncome         float64
	AnnualIncomeFromOtherSource    float64
	LifeAndSicknessInsuranceCover  float64
	MoFood                         float64
	MoChildCare                    float64
	MoDependents                   float64
	MoUtilityBills                 float64
	MoPhoneBroadbandTv             float64
	MoTravelEntertainment          float64
	MoBuildingAndContentsInsurance float64
	Net_Income                     float64
	WalletBalance                  float64
	WalletId                       string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Buyer))
}

// AddBuyer insert a new Buyer into database and returns
// last inserted Id on success.
func AddBuyer(m *Buyer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBuyerById retrieves Buyer by Id. Returns error if
// Id doesn't exist
func GetBuyerById(id int64) (v *Buyer, err error) {
	o := orm.NewOrm()
	v = &Buyer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBuyer retrieves all Buyer matches certain condition. Returns empty list if
// no records exist
func GetAllBuyer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Buyer))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Buyer
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBuyer updates Buyer by Id and returns error if
// the record to be updated doesn't exist
func UpdateBuyerById(m *Buyer) (err error) {
	o := orm.NewOrm()
	v := Buyer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBuyer deletes Buyer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBuyer(id int64) (err error) {
	o := orm.NewOrm()
	v := Buyer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Buyer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
