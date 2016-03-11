## How does your data structure allow developers to access and manipulate the data?

With `AmusementRide` Struct I am able to add names to the end of the lin using the `:addToEndOfLine` function, which uses the [Slice](http://blog.golang.org/go-slices-usage-and-internals) `:append` function to literally add to the end of the line array.
The only manipulation of the line available to a user is `:removeFromLine` and that is a standard loop and match using a regular expression library. The Go Standard library does not have a delete function for arrays, so I am using a Slice operator to append the line after the point where name(string) matched. 

## If a developer wanted to find a specific element in your data structure, how would you search for it?

I did not write a specific function to do that, but they can loop over the line attribue match for a specific 

## What other real-world data can each structure represent?
Train, Planes, and Automobiles ¯\_(ツ)_/¯

## Dicoveries in GO

I found out very quickly the choice for GO increase my potential for learning. This is because I had clear limitations with the Standard Library and had to hand build things I would normally take for granted in Ruby or JavaScript.  
