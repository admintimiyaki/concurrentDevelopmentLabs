//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Temur Rustamov C00280204
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff4(goNum int, total *atomic.Int64, max int, wg *sync.WaitGroup, theLock *sync.Mutex, c *chan struct{}) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	theLock.Lock()
	total.Add(1)
	if total.Load() == int64(max) { //last to arrive -signal others to go
		theLock.Unlock() //unlock before any potentially blocking code
		*c <- struct{}{}
		<-*c
	} else { //not all here yet we wait until signal
		theLock.Unlock() //unlock before any potentially blocking code
		<-*c
		*c <- struct{}{}
	} //end of if-else
	theLock.Lock()
	total.Add(-1)
	theLock.Unlock()
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
} //end-doStuff

func main() {
	var total atomic.Int64
	totalRoutines := 10
	//var wg sync.WaitGroup
	//n := 0
	//we will need some of these
	var theLock sync.Mutex
	//wg2.Add(n)
	theChan := make(chan struct{})       //use unbuffered channel in place of semaphore
	for k := 0; k < totalRoutines; k++ { //create the go Routines here
		var wg sync.WaitGroup // ⬅️ move inside the loop
		wg.Add(totalRoutines)
		total = atomic.Int64{}
		theChan = make(chan struct{})
		for j := 0; j < totalRoutines; j++ {
			go doStuff4(j, &total, totalRoutines, &wg, &theLock, &theChan)
		}

		wg.Wait() //wait for everyone to finish before exiting
		theLock = sync.Mutex{}
		//wg2.Done()
	} //end-main
	//wg2.Wait()
}
