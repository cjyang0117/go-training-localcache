package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	EXPIRY = 1
)

type localCacheSuite struct {
	suite.Suite
	cache *cache
}

func (s *localCacheSuite) SetupSuite() {}

func (s *localCacheSuite) TearDownSuite() {}

func (s *localCacheSuite) SetupTest() {
	s.cache = New(EXPIRY).(*cache)
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
			SetupTest: func() {
			},
			ExpError: ErrCacheMiss,
		},
		{
			Desc: "cache expiry",
			SetupTest: func() {
				time.Sleep(2 * time.Second)
			},
			ExpError: ErrCacheMiss,
		},
	}

	for _, t := range tests {
		t.SetupTest()

		res, err := s.cache.Get(t.Key)
		if err != nil {
			s.Require().Equal(t.ExpError, err, t.Desc)
			return
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
		s.cache.Set(t.Key, t.ExpResult)

		res, _ := s.cache.data[t.Key]
		s.Require().Equal(t.ExpResult, res, t.Desc)
	}
}
