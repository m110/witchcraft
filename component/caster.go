package component

import (
	"github.com/yohamta/donburi"

	"github.com/m110/witchcraft/engine"
	"github.com/m110/witchcraft/spell"
)

type CasterData struct {
	PreparedSpellIndex *int
	KnownSpells        []Spell
	IsCasting          bool
	IsChannelling      bool
}

func (d *CasterData) PrepareSpell(index int) {
	if index < 0 || index >= len(d.KnownSpells) {
		return
	}
	d.PreparedSpellIndex = &index
}

func (d *CasterData) PreparedSpell() (Spell, bool) {
	if d.PreparedSpellIndex == nil {
		return Spell{}, false
	}

	return d.KnownSpells[*d.PreparedSpellIndex], true
}

var Caster = donburi.NewComponentType[CasterData]()

type Spell struct {
	Template spell.Spell

	CastingTimer        *engine.Timer
	ChannellingTimer    *engine.Timer
	MaxChannellingTimer *engine.Timer
	CooldownTimer       *engine.Timer
}

func NewSpell(template spell.Spell) Spell {
	s := Spell{
		Template:            template,
		CastingTimer:        engine.NewTimer(template.CastingTime),
		ChannellingTimer:    engine.NewTimer(template.ChannelTickDuration),
		MaxChannellingTimer: engine.NewTimer(template.MaxChannelTime),
		CooldownTimer:       engine.NewTimer(template.Cooldown),
	}

	s.CooldownTimer.Finish()

	return s
}
