package yh

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"

	"github.com/nansystem/mysql-test/common"
	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/infra"
)

func FillPushInfo(count int64) ([]generated.YhPushInfo, error) {
	list := createPushInfoList(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(list), 2000) {
		q := sq.Insert("yh_push_info").Columns("push_id", "mod_date", "deleted_flg")
		for _, item := range list[idx.From:idx.To] {
			q = q.Values(item.PushID, item.ModDate, item.DeletedFlg)
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

func createPushInfoList(count int64) []generated.YhPushInfo {
	series := common.RandTimeSeries(time.Now(), 1*time.Minute, 10*time.Minute)
	list := make([]generated.YhPushInfo, count)
	for i := int64(0); i < count; i++ {
		list[i] = newPushInfo(
			uint(i+1),
			series(),
			common.RandValWeight([]common.ValWeight{
				{Val: int8(0), Weight: 1},
				{Val: int8(1), Weight: 9},
			}).(int8),
		)
	}
	return list
}

func newPushInfo(pushID uint, modDate time.Time, deletedFlg int8) generated.YhPushInfo {
	return generated.YhPushInfo{
		PushID:     pushID,
		ModDate:    modDate,
		DeletedFlg: deletedFlg,
	}
}
