package algs

import (
       "../arms"
       "../agent"     
)

func Sample_Mean(agent agent.Agent, arm arms.Arms, count int)([]float64){
     //各アームの標本平均を算出
     x_mean := make([]float64, count)
     for i:=0; i<count; i++{
	if arm[i].Count != 0 {
    	   for j := range agent.Reward{
	      x_mean[i] = x_mean[i] + float64(agent.Reward[i][j])
	   }
	      x_mean[i] = x_mean[i]/float64(arm[i].Count)
	   } else {
	      x_mean[i] = 0.0
	   }	
     }
     return x_mean
}