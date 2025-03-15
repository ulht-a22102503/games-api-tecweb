package db

import (
	"fmt"
	"log"
	"os"

	"github.com/martim-lusofona/games-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_DB_CONNECTION *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	_DB_CONNECTION = db
}

func Migrate() {
	if _DB_CONNECTION == nil {
		log.Println("Migrate NOK")
		return
	}
	err := _DB_CONNECTION.AutoMigrate(
		&models.Game{},
	)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Migrate OK")
	}
}

func Populate() {
	if _DB_CONNECTION == nil {
		log.Println("Populate NOK")
		return
	}

	games := []models.Game{
		{0, "The Legend of Zelda: Breath of the Wild", "An open-world adventure where you explore Hyrule.", "Action-Adventure"},
		{0, "Red Dead Redemption 2", "An epic tale of life in America’s unforgiving heartland.", "Action"},
		{0, "The Witcher 3: Wild Hunt", "A story-driven RPG set in a stunning open world.", "RPG"},
		{0, "Grand Theft Auto V", "A sprawling crime saga in Los Santos.", "Action"},
		{0, "Dark Souls III", "A challenging action RPG set in a dark fantasy world.", "RPG"},
		{0, "Minecraft", "A sandbox game where you create, explore, and survive.", "Sandbox"},
		{0, "Cyberpunk 2077", "An open-world RPG set in a cybernetic future.", "RPG"},
		{0, "God of War (2018)", "A mythological action-adventure featuring Kratos and Atreus.", "Action-Adventure"},
		{0, "Horizon Zero Dawn", "A post-apocalyptic action RPG featuring robotic creatures.", "Action"},
		{0, "DOOM Eternal", "A fast-paced shooter against demonic forces.", "Shooter"},
		{0, "Super Mario Odyssey", "A 3D platformer where Mario explores diverse worlds.", "Platformer"},
		{0, "Elden Ring", "A vast open-world RPG created by FromSoftware.", "RPG"},
		{0, "Bloodborne", "A Lovecraftian action RPG in a gothic world.", "RPG"},
		{0, "Persona 5", "A stylish JRPG about rebellious high school students.", "JRPG"},
		{0, "Metal Gear Solid V: The Phantom Pain", "A tactical espionage action game.", "Stealth"},
		{0, "Assassin's Creed Valhalla", "A Viking-themed open-world action game.", "Action-Adventure"},
		{0, "Half-Life: Alyx", "A VR-based first-person shooter.", "Shooter"},
		{0, "Celeste", "A touching platformer about climbing a mountain.", "Platformer"},
		{0, "Hollow Knight", "A deep and challenging Metroidvania.", "Metroidvania"},
		{0, "Stardew Valley", "A relaxing farming and life simulation game.", "Simulation"},
		{0, "Terraria", "A 2D sandbox adventure with crafting and exploration.", "Sandbox"},
		{0, "The Last of Us Part II", "A gripping narrative-driven action game.", "Action-Adventure"},
		{0, "Control", "A supernatural action game with telekinetic powers.", "Action"},
		{0, "Sekiro: Shadows Die Twice", "A fast-paced samurai action game.", "Action"},
		{0, "Resident Evil Village", "A survival horror game set in a mysterious village.", "Horror"},
		{0, "Monster Hunter: World", "A co-op action RPG about hunting massive creatures.", "Action-RPG"},
		{0, "Ghost of Tsushima", "A samurai action game set in feudal Japan.", "Action-Adventure"},
		{0, "Fire Emblem: Three Houses", "A tactical RPG with deep storytelling.", "Tactical RPG"},
		{0, "Xenoblade Chronicles", "A vast open-world JRPG.", "JRPG"},
		{0, "Splatoon 3", "A colorful multiplayer shooter.", "Shooter"},
		{0, "Bioshock Infinite", "A story-rich first-person shooter.", "Shooter"},
		{0, "Uncharted 4: A Thief's End", "A cinematic action-adventure game.", "Action-Adventure"},
		{0, "Ratchet & Clank: Rift Apart", "A dimension-hopping platformer.", "Platformer"},
		{0, "Halo Infinite", "The latest in the Halo shooter franchise.", "Shooter"},
		{0, "Forza Horizon 5", "An open-world racing game.", "Racing"},
		{0, "FIFA 23", "The latest iteration of the FIFA soccer series.", "Sports"},
		{0, "NBA 2K23", "A basketball simulation game.", "Sports"},
		{0, "Valorant", "A tactical hero shooter.", "Shooter"},
		{0, "League of Legends", "A competitive MOBA game.", "MOBA"},
		{0, "Counter-Strike: Global Offensive", "A competitive first-person shooter.", "Shooter"},
		{0, "Dota 2", "A deep and strategic MOBA.", "MOBA"},
		{0, "Apex Legends", "A battle royale shooter with unique characters.", "Battle Royale"},
		{0, "Fall Guys", "A fun and chaotic party battle royale.", "Party"},
		{0, "Among Us", "A social deduction game.", "Party"},
		{0, "Dead Cells", "A rogue-lite Metroidvania action game.", "Rogue-lite"},
		{0, "Slay the Spire", "A deck-building rogue-like game.", "Rogue-like"},
		{0, "The Sims 4", "A life simulation game.", "Simulation"},
		{0, "No Man’s Sky", "A space exploration game with infinite planets.", "Exploration"},
		{0, "Hades", "A rogue-like dungeon crawler.", "Rogue-like"},
		{0, "Disco Elysium", "A deep, story-driven RPG.", "RPG"},
	}

	for _, g := range games {

		err := _DB_CONNECTION.Where(models.Game{
			Title: g.Title,
		}).FirstOrCreate(&g).Error

		if err != nil {
			log.Println(err)
		} else {
			log.Println("Populate OK")
		}
	}

}

func GetDbConnection() *gorm.DB {
	if _DB_CONNECTION == nil {
		Connect()
	}

	return _DB_CONNECTION
}
