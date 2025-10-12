Simple Barrier Implementation
Author: Temur Rustamov

Overview of the lab:
Here I have demonstrated a simple barrier implementation using channels and an atomic variable. There is a predefined number of threads that first of all do the 1st process, then they all wait except the last one who reads all messages from the channel, so they all could start working on the 2nd process. 