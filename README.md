# Go Data Frames supporting np like behaviours 

A very limited implementation of np like data frames, built to support the project

https://github.com/mdcfrancis/gomnist1d

To create the data set it is important the the random number generation is stable and the code is a fairly close 
match to the original mnist1d code style.

No parts of the original numpy are used in this code base. 

* Implements a simplified version of NpArray and NpStack, including many helper methods.
* Only supports one dimensional arrays and stacks of frames.
* Uses the same random number generator as the original code from numpy (golang impl).
* Implements np.RandChoice an extensions to randomkit to exactly match intn from numpy.

## References
* https://github.com/montanaflynn/stats : stats library is lightly used but allows for additional support if needed.
* https://github.com/pa-m/randomkit : without this library the random number generation would not match.

