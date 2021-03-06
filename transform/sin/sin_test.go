package sin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"osampler/audio"
	"osampler/test"
	"osampler/transform/aiff"
	"osampler/transform/gain"
)

func TestBasics(t *testing.T) {
	ass := require.New(t)
	bufferSize := 512
	buffer := audio.NewBuffer(bufferSize)
	frequency := 440
	phase := 0

	sinTransform := New(audio.NewCDContext(), buffer, float64(frequency), float64(phase))
	ass.NotNil(sinTransform)

	gainFactor := 10000.0
	gainTransform := gain.New(buffer, gainFactor)

	outfile, err := test.TempFile("test-*.aif")
	ass.Nil(err)
	ass.NotNil(outfile)
	filename := outfile.Name()
	out := aiff.NewAiffOutput(audio.NewCDContext(), buffer, outfile)

	iterations := 100

	for i := 0; i < iterations; i++ {
		sinTransform.CalculateBuffer()
		gainTransform.CalculateBuffer()
		out.CalculateBuffer()
		ass.Nil(err)
	}
	err = out.Close()
	ass.Nil(err)

	err = outfile.Close()
	ass.Nil(err)

	t.Logf("Wrote file://%v", filename)
}
