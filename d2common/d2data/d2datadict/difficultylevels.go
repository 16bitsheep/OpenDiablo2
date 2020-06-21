package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

var DifficultyLevels map[string]*DifficultyLevelRecord

type DifficultyLevelRecord struct {
	// Difficulty name. it is hardcoded and you cannot add new ones unless you do
	// some Code Edits
	Name string // Name

	// Resistance penalty in the current difficulty.
	ResistancePenalty int // ResistPenalty

	// The percentage of experience you lose when you die on this difficulty.
	DeathExperiencePenalty int // DeathExpPenalty

	// Not Used. Pre 1.07 it was the percentage of low quality, normal, superior and
	// exceptional items dropped on this difficulty.
	DropChanceLow         int // UberCodeOddsNormal
	DropChanceNormal      int // UberCodeOddsNormal
	DropChanceSuperior    int // UberCodeOddsNormal
	DropChanceExceptional int // UberCodeOddsNormal
	// Gravestench - I'm splitting this field because I feel it's incoherent
	// to keep all of those drop chances together, even if it is that way in the
	// txt file...

	// Not used. Pre 1.07 it was the percentage of magic, rare, set and unique
	// exceptional items dropped on this difficuly.
	DropChanceMagic  int // UberCodeOddsGood
	DropChanceRare   int // UberCodeOddsGood
	DropChanceSet    int // UberCodeOddsGood
	DropChanceUnique int // UberCodeOddsGood
	// Gravestench - same as my above comment

	// Not used and didn't exist pre 1.07.
	// UltraCodeOddsNormal

	// Additional skill points added to monster skills specified in MonStats.txt
	// for this difficulty. It has nothing to do with the missile damage bonus.
	MonsterSkillBonus int // MonsterSkillBonus

	// This value is a divisor, and so never set it to 0. It applies to the monster
	// freezing length and cold length duration.
	MonsterColdDivisor   int // MonsterColdDivisor
	MonsterFreezeDivisor int // MonsterFreezeDivisor

	// These values are divisor and they're used respectively for AI altering states
	AiCurseDivisor   int // AiCurseDivisor
	LifeStealDivisor int // LifeStealDivisor
	ManaStealDivisor int // ManaStealDivisor

	// -----------------------------------------------------------------------
	// Gravestench: The rest of these are listed on PK page, but not present in
	// my copy of the txt file (patch_d2/data/global/excel/difficultylevels.txt)
	// so I am going to leave these comments

	// Effective percentage of damage and attack rating added to Extra Strong
	// Unique/Minion and Champion monsters. This field is actually a coefficient,
	// as the total bonus output is BonusFromMonUMod/100*ThisField
	// UniqueDamageBonus
	// ChampionDamageBonus

	//  This is a percentage of how much damage your mercenaries do to an Act boss.
	// HireableBossDamagePercent

	// Monster Corpse Explosion damage percent limit. Since the monsters HP grows
	// proportionally to the number of players in the game, you can set a cap via
	// this field.
	// MonsterCEDamagePercent

	// Maximum cap of the monster hit points percentage that can be damaged through
	// Static Field. Setting these columns to 0 will make Static Field work the same
	// way it did in Classic Diablo II.
	// StaticFieldMin

	// Parameters for gambling. They states the odds to find Rares, Sets, Uniques,
	// Exceptionals and Elite items when gambling. See Appendix A
	// GambleRare
	// GambleSet
	// GambleUnique
	// GambleUber
	// GambleUltra
	// -----------------------------------------------------------------------

}

func LoadDifficultyLevels(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRows := len(dict.Data)

	DifficultyLevels = make(map[string]*DifficultyLevelRecord, numRows)

	for idx, _ := range dict.Data {
		record := &DifficultyLevelRecord{
			Name:                   dict.GetString("Name", idx),
			ResistancePenalty:      dict.GetNumber("ResistPenalty", idx),
			DeathExperiencePenalty: dict.GetNumber("DeathExpPenalty", idx),
			DropChanceLow:          dict.GetNumber("UberCodeOddsNormal", idx),
			DropChanceNormal:       dict.GetNumber("UberCodeOddsNormal", idx),
			DropChanceSuperior:     dict.GetNumber("UberCodeOddsNormal", idx),
			DropChanceExceptional:  dict.GetNumber("UberCodeOddsNormal", idx),
			DropChanceMagic:        dict.GetNumber("UberCodeOddsGood", idx),
			DropChanceRare:         dict.GetNumber("UberCodeOddsGood", idx),
			DropChanceSet:          dict.GetNumber("UberCodeOddsGood", idx),
			DropChanceUnique:       dict.GetNumber("UberCodeOddsGood", idx),
			MonsterSkillBonus:      dict.GetNumber("MonsterSkillBonus", idx),
			MonsterColdDivisor:     dict.GetNumber("MonsterColdDivisor", idx),
			MonsterFreezeDivisor:   dict.GetNumber("MonsterFreezeDivisor", idx),
			AiCurseDivisor:         dict.GetNumber("AiCurseDivisor", idx),
			LifeStealDivisor:       dict.GetNumber("LifeStealDivisor", idx),
			ManaStealDivisor:       dict.GetNumber("ManaStealDivisor", idx),
		}
		DifficultyLevels[record.Name] = record
	}

	log.Printf("Loaded %d DifficultyLevel records", len(DifficultyLevels))

}