package service

import (
	"fmt"
	"github.com/xiaozefeng/apiserver/model"
	"github.com/xiaozefeng/apiserver/util"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  sync.Mutex{},
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finishChan := make(chan bool, 1)
	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:       u.Id,
				Username: u.Username,
				SayHello: fmt.Sprintf("Hello %s", shortId),
				Password: u.Password,
				CreateAt: u.CreateAt.Format("2006-01-02 15:04:05"),
				UpdateAt: u.UpdateAt.Format("2006-01-02 15:04:05"),
			}

		}(u)
	}

	go func() {
		wg.Wait()
		close(finishChan)
	}()

	select {
	case <-finishChan:
	case err := <-errChan:
		return nil, count, err

	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}
	return infos, count, nil
}
