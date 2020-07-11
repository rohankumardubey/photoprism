package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAsciiID(t *testing.T) {
	assert.False(t, IsAsciiID("lt9k3pw1wowuy3c2"))
	assert.False(t, IsAsciiID("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.False(t, IsAsciiID("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.False(t, IsAsciiID("55785BAC-9A4B-4747-B090-EE123FFEE437"))
	assert.False(t, IsAsciiID("550e8400-e29b-11d4-a716-446655440000"))
	assert.False(t, IsAsciiID("IMG_0599.JPG"))
	assert.True(t, IsAsciiID("DSC10599"))
	assert.True(t, IsAsciiID("IQVG4929"))
	assert.False(t, IsAsciiID("DSC_0599"))
	assert.False(t, IsAsciiID("iqVG4929"))
	assert.False(t, IsAsciiID("20091117_203458_ERROR000"))
	assert.False(t, IsAsciiID("20091117_203458_12345678"))
	assert.True(t, IsAsciiID("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
	assert.True(t, IsAsciiID("123"))
	assert.False(t, IsAsciiID("_"))
	assert.False(t, IsAsciiID(""))
	assert.False(t, IsAsciiID("20191117-153400-Central-Park-New-York-2019-3qy.mov"))
	assert.False(t, IsAsciiID("e98eb86480a72bd585d228a709f0622f90e86cbc.jpg"))
	assert.False(t, IsAsciiID("IMG_8115.jpg"))
	assert.False(t, IsAsciiID("01 Introduction Businessmodel.pdf"))
	assert.False(t, IsAsciiID("A regular file name with 121345678643 numbers"))
}

func TestIsGenerated(t *testing.T) {
	assert.True(t, IsGenerated("lt9k3pw1wowuy3c2"))
	assert.True(t, IsGenerated("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.True(t, IsGenerated("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.True(t, IsGenerated("55785BAC-9A4B-4747-B090-EE123FFEE437"))
	assert.True(t, IsGenerated("550e8400-e29b-11d4-a716-446655440000"))
	assert.True(t, IsGenerated("IMG_0599.JPG"))
	assert.True(t, IsGenerated("DSC10599"))
	assert.True(t, IsGenerated("IQVG4929"))
	assert.True(t, IsGenerated("49007520716_67ff0ce0ec_4k"))
	assert.True(t, IsGenerated("8263987746_d0a6055c58_o"))
	assert.True(t, IsGenerated("20091117_203458_ERROR000"))
	assert.True(t, IsGenerated("20091117_203458_12345678"))
	assert.True(t, IsGenerated("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
	assert.True(t, IsGenerated("123"))
	assert.False(t, IsGenerated("_"))
	assert.False(t, IsGenerated(""))
	assert.False(t, IsGenerated("20191117-153400-Central-Park-New-York-2019-3qy.mov"))
	assert.True(t, IsGenerated("e98eb86480a72bd585d228a709f0622f90e86cbc.jpg"))
	assert.True(t, IsGenerated("IMG_8115.jpg"))
	assert.False(t, IsGenerated("01 Introduction Businessmodel.pdf"))
	assert.False(t, IsGenerated("A regular file name with 121345678643 numbers"))
}

func TestIsInt(t *testing.T) {
	assert.True(t, IsInt("123"))
	assert.False(t, IsInt(""))
}

func TestIsUniqueName(t *testing.T) {
	assert.False(t, IsUniqueName("lt9k3pw1wowuy3c2"))
	assert.True(t, IsUniqueName("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.True(t, IsUniqueName("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.True(t, IsUniqueName("55785BAC-9A4B-4747-B090-EE123FFEE437"))
	assert.True(t, IsUniqueName("550e8400-e29b-11d4-a716-446655440000"))
	assert.False(t, IsUniqueName("IMG_0599.JPG"))
	assert.False(t, IsUniqueName("DSC10599"))
	assert.False(t, IsUniqueName("IQVG4929"))
	assert.True(t, IsUniqueName("7045C6DD-448A-405D-B2A4-6F4DA0D44F2F"))
	assert.False(t, IsUniqueName("7045C6DD-448A-405D-B2A4-6F4DA0D44F2F-FOO"))
	assert.True(t, IsUniqueName("7045C6DD-448A-405D-B2A4-6F4DA0D44F2F_L0_001"))
	assert.True(t, IsUniqueName("10298590_856989717698278_4041733760069394245_o"))
	assert.True(t, IsUniqueName("49007520716_67ff0ce0ec_4k"))
	assert.True(t, IsUniqueName("8263987746_d0a6055c58_o"))
	assert.True(t, IsUniqueName("20091117_203458_ERROR000"))
	assert.True(t, IsUniqueName("20091117_203458_12345678"))
	assert.True(t, IsUniqueName("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
	assert.False(t, IsUniqueName("123"))
	assert.False(t, IsUniqueName("_"))
	assert.False(t, IsUniqueName(""))
	assert.False(t, IsUniqueName("20191117-153400-Central-Park-New-York-2019-3qy.mov"))
	assert.False(t, IsUniqueName("e98eb86480a72bd585d228a709f0622f90e86cbc.jpg"))
	assert.False(t, IsUniqueName("IMG_8115.jpg"))
	assert.False(t, IsUniqueName("01 Introduction Businessmodel.pdf"))
	assert.False(t, IsUniqueName("A regular file name with 121345678643 numbers"))
}

func TestIsDscName(t *testing.T) {
	assert.False(t, IsDscName("lt9k3pw1wowuy3c2"))
	assert.False(t, IsDscName("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.False(t, IsDscName("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.False(t, IsDscName("55785BAC-9A4B-4747-B090-EE123FFEE437"))
	assert.False(t, IsDscName("550e8400-e29b-11d4-a716-446655440000"))
	assert.True(t, IsDscName("IMG_0599.JPG"))
	assert.True(t, IsDscName("IMG_20190721_095643.JPG"))
	assert.True(t, IsDscName("IMG_20190721_095643"))
	assert.True(t, IsDscName("DSC_20190721_095643"))
	assert.True(t, IsDscName("DSC10599"))
	assert.False(t, IsDscName("IQVG4929"))
	assert.False(t, IsDscName("49007520716_67ff0ce0ec_4k"))
	assert.False(t, IsDscName("8263987746_d0a6055c58_o"))
	assert.False(t, IsDscName("20091117_203458_ERROR000"))
	assert.False(t, IsDscName("20091117_203458_12345678"))
	assert.False(t, IsDscName("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
	assert.False(t, IsDscName("123"))
	assert.False(t, IsDscName("_"))
	assert.False(t, IsDscName(""))
	assert.False(t, IsDscName("20191117-153400-Central-Park-New-York-2019-3qy.mov"))
	assert.False(t, IsDscName("e98eb86480a72bd585d228a709f0622f90e86cbc.jpg"))
	assert.False(t, IsDscName("IMG_8115.jpg"))
	assert.False(t, IsDscName("01 Introduction Businessmodel.pdf"))
	assert.False(t, IsDscName("A regular file name with 121345678643 numbers"))
}
