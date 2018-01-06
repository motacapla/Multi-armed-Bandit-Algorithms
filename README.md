# Multi-armed-Bandit-Algorithms
Multi-armed Bandit Algorithms in Golang.

## What is this
This is the algorithm to solve multi-armed bandit problem. <br>

https://en.wikipedia.org/wiki/Multi-armed_bandit <br>


Though there are a lot of algorithms to solve this problem, I just only implemented UCB1 strategy one so far. <br>

To be simple, I suppose that all of success probabilities of arms will follow a bernoulli distribution. <br>

## How to Compile
```$ go build main.go```

## How to Use
```$ ./main ./arms/prob.d```  <br>
You can also modify the number of trials and the maximum number of arms. (They are defined in main.go as constant.)

If you want to add another arm, you need to modify "./arms/prob.d" and add success probability of the arm.

## Result
   ```
   ------------------------------
   Total Trials:  99
   Selected arm: 1
   Reward: 1
   [{Succecc probability, Number of selected}] =  [{0.2 18} {0.4 28} {0.8 36} {0.1 18}]
   ------------------------------
   ```

| Arms | Success Probability | Selected Num |
|:----:|:-------------------:|:------------:|
|Arm  A|         0.2         |      18      |
|Arm  B|         0.4         |      28      |
|Arm  C|         0.8         |      36      |
|Arm  D|         0.1         |      18      |

## Comments
This was my first time to write code in Golang. <br>

I would like to implement another algorithm like epsilon-greedy, softmax and Thompson Sampling. 
