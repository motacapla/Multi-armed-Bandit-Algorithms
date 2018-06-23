package algs

import (       
       "../arms"
       "../agent"
       
       "fmt"
       "math"
)

func UCB1(agent agent.Agent, arm arms.Arms, count int)(int){

     //まだ引いてないアームを優先的に選択
     for i:=0; i<count; i++{
     	 if arm[i].Count == 0 {
	    return i
	 }
     }

     x_mean := Sample_Mean(agent, arm, count)

     //fmt.Println(x_mean)

     //choose the best arm index
     val := make([]float64, count)     
     s := 0
     max := 0.0
     for i:=0; i<count; i++{
     	 val[i] = x_mean[i] +  math.Sqrt( math.Log(float64(agent.Trials))/ (2 * float64(arm[i].Count)) )
	 if val[i] > max{
	    max = val[i]
	    s = i	    
	 }
     }
     return s
}