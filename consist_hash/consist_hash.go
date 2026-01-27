package consist_hash

import (
	"errors"
	"hash/crc32"
	"sort"
)

type ServerName string

type ServerNo uint32

func (no ServerNo) MoreOrEqual(val uint32) bool {
	return uint32(no) >= val
}

func NewServerRing() *ServersRing {
	return &ServersRing{
		names: map[ServerNo]ServerName{},
	}
}

type ServersRing struct {
	servers []ServerNo
	names   map[ServerNo]ServerName
}

func (s *ServersRing) Add(name ServerName) error {
	if len(name) == 0 {
		return errors.New("server name must not be empty")
	}

	no := ServerNo(hash(string(name)))
	if _, ok := s.names[no]; ok {
		return errors.New("server already exists")
	}

	s.servers = append(s.servers, no)
	s.names[no] = name
	sort.Slice(s.servers, func(i, j int) bool { return s.servers[i] < s.servers[j] })

	return nil
}

func (s *ServersRing) Get(key string) (ServerName, error) {
	hash := hash(key)
	no, err := getServer(s.servers, hash)
	if err != nil {
		return "", err
	}
	return s.names[no], nil
}

func hash(v string) uint32 {
	return crc32.ChecksumIEEE([]byte(v))
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
