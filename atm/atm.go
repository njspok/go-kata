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

func (a *ATM) Put(coin Coin) error {
	a.cache[coin]++
	return nil
}

func (a *ATM) Total() Sum {
	return a.cache.Total()
}

func (a *ATM) Cache() Cache {
	if a.cache.Empty() {
		return nil
	}
	return a.cache
}

func (a *ATM) Give(sum Sum) (Cache, error) {
	if !a.enoughCache(sum) {
		return nil, ErrNotEnoughCoins
	}

	originalSum := sum
	result := make(Cache)
	iterateDescending(a.cache, func(coin Coin, count Count) {
		for {
			if sum < coin.ToSum() {
				return
			}

			if a.cache[coin] <= result[coin] {
				return
			}

			sum -= coin.ToSum()
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

func (a *ATM) enoughCache(sum Sum) bool {
	return a.Total() >= sum
}

func iterateDescending(cache Cache, f func(coin Coin, count Count)) {
	var keys []Coin
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
