package usecase

import (
	"database/sql"
	"time"

	"github.com/cheggaaa/pb/v3"

	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/utils"
)

func createMockCv() *generated.Cv {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	return &generated.Cv{
		AdID:   int(utils.RandNum(1, 100000)),
		UserID: int(utils.RandNum(1, 100000)),
		Status: int8(utils.RandNum(1, 4)),
		// OPTIMIZE
		CreatedAt: int(utils.RandUnixTime(time.Date(2021, 4, 20, 0, 0, 0, 0, loc), time.Date(2026, 4, 20, 0, 0, 0, 0, loc))),
	}
}

func createMockCvs(count int64) []*generated.Cv {
	cvs := make([]*generated.Cv, count)
	for i := int64(0); i < count; i++ {
		cvs[i] = createMockCv()
	}
	return cvs
}

func FillCvs(db *sql.DB, count int64) error {
	bar := pb.StartNew(int(count))
	cvs := createMockCvs(count)
	for _, cv := range cvs {
		err := cv.Insert(db)
		bar.Increment()
		if err != nil {
			return err
		}
	}
	bar.Finish()
	return nil
}
