package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateAvailabilityStatus(database *gorm.DB) {
	if !database.HasTable(&ent.TableAvailabilityStatus{}) {
		database.CreateTable(&ent.TableAvailabilityStatus{})
	}
}
