package algs

import (
       "../arms"
       "../agent"

       
       "fmt"       
       "math/rand"
       "time"              
)

func Softmax(agent agent.Agent, arm arms.Arms)(int){
     tau := 0.5
     
     rand.Seed(time.Now().UnixNano())
     var d float64 = rand.Float64()
     
     count := len(arm)     
     x_mean := Sample_Mean(agent, arm, count)     
     
     sum := 0.0
     for i:=0; i<count; i++ {
          x_mean[i] = x_mean[i] / tau
	  sum = sum + x_mean[i]
     }

     prob := make([]float64, count)
     if sum == 0 {
     	for i:=0; i<count; i++ {
     	    prob[i] = 1.0 / float64(count)
     	}
     } else{
     	for i:=0; i<count; i++ {
     	     prob[i] = x_mean[i] / sum
     	}
     }
     fmt.Println("P:", prob)
     
     prob_sum := make([]float64, count)     
     for i:=0; i<count; i++{
     	 for j:=0; j<=i; j++{
     	     prob_sum[i] = prob_sum[i] + prob[j]
	 }
     }

     fmt.Println("d:", d, "vs", prob_sum)

     for i:=0; i<count; i++ {
     	 if d < prob_sum[i] {
     	    return i
 	 }
     }

     return count-1
}