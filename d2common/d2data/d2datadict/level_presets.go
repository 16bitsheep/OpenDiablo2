package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"log"
	"strings"
)

// LevelPresetRecord is a representation of a row from lvlprest.txt
// these records define parameters for the preset level map generator
type LevelPresetRecord struct {
	Files        [6]string
	Name         string
	DefinitionID int
	LevelID      int
	SizeX        int
	SizeY        int
	Pops         int
	PopPad       int
	FileCount    int
	Dt1Mask      uint
	Populate     bool
	Logicals     bool
	Outdoors     bool
	Animate      bool
	KillEdge     bool
	FillBlanks   bool
	AutoMap      bool
	Scan         bool
	Beta         bool
	Expansion    bool
}

// CreateLevelPresetRecord parses a row from lvlprest.txt into a LevelPresetRecord
func createLevelPresetRecord(props []string) LevelPresetRecord {
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := LevelPresetRecord{
		Name:         props[inc()],
		DefinitionID: d2util.StringToInt(props[inc()]),
		LevelID:      d2util.StringToInt(props[inc()]),
		Populate:     d2util.StringToUint8(props[inc()]) == 1,
		Logicals:     d2util.StringToUint8(props[inc()]) == 1,
		Outdoors:     d2util.StringToUint8(props[inc()]) == 1,
		Animate:      d2util.StringToUint8(props[inc()]) == 1,
		KillEdge:     d2util.StringToUint8(props[inc()]) == 1,
		FillBlanks:   d2util.StringToUint8(props[inc()]) == 1,
		SizeX:        d2util.StringToInt(props[inc()]),
		SizeY:        d2util.StringToInt(props[inc()]),
		AutoMap:      d2util.StringToUint8(props[inc()]) == 1,
		Scan:         d2util.StringToUint8(props[inc()]) == 1,
		Pops:         d2util.StringToInt(props[inc()]),
		PopPad:       d2util.StringToInt(props[inc()]),
		FileCount:    d2util.StringToInt(props[inc()]),
		Files: [6]string{
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
			props[inc()],
		},
		Dt1Mask:   d2util.StringToUint(props[inc()]),
		Beta:      d2util.StringToUint8(props[inc()]) == 1,
		Expansion: d2util.StringToUint8(props[inc()]) == 1,
	}

	return result
}

// LevelPresets stores all of the LevelPresetRecords
var LevelPresets map[int]LevelPresetRecord //nolint:gochecknoglobals // Currently global by design

// LoadLevelPresets loads level presets from text file
func LoadLevelPresets(file []byte) {
	LevelPresets = make(map[int]LevelPresetRecord)
	data := strings.Split(string(file), "\r\n")[1:]

	for _, line := range data {
		if line == "" {
			continue
		}

		props := strings.Split(line, "\t")

		if props[1] == "" {
			continue // any line without a definition id is skipped (e.g. the "Expansion" line)
		}

		rec := createLevelPresetRecord(props)
		LevelPresets[rec.DefinitionID] = rec
	}

	log.Printf("Loaded %d level presets", len(LevelPresets))
}

// LevelPreset looks up a LevelPresetRecord by ID
func LevelPreset(id int) LevelPresetRecord {
	for i := 0; i < len(LevelPresets); i++ {
		if LevelPresets[i].DefinitionID == id {
			return LevelPresets[i]
		}
	}
	panic("Unknown level preset")
}
