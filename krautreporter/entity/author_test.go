package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthor_AddImage(t *testing.T) {
	a := Author{}
	assert.Nil(t, a.Images)
	assert.Len(t, a.Images, 0)

	i1 := Image{Width: 130, Src: "/foobar.jpg"}
	a.AddImage(i1)
	assert.Len(t, a.Images, 1)
	assert.Equal(t, i1, a.Images[0])

	i2 := Image{Width: 260, Src: "/foobaz.jpg"}
	a.AddImage(i2)
	assert.Len(t, a.Images, 2)
	assert.Equal(t, i2, a.Images[1])

	i3 := Image{Width: 130, Src: "/baz.jpg"}
	a.AddImage(i3)
	assert.Len(t, a.Images, 2)
	assert.Equal(t, i3, a.Images[0])
}
