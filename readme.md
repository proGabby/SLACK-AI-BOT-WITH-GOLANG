## Q/A AI SLACK BOT USING GO LANG, SLACK API, WITAI AND WOLFRAM API

# HOW IT WORKS
1. Bot listens to query on slack when it is tag.
2. Retrieves the query and send it to a GO server.
3. Server forwards the query to witAI which streamline the query into valid question 
4. witAI send back the streamline query back to the server
5. server redirect the query to wolfram to get an answer
6. wolfram provides the answer to the query sending it as data to the server which then redirect to the slack bot
