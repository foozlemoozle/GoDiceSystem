package dice

type IDiePoolManager interface {
	CreateDiePool(sides int, size int) uint64
	GetDiePool(poolId uint64) IDiePool
	GetFatePool() IDiePool
	ContestPools(actorId uint64, actorContext DieRollContext, defenderId uint64, defenderContext DieRollContext) int
	ContestFate(actorId uint64, actorContext DieRollContext) int
}

type IBuffManager interface {
	AddBuff(poolId int, buff IDiePoolBuff) BuffId
	BuffDiePoolSize(poolId uint64, duration int, amount int) BuffId
	BuffDieRollBonus(poolId uint64, duration int, amount int, context DieRollContext) BuffId
	BuffNumDieRolls(poolId uint64, duration int, amount int) BuffId
	BuffNumDieRollsByCost(poolId uint64, duration int) BuffId
	BuffRemoveDice(poolId uint64, duration int) BuffId

	DecrementBuffTimesAndRemove()
	DecrementBuffTimeAndRemove(id BuffId)

	ClearAllBuffsFromPool(poolId int)
	RemoveBuffFromPool(poolId int, buff IDiePoolBuff)
}

type BuffId struct {
	poolId uint64
	buffId uint64
}

var nextPoolId uint64 = 0
var nextBuffId uint64 = 0
var defaultSides int = 6

type DiePoolManager struct {
	diePools   map[uint64]IDiePool
	fatePoolId uint64
	buffs      map[BuffId]IDiePoolBuff
}

func MakeDiePoolManager() IDiePoolManager {
	diePoolMan := new(DiePoolManager)
	diePoolMan.diePools = make(map[uint64]IDiePool)
	diePoolMan.fatePoolId = diePoolMan.CreateDiePool(defaultSides, 1)
	diePoolMan.AddBuff(diePoolMan.fatePoolId, makeBuffInfiniteDice())

	return diePoolMan
}

func (man *DiePoolManager) CreateDiePool(sides int, size int) uint64 {
	id := nextPoolId
	nextPoolId++

	man.diePools[id] = DiePool(size, sides, id)
	return id
}

func (man *DiePoolManager) GetDiePool(poolId uint64) IDiePool {
	return man.diePools[poolId]
}

func (man *DiePoolManager) GetFatePool() IDiePool {
	return man.GetDiePool(man.fatePoolId)
}

func (man *DiePoolManager) ContestPools(actorId uint64, actorContext DieRollContext, defenderId uint64, defenderContext DieRollContext) int {
	actor := man.GetDiePool(actorId)
	defender := man.GetDiePool(defenderId)

	return contestPools(actor, actorContext, defender, defenderContext)
}

func (man *DiePoolManager) ContestFate(actorId uint64, actorContext DieRollContext) int {
	actor := man.GetDiePool(actorId)
	defender := man.GetFatePool()

	actorContext = actorContext | VersusFate
	var defenderContext DieRollContext = VersusFate

	return contestPools(actor, actorContext, defender, defenderContext)
}

func (man *DiePoolManager) AddBuff(poolId uint64, buff IDiePoolBuff) BuffId {
	man.GetDiePool(poolId).AddBuff(buff)

	id := BuffId{poolId: poolId, buffId: nextBuffId}
	man.buffs[id] = buff

	nextBuffId++

	return id
}

func (man *DiePoolManager) BuffDiePoolSize(poolId uint64, duration int, amount int) BuffId {
	buff := makeBuffDiePoolSize(duration, amount)
	return man.AddBuff(poolId, buff)
}

func (man *DiePoolManager) BuffDieRollBonus(poolId uint64, duration int, amount int, context DieRollContext) BuffId {
	buff := makeBuffDieRollBonus(duration, amount, context)
	return man.AddBuff(poolId, buff)
}

func (man *DiePoolManager) BuffNumDieRolls(poolId uint64, duration int, amount int) BuffId {
	buff := makeBuffNumDieRolls(duration, amount)
	return man.AddBuff(poolId, buff)
}

func (man *DiePoolManager) BuffNumDieRollsByCost(poolId uint64, duration int) BuffId {
	buff := makeBuffNumDieRollsByCost(duration)
	return man.AddBuff(poolId, buff)
}

func (man *DiePoolManager) BuffRemoveDice(poolId uint64, duration int) BuffId {
	buff := makeBuffRemoveDice(duration, man.GetFatePool(), man.GetDiePool(poolId))
	return man.AddBuff(poolId, buff)
}

func (man *DiePoolManager) DecrementBuffTimesAndRemove() {
	for id, buff := range man.buffs {
		buff.ModifyDuration(-1)
		if buff.Duration() == 0 {
			delete(man.buffs, id)
		}
	}
}

func (man *DiePoolManager) DecrementBuffTimeAndRemove(id BuffId) {
	buff := man.buffs[id]
	buff.ModifyDuration(-1)
	if buff.Duration() == 0 {
		delete(man.buffs, id)
	}
}

func (man *DiePoolManager) ClearAllBuffsFromPool(poolId uint64) {
	pool := man.GetDiePool(poolId)
	for buffId, buff := range man.buffs {
		if buffId.poolId == poolId {
			pool.RemoveBuff(buff)
			delete(man.buffs, buffId)
		}
	}
}

func (man *DiePoolManager) RemoveBuffFromPool(poolId uint64, buffId BuffId) {
	pool := man.GetDiePool(poolId)
	buff, ok := man.buffs[buffId]
	if ok {
		pool.RemoveBuff(buff)
		delete(man.buffs, buffId)
	}
}
