package controller

import (
	"fmt"
	"goPo/model"
	"goPo/services"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

// FindAll for get all po
func FindAll() []model.Po {
	var session, err = services.Connect()
	var result = []model.Po{}
	if err != nil {
		fmt.Println("Error!", err.Error())
		return result
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("po")
	err = collection.Find(nil).All(&result)

	return result
}

// FindOne for search one po by id
func FindOne(id string) []model.Po {
	var session, err = services.Connect()
	var result = make([]model.Po, 0, 2)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return result
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("po")
	var selector = bson.M{"poId": id}
	err = collection.Find(selector).All(&result)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return result
	}
	return result
}

// Insert is used for input data po
func Insert(isi model.Po) string {
	var session, err = services.Connect()
	var kosong = model.Po{}
	if err != nil || reflect.DeepEqual(isi, kosong) {
		fmt.Println("Error!", err.Error())
		return "input gagal"
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("po")
	err = collection.Insert(isi)
	return "input sukses"
}

// Update for change data po
func Update(isi model.Po) string {
	var session, err = services.Connect()
	var kosong = model.Po{}
	if err != nil || reflect.DeepEqual(isi, kosong) {
		fmt.Println("Error!", err.Error())
		return "input gagal"
	}
	defer session.Close()
	var collection = session.DB("belajar_golang").C("po")
	var id = isi.PoID
	var selector = bson.M{"poId": id}
	err = collection.Update(selector, isi)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return "update gagal"
	}
	return "update sukses"
}

// Delete for remove data from record
func Delete(id string) string {
	var session, err = services.Connect()
	var collection = session.DB("belajar_golang").C("po")
	if err != nil || len(id) < 1 {
		fmt.Println("Error!", err.Error())
		return "input gagal"
	}
	var selector = bson.M{"poId": id}
	err = collection.Remove(selector)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return "delete gagal"
	}
	fmt.Println("Remove success!")
	return "hapus sukses"
}
