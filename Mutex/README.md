<span align="center">
 <h1>Mutex In Golang</h1>
</span>


## Mutex :-
- In Golang we are using multiple goroutines to perform multiple tasks at the same time.
- A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time. 
- This is done to prevent race conditions from happening. Sync package contains the Mutex.
- Mutext is also part of sync package in gosdk.
- There are two methods provided by Mutex :
    1. Lock() :
     - When a goroutine enters the critical section by using Lock() method we put the lock on that section. Which basically will ensure that another goroutine can not enter into that section till the critical section is unlocked.
    2. Unlock() :
     - After completion of work in the critical section we will unlock it by using Unlock() method so it can be accessible by another goroutine.

- There is another type of mutex available in Golang which is ReadWrite Mutex(RWMutex).
- Basically RWMutex will allow multiple goroutines to read the data of critical section. It will block other goroutines to get entered into the critical section when one goroutine is writing.

## Refrences

- [Golang Tutorial by Hitesh Chodhary](https://youtu.be/V-0ifUKCkBI)

- [sync package documentation](https://pkg.go.dev/sync#WaitGroup)
