// Code generated by ent, DO NOT EDIT.

package ent

import (
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/schema"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/season"
	"capstone-cs.eng.utah.edu/mapeleven/mapeleven/ent/team"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	seasonFields := schema.Season{}.Fields()
	_ = seasonFields
	// seasonDescCurrent is the schema descriptor for current field.
	seasonDescCurrent := seasonFields[3].Descriptor()
	// season.DefaultCurrent holds the default value on creation for the current field.
	season.DefaultCurrent = seasonDescCurrent.Default.(bool)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCode is the schema descriptor for code field.
	teamDescCode := teamFields[2].Descriptor()
	// team.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	team.CodeValidator = teamDescCode.Validators[0].(func(string) error)
}
