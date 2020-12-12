package bloom

import (
	corebloom "github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"sync"
)

const (
	bloomUserKey = "bloom:user#key"
)

type BloomerUserAbility interface {
	AddBloomRules(rules ...string) error
	Exists(data string) (bool, error)
}
type BloomerUser struct {
	bloomFilter *corebloom.BloomFilter
}

func (f *BloomerUser) AddBloomRules(rules ...string) error {
	for _, value := range rules {
		key := []byte(value)
		exists, err := f.Exists(value)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		err = f.add(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *BloomerUser) Exists(data string) (bool, error) {
	exists, err := f.bloomFilter.Exists([]byte(data))
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	return true, nil
}

func (f *BloomerUser) add(data []byte) error {
	err := f.bloomFilter.Add(data)
	if err != nil {
		return err
	}
	return nil
}

func MustNewBloom(store *redis.Redis) *BloomerUser {
	var one sync.Once
	var bloomUser *BloomerUser
	one.Do(func() {
		bloomUser = oneNewBloom(store)
	})
	return bloomUser
}

func oneNewBloom(store *redis.Redis) *BloomerUser {
	result := new(BloomerUser)
	filter := corebloom.New(store, bloomUserKey, uint(1024))
	result.bloomFilter = filter
	return result
}
