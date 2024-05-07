package service

import (
	"strconv"
	"testing"
	"time"

	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/dto"
	"github.com/agusheryanto182/go-inventory-management/module/feature/staff/repository"
	"github.com/agusheryanto182/go-inventory-management/utils/database"
	"github.com/stretchr/testify/assert"
)

func TestRegister1000(t *testing.T) {
	db, err := database.InitDatabase()
	if err != nil {
		t.Fatal(err)
	}

	staffRepo := repository.NewStaffRepository(db)
	staffSvc := NewStaffService(staffRepo)

	start := time.Now()
	for i := 0; i < 500; i++ {
		_, err = staffSvc.Register(&dto.StaffRegisterReq{
			Name:        "agus" + strconv.Itoa(i),
			PhoneNumber: "082191000" + strconv.Itoa(i),
			Password:    "password",
		})
		assert.NoError(t, err)
	}
	end := time.Now()

	println("elapsed:", end.Sub(start).Milliseconds(), "ms")
}
