package algs

import (
       "../arms"
       "../agent"

       "fmt"
       "math/rand"
       "time"
)

func Epsilon_Greedy(agent agent.Agent, arm arms.Arms)(int){
     epsilon := 0.5

     rand.Seed(time.Now().UnixNano())
     var d float64 = rand.Float64()

     if d < epsilon {     	
     	fmt.Println("Randomly selected.")
        return rand.Intn(len(arm))
     } else {     
	fmt.Println("Best arm selected (if they are known).")
	
     count := len(arm)

     //各アームの標本平均を算出
     x_mean := Sample_Mean(agent, arm, count)

     //確認
     fmt.Println(x_mean)
	
     max := 0.0
     for i:=0; i<count; i++ {
     	 if max < x_mean[i] {
	    max = x_mean[i]
	 }
     }	
     var best_arms []int	    
	for i:=0; i<count; i++ {
	    if max == x_mean[i] {
	       best_arms = append(best_arms, i)
	    }
	}
	//これまでの値が最高となった候補アームの確認
	fmt.Println(best_arms)

	return best_arms[rand.Intn(len(best_arms))]
    }       
}

