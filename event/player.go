package event

import (
	"net"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type PlayerJoinEvent struct {
	Player *player.Player
}
type PlayerMoveEvent struct {
	Ctx    *player.Context
	NewPos mgl64.Vec3
	NewRot cube.Rotation
}

type PlayerJumpEvent struct {
	Player *player.Player
}

type PlayerTeleportEvent struct {
	Ctx *player.Context
	Pos mgl64.Vec3
}

type PlayerChangeWorldEvent struct {
	Player *player.Player
	Before *world.World
	After  *world.World
}

type PlayerToggleSprintEvent struct {
	Ctx   *player.Context
	After bool
}

type PlayerToggleSneakEvent struct {
	Ctx   *player.Context
	After bool
}

type PlayerChatEvent struct {
	Ctx     *player.Context
	Message *string
}

type PlayerFoodLossEvent struct {
	Ctx  *player.Context
	From int
	To   *int
}

type PlayerHealEvent struct {
	Ctx    *player.Context
	Health *float64
	Src    world.HealingSource
}

type PlayerHurtEvent struct {
	Ctx            *player.Context
	Damage         *float64
	Immune         bool
	AttackImmunity *time.Duration
	Src            world.DamageSource
}

type PlayerSetOnFireEvent struct {
	Ctx      *player.Context
	Duration *time.Duration
}

type PlayerDeathEvent struct {
	Player  *player.Player
	Src     world.DamageSource
	KeepInv *bool
}

type PlayerRespawnEvent struct {
	Player *player.Player
	Pos    *mgl64.Vec3
	World  **world.World
}

type PlayerSkinChangeEvent struct {
	Ctx  *player.Context
	Skin *skin.Skin
}

type PlayerFireExtinguishEvent struct {
	Ctx *player.Context
	Pos cube.Pos
}

type PlayerStartBreakEvent struct {
	Ctx *player.Context
	Pos cube.Pos
}

type PlayerBlockBreakEvent struct {
	Ctx   *player.Context
	Pos   cube.Pos
	Drops *[]item.Stack
	Xp    *int
}

type PlayerBlockPlaceEvent struct {
	Ctx *player.Context
	Pos cube.Pos
	B   world.Block
}

type PlayerBlockPickEvent struct {
	Ctx *player.Context
	Pos cube.Pos
	B   world.Block
}

type PlayerItemUseEvent struct {
	Ctx *player.Context
}

type PlayerItemUseOnBlockEvent struct {
	Ctx      *player.Context
	Pos      cube.Pos
	Face     cube.Face
	ClickPos mgl64.Vec3
}

type PlayerItemUseOnEntityEvent struct {
	Ctx *player.Context
	E   world.Entity
}

type PlayerItemReleaseEvent struct {
	Ctx  *player.Context
	Item item.Stack
	Dur  time.Duration
}

type PlayerItemConsumeEvent struct {
	Ctx  *player.Context
	Item item.Stack
}

type PlayerAttackEntityEvent struct {
	Ctx      *player.Context
	E        world.Entity
	Force    *float64
	Height   *float64
	Critical *bool
}

type PlayerExperienceGainEvent struct {
	Ctx    *player.Context
	Amount *int
}

type PlayerPunchAirEvent struct {
	Ctx *player.Context
}

type PlayerSignEditEvent struct {
	Ctx       *player.Context
	Pos       cube.Pos
	FrontSide bool
	OldText   string
	NewText   string
}

type PlayerSleepEvent struct {
	Ctx          *player.Context
	SendReminder *bool
}

type PlayerLecternPageTurnEvent struct {
	Ctx     *player.Context
	Pos     cube.Pos
	OldPage int
	NewPage *int
}

type PlayerItemDamageEvent struct {
	Ctx    *player.Context
	Item   item.Stack
	Damage *int
}

type PlayerItemPickupEvent struct {
	Ctx *player.Context
	I   *item.Stack
}

type PlayerHeldSlotChangeEvent struct {
	Ctx  *player.Context
	From int
	To   int
}

type PlayerItemDropEvent struct {
	Ctx *player.Context
	S   item.Stack
}

type PlayerTransferEvent struct {
	Ctx  *player.Context
	Addr *net.UDPAddr
}

type PlayerCommandExecutionEvent struct {
	Ctx     *player.Context
	Command cmd.Command
	Args    []string
}

type PlayerQuitEvent struct {
	Player *player.Player
}

type PlayerDiagnosticsEvent struct {
	Player      *player.Player
	Diagnostics session.Diagnostics
}

var _ player.Handler = (*PlayerHandler)(nil)

// PlayerHandler publishes all events to the internal buses.
type PlayerHandler struct {
	buses       []*Bus
	busProvider WorldBusProvider
}

func NewPlayerHandler(busProvider WorldBusProvider, buses ...*Bus) *PlayerHandler {
	return &PlayerHandler{buses: buses, busProvider: busProvider}
}

func PublishToWorldBus[T any](pl *player.Player, h *PlayerHandler, event T) {
	if h.busProvider == nil {
		return
	}
	bus, ok := h.busProvider.WorldBus(pl.Tx().World())
	if ok {
		Publish(bus, event)
	}
}

func (h *PlayerHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	ev := PlayerMoveEvent{Ctx: ctx, NewPos: newPos, NewRot: newRot}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleJump(p *player.Player) {
	ev := PlayerJumpEvent{Player: p}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	ev := PlayerTeleportEvent{Ctx: ctx, Pos: pos}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleChangeWorld(p *player.Player, before, after *world.World) {
	ev := PlayerChangeWorldEvent{Player: p, Before: before, After: after}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleToggleSprint(ctx *player.Context, after bool) {
	ev := PlayerToggleSprintEvent{Ctx: ctx, After: after}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleToggleSneak(ctx *player.Context, after bool) {
	ev := PlayerToggleSneakEvent{Ctx: ctx, After: after}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleChat(ctx *player.Context, message *string) {
	ev := PlayerChatEvent{Ctx: ctx, Message: message}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	ev := PlayerFoodLossEvent{Ctx: ctx, From: from, To: to}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	ev := PlayerHealEvent{Ctx: ctx, Health: health, Src: src}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	ev := PlayerHurtEvent{Ctx: ctx, Damage: damage, Immune: immune, AttackImmunity: attackImmunity, Src: src}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleSetOnFire(ctx *player.Context, duration *time.Duration) {
	ev := PlayerSetOnFireEvent{Ctx: ctx, Duration: duration}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
	ev := PlayerDeathEvent{Player: p, Src: src, KeepInv: keepInv}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {
	ev := PlayerRespawnEvent{Player: p, Pos: pos, World: w}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	ev := PlayerSkinChangeEvent{Ctx: ctx, Skin: skin}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	ev := PlayerFireExtinguishEvent{Ctx: ctx, Pos: pos}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	ev := PlayerStartBreakEvent{Ctx: ctx, Pos: pos}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	ev := PlayerBlockBreakEvent{Ctx: ctx, Pos: pos, Drops: drops, Xp: xp}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	ev := PlayerBlockPlaceEvent{Ctx: ctx, Pos: pos, B: b}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	ev := PlayerBlockPickEvent{Ctx: ctx, Pos: pos, B: b}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemUse(ctx *player.Context) {
	ev := PlayerItemUseEvent{Ctx: ctx}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	ev := PlayerItemUseOnBlockEvent{Ctx: ctx, Pos: pos, Face: face, ClickPos: clickPos}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	ev := PlayerItemUseOnEntityEvent{Ctx: ctx, E: e}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {
	ev := PlayerItemReleaseEvent{Ctx: ctx, Item: item, Dur: dur}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemConsume(ctx *player.Context, item item.Stack) {
	ev := PlayerItemConsumeEvent{Ctx: ctx, Item: item}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool) {
	ev := PlayerAttackEntityEvent{Ctx: ctx, E: e, Force: force, Height: height, Critical: critical}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleExperienceGain(ctx *player.Context, amount *int) {
	ev := PlayerExperienceGainEvent{Ctx: ctx, Amount: amount}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandlePunchAir(ctx *player.Context) {
	ev := PlayerPunchAirEvent{Ctx: ctx}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string) {
	ev := PlayerSignEditEvent{Ctx: ctx, Pos: pos, FrontSide: frontSide, OldText: oldText, NewText: newText}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleSleep(ctx *player.Context, sendReminder *bool) {
	ev := PlayerSleepEvent{Ctx: ctx, SendReminder: sendReminder}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	ev := PlayerLecternPageTurnEvent{Ctx: ctx, Pos: pos, OldPage: oldPage, NewPage: newPage}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage *int) {
	ev := PlayerItemDamageEvent{Ctx: ctx, Item: i, Damage: damage}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	ev := PlayerItemPickupEvent{Ctx: ctx, I: i}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleHeldSlotChange(ctx *player.Context, from, to int) {
	ev := PlayerHeldSlotChangeEvent{Ctx: ctx, From: from, To: to}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleItemDrop(ctx *player.Context, s item.Stack) {
	ev := PlayerItemDropEvent{Ctx: ctx, S: s}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	ev := PlayerTransferEvent{Ctx: ctx, Addr: addr}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	ev := PlayerCommandExecutionEvent{Ctx: ctx, Command: command, Args: args}
	PublishToWorldBus(ctx.Player(), h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleQuit(p *player.Player) {
	ev := PlayerQuitEvent{Player: p}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}

func (h *PlayerHandler) HandleDiagnostics(p *player.Player, d session.Diagnostics) {
	ev := PlayerDiagnosticsEvent{Player: p, Diagnostics: d}
	PublishToWorldBus(p, h, ev)
	PublishMultiple(h.buses, ev)
}
