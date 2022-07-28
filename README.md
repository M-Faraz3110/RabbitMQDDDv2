## RabbitMQ Listener 
RabbitMQ Listener app with a Domain Driven Design structure based loosely on the template created [here](https://github.com/BetaLixT/goplates/tree/86876867228611b83274428687c6dc3235a38768/templates/dddv2).

## Structure 
Consists of two folders cmd and pkg. cmd contains the main.go and pkg contains just about everything else to do with the app.

### cmd/server 
Has main.go that starts the app and waits for all goroutines to finish.

### pkg 
Contains three folders namely app, domain and infra. App has the constructor, the Start() function called by main which is also where you can change the routingkey to listen to and the dependencyset for wire.  We have two folders within domain which are rabbitmq (basically sets up the channel rabbitmq will be listening on, while also using the logger) and tablestorage (which has everything we need to store the logs on azure table storage). 

#### app 
Contains the controllers folder (which will call the service), app.go, and since we're using wire for DI we also have wire.go and wire_gen.go. Wire.go uses the dependencyset(s) needed and wire_gen.go is automatically generated.

#### domain 
domain/rabbitmq has the service that will be exposed and it calls the function declared in tablestorage/externals. This means /tablestorage does not have a service of its own that needs to be exposed while /rabbitmq does not need anything in its externals.go. rabbitmqservice first creates the channel and then calls the function required to store the log in the db. Also starts the goroutine we need to make sure the channel stays open as long as the app stays running, which is until the app is stopped manually. domain.go creates a new dependencyset with the rabbitmqservice. 

#### infra 
Conatins the implementation of everything to be used by the above mentioned and a few other things, like the db setup, the logger and definition of the repo function used to store it in the db. Also has infra.go which again creates a dependency set.







