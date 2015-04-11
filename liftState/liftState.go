package liftState

import (
	"net"
	."fmt"
	."time"
	."../network"
)

type elevator struct {
	floorNum int
	floorTarget int
	state string
}

func aliveBroadcast(commanderChan chan Message) {
	message := Message{}
	message.Type = "imAlive"
	for {
		Sleep(100 * Millisecond)
		commanderChan <- message
	}
}

func LiftState(networkReceive chan Message, commanderChan chan Message, aliveChan chan Message) {
	message := Message{}
	elev := make([]elevator, 1, ELEV_COUNT + 1)
	elev[0] = {0, 0, "Idle"}
	inside 	:= make([]int, FLOOR_COUNT+1)
	outUp 	:= make([]int, FLOOR_COUNT+1)
	outDown	:= make([]int, FLOOR_COUNT+1)

	message.Type "newID"
	commanderChan <- message
	
	for{
		select{
			case message = <- networkReceive:
				switch{
				case message.Type == "broadcast":
					go aliveBroadcast(commanderChan, &elev)

				case message.Type == "imAlive":
					aliveChan <- message

				case message.Type == "command":
					commanderChan <- message

				case message.Type == "newID":
					elev = append(elev, elevator{0, 0, "Idle"})
					
				case message.Type == "connectionChange":
					(elev)[message.From].onlineStatus = message.Online

				case message.Type == "rankChange":
					(elev)[message.From].rank = message.Value

				case message.Type == "newOrder":
					switch{
					case message.Content == "inside":
						inside[message.Floor] = 1
						
					case message.Content == "outsideUp":
						outUp[message.Floor] = 1
					case message.Content == "outsideDown":
						outDown[message.Floor] = 1
					}
						//CHECK IF NOT OWN ELEVATOR && INSIDE DON'T SEND SIGNALCHAN
					commanderChan <- message

					// Kjør kostfunksjon (hvis noen er idle)

				case message.Type == "deleteOrder":
					switch{
					case message.Content == "inside":
						inside[message.Floor] = 0
					case message.Content == "outsideUp":
						outUp[message.Floor] = 0
					case message.Content == "outsideDown":
						outDown[message.Floor] = 0
					}
						//CHECK IF NOT OWN ELEVATOR && INSIDE DON'T SEND SIGNALCHAN
					commanderChan <- message

				case message.Type == "newTarget":
					(elev)[message.From].floorTarget = message.Floor

				case message.Type == "stateUpdate":
					(elev)[message.From].state = message.Content
					//If State == "Idle": Kjør kostfunksjon (hvis flere bestillinger)
					//Legg til kode som gjør at Masters IP blir sendt om RecipientID pga timeOut for døra
				}
		}
	}
}