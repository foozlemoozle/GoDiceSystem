package dice

import (
	"github.com/foozlemoozle/LinkedList"
	"math/rand"
	"time"
)

/*DIE DEFINITION*/

type Die struct {
	sides int
}

func (die *Die) Roll() int {
	rand.NewSource(time.Now().Unix())
	return rand.Intn(die.sides) + 1
}

/*END*/

/* DIE POOL DEFINITION*/

type IDiePool interface {
	Id() uint64
	HasDice() bool
	RollDice(context DieRollContext, diceCost int) int
	RemoveDice(amount int)
	AddDice(maxToAdd int)
	ResetToMaxSize()
	AddBuff(buff IDiePoolBuff)
	RemoveBuff(buff IDiePoolBuff)
}

func contestPools(actor IDiePool, actorContext DieRollContext, defender IDiePool, defenderContext DieRollContext) int {
	actorResult := actor.RollDice(actorContext, 1)
	defenderResult := defender.RollDice(defenderContext, 1)

	if actorResult > defenderResult {
		return Actor
	} else {
		return Target
	}
}

type DiePoolAction func(input int) int

const (
	Actor  = iota
	Target = iota
)

type DieRollContext int

const (
	Normal     = 0
	Defender   = 1
	Offender   = 1 << 1
	Interupt   = 1 << 2
	Fate       = 1 << 3
	VersusFate = 1 << 4
)

type diePool struct {
	id      uint64
	dice    LinkedList.IStack
	maxSize int
	sides   int
	buffs   LinkedList.IList
}

func DiePool(size int, sides int, id uint64) IDiePool {
	pool := &diePool{
		id:      id,
		maxSize: size,
		sides:   sides,
		dice:    LinkedList.Stack(),
		buffs:   LinkedList.List(),
	}

	for i := 0; i < size; i++ {
		pool.dice.Push(&Die{sides: sides})
	}

	return pool
}

func (pool *diePool) Id() uint64 {
	return pool.id
}

func (pool *diePool) HasDice() bool {
	return pool.dice.Count() > 0
}

func (pool *diePool) ChangePoolSize(newSize int) {
	diceAdded := newSize - pool.maxSize
	pool.maxSize = newSize

	if diceAdded > 0 {
		for i := 0; i < diceAdded; i++ {
			pool.dice.Push(&Die{sides: pool.sides})
		}
	} else {
		for i := 0; i < -1*diceAdded; i++ {
			pool.dice.Pop()
		}
	}
}

func (pool *diePool) RollDice(context DieRollContext, dieCost int) int {

	die := pool.dice.Peek().(*Die)
	numRollsContext := &BuffContextNumDiceRolls{
		roll:        die.Roll,
		highestRoll: 0,
		diceCost:    dieCost,
	}

	rollBonusContext := &BuffContextRollDice{
		bonus:   0,
		context: context,
	}

	pool.calculateBuff(numRollsContext)
	pool.calculateBuff(rollBonusContext)

	return numRollsContext.highestRoll + rollBonusContext.bonus
}

func (pool *diePool) RemoveDice(amount int) {
	removeContext := &BuffContextRemoveDice{
		amountToRemove: amount,
	}

	pool.calculateBuff(removeContext)
	amount = removeContext.amountToRemove
	for i := 0; i < amount; i++ {
		pool.dice.Pop()
	}
}

func (pool *diePool) AddDice(maxToAdd int) {
	maxSize := pool.calculateMaxSize()

	availableSpace := maxSize - pool.dice.Count()
	if availableSpace < maxToAdd {
		maxToAdd = availableSpace
	}

	for i := 0; i < maxToAdd; i++ {
		pool.dice.Push(&Die{sides: pool.sides})
	}
}

func (pool *diePool) ResetToMaxSize() {
	maxSize := pool.calculateMaxSize()
	diff := pool.dice.Count() - maxSize

	if diff < 0 {
		for i := 0; i < -1*diff; i++ {
			pool.dice.Push(&Die{sides: pool.sides})
		}
	} else {
		for i := 0; i < diff; i++ {
			pool.dice.Pop()
		}
	}
}

func (pool *diePool) calculateMaxSize() int {
	addContext := &BuffContextAddDice{
		maxSize: pool.maxSize,
	}

	pool.calculateBuff(addContext)

	return addContext.maxSize
}

func (pool *diePool) AddBuff(buff IDiePoolBuff) {
	pool.buffs.Enqueue(buff)
}

func (pool *diePool) RemoveBuff(buff IDiePoolBuff) {
	pool.buffs.Remove(buff)
}

func (pool *diePool) calculateBuff(buffContext interface{}) {
	iter := pool.buffs.Iterator()
	for cur, ok := iter.Current(); ok; cur, ok = iter.MoveNext() {
		cur.(IDiePoolBuff).Buff(buffContext)
	}
}

/*END*/
