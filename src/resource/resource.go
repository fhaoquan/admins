/**********************************************************
 * Author : Michael
 * Email : dolotech@163.com
 * Last modified : 2016-03-28 20:44
 * Filename :resource.go
 * Description :变更玩家的经济资源,并推变更消息送给玩家
 * *******************************************************/

package resource

import (
	"basic/ssdb/gossdb"
	"basic/utils"
	"csv"
	"data"

	"github.com/golang/glog"
)

const (
	COIN     uint32 = 1
	EXCHANGE uint32 = 2
	TICKET   uint32 = 3
	DIAMOND  uint32 = 4

	VIP0 uint32 = 21
	VIP1 uint32 = 22
	VIP2 uint32 = 23
	VIP3 uint32 = 24

	SOUND uint32 = 15

	EXP  uint32 = 100
	WIN  uint32 = 101
	LOST uint32 = 102
	PING uint32 = 103

	RMB uint32 = 1
	DIA uint32 = 2
)

// 更改单个资源
func ChangeRes(userid string, id uint32, count int32) error {
	m := make(map[uint32]int32)
	m[id] = count
	return ChangeMulti(userid, m)
}

// 更改多个资源
func ChangeMulti(userid string, res map[uint32]int32) error {
	userdata := &data.User{Userid: userid}
	if err := userdata.Get(); err != nil {
		return err
	}
	glog.Infoln(res)
	for id, count := range res {
		var current int32
		var err error
		if id == COIN {
			current = int32(userdata.Coin) + count
			if current < 0 {
				current = 0
			}
			err = gossdb.C().Hset(data.KEY_USER+userid, "Coin", current)
			if err == nil {
			}
		} else if id == EXCHANGE {
			current = int32(userdata.Exchange) + count
			if current < 0 {
				current = 0
			}
			err = gossdb.C().Hset(data.KEY_USER+userid, "Exchange", current)
			if err == nil {

			}
		} else if id == TICKET {
			current = int32(userdata.Ticket) + count
			if current < 0 {
				current = 0
			}
			err = gossdb.C().Hset(data.KEY_USER+userid, "Ticket", current)
			if err == nil {

			}

		} else if id == DIAMOND {
			current = int32(userdata.Diamond) + count
			if current < 0 {
				current = 0
			}
			err = gossdb.C().Hset(data.KEY_USER+userid, "Diamond", int64(current))
			if err == nil {

			}

		} else if id == EXP {
			current = int32(userdata.Exp) + count
			if current < 0 {
				current = 0
			}
			err = gossdb.C().Hset(data.KEY_USER+userid, "Exp", current)
			if err == nil {

			}

		} else if id >= VIP0 && id <= VIP3 {
			vipdata := csv.GetVip(id)
			glog.Infoln(vipdata, userid, id)
			if vipdata != nil {
				viplev := vipdata.Viptype
				vip := userdata.Vip
				if viplev < vip {
					viplev = vip
				}
				expire := userdata.VipExpire
				now := uint32(utils.Timestamp())
				add := uint32(vipdata.Expire * 86400)
				if expire < now {
					expire = now + add
				} else {
					expire += add
				}
				kvs := map[string]interface{}{
					"Vip":       viplev,
					"VipExpire": expire,
				}
				err := gossdb.C().MultiHset(data.KEY_USER+userid, kvs)

				if err != nil {
					glog.Errorln(err)
				}

			}
		}

		if err != nil {
			glog.Errorln(err)
		}
	}
	return nil

}
