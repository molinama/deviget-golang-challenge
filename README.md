# Golang-Challenge
Challenge test

We ask that you complete the following challenge to evaluate your development skills.

## The Challenge
Finish the implementation of the provided Transparent Cache package.

## Show your work

1.  Create a **Private** repository and share it with the recruiter ( please dont make a pull request, clone the private repository and create a new private one on your profile)
2.  Commit each step of your process so we can follow your thought process.
3.  Give your interviewer access to the private repo

## What to build
Take a look at the current TransparentCache implementation.

You'll see some "TODO" items in the project for features that are still missing.

The solution can be implemented either in Golang or Java ( but you must be able to read code in Golang to realize the exercise ) 

Also, you'll see that some of the provided tests are failing because of that.

The following is expected for solving the challenge:
* Design and implement the missing features in the cache
* Make the failing tests pass, trying to make none (or minimal) changes to them
* Add more tests if needed to show that your implementation really works
 
## Deliverables we expect:
* Your code in a private Github repo
* README file with the decisions taken and important notes

## Time Spent
We suggest not to spend more than 2 hours total, which can be done over the course of 2 days.  Please make commits as often as possible so we can see the time you spent and please do not make one commit.  We will evaluate the code and time spent.
 
What we want to see is how well you handle yourself given the time you spend on the problem, how you think, and how you prioritize when time is insufficient to solve everything.

Please email your solution as soon as you have completed the challenge or the time is up.

## Decisions Taken
* In order to control de maxage, I created new struct that will store the timestamp so I can use it to compare against max age.
I will setup that timestamp inside the getPrice method just after storing it into the map. I will preserv the getPrice original signature, so there is no need to modify the tests.
* In order to add asyc, I created a new function in the transparentChange so I can call it using a Go routine. Also I created a 
new buffered channel (to avoid blocking) to share the price and error. Finally I use a counter to determine when I have to close the channel.
* I update parallel test to valite the async call.
* Run a test coverage and create new tests.

## Important Notes
* First of all I will run the tests in order to validate what is going on.
* After any significant change I run the tests.
* After first iterate over the sync TODO and running tests I notice that I need to apply some condition to end the channel range, because it blocks.
* After running the parallel test modification I detected some issues to solve.
* In NewTransparentCache there should be a validation on the maxage to be sure not negative value, and also a Test.