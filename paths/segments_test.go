package paths

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestGetSegmentType(t *testing.T) {

	assert.Equal(t, segmentType(segmentTypeLiteral), getSegmentType("people"))
	assert.Equal(t, segmentType(segmentTypeDynamic), getSegmentType("{id}"))
	assert.Equal(t, segmentType(segmentTypeDynamicOptional), getSegmentType("[id]"))
	assert.Equal(t, segmentType(segmentTypeWildcard), getSegmentType("*"))
	assert.Equal(t, segmentType(segmentTypeCatchall), getSegmentType("..."))

}