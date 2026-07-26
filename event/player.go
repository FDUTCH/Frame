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
	buses []*Bus
}

func NewPlayerHandler(buses ...*Bus) *PlayerHandler {
	return &PlayerHandler{buses: buses}
}

func (h *PlayerHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	PublishMultiple(h.buses, PlayerMoveEvent{Ctx: ctx, NewPos: newPos, NewRot: newRot})
}

func (h *PlayerHandler) HandleJump(p *player.Player) {
	PublishMultiple(h.buses, PlayerJumpEvent{Player: p})
}

func (h *PlayerHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	PublishMultiple(h.buses, PlayerTeleportEvent{Ctx: ctx, Pos: pos})
}

func (h *PlayerHandler) HandleChangeWorld(p *player.Player, before, after *world.World) {
	PublishMultiple(h.buses, PlayerChangeWorldEvent{Player: p, Before: before, After: after})
}

func (h *PlayerHandler) HandleToggleSprint(ctx *player.Context, after bool) {
	PublishMultiple(h.buses, PlayerToggleSprintEvent{Ctx: ctx, After: after})
}

func (h *PlayerHandler) HandleToggleSneak(ctx *player.Context, after bool) {
	PublishMultiple(h.buses, PlayerToggleSneakEvent{Ctx: ctx, After: after})
}

func (h *PlayerHandler) HandleChat(ctx *player.Context, message *string) {
	PublishMultiple(h.buses, PlayerChatEvent{Ctx: ctx, Message: message})
}

func (h *PlayerHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	PublishMultiple(h.buses, PlayerFoodLossEvent{Ctx: ctx, From: from, To: to})
}

func (h *PlayerHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	PublishMultiple(h.buses, PlayerHealEvent{Ctx: ctx, Health: health, Src: src})
}

func (h *PlayerHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	PublishMultiple(h.buses, PlayerHurtEvent{Ctx: ctx, Damage: damage, Immune: immune, AttackImmunity: attackImmunity, Src: src})
}

func (h *PlayerHandler) HandleSetOnFire(ctx *player.Context, duration *time.Duration) {
	PublishMultiple(h.buses, PlayerSetOnFireEvent{Ctx: ctx, Duration: duration})
}

func (h *PlayerHandler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
	PublishMultiple(h.buses, PlayerDeathEvent{Player: p, Src: src, KeepInv: keepInv})
}

func (h *PlayerHandler) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {
	PublishMultiple(h.buses, PlayerRespawnEvent{Player: p, Pos: pos, World: w})
}

func (h *PlayerHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	PublishMultiple(h.buses, PlayerSkinChangeEvent{Ctx: ctx, Skin: skin})
}

func (h *PlayerHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	PublishMultiple(h.buses, PlayerFireExtinguishEvent{Ctx: ctx, Pos: pos})
}

func (h *PlayerHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	PublishMultiple(h.buses, PlayerStartBreakEvent{Ctx: ctx, Pos: pos})
}

func (h *PlayerHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	PublishMultiple(h.buses, PlayerBlockBreakEvent{Ctx: ctx, Pos: pos, Drops: drops, Xp: xp})
}

func (h *PlayerHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	PublishMultiple(h.buses, PlayerBlockPlaceEvent{Ctx: ctx, Pos: pos, B: b})
}

func (h *PlayerHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	PublishMultiple(h.buses, PlayerBlockPickEvent{Ctx: ctx, Pos: pos, B: b})
}

func (h *PlayerHandler) HandleItemUse(ctx *player.Context) {
	PublishMultiple(h.buses, PlayerItemUseEvent{Ctx: ctx})
}

func (h *PlayerHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	PublishMultiple(h.buses, PlayerItemUseOnBlockEvent{Ctx: ctx, Pos: pos, Face: face, ClickPos: clickPos})
}

func (h *PlayerHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	PublishMultiple(h.buses, PlayerItemUseOnEntityEvent{Ctx: ctx, E: e})
}

func (h *PlayerHandler) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {
	PublishMultiple(h.buses, PlayerItemReleaseEvent{Ctx: ctx, Item: item, Dur: dur})
}

func (h *PlayerHandler) HandleItemConsume(ctx *player.Context, item item.Stack) {
	PublishMultiple(h.buses, PlayerItemConsumeEvent{Ctx: ctx, Item: item})
}

func (h *PlayerHandler) HandleAttackEntity(ctx *player.Context, e world.Entity, force, height *float64, critical *bool) {
	PublishMultiple(h.buses, PlayerAttackEntityEvent{Ctx: ctx, E: e, Force: force, Height: height, Critical: critical})
}

func (h *PlayerHandler) HandleExperienceGain(ctx *player.Context, amount *int) {
	PublishMultiple(h.buses, PlayerExperienceGainEvent{Ctx: ctx, Amount: amount})
}

func (h *PlayerHandler) HandlePunchAir(ctx *player.Context) {
	PublishMultiple(h.buses, PlayerPunchAirEvent{Ctx: ctx})
}

func (h *PlayerHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText, newText string) {
	PublishMultiple(h.buses, PlayerSignEditEvent{Ctx: ctx, Pos: pos, FrontSide: frontSide, OldText: oldText, NewText: newText})
}

func (h *PlayerHandler) HandleSleep(ctx *player.Context, sendReminder *bool) {
	PublishMultiple(h.buses, PlayerSleepEvent{Ctx: ctx, SendReminder: sendReminder})
}

func (h *PlayerHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	PublishMultiple(h.buses, PlayerLecternPageTurnEvent{Ctx: ctx, Pos: pos, OldPage: oldPage, NewPage: newPage})
}

func (h *PlayerHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage *int) {
	PublishMultiple(h.buses, PlayerItemDamageEvent{Ctx: ctx, Item: i, Damage: damage})
}

func (h *PlayerHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	PublishMultiple(h.buses, PlayerItemPickupEvent{Ctx: ctx, I: i})
}

func (h *PlayerHandler) HandleHeldSlotChange(ctx *player.Context, from, to int) {
	PublishMultiple(h.buses, PlayerHeldSlotChangeEvent{Ctx: ctx, From: from, To: to})
}

func (h *PlayerHandler) HandleItemDrop(ctx *player.Context, s item.Stack) {
	PublishMultiple(h.buses, PlayerItemDropEvent{Ctx: ctx, S: s})
}

func (h *PlayerHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	PublishMultiple(h.buses, PlayerTransferEvent{Ctx: ctx, Addr: addr})
}

func (h *PlayerHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	PublishMultiple(h.buses, PlayerCommandExecutionEvent{Ctx: ctx, Command: command, Args: args})
}

func (h *PlayerHandler) HandleQuit(p *player.Player) {
	PublishMultiple(h.buses, PlayerQuitEvent{Player: p})
}

func (h *PlayerHandler) HandleDiagnostics(p *player.Player, d session.Diagnostics) {
	PublishMultiple(h.buses, PlayerDiagnosticsEvent{Player: p, Diagnostics: d})
}
