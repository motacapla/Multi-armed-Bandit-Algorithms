package arms

import (
       "math/rand"
       "time"
)

type Arm struct{
     Prob float64
     Count int     
}

type Arms []Arm 

//
// ベルヌーイ分布に従うアーム
//
// 不適切な引数なら-1, それ以外は1か0を返す
//
func Bernoulli_try(arm *Arm)(int){

     if arm.Prob < 0 || arm.Prob > 1 {
     	return(-1)
     }

     //Seed値を設定, 乱数生成
     rand.Seed(time.Now().UnixNano())
     var d float64 = rand.Float64()

     arm.Count++

     if arm.Prob > d {
     	return 1
     } else{
	return 0
     }

}