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
//--------------------------------------------

package main

import (
	//"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	//"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(totalRoutines int, goNum int, wg *sync.WaitGroup, total *atomic.Int64, theLock *sync.Mutex, c *chan struct{}) bool {

	fmt.Println("Part A", goNum)
	time.Sleep(time.Second)
	theLock.Lock()
	total.Add(1)

	//we wait here until everyone has completed part A
	if total.Load() == int64(totalRoutines) {
		theLock.Unlock() // we need to unlock so comparison would be fair and go routine would not increment when we do the comparison
		i := 0
		for int64(i) < total.Load()-1 {
			i++
			<-*c
		}
	} else {
		theLock.Unlock()
		*c <- struct{}{}
	}

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	var total atomic.Int64
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	var theLock sync.Mutex

	theLock.Lock()

	c := make(chan struct{})
	for i := 0; i < totalRoutines; i++ {
		go doStuff(totalRoutines, i, &wg, &total, &theLock, &c)
	}
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting

}
