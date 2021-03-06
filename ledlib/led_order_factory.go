package ledlib

import (
	"errors"
	"fmt"
	"3d_led_cube_go/ledlib/util"
	"reflect"
)

func CreateObject(order map[string]interface{}, ledCanvas LedCanvas) (interface{}, float64, error) {

	value, err := GetJSONValue(order, "id")
	ilifetime, _ := GetJSONValueOrDefault(order, "lifetime", float64(16.0))
	if err != nil {
		return nil, 0, err
	}
	if id, ok := value.(string); ok {

		if lifetime, ok := ilifetime.(float64); ok {
			switch id {
			/*
				Objects
			*/
			case "object-blank":
				return NewObjectFill(util.NewColorFromRGB(0, 0, 0)), lifetime, nil
			case "object-rocket":
				return NewObjectRocket(), lifetime, nil
			case "object-stickman":
				return NewObjectStickman(), lifetime, nil
			case "object-ghost":
				return NewObjectGhost(), lifetime, nil
			case "object-yacht":
				return NewObjectYacht(), lifetime, nil
			case "object-heart":
				return NewObjectHeart(), lifetime, nil
			case "object-painting":
				return NewObjectPainting(), lifetime, nil
			case "object-fireworks":
				return NewObjectFireworks(), lifetime, nil
			case "object-saboten":
				return NewObjectSaboten(), lifetime, nil
			case "object-snowman":
				return NewObjectSnowman(), lifetime, nil
			case "object-star":
				return NewObjectStar(), lifetime, nil
			case "object-tree":
				return NewObjectTree(), lifetime, nil
			case "object-tulip":
				return NewObjectTulip(), lifetime, nil
			case "object-note":
				return NewObjectNote(), lifetime, nil
			case "object-realsense":
				return NewObjectRealsense(), lifetime, nil

			/*
				Filters
			*/
			case "filter-bk-cloud":
				return NewFilterBkCloudsLow(ledCanvas), 0, nil
			case "filter-rolling":
				return NewFilterRolling(ledCanvas), 0, nil
			case "filter-skewed":
				return NewFilterSkewed(ledCanvas), 0, nil
			case "filter-jump":
				return NewFilterJump(ledCanvas), 0, nil
			case "filter-bk-snows":
				return NewFilterBkSnowsLow(ledCanvas), 0, nil
			case "filter-bk-stars":
				return NewFilterBkStars(ledCanvas), 0, nil
			case "filter-bk-rains":
				return NewFilterBkRains(ledCanvas), 0, nil
			case "filter-bk-mountain":
				return NewFilterBkMountain(ledCanvas), 0, nil
			case "filter-bk-grass":
				return NewFilterBkGrass(ledCanvas), 0, nil
			case "filter-bk-wave":
				return NewFilterBkWave(ledCanvas), 0, nil
			case "filter-bk-fireworks":
				return NewFilterBkFireworks(ledCanvas), 0, nil
			case "filter-explosion":
				return NewFilterExplosion(ledCanvas, 2), 0, nil
			case "filter-3d-explosion":
				return NewFilterExplosion(ledCanvas, 3), 0, nil
			case "filter-wakame":
				return NewFilterWakame(ledCanvas), 0, nil
			case "filter-exile":
				return NewFilterExile(ledCanvas), 0, nil
			case "filter-rainbow":
				return NewFilterRainbow(ledCanvas), 0, nil
			case "filter-swaying":
				return NewFilterSwaying(ledCanvas), 0, nil
			case "filter-wave":
				return NewFilterWave(ledCanvas), 0, nil
			case "filter-elastic":
				return NewFilterElastic(ledCanvas), 0, nil
			case "filter-zanzo":
				return NewFilterZanzo(ledCanvas), 0, nil

			default:
				return nil, 0, errors.New("Unnown Object Id")
			}
		} else {
			return nil, 0, fmt.Errorf("Unknown Error: %s", reflect.TypeOf(ilifetime))
		}
	} else {
		return nil, 0, fmt.Errorf("Unknown Error: %s", reflect.TypeOf(value))
	}
}
