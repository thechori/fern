package generatorexec

import (
	json "encoding/json"
	fmt "fmt"
)

type GeneratorUpdate struct {
	Type             string
	Init             *InitUpdate
	InitV2           *InitUpdateV2
	Log              *LogUpdate
	Publishing       *PackageCoordinate
	Published        *PackageCoordinate
	ExitStatusUpdate *ExitStatusUpdate
}

func (x *GeneratorUpdate) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type             string             `json:"_type"`
		Publishing       *PackageCoordinate `json:"publishing"`
		Published        *PackageCoordinate `json:"published"`
		ExitStatusUpdate *ExitStatusUpdate  `json:"exitStatusUpdate"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	x.Type = unmarshaler.Type
	switch unmarshaler.Type {
	case "init":
		value := new(InitUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Init = value
	case "initV2":
		value := new(InitUpdateV2)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.InitV2 = value
	case "log":
		value := new(LogUpdate)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		x.Log = value
	case "publishing":
		x.Publishing = unmarshaler.Publishing
	case "published":
		x.Published = unmarshaler.Published
	case "exitStatusUpdate":
		x.ExitStatusUpdate = unmarshaler.ExitStatusUpdate
	}
	return nil
}

type GeneratorUpdateVisitor interface {
	VisitInit(*InitUpdate) error
	VisitInitV2(*InitUpdateV2) error
	VisitLog(*LogUpdate) error
	VisitPublishing(*PackageCoordinate) error
	VisitPublished(*PackageCoordinate) error
	VisitExitStatusUpdate(*ExitStatusUpdate) error
}

func (x *GeneratorUpdate) Accept(v GeneratorUpdateVisitor) error {
	switch x.Type {
	default:
		return fmt.Errorf("invalid type %s in %T", x.Type, x)
	case "init":
		return v.VisitInit(x.Init)
	case "initV2":
		return v.VisitInitV2(x.InitV2)
	case "log":
		return v.VisitLog(x.Log)
	case "publishing":
		return v.VisitPublishing(x.Publishing)
	case "published":
		return v.VisitPublished(x.Published)
	case "exitStatusUpdate":
		return v.VisitExitStatusUpdate(x.ExitStatusUpdate)
	}
}
