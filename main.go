package main

import (
       "./arms"
       "./algs"
       "./agent"
       "fmt"
       "os"
       "bufio"
       "strconv"
)

const BUFSIZE = 1024
const MAX_NUM_OF_ARMS = 1024
const MAX_NUM_OF_TRIALS = 100

func main(){

     //ファイル読み込み, 第2引数にファイルパス
     var fp *os.File
     var err error
     if len(os.Args) < 2 {
     	fp = os.Stdin
     } else{
	fmt.Printf(">> read file: %s\n", os.Args[1])
	fp, err = os.Open(os.Args[1])
	if err != nil {
	   panic(err)
	}
	defer fp.Close()
     }
     scanner := bufio.NewScanner(fp)
     p := make([]float64, MAX_NUM_OF_ARMS)

     lineCount := 0
     for scanner.Scan(){
    	p[lineCount], err = strconv.ParseFloat(scanner.Text(),64)
	lineCount++
     }
     if err := scanner.Err(); err != nil{
     	panic(err)
     }
     
     //AIの総試行回数と各アームの報酬を初期化
     agent := agent.Agent{}
     agent.Trials = 0
     agent.Reward = make([][]int, lineCount)
     for i := range agent.Reward {
     	 agent.Reward[i] = make([]int, MAX_NUM_OF_TRIALS)
     }     
     for i:=0; i<lineCount; i++{
     	 for j:=0; j<MAX_NUM_OF_TRIALS; j++{
	     agent.Reward[i][j] = 0
	 }
     }     
          
     //各アームの成功確率と選択数を設定
     arm := make(arms.Arms, lineCount)     
     for i:=0; i<lineCount; i++{
          arm[i].Prob = p[i]
	  arm[i].Count = 0
     }     

     /*
     for ; agent.Trials<MAX_NUM_OF_TRIALS; agent.Trials++{     
         //UCB1で引くアームを選択
     	 s := algs.UCB1(agent, arm, lineCount)

	 //アームを引いて報酬を取得
	 reward := arms.Bernoulli_try(&arm[s])
	 //agent.Reward[s][arm[s].Count - 1] = reward
	 agent.Reward[s][agent.Trials] = reward
	 
	 //状態の確認
	 fmt.Println("------------------------------")
	 fmt.Println("Total Trials: ", agent.Trials)
         fmt.Println("Selected arm:", s)
	 fmt.Println("Reward:", reward)	 
	 fmt.Println("[{Succecc probability, Number of selected}] = ", arm)
	 fmt.Println("[Count]")
	 for i:=0; i<lineCount; i++{
	     fmt.Println("Arm:", i, "=>", arm[i].Count, "times")
	 }
	 fmt.Println("------------------------------\n")
      }
      */
      
     for ; agent.Trials<MAX_NUM_OF_TRIALS; agent.Trials++{     
     	 s := algs.Epsilon_Greedy(agent, arm)	 
	 reward := arms.Bernoulli_try(&arm[s])
	 agent.Reward[s][agent.Trials] = reward
	 //状態の確認
	 fmt.Println("------------------------------")
	 fmt.Println("Total Trials: ", agent.Trials)
         fmt.Println("Selected arm:", s)
	 fmt.Println("Reward:", reward)	 
	 fmt.Println("[{Succecc probability, Number of selected}] = ", arm)
	 fmt.Println("[Count]")
	 for i:=0; i<lineCount; i++{
	     fmt.Println("Arm:", i, "=>", arm[i].Count, "times")
	 }
	 fmt.Println("------------------------------\n")
     }
     
}