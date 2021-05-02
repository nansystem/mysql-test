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

func FillProducts(count int64) ([]generated.Product, error) {
	list := createMockProducts(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(list), 2000) {
		q := sq.Insert("products").Columns("id", "shop_id", "name", "price", "starts_at", "ends_at")
		for _, item := range list[idx.From:idx.To] {
			q = q.Values(item.ID, item.ShopID, item.Name, item.Price, item.StartsAt, item.EndsAt)
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

func createMockProducts(count int64) []generated.Product {
	baseMin := time.Date(2018, 5, 1, 0, 0, 0, 0, common.JP())
	baseMax := time.Date(2022, 5, 1, 0, 0, 0, 0, common.JP())

	list := make([]generated.Product, count)
	for i := int64(0); i < count; i++ {
		p := common.RandPeriod(baseMin, baseMax, 1, 180)
		list[i] = newMockProduct(
			uint(i+1),
			uint(common.RandNum(1, 1000)),
			common.RandomString(10, 30),
			uint(common.RandNum(5000, 100000)),
			p.Start,
			p.End,
		)
	}
	return list
}

func newMockProduct(ID, shopID uint, name string, price uint, start, end time.Time) generated.Product {
	return generated.Product{
		ID:       ID,
		ShopID:   shopID,
		Name:     name,
		Price:    price,
		StartsAt: start,
		EndsAt:   end,
	}
}
