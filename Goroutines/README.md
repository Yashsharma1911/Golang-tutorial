<span align="center">
 <h1>Goroutines In Golang</h1>
</span>

## Concurrency vs Parallelism?
- Concurrency :-
    - Basically concurrency means that you suppose to perform multiple tasks but actually your performing a single task at a time.
    - For, Example Let assume that your are eating some food in front of your TV. Suddenly you started feeling hot so you have to start AC. Now in this scenario If you want to start the AC than you have to stop eating in order to turn on the AC. Similarly if you want to change you TV channel than you have to stop eating.
    - Basically from our story we can say that you are suppose to perform multiple tasks but at a particular moment you can perform only task.This is known as Concurrency.

- Parallelism :-
    - Parallelism means you are performing all tasks at a same time. Like you are eating, changing TV channel and also Turning on AC at the same time.


## Goroutines :-
- To achieve concurrency and parallelism in Golang, we need to understand the concept of Goroutines.
- Goroutines basically similar concept like Thread but, goroutines are not same as thread. Goroutines are actually small threads which are handled by GO runtime.
- Go runtime has the responsibility to assign or withdraw memory resources from Goroutines.
- Advantages of Goroutines :
    - They are lightweight.
    - Ability scale Seamlessly.
    - They are virtual threads.
    - Less memory requirement (2KB).
    - Provide additional memory to goroutine at runtime.

## Refrences

- [Golang Tutorial by Hitesh Chodhary](https://youtu.be/V-0ifUKCkBI)

- [Uderstatnding Golang and Goroutines](https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d)
