package archetype

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"golang.org/x/image/colornames"

	"github.com/m110/witchcraft/component"
	"github.com/m110/witchcraft/engine"
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
		Body:       engine.RandomFrom(component.AllBodyParts.Bodies),
		Hair:       engine.RandomFromOrEmpty(component.AllBodyParts.Hairs),
		FacialHair: engine.RandomFromOrEmpty(component.AllBodyParts.FacialHairs),
		Equipment: component.Equipment{
			Head:     engine.RandomFromOrEmpty(component.AllBodyParts.HeadArmors),
			Chest:    engine.RandomFromOrEmpty(component.AllBodyParts.ChestArmors),
			Legs:     engine.RandomFromOrEmpty(component.AllBodyParts.LegsArmors),
			Feet:     engine.RandomFromOrEmpty(component.AllBodyParts.FeetArmors),
			MainHand: engine.RandomFromOrEmpty(component.AllBodyParts.MainHandWeapons),
			OffHand:  engine.RandomFromOrEmpty(component.AllBodyParts.OffHandWeapons),
		},
	}

	component.Character.Set(c, &character)

	component.Sprite.SetValue(c, component.SpriteData{
		Image: character.Image(),
	})
}

func PlayerUIBasePosition(index int, screenWidth int, screenHeight int) math.Vec2 {
	positions := []math.Vec2{
		{X: 5, Y: 5},
		{X: float64(screenWidth) - 200, Y: 5},
		{X: 5, Y: float64(screenHeight) - 100},
		{X: float64(screenWidth) - 200, Y: float64(screenHeight) - 100},
	}

	return positions[index]
}

const (
	playerNextSkillPointExperience = 10
)

