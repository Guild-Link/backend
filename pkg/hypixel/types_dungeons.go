package hypixel

var catacombsXPTable = [...]float64{
	50, 75, 110, 160, 230, 330, 470, 670, 950, 1340,
	1890, 2665, 3760, 5260, 7380, 10300, 14400, 20000, 27600, 38000,
	52500, 71500, 97000, 132000, 180000, 243000, 328000, 445000, 600000, 800000,
	1065000, 1410000, 1900000, 2500000, 3300000, 4300000, 5600000, 7200000, 9200000, 12000000,
	15000000, 19000000, 24000000, 30000000, 38000000, 48000000, 60000000, 75000000, 93000000, 116250000,
}

type DungeonClass uint8

type DungeonClassLevel struct {
	Class DungeonClass
	Level float64
}

const (
	DungeonClassBerserk DungeonClass = iota
	DungeonClassHealer
	DungeonClassArcher
	DungeonClassMage
	DungeonClassTank
)

var dungeonClassNames = [...]string{
	"berserk",
	"healer",
	"archer",
	"mage",
	"tank",
}

type DungeonMode uint8

const (
	DungeonModeCatacombs DungeonMode = iota
	DungeonModeMasterCatacombs
)

type DungeonFloorStats struct {
	Mode         DungeonMode
	Floor        uint32
	Completions  float64
	PersonalBest *float64
}

type DungeonStats struct {
	Profile            *SkyBlockProfile
	ClassAverage       float64
	CatacombsLevel     float64
	SelectedClassLevel float64
	SecretsFound       float64
	SelectedClass      *DungeonClass
	Classes            []DungeonClassLevel
	Floors             []DungeonFloorStats
}
