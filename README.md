# Prime Sieve
This is a repo containing a simple project written in Go. This was a starter project I undertook without the help of tutorials to help me learn Go.

The project is simple, it works out the number of Prime number below the given value. To accomplish this, it uses the "# [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)" algorithm to accomplish this. I chose this method as it's more of a programming "challenge" than just iterating over every value and calculating them all individually.

If you wish to run the project, follow these steps:
-   Ensure you have Go installed on your machine
    -   You can run  `go version`  in your terminal/command prompt to see if this is the case
    -   _if you don't_  you can download it from  [golang.org](https://golang.org/)
-   Clone the repo
-   Navigate to the root directory using terminal/command prompt
-   Run  `go build .`
-   By default it will calculate the number of Primes up to 1 million!

To calculate Primes to a different value, you can change the value of ``sieveSize`` to any value matching one of the keys in ``expectedCount``, or if you're feeling adventurous add the desired value to the slice, change the variable, run the script, and see if it's correct!

