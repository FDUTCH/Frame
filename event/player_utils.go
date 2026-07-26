package event

func CancelChat(ev PlayerChatEvent) {
	ev.Ctx.Cancel()
}

func CancelFoodLoss(ev PlayerFoodLossEvent) {
	ev.Ctx.Cancel()
}

func CancelHeal(ev PlayerHealEvent) {
	ev.Ctx.Cancel()
}

func CancelHurt(ev PlayerHurtEvent) {
	ev.Ctx.Cancel()
}

func CancelSetOnFire(ev PlayerSetOnFireEvent) {
	ev.Ctx.Cancel()
}

func CancelSkinChange(ev PlayerSkinChangeEvent) {
	ev.Ctx.Cancel()
}

func CancelFireExtinguish(ev PlayerFireExtinguishEvent) {
	ev.Ctx.Cancel()
}

func CancelStartBreak(ev PlayerStartBreakEvent) {
	ev.Ctx.Cancel()
}

func CancelBlockBreak(ev PlayerBlockBreakEvent) {
	ev.Ctx.Cancel()
}

func CancelBlockPlace(ev PlayerBlockPlaceEvent) {
	ev.Ctx.Cancel()
}

func CancelBlockPick(ev PlayerBlockPickEvent) {
	ev.Ctx.Cancel()
}

func CancelItemUse(ev PlayerItemUseEvent) {
	ev.Ctx.Cancel()
}

func CancelItemUseOnBlock(ev PlayerItemUseOnBlockEvent) {
	ev.Ctx.Cancel()
}

func CancelItemUseOnEntity(ev PlayerItemUseOnEntityEvent) {
	ev.Ctx.Cancel()
}

func CancelItemRelease(ev PlayerItemReleaseEvent) {
	ev.Ctx.Cancel()
}

func CancelItemConsume(ev PlayerItemConsumeEvent) {
	ev.Ctx.Cancel()
}

func CancelAttackEntity(ev PlayerAttackEntityEvent) {
	ev.Ctx.Cancel()
}

func CancelExperienceGain(ev PlayerExperienceGainEvent) {
	ev.Ctx.Cancel()
}

func CancelPunchAir(ev PlayerPunchAirEvent) {
	ev.Ctx.Cancel()
}

func CancelSignEdit(ev PlayerSignEditEvent) {
	ev.Ctx.Cancel()
}

func CancelSleep(ev PlayerSleepEvent) {
	ev.Ctx.Cancel()
}

func CancelLecternPageTurn(ev PlayerLecternPageTurnEvent) {
	ev.Ctx.Cancel()
}

func CancelItemDamage(ev PlayerItemDamageEvent) {
	ev.Ctx.Cancel()
}

func CancelItemPickup(ev PlayerItemPickupEvent) {
	ev.Ctx.Cancel()
}

func CancelHeldSlotChange(ev PlayerHeldSlotChangeEvent) {
	ev.Ctx.Cancel()
}

func CancelItemDrop(ev PlayerItemDropEvent) {
	ev.Ctx.Cancel()
}

func CancelTransfer(ev PlayerTransferEvent) {
	ev.Ctx.Cancel()
}

func CancelCommandExecution(ev PlayerCommandExecutionEvent) {
	ev.Ctx.Cancel()
}
