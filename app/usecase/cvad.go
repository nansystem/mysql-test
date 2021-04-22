package usecase

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"

	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/utils"
)

func createMockCv() generated.Cv {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	return generated.Cv{
		AdID:   int(utils.RandNum(1, 100000)),
		UserID: int(utils.RandNum(1, 100000)),
		Status: int8(utils.RandNum(1, 4)),
		// OPTIMIZE
		CreatedAt: int(utils.RandUnixTime(time.Date(2021, 4, 20, 0, 0, 0, 0, loc), time.Date(2026, 4, 20, 0, 0, 0, 0, loc))),
	}
}

func createMockCvs(count int64) []generated.Cv {
	cvs := make([]generated.Cv, count)
	for i := int64(0); i < count; i++ {
		cvs[i] = createMockCv()
	}
	return cvs
}

func FillCvs(db *sql.DB, count int64) error {
	bar := pb.StartNew(int(count))
	cvs := createMockCvs(count)

	for idx := range utils.IndexChunks(len(cvs), 2000) {
		q := sq.Insert("cv").Columns("id", "ad_id", "user_id", "status", "created_at")
		for _, cv := range cvs[idx.From:idx.To] {
			bar.Increment()
			q = q.Values(cv.ID, cv.AdID, cv.UserID, cv.Status, cv.CreatedAt)
		}
		sql, args, err := q.ToSql()
		if err != nil {
			return err
		}
		_, err = db.Exec(sql, args...)
		if err != nil {
			// Error 1390: Prepared statement contains too many placeholders
			log.Fatal(err)
			return err
		}
	}
	bar.Finish()
	return nil
	// for _, cv := range cvs {
	// 	err := cv.Insert(db)
	// 	bar.Increment()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// bar.Finish()
	// return nil
}

// func OutputCvs(db *sql.DB, count int64) error {
// 	cvs := createMockCvs(count)

// 	f, err := os.Create("file.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	w := csv.NewWriter(f)

// 	w.WriteAll(records)

// 	w.Flush()
// 	return nil
// }
