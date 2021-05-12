package yh

import (
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/cheggaaa/pb"
	"github.com/mattn/go-gimei"

	"github.com/nansystem/mysql-test/common"
	"github.com/nansystem/mysql-test/generated"
	"github.com/nansystem/mysql-test/infra"
)

func FillPeople(count int64) ([]generated.YhPerson, error) {
	list := createPeopleList(count)

	bar := pb.StartNew(int(count))
	for idx := range common.IndexChunks(len(list), 2000) {
		q := sq.Insert("yh_people").Columns("pref_id", "name", "age")
		for _, item := range list[idx.From:idx.To] {
			q = q.Values(item.PrefID, item.Name, item.Age)
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

func createPeopleList(count int64) []generated.YhPerson {
	list := make([]generated.YhPerson, count)
	for i := int64(0); i < count; i++ {
		name := gimei.NewName()
		list[i] = newPerson(
			uint(common.RandNum(1, 47)),
			fmt.Sprintf("%s_%d", name.Katakana(), i+1),
			uint(common.RandNum(1, 100)),
		)
	}
	return list
}

func newPerson(prefID uint, name string, age uint) generated.YhPerson {
	return generated.YhPerson{
		PrefID: prefID,
		Name:   name,
		Age:    age,
	}
}
