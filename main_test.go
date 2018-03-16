package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func mockupTestReader() (io.Reader, error) {
	data, err := ioutil.ReadFile("./tests/test1.yml")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}

func TestNewRouters(t *testing.T) {
	rd, err := mockupTestReader()
	assert.NoError(t, err)

	rs, err := NewRouters(rd)
	assert.NoError(t, err)
	logrus.Infof("routers: %v", rs)
	assert.Equal(t, 2, len(rs.List))

	for i := range rs.List {
		assert.Equal(t, fmt.Sprintf("router%d", i), rs.List[i].Name)
		assert.Equal(t, fmt.Sprintf("file%d", i), rs.List[i].File)
	}
}