func NewPlayer(w donburi.World, playerID int, gamepadID *ebiten.GamepadID, position math.Vec2, class Class) *donburi.Entry {
	player := w.Entry(
		w.Create(
			transform.Transform,
			component.Player,
			component.Team,
			component.Velocity,
			component.Direction,
			component.Sprite,
			component.Character,
			component.Input,
			component.Health,
			component.Mana,
			component.Caster,
			component.AuraHolder,
			component.Collider,
			component.Mover,
			component.Experience,
		),
	)

	component.Player.Set(player, &component.PlayerData{
		PlayerID: playerID,
	})
	component.Team.Set(player, &component.TeamData{
		TeamID: component.TeamID(playerID),
	})

	component.Input.Set(player, &component.InputData{
		GamepadID:    gamepadID,
		MoveUpKey:    ebiten.KeyW,
		MoveRightKey: ebiten.KeyD,
		MoveDownKey:  ebiten.KeyS,
		MoveLeftKey:  ebiten.KeyA,
		SpellKeyA:    ebiten.Key1,
		SpellKeyB:    ebiten.Key2,
		SpellKeyC:    ebiten.Key3,
		CastKey:      ebiten.KeySpace,
	})

	component.Health.Set(player, &component.HealthData{
		Health:    100,
		MaxHealth: 100,
	})

	component.Mana.Set(player, &component.ManaData{
		Mana:           100,
		MaxMana:        100,
		ManaRegenTimer: engine.NewTimer(time.Millisecond * 100),
		ManaRegen:      1,
	})

	component.Caster.Set(player, &component.CasterData{
		KnownSpells: class.Spells,
	})
	component.Caster.Get(player).PrepareSpell(0)

	scale := 2.0

	transform.Transform.Get(player).LocalPosition = position
	transform.Transform.Get(player).LocalScale = math.Vec2{X: scale, Y: scale}

	component.Character.Set(player, &class.Character)

	component.Sprite.SetValue(player, component.SpriteData{
		Image: class.Character.Image(),
		Layer: component.SpriteLayerUnits,
		Pivot: component.SpritePivotCenter,
	})

	bounds := class.Character.Image().Bounds()
	spriteWidth := float64(bounds.Dx())
	spriteHeight := float64(bounds.Dy())

	component.Collider.SetValue(player, component.ColliderData{
		Offset: math.Vec2{
			X: -spriteWidth / 2,
			Y: -spriteHeight / 2,
		},
		Width:  spriteWidth,
		Height: spriteHeight,
		Layer:  component.CollisionLayerPlayers,
	})

	component.Mover.SetValue(player, component.MoverData{
		Speed: 3,
	})

	component.Experience.SetValue(player, component.ExperienceData{
		NextSkillPointExperience: playerNextSkillPointExperience,
	})

	settings := component.MustFindGame(w).Settings
	baseUIPosition := PlayerUIBasePosition(playerID, settings.ScreenWidth, settings.ScreenHeight)

	uiParent := w.Entry(w.Create(transform.Transform))
	transform.Transform.Get(uiParent).LocalPosition = baseUIPosition

	healthBar := NewProgressBar(
		w,
		math.Vec2{X: 10, Y: 10},
		100, 10,
		colornames.Red,
		func(bar *component.ProgressBarData) {
			h := component.Health.Get(player)
			bar.Value = h.Health
			bar.MaxValue = h.MaxHealth
		},
	)
	transform.AppendChild(uiParent, healthBar, false)

	manaBar := NewProgressBar(
		w,
		math.Vec2{X: 10, Y: 25},
		100, 10,
		colornames.Blue,
		func(bar *component.ProgressBarData) {
			h := component.Mana.Get(player)
			bar.Value = h.Mana
			bar.MaxValue = h.MaxMana
		},
	)
	transform.AppendChild(uiParent, manaBar, false)

	experienceBar := NewProgressBar(
		w,
		math.Vec2{X: 10, Y: 40},
		100, 5,
		colornames.Yellow,
		func(bar *component.ProgressBarData) {
			exp := component.Experience.Get(player)
			bar.Value = exp.Experience
			bar.MaxValue = exp.NextSkillPointExperience
		},
	)
	transform.AppendChild(uiParent, experienceBar, false)

	skillPointsText := w.Entry(
		w.Create(
			transform.Transform,
			component.Text,
		),
	)
	transform.Transform.Get(skillPointsText).LocalPosition = math.Vec2{X: 125, Y: 25}
	component.Text.Set(skillPointsText, &component.TextData{
		Size:  component.TextSizeLarge,
		Text:  "",
		Color: color.White,
		Update: func(text *component.TextData) {
			exp := component.Experience.Get(player)
			if exp.SkillPoints == 0 {
				text.Text = ""
			} else {
				text.Text = fmt.Sprint(exp.SkillPoints)
			}
		},
	})

	transform.AppendChild(uiParent, skillPointsText, false)

	caster := component.Caster.Get(player)

	for i := range caster.KnownSpells {
		i := i
		s := caster.KnownSpells[i]

		text := w.Entry(w.Create(
			transform.Transform,
			component.Text,
		))

		pos := math.Vec2{X: 10, Y: 65 + float64(i*15)}

		spellBar := NewProgressBar(
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
		transform.AppendChild(uiParent, spellBar, false)

		transform.GetTransform(text).LocalPosition = math.Vec2{X: 55, Y: 8}
		component.Text.Set(text, &component.TextData{
			Text: s.Template.Name,
		})
		transform.AppendChild(spellBar, text, false)
	}

	aurasUI := w.Entry(
		w.Create(
			transform.Transform,
			component.List,
		),
	)
	component.List.Set(aurasUI, &component.ListData{
		Offset: math.Vec2{X: 36, Y: 0},
	})

	transform.Transform.Get(aurasUI).LocalPosition = math.Vec2{X: 10, Y: 100}
	transform.AppendChild(uiParent, aurasUI, false)

	component.AuraHolder.Set(player, &component.AuraHolderData{
		UI: aurasUI,
	})

	castingPB := NewProgressBar(
		w,
		math.Vec2{X: 0, Y: -25},
		30, 3,
		colornames.Green,
		func(bar *component.ProgressBarData) {
			caster := component.Caster.Get(player)
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
	transform.AppendChild(player, castingPB, false)

	crosshair := w.Entry(
		w.Create(
			transform.Transform,
			component.Sprite,
			component.Crosshair,
		),
	)

	crosshairImage := ebiten.NewImage(3, 3)
	crosshairImage.Fill(colornames.Orange)
	component.Sprite.Set(crosshair, &component.SpriteData{
		Image: crosshairImage,
		Pivot: component.SpritePivotCenter,
		Layer: component.SpriteLayerUI,
	})

	transform.AppendChild(player, crosshair, false)

	collector := w.Entry(
		w.Create(
			transform.Transform,
			component.Collector,
			component.Sprite,
			component.Collider,
		),
	)

	transform.Transform.Get(collector).LocalPosition = math.Vec2{X: 0, Y: 0}

	collectorColliderWidth := spriteWidth * 2.5
	collectorColliderHeight := spriteHeight * 2.5

	component.Collider.Set(collector, &component.ColliderData{
		Offset: math.Vec2{
			X: collectorColliderWidth / -2,
			Y: collectorColliderHeight / -2,
		},
		Width:  collectorColliderWidth,
		Height: collectorColliderHeight,
		Layer:  component.CollisionLayerCollector,
	})

	// TODO for debug purposes
	component.Sprite.Set(collector, &component.SpriteData{
		Image: ebiten.NewImage(1, 1),
		Layer: component.SpriteLayerUI,
	})

	transform.AppendChild(player, collector, false)

	return player
}
