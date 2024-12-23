package ebicanvas

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed filters/aa_sampling_soft.kage
var _aaSamplingSoft []byte

//go:embed filters/aa_sampling_sharp.kage
var _aaSamplingSharp []byte

//go:embed filters/nearest.kage
var _nearest []byte

//go:embed filters/hermite.kage
var _hermite []byte

//go:embed filters/bicubic.kage
var _bicubic []byte

//go:embed filters/bilinear.kage
var _bilinear []byte

//go:embed filters/src_hermite.kage
var _srcHermite []byte

//go:embed filters/src_bicubic.kage
var _srcBicubic []byte

//go:embed filters/src_bilinear.kage
var _srcBilinear []byte

var (
	shaderSources     [scalingFilterEndSentinel][]byte
	shaders           [scalingFilterEndSentinel]*ebiten.Shader
	shaderOpts        ebiten.DrawTrianglesShaderOptions
	shaderVertices    []ebiten.Vertex
	shaderVertIndices []uint16
)

func init() {
	shaderSources[Nearest] = _nearest
	shaderSources[AASamplingSoft] = _aaSamplingSoft
	shaderSources[AASamplingSharp] = _aaSamplingSharp
	shaderSources[Hermite] = _hermite
	shaderSources[Bicubic] = _bicubic
	shaderSources[Bilinear] = _bilinear
	shaderSources[SrcHermite] = _srcHermite
	shaderSources[SrcBicubic] = _srcBicubic
	shaderSources[SrcBilinear] = _srcBilinear
}

func compileShader(filter ScalingFilter) {
	var err error
	shaders[filter], err = ebiten.NewShader(shaderSources[filter])
	if err != nil {
		panic("Failed to compile shader filter: " + err.Error())
	}
	if shaderOpts.Uniforms == nil {
		initShaderProperties()
	}
}

func initShaderProperties() {
	shaderVertices = make([]ebiten.Vertex, 4)
	shaderVertIndices = []uint16{0, 1, 3, 3, 1, 2}
	shaderOpts.Uniforms = make(map[string]interface{}, 2)
	for i := range 4 { // doesn't matter unless I start doing color scaling
		shaderVertices[i].ColorR = 1.0
		shaderVertices[i].ColorG = 1.0
		shaderVertices[i].ColorB = 1.0
		shaderVertices[i].ColorA = 1.0
	}
}
