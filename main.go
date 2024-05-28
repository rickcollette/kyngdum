package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Resources struct represents the available resources
type Resources struct {
	Food     int
	Wood     int
	Ore      int
	Livestock int
	Wheat    int
}

// Buildings struct represents the buildings in the kingdom
type Buildings struct {
	House    int
	Farm     int
	Townhall int
	Mill     int
	Blacksmith int
	Barracks int
}

// Citizens struct represents the citizens in the kingdom
type Citizens struct {
	Peasants   int
	Merchants  int
	Farmers    int
	Footsoldiers int
	Commanders int
}

// Game represents the state of the game
type Game struct {
	Resources
	Buildings
	Citizens
}

// NewGame initializes a new game state
func NewGame() *Game {
	return &Game{
		Resources: Resources{
			Food:     100,
			Wood:     50,
			Ore:      20,
			Livestock: 30,
			Wheat:    40,
		},
		Buildings: Buildings{
			House:    5,
			Farm:     2,
			Townhall: 1,
			Mill:     1,
			Blacksmith: 1,
			Barracks: 1,
		},
		Citizens: Citizens{
			Peasants:   10,
			Merchants:  2,
			Farmers:    3,
			Footsoldiers: 5,
			Commanders: 1,
		},
	}
}

// PrintStatus prints the current game status
func (g *Game) PrintStatus() {
	fmt.Println("Resources:")
	fmt.Printf("Food: %d, Wood: %d, Ore: %d, Livestock: %d, Wheat: %d\n", g.Food, g.Wood, g.Ore, g.Livestock, g.Wheat)
	fmt.Println("\nBuildings:")
	fmt.Printf("House: %d, Farm: %d, Townhall: %d, Mill: %d, Blacksmith: %d, Barracks: %d\n", g.House, g.Farm, g.Townhall, g.Mill, g.Blacksmith, g.Barracks)
	fmt.Println("\nCitizens:")
	fmt.Printf("Peasants: %d, Merchants: %d, Farmers: %d, Footsoldiers: %d, Commanders: %d\n", g.Peasants, g.Merchants, g.Farmers, g.Footsoldiers, g.Commanders)
}

// CollectTaxes collects taxes from peasants based on the number of buildings
func (g *Game) CollectTaxes() {
	maxTax := g.Peasants
	if maxTax > g.House {
		maxTax = g.House
	}

	taxAmount := rand.Intn(maxTax + 1)
	g.Food += taxAmount
	fmt.Printf("Collected %d food in taxes.\n", taxAmount)
}

// BuildBuilding builds a new building
func (g *Game) BuildBuilding(buildingType string) {
	switch buildingType {
	case "House":
		if g.Food >= 50 && g.Wood >= 20 {
			g.Food -= 50
			g.Wood -= 20
			g.House++
			fmt.Println("Built a new house!")
		} else {
			fmt.Println("Not enough resources to build a new house.")
		}
	case "Farm":
		if g.Wood >= 30 && g.Ore >= 20 {
			g.Wood -= 30
			g.Ore -= 20
			g.Farm++
			fmt.Println("Built a new farm!")
		} else {
			fmt.Println("Not enough resources to build a new farm.")
		}
	
	case "Townhall":
		if g.Wood >= 100 && g.Ore >= 50 && g.Livestock >= 30 {
			g.Wood -= 100
			g.Ore -= 50
			g.Livestock -= 30
			g.Townhall++
			fmt.Println("Built a new townhall!")
		} else {
			fmt.Println("Not enough resources to build a new townhall.")
		}
	
	case "Mill":
		if g.Wood >= 40 && g.Ore >= 30 {
			g.Wood -= 40
			g.Ore -= 30
			g.Mill++
			fmt.Println("Built a new mill!")
		} else {
			fmt.Println("Not enough resources to build a new mill.")
		}
	
	case "Blacksmith":
		if g.Wood >= 50 && g.Ore >= 40 {
			g.Wood -= 50
			g.Ore -= 40
			g.Blacksmith++
			fmt.Println("Built a new blacksmith!")
		} else {
			fmt.Println("Not enough resources to build a new blacksmith.")
		}
	
	case "Barracks":
		if g.Wood >= 60 && g.Ore >= 50 {
			g.Wood -= 60
			g.Ore -= 50
			g.Barracks++
			fmt.Println("Built a new barracks!")
		} else {
			fmt.Println("Not enough resources to build a new barracks.")
		}
	
	default:
		fmt.Println("Invalid building type.")
	}
}

