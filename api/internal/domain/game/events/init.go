package events

import (
	"fmt"
	"math"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/kodmain/thetiptop/api/config"
	"github.com/kodmain/thetiptop/api/internal/application/hook"
	transfert "github.com/kodmain/thetiptop/api/internal/application/transfert/game"
	"github.com/kodmain/thetiptop/api/internal/domain/game/repositories"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/providers/database"
	"github.com/schollz/progressbar/v3"
)

const (
	DomainGameInit = "domain_game_init"
)

func init() {
	hook.Register(DomainGameInit, func() {
		HydrateDBWithTickets()
	})
}

func HydrateDBWithTickets() {
	repo := repositories.NewGameRepository(database.Get(config.GetString("services.game.database", config.DEFAULT)))
	require := 1500000
	dispatch := map[string]int{
		"X": 10,
		"Y": 40,
		"Z": 50,
	}

	// Compter le nombre actuel de tickets pour chaque `prize`
	existingCounts := make(map[string]int)
	for prize := range dispatch {
		count, err := repo.CountTicket(&transfert.Ticket{
			Prize: aws.String(prize),
		})
		if err != nil {
			panic(fmt.Sprintf("Failed to count tickets for %s: %v", prize, err))
		}

		existingCounts[prize] = count
	}

	// Calculer le nombre total de tickets déjà insérés
	totalExisting := 0
	for _, count := range existingCounts {
		totalExisting += count
	}

	// Vérifier s'il y a encore des tickets à insérer
	if totalExisting >= require {
		fmt.Printf("%d tickets are already ready\n", totalExisting)
		return
	}

	// Calculer le nombre de tickets nécessaires pour chaque `prize`
	remaining := require - totalExisting
	ticketsPerPrize := make(map[string]int)
	for prize, percent := range dispatch {
		expected := int(math.Round(float64(require) * float64(percent) / 100.0))
		ticketsPerPrize[prize] = expected - existingCounts[prize]
		fmt.Println("Ticket per prize", ticketsPerPrize[prize])
	}

	modulo := 1000
	fmt.Println("We need", remaining, "more tickets")

	// Initialiser la barre de progression avec le nombre total de tickets requis
	bar := progressbar.NewOptions(require,
		progressbar.OptionSetDescription("Inserting tickets..."),
		progressbar.OptionSetPredictTime(false),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionEnableColorCodes(true),
	)

	bar.Add(totalExisting)

	for prize, numTickets := range ticketsPerPrize {
		if numTickets <= 0 {
			continue // Passer au prochain `prize` si aucun ticket supplémentaire n'est nécessaire
		}

		tickets := []*transfert.Ticket{}
		for i := 0; i < numTickets; i++ {
			tickets = append(tickets, &transfert.Ticket{
				Prize: aws.String(prize),
			})

			// Insérer par lot lorsque le modulo est atteint ou à la fin de la boucle
			if len(tickets) >= modulo || i == numTickets-1 {
				if err := repo.CreateTickets(tickets); err != nil {
					panic(fmt.Sprintf("Failed to insert tickets for %s: %v", prize, err))
				}
				tickets = []*transfert.Ticket{}
				bar.Add(modulo)
			}
		}
	}

	fmt.Printf("\n%d tickets are ready\n", require)

}
