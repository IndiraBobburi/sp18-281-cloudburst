# CMPE 281 - Team Hackathon Project<br/>

## TEAM CLOUDBURST<br/>

1.Indira Bobburi<br/>
2.Manasa Yedire<br/>
3.Aman Ojha <br/>
4.Kalikalyan Dash<br/>
5.Viraj Nilakh<br/>

## ARCHITECTURE DIAGRAM<br/>
![Architecture Diagram](Architecture_new.png)

## DESCRIPTION:<br/>
1.The client<br/>
Technology Stack: Nodejs, React, Redux<br/>
The client will take the user input and cascade the request to service A

2.Service A<br/>
Technology Stack:Nodejs<br/>
This layer will resolve the nature of the request and cascade it to appropriate microservice.

3.The gateway<br/>
Technology Stack: KONG/ AMAZON<br/>
The gateway will redirect the request to appropriate datastore as per the shard key

4.The datastore
Technology Stack: Cassandra<br/>
The customer information will be persisted in the datastores. Each store will be collection of 5 nodes capable of handling network partition.

**AKF Scale Cube**<br/>
**X-axis scaling:** The horizontal duplication involves scaling an application by running clones of the application on AWS behind an AWS elastic load balancer <br/>
**Y-axis scaling:** The functional decomposition is achieved by using microservices architecture. The verb-based decomposition approach is used where each service is implemented independently. <br/>
**Z-axis scaling:** The data is shared based on the zipcodes. The datastore will be shared by all the microservices. In order to retain the isolation of each service, the tables will be granted appropraite permissions<br/>

# Week 1

* Researched on database to be used.
* Researched on Go API.
* Researched on front end technology to be used.
* Discussed the architecture of the project.

# Week 2

* Developed Go APIs for user login.
* Developed Go API for user data.
* Developed front end APIs for menu and orders.


# Week 3
* Developed Go APIs for menu.
* Integrated Riak.
* Developed front end pages for menu and user login

# Week 4

* Developed Go APIs for user login.
* Developed Go APIs for user login.
* Developed front end pages for order placement.

# Week 5

* Deployment various AWS instances for front end, Riak and Casandra.
* Testing application.
* Tested sharding by creating cluster of Riak.