// FeedPeasants feeds the peasants
func (g *Game) FeedPeasants() {
	if g.Food >= g.Peasants {
		g.Food -= g.Peasants
		fmt.Println("Peasants are well-fed!")
	} else {
		fmt.Println("Not enough food to feed the peasants.")
	}
}

// EndTurn simulates the end of a turn
func (g *Game) EndTurn() {
	// Implement any end of turn logic here
	fmt.Println("End of turn.")
}

// TitlePage displays the title page and handles game options
func TitlePage() {
    fmt.Println("Welcome to Kyngdum!")
    fmt.Println("1. New Game")
    fmt.Println("2. Load Game")
    fmt.Println("3. Delete Save File")
    fmt.Println("4. Exit")

    var choice int
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        // Create a new game
        var saveName string
        fmt.Println("Enter the name for the new game:")
        fmt.Scanln(&saveName)
        game := NewGame()
        err := SaveGame(game, saveName)
        if err != nil {
            fmt.Println("Failed to save the game:", err)
        } else {
            RunGame(game, saveName)
        }
    case 2:
        // Load an old game
        ListGames()
        fmt.Println("Enter the name of the save file you want to load:")
        var saveName string
        fmt.Scanln(&saveName)
        game, err := LoadGame(saveName)
        if err != nil {
            fmt.Println("Failed to load the game:", err)
        } else {
            RunGame(game, saveName)
        }
    case 3:
        // Delete a game save file
		ListGames()
        fmt.Println("Enter the name of the save file you want to delete:")
        var saveName string
        fmt.Scanln(&saveName)
        err := DeleteGame(saveName)
        if err != nil {
            fmt.Println("Failed to delete the save file:", err)
        }
    case 4:
        // Exit the game
        fmt.Println("Exiting game.")
        os.Exit(0)
    default:
        fmt.Println("Invalid choice. Please choose a valid option.")
        TitlePage() // Show title page again if choice is invalid
    }
}

// Main game loop
func RunGame(game *Game, saveFileName string) {
    fmt.Println("Welcome to Kingdom!")

    for {
        fmt.Println("\nWhat would you like to do?")
        fmt.Println("0. Current Status")
        fmt.Println("1. Collect taxes")
        fmt.Println("2. Build a new building")
        fmt.Println("3. Feed peasants")
        fmt.Println("4. Forage")
        fmt.Println("5. Craft")
        fmt.Println("6. End turn")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 0:
            game.PrintStatus()
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 1:
            game.CollectTaxes()
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 2:
            fmt.Println("Enter the building type (e.g., House, Farm, Townhall, etc.):")
            var buildingType string
            fmt.Scanln(&buildingType)
            game.BuildBuilding(buildingType)
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 3:
            game.FeedPeasants()
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 4:
            fmt.Println("\nWhat resource do you want to forage?")
            fmt.Println("1. Wood")
            fmt.Println("2. Wheat")
            fmt.Println("3. Ore")

            var forageChoice int
            fmt.Scanln(&forageChoice)
            switch forageChoice {
            case 1:
                game.ForageWood()
            case 2:
                game.ForageWheat()
            case 3:
                game.ForageOre()
            default:
                fmt.Println("Invalid choice.")
            }
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 5:
            fmt.Println("\nWhat do you want to craft?")
            fmt.Println("1. Food from Livestock")
            fmt.Println("2. Food from Wheat")

            var craftChoice int
            fmt.Scanln(&craftChoice)
            switch craftChoice {
            case 1:
                game.CraftFoodFromLivestock()
            case 2:
                game.CraftFoodFromWheat()
            default:
                fmt.Println("Invalid choice.")
            }
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        case 6:
            game.EndTurn()
            err := SaveGame(game, saveFileName) // Save game state after each action
            if err != nil {
                fmt.Println("Failed to save the game:", err)
            }
        default:
            fmt.Println("Invalid choice. Please choose a valid option.")
        }
    }
}


