A major television network wants to conduct a poll among its viewers to determine the best player after each game. To do this, it is necessary to develop a program that will be used by telephone operators to count the votes. Your team has been hired to develop this program using the C++ programming language. To count each vote, the operator will enter a number between 1 and 23, corresponding to the player's jersey number. A player number of zero indicates that the voting has ended. If an invalid number is entered, the program should ignore it, displaying a brief warning message, and asking for another number. After the voting ends, the program should display:

→ The total number of votes counted;

→ The numbers and respective votes of all players who received votes;

→ The percentage of votes for each of these players;

→ The number of the player chosen as the best player of the match, along with the number of votes and the percentage of votes given to him.

→ Note that invalid votes and the final zero should not be counted as votes. The result appears ordered by the player's number. The program should use arrays. The program should calculate the percentage of each player using a function. This function will receive two parameters: the number of votes for a player and the total number of votes.

The function will calculate the percentage and return the calculated value. Below is an example screen. The layout of the information should be as close as possible to the example. The data is fictitious and may change with each execution of the program. At the end, the program should also save the data relating to the voting result in a text file on disk, following the same layout shown on the screen.

Poll: Who was the best player?

Player Number (0=end): 9
Player Number (0=end): 10
Player Number (0=end): 9
Player Number (0=end): 10
Player Number (0=end): 11
Player Number (0=end): 10
Player Number (0=end): 50
Enter a value between 1 and 23 or 0 to exit!

Player Number (0=end): 9
Player Number (0=end): 9
Player Number (0=end): 0

Voting Result:

8 votes were counted.

Player Votes %
9 4 50.0%
10 3 37.5%
11 1 12.5%
The best player was number 9, with 4 votes, corresponding to 50% of the total votes.