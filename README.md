# [TTK4145 - Real-time Programming](https://www.ntnu.edu/studies/courses/TTK4145)

*A collaborative project with [Simon Blindheim](https://github.com/simbli)*

**Course content:** Programming formalisms for real-time systems; POSIX, Ada, Java and Go. Threads/processes, synchronization and communication. Shared variable-based synchronization and resource control. Fault Tolerance, availability and consistency. Message-based synchronization, CSP and formal methods. Exercises and project.

**Project:** Create software for controlling `n` elevators working in parallel across `m` floors, with the following requirements:
- No orders are lost
- Multiple elevators should be more efficient than one
- An individual elevator should behave sensibly and efficiently
- The lights and buttons should function as expected

Read a detailed description of all requirements at the [TTK4145 repository](https://github.com/TTK4145/Project#elevator-project) (the description might change as the course is held every year).

**Solution:** Our solution was written in Google Go. We used TCP communication for all network communication, and a master/slave model for decision making between the elevators. The code consists of 6 modules:
- Network: A message that is to be sent over the network is converted to JSON and sent to the receiving elevator through the network.
- LiftState: The "brain" of the system. All information of the elevators is stored here, and decisions on what to do in every situation is taken here.
- FileManager: The FileManager writes down the IP address of all other elevators in the network in case the elevator should lose connection and then try to reconnect. Also, the internal orders in an elevator is written down in case an elevator malfunctions (is turned off and on again), so that these orders are never lost.
- Timer: Used to time how long a door should be kept open and how long it has been since a message has been received on the network.
- Driver: Changing the parameters against the physical components such as the motor and the door lights.
- Commander: This module is the organizer. It receives messages from LiftState, communicates with Driver and Timer, and passes messages to Network.
