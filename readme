# HTTP Requests and Go Routines
Learing how to use the http library in golang
to make Get and Post requests and also concurrency
and paralelism in go.

## Summary
Basically, it is a bot to check the status
of online websites periodically. Program operation:
 - Get the list of websites either hardcoded on the initial slice or any number of urls provided through the program's arguments.
 - Iterate through that list over and over to check if the website is up.
 - The end.

## What to have in mind
Accidentally made it a DDoS attack bot midway through
development. 
 - Don't use on small businesses and/or small websites as they likely won't have the means to block your high traffic and will crash. That is a crime.
 - Apply a Sleep() step between iterations if I'm ever going to use it as an actual status check bot.
 - Use multiple websites instead of just a few in the websites slice. That way the request load is smaller for the individual websites.
 - **Even if accidentally, DDoS attacks are a crime. Don't play with serious stuff.**