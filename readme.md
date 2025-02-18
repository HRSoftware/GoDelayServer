A simple Go server that can mimic response times configured by the user


To use:
Lauch Go app

(Using postman or equiv) send a GET request to: (serveraddress):8080/delay?delayMs=2000

This response will take 2 seconds to responsd

Params:
delaysMs - An interger value representing the milliseconds the request to take to respond.
           If param is omitted(or not valid) the server will use the default response time.

Configuring the default response time
            Rather than providing a response time for each request, the 
            default response time can be set by sending the same request (and providing the delayMs 
            parameter) to a POST endpoint, ie:

            POST (serveraddress):8080/delay?delayMs=5000 

            This will set the default response time to 5 seconds for all following requests (Note - request that               provide a valid delayMs value will override the default value)
