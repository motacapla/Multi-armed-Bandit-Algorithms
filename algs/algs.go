package algs

import (       
       "../arms"
       "../agent"
       
       //"fmt"
       "math"
)


func Epsilon_greedy(){

}

func UCB1(agent agent.Agent, arm arms.Arms, count int)(int){

     //まだ引いてないアームを優先的に選択
     for i:=0; i<count; i++{
     	 if arm[i].Count == 0 {
	    return i
	 }
     }

     //各アームの標本平均を算出
     x_mean := make([]float64, count)
     for i:=0; i<count; i++{
     	 for j := range agent.Reward{
	     x_mean[i] = x_mean[i] + float64(agent.Reward[i][j])
	 }
	 x_mean[i] = x_mean[i]/float64(arm[i].Count)
     }

     //確認
     //fmt.Println(x_mean)

     //UCB1の評価式より, アームを選定
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