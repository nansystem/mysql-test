package usecase

import (
	"log"
	"math/rand"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"

	"github.com/nansystem/mysql-test/common"
	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/infra"
)

func FillCvs(count int64, ads []generated.Ad) ([]generated.Cv, error) {
	cvs := createMockCvs(count, ads)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(cvs), 2000) {
		q := sq.Insert("cv").Columns("id", "ad_id", "user_id", "status", "created_at")
		for _, cv := range cvs[idx.From:idx.To] {
			q = q.Values(cv.ID, cv.AdID, cv.UserID, cv.Status, cv.CreatedAt)
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
	return cvs, nil
}

func createMockCvs(count int64, ads []generated.Ad) []generated.Cv {
	start := time.Date(2021, 3, 23, 0, 0, 0, 0, common.JP())
	end := time.Date(2021, 4, 23, 0, 0, 0, 0, common.JP())

	cvs := make([]generated.Cv, count)
	for i := int64(0); i < count; i++ {
		cvs[i] = newMockCv(
			int(i+1),
			ads[rand.Intn(len(ads))].ID,
			int(common.RandNum(1, 100000)),
			int8(common.RandNum(1, 10)),
			int(common.RandUnixTime(start, end)),
		)
	}
	return cvs
}

func newMockCv(ID, adID, userID int, status int8, createdAt int) generated.Cv {
	return generated.Cv{
		ID:        ID,
		AdID:      adID,
		UserID:    userID,
		Status:    status,
		CreatedAt: createdAt,
	}
}
