// Proto-based function interface for the simulator
package core

import (
	"github.com/Tereneckla/wotlk70/sim/core/proto"
	"github.com/Tereneckla/wotlk70/sim/core/stats"
)

/**
 * Returns character stats taking into account gear / buffs / consumes / etc
 */
func ComputeStats(csr *proto.ComputeStatsRequest) *proto.ComputeStatsResult {
	_, raidStats := NewEnvironment(csr.Raid, &proto.Encounter{})

	return &proto.ComputeStatsResult{
		RaidStats: raidStats,
	}
}

/**
 * Returns stat weights and EP values, with standard deviations, for all stats.
 */
func StatWeights(request *proto.StatWeightsRequest) *proto.StatWeightsResult {
	result := CalcStatWeight(request, stats.Stat(request.EpReferenceStat), nil)
	return result.ToProto()
}

func StatWeightsAsync(request *proto.StatWeightsRequest, progress chan *proto.ProgressMetrics) {
	go func() {
		result := CalcStatWeight(request, stats.Stat(request.EpReferenceStat), progress)
		progress <- &proto.ProgressMetrics{
			FinalWeightResult: result.ToProto(),
		}
	}()
}

/**
 * Runs multiple iterations of the sim with a full raid.
 */
func RunRaidSim(request *proto.RaidSimRequest) *proto.RaidSimResult {
	return RunSim(request, nil)
}

func RunRaidSimAsync(request *proto.RaidSimRequest, progress chan *proto.ProgressMetrics) {
	go RunSim(request, progress)
}
