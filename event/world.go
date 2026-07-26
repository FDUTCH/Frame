package event

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type LiquidFlowEvent struct {
	Ctx      *world.Context
	From     cube.Pos
	Into     cube.Pos
	Liquid   world.Liquid
	Replaced world.Block
}

type LiquidDecayEvent struct {
	Ctx    *world.Context
	Pos    cube.Pos
	Before world.Liquid
	After  world.Liquid
}

type LiquidHardenEvent struct {
	Ctx            *world.Context
	HardenedPos    cube.Pos
	LiquidHardened world.Block
	OtherLiquid    world.Block
	NewBlock       world.Block
}

type SoundEvent struct {
	Ctx *world.Context
	S   world.Sound
	Pos mgl64.Vec3
}

type FireSpreadEvent struct {
	Ctx  *world.Context
	From cube.Pos
	To   cube.Pos
}

type BlockBurnEvent struct {
	Ctx *world.Context
	Pos cube.Pos
}

type CropTrampleEvent struct {
	Ctx *world.Context
	Pos cube.Pos
}

type LeavesDecayEvent struct {
	Ctx *world.Context
	Pos cube.Pos
}

type PortalCreateEvent struct {
	Ctx        *world.Context
	PortalType world.Dimension
	Positions  []cube.Pos
}

type PortalActivateEvent struct {
	Ctx        *world.Context
	PortalType world.Dimension
	Positions  []cube.Pos
}

type EntitySpawnEvent struct {
	Tx *world.Tx
	E  world.Entity
}

type EntityDespawnEvent struct {
	Tx *world.Tx
	E  world.Entity
}

type ExplosionEvent struct {
	Ctx            *world.Context
	Src            world.ExplosionSource
	Entities       *[]world.Entity
	Blocks         *[]cube.Pos
	ItemDropChance *float64
	SpawnFire      *bool
}

type RedstoneUpdateEvent struct {
	Ctx    *world.Context
	Update world.RedstoneUpdate
}

type WorldCloseEvent struct {
	Tx *world.Tx
}

var _ world.Handler = (*WorldHandler)(nil)

// WorldHandler publishes all events to the internal buses.
type WorldHandler struct {
	buses []*Bus
}

func NewWorldHandler(buses ...*Bus) *WorldHandler {
	return &WorldHandler{buses: buses}
}

func (w *WorldHandler) HandleLiquidFlow(ctx *world.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	PublishMultiple(w.buses, LiquidFlowEvent{
		Ctx:      ctx,
		From:     from,
		Into:     into,
		Liquid:   liquid,
		Replaced: replaced,
	})
}

func (w *WorldHandler) HandleLiquidDecay(ctx *world.Context, pos cube.Pos, before, after world.Liquid) {
	PublishMultiple(w.buses, LiquidDecayEvent{
		Ctx:    ctx,
		Pos:    pos,
		Before: before,
		After:  after,
	})
}

func (w *WorldHandler) HandleLiquidHarden(ctx *world.Context, hardenedPos cube.Pos, liquidHardened, otherLiquid, newBlock world.Block) {
	PublishMultiple(w.buses, LiquidHardenEvent{
		Ctx:            ctx,
		HardenedPos:    hardenedPos,
		LiquidHardened: liquidHardened,
		OtherLiquid:    otherLiquid,
		NewBlock:       newBlock,
	})
}

func (w *WorldHandler) HandleSound(ctx *world.Context, s world.Sound, pos mgl64.Vec3) {
	PublishMultiple(w.buses, SoundEvent{
		Ctx: ctx,
		S:   s,
		Pos: pos,
	})
}

func (w *WorldHandler) HandleFireSpread(ctx *world.Context, from, to cube.Pos) {
	PublishMultiple(w.buses, FireSpreadEvent{
		Ctx:  ctx,
		From: from,
		To:   to,
	})
}

func (w *WorldHandler) HandleBlockBurn(ctx *world.Context, pos cube.Pos) {
	PublishMultiple(w.buses, BlockBurnEvent{
		Ctx: ctx,
		Pos: pos,
	})
}

func (w *WorldHandler) HandleCropTrample(ctx *world.Context, pos cube.Pos) {
	PublishMultiple(w.buses, CropTrampleEvent{
		Ctx: ctx,
		Pos: pos,
	})
}

func (w *WorldHandler) HandleLeavesDecay(ctx *world.Context, pos cube.Pos) {
	PublishMultiple(w.buses, LeavesDecayEvent{
		Ctx: ctx,
		Pos: pos,
	})
}

func (w *WorldHandler) HandlePortalCreate(ctx *world.Context, portalType world.Dimension, positions []cube.Pos) {
	PublishMultiple(w.buses, PortalCreateEvent{
		Ctx:        ctx,
		PortalType: portalType,
		Positions:  positions,
	})
}

func (w *WorldHandler) HandlePortalActivate(ctx *world.Context, portalType world.Dimension, positions []cube.Pos) {
	PublishMultiple(w.buses, PortalActivateEvent{
		Ctx:        ctx,
		PortalType: portalType,
		Positions:  positions,
	})
}

func (w *WorldHandler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	PublishMultiple(w.buses, EntitySpawnEvent{
		Tx: tx,
		E:  e,
	})
}

func (w *WorldHandler) HandleEntityDespawn(tx *world.Tx, e world.Entity) {
	PublishMultiple(w.buses, EntityDespawnEvent{
		Tx: tx,
		E:  e,
	})
}

func (w *WorldHandler) HandleExplosion(ctx *world.Context, src world.ExplosionSource, entities *[]world.Entity, blocks *[]cube.Pos, itemDropChance *float64, spawnFire *bool) {
	PublishMultiple(w.buses, ExplosionEvent{
		Ctx:            ctx,
		Src:            src,
		Entities:       entities,
		Blocks:         blocks,
		ItemDropChance: itemDropChance,
		SpawnFire:      spawnFire,
	})
}

func (w *WorldHandler) HandleRedstoneUpdate(ctx *world.Context, update world.RedstoneUpdate) {
	PublishMultiple(w.buses, RedstoneUpdateEvent{
		Ctx:    ctx,
		Update: update,
	})
}

func (w *WorldHandler) HandleClose(tx *world.Tx) {
	PublishMultiple(w.buses, WorldCloseEvent{
		Tx: tx,
	})
}
