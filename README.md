# Vehicle Routing Problem with Time Window Solver with GA

Language : Go

## Assumption

- One single depot
- All vehicles have the same capacity

## Dataset

http://w.cba.neu.edu/~msolomon/problems.htm

## Technique

### Selection
- Weight Sum method
- Rank Sum method
- Pareto Ranking Selection

### Crossover

- Uniform Order Crossover (UOX)
- Partially Mapped Crossover (PMX)
- Route Crossover (RC)
- Best Cost Route Crossover (BCRC)

### Mutation

- Inversion Mutation

## Execution

Edit the parameters of run.sh and excute
```
./run.sh
```
