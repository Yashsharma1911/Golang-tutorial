<span align="center">
 <h1>WaitGruops In Golang</h1>
</span>


## Waitgroups :-
- As we seen in last tutorial that if we launch a goroutine from a function than a function will not wait for groutine to comeback and function will complete it’s execution.
- To resolve this issue we used sleep() method. But for using sleep() method we have to explicitly tell that for how long goroutine will sleep which is headache.
- To overcome this issue we use Waitgroup which is part of ‘sync’ package. Basically waitgroup provides us some method and by using this we can control our goroutines.
- There are three methods provided by waitgroup :
    1. Add() :
     - Add() method basically takes number of goroutines which are going to be managed by waitgroup as function parameter.
    2. Wait() :
     - Wait() method is generally for waiting of goroutine. This method makes sure that all goroutines of wg are completed its execution and till then it will not allow the function to complete its execution.
    3. Done() :
     - Done() method notifies that goroutine has completed its execution.

## Refrences

- [Golang Tutorial by Hitesh Chodhary](https://youtu.be/V-0ifUKCkBI)

- [sync package documentation](https://pkg.go.dev/sync#WaitGroup)
