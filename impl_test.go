package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	expiry = 5 * time.Millisecond
)

type localCacheSuite struct {
	suite.Suite
	cache *cache
}

func (s *localCacheSuite) SetupSuite() {}

func (s *localCacheSuite) TearDownSuite() {}

func (s *localCacheSuite) SetupTest() {
	s.cache = New(expiry).(*cache)
}

func (s *localCacheSuite) TearDownTest() {}

func TestLocalcacheSuite(t *testing.T) {
	suite.Run(t, new(localCacheSuite))
}

func (s *localCacheSuite) TestGet() {
	tests := []struct {
		Desc      string
		Key       string
		SetupTest func()
		ExpError  error
		ExpResult interface{}
	}{
		{
			Desc:      "normal Get",
			Key:       "key",
			ExpResult: "data",
			SetupTest: func() {
				s.cache.Set("key", "data")
			},
		},
		{
			Desc: "data not found",
			Key:  "key2",
			SetupTest: func() {
			},
			ExpError: ErrCacheMiss,
		},
		{
			Desc: "cache expiry",
			Key:  "key3",
			SetupTest: func() {
				s.cache.Set("key3", "data3")
				time.Sleep(expiry + (5 * time.Millisecond))
			},
			ExpError: ErrCacheMiss,
		},
	}

	for _, t := range tests {
		t.SetupTest()

		res, err := s.cache.Get(t.Key)
		if err != nil {
			s.Require().Equal(t.ExpError, err, t.Desc)
			continue
		}
		s.Require().Equal(t.ExpResult, res, t.Desc)
	}
}

func (s *localCacheSuite) TestSet() {
	tests := []struct {
		Desc      string
		Key       string
		SetupTest func()
		ExpResult interface{}
	}{
		{
			Desc: "normal set",
			Key:  "key1",
			SetupTest: func() {
			},
			ExpResult: "data",
		},
		{
			Desc: "update existed cache",
			Key:  "key2",
			SetupTest: func() {
				s.cache.Set("key2", "data")
			},
			ExpResult: "data2",
		},
	}

	for _, t := range tests {
		t.SetupTest()
		err := s.cache.Set(t.Key, t.ExpResult)
		res, _ := s.cache.data[t.Key]

		s.Require().Equal(nil, err, t.Desc)
		s.Require().Equal(t.ExpResult, res, t.Desc)
	}
}
