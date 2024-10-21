package logger

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	pb "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/proto" //nolint:goanalysis_metalinter
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

type EventType string

const (
	StartEvent      EventType = "start"
	PlaceTileEvent  EventType = "place"
	ScoreEvent      EventType = "score"
	FinalScoreEvent EventType = "final_score"
)

func NewStartEntry(gameID uint32, gameSeed uint32, playerCount uint8, startingTile tiles.Tile) pb.Entry {
	tile := &pb.Tile{
		Position: &pb.Position{
			X: 0,
			Y: 0,
		},
		Features: []*pb.Feature{},
	}

	for _, f := range startingTile.Features {
		feature := &pb.Feature{
			Type:     pb.FeatureType(f.FeatureType),
			Modifier: pb.ModifierType(f.ModifierType),
			Side:     pb.Side(f.Sides),
			Meeple: &pb.Meeple{
				PlayerId:   uint32(elements.NoneMeeple),
				MeepleType: pb.MeepleType(elements.NonePlayer),
			},
		}
		tile.Features = append(tile.Features, feature)
	}

	return pb.Entry{
		Event: pb.EventType_EVENT_TYPE_START_EVENT,
		Content: &pb.Entry_StartEntryContent{
			StartEntryContent: &pb.StartEntryContent{
				GameId:      gameID,
				GameSeed:    gameSeed,
				PlayerCount: uint32(playerCount),
				StartTile:   tile,
			},
		},
	}
}

func NewPlaceTileEntry(playerID elements.ID, tile elements.PlacedTile) pb.Entry {
	move := &pb.Tile{
		Position: &pb.Position{
			X: int32(tile.Position.X()),
			Y: int32(tile.Position.Y()),
		},
		Features: []*pb.Feature{},
	}
	for _, f := range tile.Features {
		feature := &pb.Feature{
			Type:     pb.FeatureType(f.FeatureType),
			Modifier: pb.ModifierType(f.ModifierType),
			Side:     pb.Side(f.Sides),
			Meeple: &pb.Meeple{
				PlayerId:   uint32(f.Meeple.PlayerID),
				MeepleType: pb.MeepleType(f.Meeple.Type),
			},
		}
		move.Features = append(move.Features, feature)
	}

	return pb.Entry{
		Event: pb.EventType_EVENT_TYPE_PLACE_TILE_EVENT,
		Content: &pb.Entry_PlaceTileEntryContent{
			PlaceTileEntryContent: &pb.PlaceTileEntryContent{
				PlayerId: uint32(playerID),
				Move:     move,
			},
		},
	}
}

func NewScoreEntry(event EventType, scoreReport elements.ScoreReport) pb.Entry {
	scores := &pb.ScoreReport{
		ReceivedPoints:  []*pb.ReceivedPoints{},
		ReturnedMeeples: []*pb.ReturnedMeeple{},
	}

	for playerID, points := range scoreReport.ReceivedPoints {
		scores.ReceivedPoints = append(scores.ReceivedPoints, &pb.ReceivedPoints{
			PlayerId: uint32(playerID),
			Score:    points,
		})
	}

	for playerID, meeples := range scoreReport.ReturnedMeeples {
		for _, meeple := range meeples {
			scores.ReturnedMeeples = append(scores.ReturnedMeeples, &pb.ReturnedMeeple{
				Meeple: &pb.Meeple{
					PlayerId:   uint32(playerID),
					MeepleType: pb.MeepleType(meeple.Type),
				},
				Position: &pb.Position{
					X: int32(meeple.Position.X()),
					Y: int32(meeple.Position.Y()),
				},
			})
		}
	}

	var eventType pb.EventType

	if event == FinalScoreEvent {
		eventType = pb.EventType_EVENT_TYPE_FINAL_SCORE_EVENT
	} else {
		eventType = pb.EventType_EVENT_TYPE_SCORE_EVENT
	}

	return pb.Entry{
		Event: eventType,
		Content: &pb.Entry_ScoreEntryContent{
			ScoreEntryContent: &pb.ScoreEntryContent{
				ScoreReport: scores,
			},
		},
	}
}
