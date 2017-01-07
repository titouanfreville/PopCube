// This file is used to test if user model is working correctly.
// A user is always linked to an organisation
// He has basic channel to join
package data_stores

import (
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Test correction test for user ;)

// Test Password functionalities from user Model
func TestDataStores(t *testing.T) {
	Convey("Testing DB initialisation and closure", t, func() {
		Convey("Given a datastore", func() {
			ds := DataStore{}
			Convey("Initialising should provide a db", func() {
				ds.initConnection("root", "popcube_test", "popcube_dev")
				db, _ := gorm.Open("mysql", "root:popcube_dev@/?charset=utf8&parseTime=True&loc=Local")
				// Should nor resemble cause of Channel IDs. 2 Objects have 2 Ids :'('
				So(ds.Db, ShouldNotResemble, db)
				So(ds.Db.Value, ShouldEqual, db.Value)
				So(ds.Db.Error, ShouldEqual, db.Error)
				So(ds.Db.RowsAffected, ShouldEqual, db.RowsAffected)
				// Should not be nill as long as we don't have db on ...
				So(ds.Err, ShouldBeNil)

				ds.initConnection("test_user", "popcube_test", "test")
				db, _ = gorm.Open("mysql", "test_user:test@/?popcube_test?charset=utf8&parseTime=True&loc=Local")
				// Should nor resemble cause of Channel IDs. 2 Objects have 2 Ids :'('
				So(ds.Db, ShouldNotResemble, db)
				So(ds.Db.Value, ShouldEqual, db.Value)
				So(ds.Db.Error, ShouldEqual, db.Error)
				So(ds.Db.RowsAffected, ShouldEqual, db.RowsAffected)
			})
			Convey("Stoping the connection should destroy the Db stored.", func() {
				ds.initConnection("test_user", "popcube_test", "test")
				ds.closeConnection()
				So(ds.Db, ShouldResemble, &gorm.DB{})
			})
		})
	})
}
