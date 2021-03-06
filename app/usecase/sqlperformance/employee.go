package sqlperformance

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"
	"github.com/mattn/go-gimei"

	"github.com/nansystem/mysql-test/common"
	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/infra"
)

func FillEmployees(count int64) ([]generated.Employee, error) {
	list := createEmployees(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(list), 3000) {
		q := sq.Insert("employees").Columns("employee_id", "subsidiary_id", "first_name", "last_name", "date_of_birth", "phone_number")
		for _, item := range list[idx.From:idx.To] {
			q = q.Values(item.EmployeeID, item.SubsidiaryID, item.FirstName, item.LastName, item.DateOfBirth, item.PhoneNumber)
		}
		sql, args, err := q.ToSql()
		if err != nil {
			return nil, err
		}
		_, err = infra.DB.Exec(sql, args...)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		bar.Set(idx.To)
	}
	bar.Finish()
	return list, nil
}

func createEmployees(count int64) []generated.Employee {
	minDate := common.CreateDate(1960, 1, 1)
	maxDate := common.CreateDate(2021, 12, 31)

	list := make([]generated.Employee, count)
	for i := int64(0); i < count; i++ {
		name := gimei.NewName()
		d := common.RandTime(minDate, maxDate)

		list[i] = newEmployee(
			uint(i+1),
			uint(common.RandNum(1, 1000)),
			name.First.Kanji(),
			name.Last.Kanji(),
			d,
			common.RandPhoneNumber(),
		)
	}
	return list
}

func newEmployee(ID, subsidiaryID uint, firstName, lastName string, dateOfBirth time.Time, phoneNumber string) generated.Employee {
	return generated.Employee{
		EmployeeID:   ID,
		SubsidiaryID: subsidiaryID,
		FirstName:    firstName,
		LastName:     lastName,
		DateOfBirth:  dateOfBirth,
		PhoneNumber:  phoneNumber,
	}
}