// LoadAndPrintTextFile loads and prints the contents of a text file
func LoadAndPrintTextFile(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	fmt.Println("Contents of", fileName+":")
	fmt.Println(string(content))
	return nil
}

// ForageWheat simulates foraging for wheat
func (g *Game) ForageWheat() {
 

    wheatFound := rand.Intn(11) // Randomly find between 0 to 10 units of wheat
    g.Wheat += wheatFound
    fmt.Printf("Foraged %d units of wheat.\n", wheatFound)
}

// ForageWood simulates foraging for wood
func (g *Game) ForageWood() {
 
    woodFound := rand.Intn(21) // Randomly find between 0 to 20 units of wood
    g.Wood += woodFound
    fmt.Printf("Foraged %d units of wood.\n", woodFound)
}

// ForageOre simulates foraging for ore
func (g *Game) ForageOre() {
    oreFound := rand.Intn(16) // Randomly find between 0 to 15 units of ore
    g.Ore += oreFound
    fmt.Printf("Foraged %d units of ore.\n", oreFound)
}
// CraftFoodFromLivestock crafts food from livestock
func (g *Game) CraftFoodFromLivestock() {
    if g.Livestock >= 5 {
        g.Livestock -= 5 // Consume 5 units of livestock
        foodProduced := rand.Intn(11) + 10 // Produce between 10 to 20 units of food
        g.Food += foodProduced
        fmt.Printf("Crafted %d units of food from livestock.\n", foodProduced)
    } else {
        fmt.Println("Not enough livestock to craft food.")
    }
}

// CraftFoodFromWheat crafts food from wheat
func (g *Game) CraftFoodFromWheat() {
    if g.Wheat >= 10 {
        g.Wheat -= 10 // Consume 10 units of wheat
        foodProduced := rand.Intn(21) + 20 // Produce between 20 to 40 units of food
        g.Food += foodProduced
        fmt.Printf("Crafted %d units of food from wheat.\n", foodProduced)
    } else {
        fmt.Println("Not enough wheat to craft food.")
    }
}

// SaveGame saves the game state to a file
func SaveGame(game *Game, saveName string) error {
	saveFileName := saveName + ".kdm"
	file, err := os.Create(saveFileName)
	if err != nil {
		return fmt.Errorf("failed to create save file: %v", err)
	}
	defer file.Close()

	// Encode game state to JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(game); err != nil {
		return fmt.Errorf("failed to encode game state: %v", err)
	}

	fmt.Println("Game saved successfully as", saveFileName)
	return nil
}

// LoadGame loads the game state from a file
func LoadGame(saveName string) (*Game, error) {
	saveFileName := saveName + ".kdm"
	file, err := os.Open(saveFileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open save file: %v", err)
	}
	defer file.Close()

	// Decode game state from JSON
	var game Game
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&game); err != nil {
		return nil, fmt.Errorf("failed to decode game state: %v", err)
	}

	fmt.Println("Game loaded successfully from", saveFileName)
	return &game, nil
}

// ListGames lists all save game files with the ".kdm" extension in the current directory
func ListGames() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Failed to list save game files:", err)
		return
	}

	fmt.Println("List of save game files:")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".kdm") {
			fmt.Println(strings.TrimSuffix(file.Name(), ".kdm"))
		}
	}
}


// DeleteGame deletes a save game file with the specified name
func DeleteGame(saveName string) error {
	saveFileName := saveName + ".kdm"

	// Check if the file exists
	_, err := os.Stat(saveFileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("save file '%s' does not exist", saveFileName)
	}

	// Prompt for confirmation
	fmt.Printf("Are you sure you want to delete '%s'? (yes/no): ", saveFileName)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read user input: %v", err)
	}

	answer = strings.ToLower(strings.TrimSpace(answer))
	if answer != "yes" {
		fmt.Println("Deletion canceled.")
		return nil
	}

	// Delete the file
	err = os.Remove(saveFileName)
	if err != nil {
		return fmt.Errorf("failed to delete save file: %v", err)
	}

	fmt.Println("Save file deleted successfully.")
	return nil
}

func main() {
	TitlePage()
}
