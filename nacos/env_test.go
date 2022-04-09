package nacos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEnvFunc test env func
func TestEnvFunc(t *testing.T) {
	assert.Equal(t, int64(8848), NacosPort())
	assert.Equal(t, "127.0.0.1", NacosAddr())
	assert.Equal(t, "", NacosNameSpaceId())
}
