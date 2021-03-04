package atm

import (
	"errors"
	"sort"
)

var ErrNotEnoughCoins = errors.New("not enough coins")

func NewATM() *ATM {
	return &ATM{
		cache: make(Cache),
	}
}

type ATM struct {
	cache Cache
}

func (a *ATM) Put(coin uint) error {
	a.cache[coin]++
	return nil
}

func (a *ATM) Total() uint {
	return a.cache.Total()
}

func (a *ATM) Cache() Cache {
	return a.cache
}

func (a *ATM) Give(sum uint) (Cache, error) {
	if a.Total() < sum {
		return nil, ErrNotEnoughCoins
	}

	originalSum := sum
	result := make(Cache)
	iterateDescending(a.cache, func(coin, count uint) {
		for {
			if sum < coin {
				return
			}

			if a.cache[coin] <= result[coin] {
				return
			}

			sum -= coin
			result[coin]++
		}
	})

	a.withdrawCoins(result)

	if originalSum > result.Total() {
		return result, ErrNotEnoughCoins
	}

	return result, nil
}

func (a *ATM) Load(cache Cache) {
	a.cache = cache

}

func (a *ATM) withdrawCoins(c Cache) {
	for coin, count := range c {
		a.cache[coin] -= count
	}
}

func iterateDescending(cache Cache, f func(coin, count uint)) {
	var keys []uint
	for k, _ := range cache {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, coin := range keys {
		f(coin, cache[coin])
	}
}
