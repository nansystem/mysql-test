package usecase

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"
	"github.com/nansystem/mysql-test/common"
	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/infra"
)

func FillAds(count int64) ([]generated.Ad, error) {
	ads := createMockAds(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(ads), 2000) {
		q := sq.Insert("ad").
			Columns("id", "title", "description", "type", "start_at", "end_at", "updated_at")
		for _, ad := range ads[idx.From:idx.To] {
			q = q.Values(ad.ID, ad.Title, ad.Description, ad.Type, ad.StartAt, ad.EndAt, ad.UpdatedAt)
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
	return ads, nil
}

func createMockAds(count int64) []generated.Ad {
	start := time.Date(2021, 3, 23, 0, 0, 0, 0, common.JP())
	end := time.Date(2021, 4, 23, 0, 0, 0, 0, common.JP())

	ads := make([]generated.Ad, count)
	for i := int64(0); i < count; i++ {
		ads[i] = newMockAd(
			int(i+1),
			common.RandomString(10, 25),
			common.RandomString(50, 100),
			int8(common.RandNum(1, 10)),
			int8(common.RandNum(1, 10)),
			int(common.RandUnixTime(start, end)),
			int(common.RandUnixTime(start, end)),
			int(common.RandUnixTime(start, end)),
		)
	}
	return ads
}

func newMockAd(id int, title, desc string, tp int8, status int8, start, end, updated int) generated.Ad {
	return generated.Ad{
		ID:          id,
		Title:       title,
		Description: desc,
		Type:        tp,
		StartAt:     start,
		EndAt:       end,
		UpdatedAt:   updated,
	}
}
