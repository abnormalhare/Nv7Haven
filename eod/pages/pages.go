package pages

import (
	"github.com/Nv7-Github/Nv7Haven/eod/base"
	"github.com/Nv7-Github/Nv7Haven/eod/categories"
	"github.com/Nv7-Github/Nv7Haven/eod/types"
	"github.com/Nv7-Github/sevcord/v2"
	"github.com/jmoiron/sqlx"
)

type Pages struct {
	base       *base.Base
	db         *sqlx.DB
	categories *categories.Categories
	s          *sevcord.Sevcord
}

func (p *Pages) Init() {
	// Inv
	p.s.RegisterSlashCommand(sevcord.NewSlashCommand(
		"inv",
		"View your inventory!",
		p.Inv,
		sevcord.NewOption("user", "The user to view the inventory of!", sevcord.OptionKindUser, false),
		sevcord.NewOption("sort", "The sort order of the inventory!", sevcord.OptionKindString, false).
			AddChoices(types.Sorts...),
	))
	p.s.AddButtonHandler("inv", p.InvHandler)

	// Lb
	p.s.RegisterSlashCommand(sevcord.NewSlashCommand(
		"lb",
		"View the leaderboard!",
		p.Lb,
		sevcord.NewOption("sort", "The sort order of the leaderboard!", sevcord.OptionKindString, false).
			AddChoices(lbSorts...),
		sevcord.NewOption("user", "The user to view the leaderboard from the point of view of!", sevcord.OptionKindUser, false),
	))
	p.s.AddButtonHandler("lb", p.LbHandler)

	// Categories
	p.s.RegisterSlashCommand(sevcord.NewSlashCommandGroup("cat", "View categories!", sevcord.NewSlashCommand(
		"list",
		"View a list of every categories!",
		p.CatList,
		sevcord.NewOption("sort", "How to order the categories!", sevcord.OptionKindString, false).AddChoices(catListSorts...),
	), sevcord.NewSlashCommand(
		"view",
		"View a category's elements",
		p.Cat,
		sevcord.NewOption("category", "The category to view!", sevcord.OptionKindString, true).AutoComplete(p.categories.Autocomplete),
		sevcord.NewOption("sort", "How to order the categories!", sevcord.OptionKindString, false).AddChoices(catListSorts...),
	)))
	p.s.AddButtonHandler("catlist", p.CatListHandler)
	p.s.AddButtonHandler("cat", p.CatHandler)
}

func NewPages(base *base.Base, db *sqlx.DB, s *sevcord.Sevcord, categories *categories.Categories) *Pages {
	p := &Pages{
		base:       base,
		db:         db,
		categories: categories,
		s:          s,
	}
	p.Init()
	return p
}