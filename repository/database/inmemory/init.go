package inmemory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ak-17/ipl_stats/usecase/util"

	"github.com/ak-17/ipl_stats/model"
	"github.com/ak-17/ipl_stats/repository"
)

func New() (repository.Repository, error) {

	players, err := initPlayers()
	if err != nil {
		log.Printf("[inmemory][New][Error] error occured during initialization. err:%s", err.Error())
		return nil, err
	}

	mem := &inMemory{
		Players: players,
	}

	return mem, nil

}

func initPlayers() ([]model.Player, error) {
	var players []model.Player
	file, err := ioutil.ReadFile("repository/database/inmemory/ipl-2019.json")
	if err != nil {
		fmt.Printf(" error occured %s", err.Error())
		return players, err
	}

	err = json.Unmarshal(file, &players)
	if err != nil {
		fmt.Printf(" error occured %s", err.Error())
		return players, err
	}

	for i := 0; i < len(players); i++ {
		players[i].Points = util.GetFantasyPoints(players[i])
	}

	return players, nil

}
