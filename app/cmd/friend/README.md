# マストドン
[とあるクエリを2万倍速にした話 -データベースの気持ちになる- 前編](https://dwango.github.io/articles/mastodon-database-index-1/)  

## テーブル定義
``` sql
  create_table "notifications", force: :cascade do |t|
    t.bigint "activity_id", null: false
    t.string "activity_type", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.bigint "account_id", null: false
    t.bigint "from_account_id", null: false
    t.string "type"
    t.index ["account_id", "id", "type"], name: "index_notifications_on_account_id_and_id_and_type", order: { id: :desc }
    t.index ["activity_id", "activity_type"], name: "index_notifications_on_activity_id_and_activity_type"
    t.index ["from_account_id"], name: "index_notifications_on_from_account_id"
  end
https://github.com/tootsuite/mastodon/blob/main/db/schema.rb
```

## 通知を取得するSQL

マストドンにおけるnotificationsテーブルは被フォロー、返信など各ユーザに対しての通知が格納されているテーブルで、 friends.nicoではこのnotificationsテーブルが非常に巨大になる利用傾向となっています。

``` sql
SELECT  "notifications"."id", "notifications"."updated_at", "notifications"."activity_type", "notifications"."activity_id" 
FROM "notifications"
WHERE "notifications"."account_id" = $1 
AND "notifications"."activity_type" IN ('Mention', 'Status', 'Follow', 'Favourite')
ORDER BY "notifications"."id" DESC LIMIT $2;
```

notificationsテーブルは前述しましたとおり各ユーザに対しての各種通知が格納されているテーブルですから、59,699人(2018/01/10時点)のユーザがいるfriends.nicoではこれを用いて絞り込むだけで単純計算で6万分の1程度にまで範囲が絞りこまれ、取得して調べる必要のあるレコードが劇的に少なくなります。

また、account_idカラムに指定される条件はその内容が通知対象ユーザのIDであることからも想像できる通り必ず単一の値であって、 BETWEENなどを用いた範囲検索やIN演算子を用いた複数指定は行われないため、その点からしてもindexの最初に書くカラムとして適しています。

逆にactivity_typeについては今回は結局indexの対象にしませんでした。 なぜならactivity_typeに入り得る値はマストドンのコードより'Mention', 'Status', 'Follow', 'Favourite', 'FollowRequest'の5種類ですが、 文字列をindex対象に含めるコストに見合うほどの範囲の狭まり方をするようには思えなかったからです。

``` sql
CREATE INDEX CONCURRENTLY new_index ON notifications (account_id, id DESC);
```
