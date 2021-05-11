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

func FillOrderMaster(count int64) ([]generated.YhOrderMaster, error) {
	list := createYhOrderMasters(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(list), 2000) {
		q := sq.Insert("yh_order_master").Columns("id", "order_time", "seller_id", "image_id", "item_id", "is_hidden_page")
		for _, item := range list[idx.From:idx.To] {
			q = q.Values(item.ID, item.OrderTime, item.SellerID, item.ImageID, item.ItemID, item.IsHiddenPage)
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

func createYhOrderMasters(count int64) []generated.YhOrderMaster {
	series := common.RandTimeSeries(time.Now(), 1*time.Minute, 10*time.Minute)
	list := make([]generated.YhOrderMaster, count)
	for i := int64(0); i < count; i++ {
		list[i] = newYhOrderMaster(
			uint(i+1),
			series(),
			uint(common.RandNum(1, 10000)),
			uint(common.RandNum(1, 10000)),
			uint(common.RandNum(1, 10000)),
			common.RandValWeight([]common.ValWeight{
				{Val: int8(0), Weight: 90},
				{Val: int8(1), Weight: 10},
			}).(int8),
		)
	}
	return list
}

func newYhOrderMaster(ID uint, orderTime time.Time, sellerID, imageID, itemID uint, isHiddenPage int8) generated.YhOrderMaster {
	return generated.YhOrderMaster{
		ID:           ID,
		OrderTime:    orderTime,
		SellerID:     sellerID,
		ImageID:      imageID,
		ItemID:       itemID,
		IsHiddenPage: isHiddenPage,
	}
}
