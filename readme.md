## Q/A AI SLACK BOT USING GO LANG, SLACK API, WITAI AND WOLFRAM API

# HOW IT WORKS
Bot listens to query on slack when it is tag
Retrieves the query and send it to a GO server
Server forwards the query to witAI which streamline the query into valid question 
witAI send back the streamline query back to the server
server redirect the query to wolfram to get an answer
wolfram provides the answer to the query sending it as data to the server which then redirect to the slack bot