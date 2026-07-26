package event

func CancelExplosions(ev ExplosionEvent) {
	ev.Ctx.Cancel()
}

func CancelLiquidFlow(ev LiquidFlowEvent) {
	ev.Ctx.Cancel()
}

func CancelLiquidDecay(ev LiquidDecayEvent) {
	ev.Ctx.Cancel()
}

func CancelLiquidHarden(ev LiquidHardenEvent) {
	ev.Ctx.Cancel()
}

func CancelFireSpread(ev FireSpreadEvent) {
	ev.Ctx.Cancel()
}

func CancelBlockBurn(ev BlockBurnEvent) {
	ev.Ctx.Cancel()
}

func CancelCropTrample(ev CropTrampleEvent) {
	ev.Ctx.Cancel()
}

func CancelLeavesDecay(ev LeavesDecayEvent) {
	ev.Ctx.Cancel()
}

func CancelPortalCreate(ev PortalCreateEvent) {
	ev.Ctx.Cancel()
}

func CancelPortalActivate(ev PortalActivateEvent) {
	ev.Ctx.Cancel()
}

func CancelRedstoneUpdate(ev RedstoneUpdateEvent) {
	ev.Ctx.Cancel()
}
