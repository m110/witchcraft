package archetype

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/engine"

	"github.com/m110/witchcraft/assets"
	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/spell"
)

func NewRandomCharacter(w donburi.World, position math.Vec2) {
	c := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Character,
		),
	)

	transform.Transform.Get(c).LocalPosition = position
	transform.Transform.Get(c).LocalScale = math.Vec2{X: 2, Y: 2}

	character := component.CharacterData{
		Body:       assets.RandomFrom(assets.Bodies),
		Hair:       assets.RandomFromOrEmpty(assets.Hairs),
		FacialHair: assets.RandomFromOrEmpty(assets.FacialHairs),
		Equipment: component.Equipment{
			Head:     assets.RandomFromOrEmpty(assets.HeadArmors),
			Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
			Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
			Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
			MainHand: assets.RandomFromOrEmpty(assets.MainHandWeapons),
			OffHand:  assets.RandomFromOrEmpty(assets.OffHandWeapons),
		},
	}

	component.Character.Set(c, &character)

	component.Sprite.SetValue(c, component.SpriteData{
		Image: character.Image(),
	})
}

func NewCharacter(w donburi.World, position math.Vec2) {
	c := w.Entry(
		w.Create(
			transform.Transform,
			component.Velocity,
			component.Sprite,
			component.Character,
			component.Input,
			component.Health,
			component.Mana,
			component.Caster,
		),
	)

	component.Input.Set(c, &component.InputData{
		MoveUpKey:    ebiten.KeyW,
		MoveRightKey: ebiten.KeyD,
		MoveDownKey:  ebiten.KeyS,
		MoveLeftKey:  ebiten.KeyA,
		SpellKeyA:    ebiten.Key1,
		SpellKeyB:    ebiten.Key2,
		SpellKeyC:    ebiten.Key3,
		CastKey:      ebiten.KeySpace,
	})

	component.Health.Set(c, &component.HealthData{
		Health:    100,
		MaxHealth: 100,
	})

	component.Mana.Set(c, &component.ManaData{
		Mana:           100,
		MaxMana:        100,
		ManaRegenTimer: engine.NewTimer(time.Millisecond * 100),
		ManaRegen:      1,
	})

	component.Caster.Set(c, &component.CasterData{
		KnownSpells: []component.Spell{
			component.NewSpell(spell.FireBall),
			component.NewSpell(spell.LightningBolt),
			component.NewSpell(spell.Spark),
		},
	})
	component.Caster.Get(c).PrepareSpell(0)

	transform.Transform.Get(c).LocalPosition = position
	transform.Transform.Get(c).LocalScale = math.Vec2{X: 2, Y: 2}

	character := component.CharacterData{
		Body:       assets.RandomFrom(assets.Bodies),
		Hair:       &assets.Hairs[37],
		FacialHair: &assets.FacialHairs[18],
		Equipment: component.Equipment{
			Chest:    assets.RandomFromOrEmpty(assets.ChestArmors),
			Legs:     assets.RandomFromOrEmpty(assets.LegsArmors),
			Feet:     assets.RandomFromOrEmpty(assets.FeetArmors),
			MainHand: &assets.MainHandWeapons[1],
		},
	}

	component.Character.Set(c, &character)

	component.Sprite.SetValue(c, component.SpriteData{
		Image: character.Image(),
	})

	NewProgressBar(
		w,
		math.Vec2{X: 10, Y: 10},
		100, 10,
		colornames.Red,
		func(bar *component.ProgressBarData) {
			h := component.Health.Get(c)
			bar.Value = h.Health
			bar.MaxValue = h.MaxHealth
		},
	)

	NewProgressBar(
		w,
		math.Vec2{X: 10, Y: 25},
		100, 10,
		colornames.Blue,
		func(bar *component.ProgressBarData) {
			h := component.Mana.Get(c)
			bar.Value = h.Mana
			bar.MaxValue = h.MaxMana
		},
	)

	caster := component.Caster.Get(c)

	for i := range caster.KnownSpells {
		i := i
		s := caster.KnownSpells[i]

		text := w.Entry(w.Create(
			transform.Transform,
			component.Text,
		))

		pos := math.Vec2{X: 10, Y: 100 + float64(i*15)}

		transform.GetTransform(text).LocalPosition = pos
		transform.GetTransform(text).LocalPosition.X += 55
		transform.GetTransform(text).LocalPosition.Y += 8
		component.Text.Set(text, &component.TextData{
			Text: s.Template.Name,
		})

		NewProgressBar(
			w,
			pos,
			50, 10,
			colornames.Gray,
			func(bar *component.ProgressBarData) {
				if s.CooldownTimer.TargetFrames() == 0 || s.CooldownTimer.IsReady() {
					if caster.PreparedSpellIndex != nil && *caster.PreparedSpellIndex == i {
						bar.BackgroundColor = colornames.Lightgreen
					} else {
						bar.BackgroundColor = colornames.White
					}
				} else {
					bar.BackgroundColor = colornames.Darkgrey
				}

				if s.CooldownTimer.TargetFrames() == 0 {
					bar.Value = 100
				} else {
					bar.Value = int(s.CooldownTimer.PercentDone() * 100)
				}

				bar.MaxValue = 100
			},
		)
	}

	castingPB := NewProgressBar(
		w,
		math.Vec2{X: 0, Y: -25},
		30, 3,
		colornames.Green,
		func(bar *component.ProgressBarData) {
			caster := component.Caster.Get(c)
			preparedSpell, ok := caster.PreparedSpell()
			if !ok {
				bar.Value = 0
				bar.MaxValue = 0
				return
			}

			bar.Value = int(preparedSpell.CastingTimer.PercentDone() * 100)
			bar.MaxValue = 100
		},
	)
	component.Sprite.Get(castingPB).Pivot = component.SpritePivotCenter
	transform.AppendChild(c, castingPB, false)
}
