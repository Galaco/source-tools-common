package entitylist

import (
	"math"
	"log"
	"errors"
	"github.com/go-gl/mathgl/mgl32"
	"fmt"
	"strconv"
	"strings"
	"github.com/galaco/source-tools-common/vmath/vector"
)

// "Abstract" Engine entity.
type Entity struct {
	Origin mgl32.Vec3
	FirstBrush int
	NumBrushes int
	EPairs *EPair
}

// ValueForKey
// Get string value for a key
// Returns empty string "" if not found
func (ent *Entity) ValueForKey(key string) string {
	return ent.ValueForKeyWithDefault(key, "")
}

// ValueForKeyWithDefault
// Get string value for key
// Returns a specified default if not found
func (ent *Entity) ValueForKeyWithDefault(key string, defaultValue string) string {
	var ep *EPair

	for ep = ent.EPairs; ep != nil; ep = ep.Next {
		if ep.Key == key {
			return ep.Value
		}

	}
	return defaultValue
}

// VectorForKey
// Returns a vector for a specified key's value
// Default (0,0,0) returned
func (ent *Entity) VectorForKey(key string) mgl32.Vec3 {
	k := ent.ValueForKey(key)
	var v1, v2, v3 = float32(0), float32(0), float32(0)
	fmt.Sscanf(k, "%f %f %f", &v1, &v2, &v3)

	return mgl32.Vec3{v1, v2, v3}
}

// IntForKey
// Returns int value for key.
// 0 returned if key not found
func (ent *Entity) IntForKey(key string) int {
	k := ent.ValueForKey(key)
	result,_ := strconv.ParseInt(k, 10, 32)
	return int(result)
}

// FloatForKey
// Returns float value for key.
// 0 returned if key not found
func (ent *Entity) FloatForKey(key string) float32 {
	k := ent.ValueForKey(key)
	result,_ := strconv.ParseFloat(k, 32)
	return float32(result)
}

// FloatForKeyWithDefault
// Returns float value for key.
// Specified default value returned if key not found
func (ent *Entity) FloatForKeyWithDefault(key string, defaultValue float32) float32 {
	for ep := ent.EPairs; ep != nil; ep = ep.Next {
		if strings.EqualFold(ep.Key, key) {
			ret, err := strconv.ParseFloat(ep.Value, 32)
			if err != nil {
				return defaultValue
			}
			return float32(ret)
		}
	}

	return defaultValue
}

// LightForKey
// Returns a Light-specific Vec3 for a key
func (ent *Entity) LightForKey(key string, useHDR bool, lightScale float32) (mgl32.Vec3,error) {
	light := ent.ValueForKey(key)

	return ent.LightForString(light, useHDR, lightScale)
}

// LightForString
// Returns a light-specific Vec3 for a string representation of a light
// This takes into account configuration of parameters:
// HDR, lightScale
func (ent *Entity) LightForString(light string, useHDR bool, lightScale float32) (mgl32.Vec3,error) {
	var r, g, b, scaler = 0.0, 0.0, 0.0, 0.0
	var argCnt int
	var intensity = mgl32.Vec3{0,0,0}

	// scanf into doubles, then assign, so it is vec_t size independent
	var r_hdr, g_hdr, b_hdr, scaler_hdr float64
	argCnt,_ = fmt.Sscanf(light, "%f %f %f %f %f %f %f %f",
		&r, &g, &b, &scaler, &r_hdr,&g_hdr,&b_hdr,&scaler_hdr)

	if argCnt == 8 { 											// 2 4-tuples
		if useHDR == true {
			r = r_hdr
			g = g_hdr
			b = b_hdr
			scaler = scaler_hdr
		}
		argCnt = 4
	}

	// make sure light is legal
	if r < 0.0 || g < 0.0 || b < 0.0 || scaler < 0.0 {
		intensity = mgl32.Vec3{0.0,0.0,0.0}
		return intensity, errors.New("invalid colour for light")
	}

	intensity[0] = float32(math.Pow( r / 255.0, 2.2 ) * 255)				// convert to linear

	switch argCnt {
	case 1:
		// The R,G,B values are all equal.
		intensity[2] = intensity[0]
		intensity[1] = intensity[2]
		break
	case 3:
	case 4:
		// Save the other two G,B values.
		intensity[1] = float32(math.Pow( float64(g / 255.0), 2.2 ) * 255)
		intensity[2] = float32(math.Pow( float64(b / 255.0), 2.2 ) * 255)

		// Did we also get an "intensity" scaler value too?
		if argCnt == 4 {
			// Scale the normalized 0-255 R,G,B values by the intensity scaler
			vector.Scale(&intensity, float32(scaler / 255.0), &intensity )
		}
		break
	default:
		log.Printf("unknown light specifier type - %s\n", light)
		return intensity, errors.New("unknown light specifier type")
	}
	// scale up source lights by scaling factor
	vector.Scale(&intensity, lightScale, &intensity)
	return intensity, nil
}


// EPair
// Linked List of entity key-value pairs
type EPair struct {
	Next *EPair
	Key string
	Value string
}
