// Code generated by ent, DO NOT EDIT.

package ent

import (
	"mapeleven/db/ent/country"
	"mapeleven/db/ent/schema"
	"mapeleven/db/ent/season"
	"mapeleven/db/ent/standings"
	"mapeleven/db/ent/team"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescCode is the schema descriptor for code field.
	countryDescCode := countryFields[0].Descriptor()
	// country.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	country.CodeValidator = countryDescCode.Validators[0].(func(string) error)
	seasonFields := schema.Season{}.Fields()
	_ = seasonFields
	// seasonDescCurrent is the schema descriptor for current field.
	seasonDescCurrent := seasonFields[3].Descriptor()
	// season.DefaultCurrent holds the default value on creation for the current field.
	season.DefaultCurrent = seasonDescCurrent.Default.(bool)
	standingsFields := schema.Standings{}.Fields()
	_ = standingsFields
	// standingsDescPoints is the schema descriptor for points field.
	standingsDescPoints := standingsFields[1].Descriptor()
	// standings.DefaultPoints holds the default value on creation for the points field.
	standings.DefaultPoints = standingsDescPoints.Default.(int)
	// standingsDescGoalsDiff is the schema descriptor for goalsDiff field.
	standingsDescGoalsDiff := standingsFields[2].Descriptor()
	// standings.DefaultGoalsDiff holds the default value on creation for the goalsDiff field.
	standings.DefaultGoalsDiff = standingsDescGoalsDiff.Default.(int)
	// standingsDescGroup is the schema descriptor for group field.
	standingsDescGroup := standingsFields[3].Descriptor()
	// standings.DefaultGroup holds the default value on creation for the group field.
	standings.DefaultGroup = standingsDescGroup.Default.(string)
	// standingsDescForm is the schema descriptor for form field.
	standingsDescForm := standingsFields[4].Descriptor()
	// standings.DefaultForm holds the default value on creation for the form field.
	standings.DefaultForm = standingsDescForm.Default.(string)
	// standingsDescLastUpdated is the schema descriptor for LastUpdated field.
	standingsDescLastUpdated := standingsFields[25].Descriptor()
	// standings.DefaultLastUpdated holds the default value on creation for the LastUpdated field.
	standings.DefaultLastUpdated = standingsDescLastUpdated.Default.(func() time.Time)
	// standings.UpdateDefaultLastUpdated holds the default value on update for the LastUpdated field.
	standings.UpdateDefaultLastUpdated = standingsDescLastUpdated.UpdateDefault.(func() time.Time)
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCode is the schema descriptor for code field.
	teamDescCode := teamFields[3].Descriptor()
	// team.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	team.CodeValidator = teamDescCode.Validators[0].(func(string) error)
}
