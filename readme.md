A simple Go server that response time is configable by the user


To use:
Lauch Go app

(Using postman or equiv) send a GET request to: (serveraddress):8080/delay?delayMs=2000

Params:
delaysMs - An interger value representing the milliseconds the request to take to respond.
           If param is omitted(or not valid) the server will use the default response time.

Configuring the default response time
            Rather than providing a response time for each request, you can change the 
            default response time by sending the same request (and providing the delayMs 
            parameter) to a POST endpoint, ie:

            POST (serveraddress):8080/delay?delayMs=5000 

            will set the default response time to 5 seconds for all following requests(that
            omit the parameter value)
