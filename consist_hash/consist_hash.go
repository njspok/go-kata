package consist_hash

import (
	"errors"
	"hash/crc32"
	"sort"
)

type ServerNo uint32

func (no ServerNo) MoreOrEqual(val uint32) bool {
	return uint32(no) >= val
}

func NewServerRing() *ServersRing {
	return &ServersRing{
		names: map[ServerNo]string{},
	}
}

type ServersRing struct {
	servers []ServerNo
	names   map[ServerNo]string
}

func (s *ServersRing) Add(serverName string) {
	no := ServerNo(crc32.ChecksumIEEE([]byte(serverName)))
	s.servers = append(s.servers, no)
	s.names[no] = serverName
	sort.Slice(s.servers, func(i, j int) bool { return s.servers[i] < s.servers[j] })
}

func (s *ServersRing) Get(key string) (string, error) {
	hash := crc32.ChecksumIEEE([]byte(key))
	no, err := getServer(s.servers, hash)
	if err != nil {
		return "", err
	}
	return s.names[no], nil
}

func getServer(servers []ServerNo, hash uint32) (ServerNo, error) {
	if len(servers) == 0 {
		return 0, errors.New("no servers")
	}

	var res ServerNo
	found := false

	for _, n := range servers {
		if n.MoreOrEqual(hash) {
			res = n
			found = true
			break
		}
	}

	if !found {
		res = servers[0]
	}

	return res, nil
}
