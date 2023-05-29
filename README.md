# Names counter
It's a CLI utility written in Go for counting name repetitions and printing result into `Stdout`.
## Usage
Using `--in` flag is necessary.
To read names from file do
````
.\cmd\namecounter\main --in <path-to-file>
````
To read names from `Stdin` do
````
.\cmd\namecounter\main --in -
````
Names can be printed in ascending or descending order of frequency. Names are printed randomly by default.
````
.\cmd\namecounter\main --in <path-to-file> --sort asc
or
.\cmd\namecounter\main --in <path-to-file> --sort desc
````