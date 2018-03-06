package algs

import (
       "../arms"
       "../agent"

       "github.com/leesper/go_rng"
       "time"
)

var (
    seed     int64
    alpha    float64
    beta     float64
)

type Rng struct {
    beta      *rng.BetaGenerator
}

func NewRng(seed int64) *Rng {
    return &Rng{        
        beta:      rng.NewBetaGenerator(seed),
    }
}

func (rng *Rng) Beta(alpha, beta float64) float64 {
    return rng.beta.Beta(alpha, beta)
}

func Thompson_Sampling(agent agent.Agent, arm arms.Arms)(int){

     //Set parameters for Uniform dist.
     alpha = 1.0
     beta = 1.0

     seed = time.Now().UnixNano()     
     rng := NewRng(seed)     

     count := len(arm)
     m := make([]float64, count)
     m = Sample_Mean(agent, arm, count)
     n := make([]int, count)
     u := make([]float64, count)

     //Generate random numbers that follow a beta dist.
     for i:=0; i<count; i++{
     	 n[i] = arm[i].Count
     	 u[i] = rng.Beta(alpha + m[i], beta + float64(n[i]) - m[i])
     }

     max := 0.0
     s := 0
     for i:=0; i<count; i++{
     	 if max < u[i] {
	    max = u[i]
	    s = i
	 }
     }
     return s
}

